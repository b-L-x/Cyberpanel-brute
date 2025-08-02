<!DOCTYPE html>
<html>
<h1 class="center">🔐 Cyberpanel Login Scanner </h1>

<p class="center">
    <strong>Cyberpanel credential testing tool</strong> combining brute-force protection bypass and high-performance scanning
</p>

<div class="center">
    <img src="https://img.shields.io/badge/Go-1.16%2B-blue" alt="Go Version">
    <img src="https://img.shields.io/badge/Threads-100%2B-green" alt="Multi-threaded">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License">
    <img src="https://img.shields.io/badge/Proxy-SOCKS5-red" alt="SOCKS5 Support">
</div>

<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">

![Diagram](https://i.imgur.com/AhhI4eN.png)

</body>
</html>

</div>

</body>
</html>
<h2>✨ Key Features</h2>

<ul>
  <li>⚡ <strong>High-Speed Scanning</strong> with goroutine worker pools</li>
  <li>🔓 <strong>CSRF Token Handling</strong> automatic detection and usage</li>
  <li>🛡️ <strong>SSL Bypass</strong> with configurable certificate verification</li>
  <li>🌐 <strong>Proxy Support</strong> SOCKS5 (Tor compatible)</li>
  <li>📊 <strong>Real-Time Stats</strong> success/failure tracking</li>
  <li>⏱️ <strong>Adaptive Timeouts</strong> for various network conditions</li>
  <li>📁 <strong>Results Export</strong> multiple output formats</li>
</ul>


<div class="features-grid">
    <div class="feature-card">
        <h3>⚡ Performance</h3>
        <ul>
            <li>Goroutine worker pool (100+ threads)</li>
            <li>HTTP connection reuse</li>
            <li>Connection pooling</li>
        </ul>
    </div>
    
<div class="features-container">
    <div class="feature-card">
        <h3>🔒 Security</h3>
        <ul>
            <li>Automatic CSRF token handling</li>
            <li>SSL verification bypass option</li>
            <li>SOCKS5 proxy support (Tor compatible)</li>
        </ul>
    </div>

Reporting</h3>
        <ul>
            <li>Real-time scanning statistics dashboard</li>
            <li>Detailed success/failure tracking</li>
            <li>Multiple export formats (TXT, JSON, CSV)</li>
            <li>Color-coded results for quick analysis</li>
        </ul>
    </div>
  
<h2>📦 Installation</h2>

<pre><code># Build from source
go mod init
go mod tidy
go build -o http_scanner main.go

# Or install directly
go install github.com/your-repo/http-scanner@latest</code></pre>

<h2>🚀 Usage Examples</h2>

<h3>Basic Scan</h3>
<pre><code>./http_scanner -url targets.txt -user usernames.txt -pass passwords.txt</code></pre>

<h3>Advanced Scan with Proxy</h3>
<pre><code>./http_scanner \
  -url targets.txt \
  -user admins.txt \
  -pass rockyou.txt \
  -threads 50 \
  -timeout 15 \
  -proxy 127.0.0.1:9050 \
  -output valid.txt</code></pre>

<h2>⚙️ Configuration Options</h2>

<table>
    <thead>
        <tr>
            <th>Parameter</th>
            <th>Description</th>
            <th>Default</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><code>-url</code></td>
            <td>File containing target URLs (one per line)</td>
            <td><em>Required</em></td>
        </tr>
        <tr>
            <td><code>-user</code></td>
            <td>Username dictionary file</td>
            <td><em>Required</em></td>
        </tr>
        <tr>
            <td><code>-pass</code></td>
            <td>Password dictionary file</td>
            <td><em>Required</em></td>
        </tr>
        <tr>
            <td><code>-threads</code></td>
            <td>Number of concurrent workers</td>
            <td>10</td>
        </tr>
        <tr>
            <td><code>-timeout</code></td>
            <td>Request timeout in seconds</td>
            <td>10</td>
        </tr>
        <tr>
            <td><code>-proxy</code></td>
            <td>SOCKS5 proxy address (ip:port)</td>
            <td>None</td>
        </tr>
        <tr>
            <td><code>-output</code></td>
            <td>File to save valid credentials</td>
            <td>None</td>
        </tr>
    </tbody>
</table>

<h2>📝 Sample Output</h2>

<pre><code>[START] Scanning initiated with 50 workers
[PROGRESS] 1425/5800 combinations tested (24.57%)
[SUCCESS] https://admin.example.com | superadmin | Admin@1234
[FAILED] https://test.site.net | guest | password123 (Invalid credentials)
[STATS] Success: 8 | Failed: 3124 | Elapsed: 6m22s
[COMPLETE] Scan finished! Valid credentials saved to results.txt</code></pre>

<h2>🛠️ Technical Implementation</h2>

<h3>Core Architecture</h3>
<pre><code class="language-go">type Scanner struct {
    stopScan      chan struct{}       // Graceful shutdown channel
    clientPool    sync.Pool           // HTTP client reuse
    httpTransport *http.Transport     // Custom transport config
    successCombos []Credential        // Valid credentials storage
    successLock   sync.Mutex          // Thread-safe access
}</code></pre>

<h3>Authentication Flow</h3>
<ol>
    <li>Fetch login page and extract CSRF token</li>
    <li>Prepare POST request with credentials</li>
    <li>Submit with proper headers and cookies</li>
    <li>Analyze response for success indicators</li>
    <li>Log and store successful attempts</li>
</ol>

<h2>📜 License</h2>
<p>MIT License - See <a href="LICENSE">LICENSE</a> for details.</p>

<div class="notice">
    <p>⚠️ <strong>Legal Notice:</strong> This tool is provided for <strong>authorized security testing only</strong>. Unauthorized use against systems you don't own is strictly prohibited.</p>
</div>

<h2>📞 Support</h2>
<p>Found an issue? <a href="https://github.com/your-repo/issues">Open a ticket</a></p>

</body>
</html>
