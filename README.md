<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>HTTP Login Scanner Pro</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
            line-height: 1.6;
            color: #24292e;
            max-width: 1012px;
            margin: 0 auto;
            padding: 20px;
        }
        h1, h2, h3 {
            margin-top: 24px;
            margin-bottom: 16px;
            font-weight: 600;
        }
        h1 {
            font-size: 2em;
            border-bottom: 1px solid #eaecef;
            padding-bottom: 0.3em;
        }
        h2 {
            font-size: 1.5em;
            border-bottom: 1px solid #eaecef;
            padding-bottom: 0.3em;
        }
        h3 {
            font-size: 1.25em;
        }
        pre {
            background-color: #f6f8fa;
            border-radius: 6px;
            padding: 16px;
            overflow: auto;
        }
        code {
            font-family: SFMono-Regular, Consolas, "Liberation Mono", Menlo, monospace;
            background-color: rgba(27, 31, 35, 0.05);
            border-radius: 3px;
            padding: 0.2em 0.4em;
            font-size: 85%;
        }
        table {
            border-spacing: 0;
            border-collapse: collapse;
            display: block;
            width: 100%;
            overflow: auto;
            margin-bottom: 16px;
        }
        th {
            font-weight: 600;
            padding: 6px 13px;
            border: 1px solid #dfe2e5;
            background-color: #f6f8fa;
        }
        td {
            padding: 6px 13px;
            border: 1px solid #dfe2e5;
        }
        img {
            max-width: 100%;
            box-sizing: content-box;
            background-color: #fff;
        }
        .center {
            text-align: center;
        }
        .notice {
            background-color: #f8f8f8;
            padding: 10px;
            border-left: 4px solid #ff9800;
            margin: 20px 0;
        }
        .features-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 15px;
            margin: 20px 0;
        }
        .feature-card {
            background: #f5f5f5;
            padding: 15px;
            border-radius: 5px;
        }
    </style>
</head>
<body>

<h1 class="center">🔐 HTTP Login Scanner Pro</h1>

<p class="center">
    <strong>Professional credential testing tool</strong> combining brute-force protection bypass and high-performance scanning
</p>

<div class="center">
    <img src="https://img.shields.io/badge/Go-1.16%2B-blue" alt="Go Version">
    <img src="https://img.shields.io/badge/Threads-100%2B-green" alt="Multi-threaded">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License">
    <img src="https://img.shields.io/badge/Proxy-SOCKS5-red" alt="SOCKS5 Support">
</div>

<h2>✨ Key Features</h2>

<div class="features-grid">
    <div class="feature-card">
        <h3>⚡ Performance</h3>
        <ul>
            <li>Goroutine worker pool (100+ threads)</li>
            <li>HTTP connection reuse</li>
            <li>Connection pooling</li>
        </ul>
    </div>
    
    <div class="feature-card">
        <h3>🔒 Security</h3>
        <ul>
            <li>CSRF token handling</li>
            <li>SSL verification bypass</li>
            <li>SOCKS5 proxy support</li>
        </ul>
    </div>

    <div class="feature-card">
        <h3>📊 Reporting</h3>
        <ul>
            <li>Real-time statistics</li>
            <li>Success/failure tracking</li>
            <li>File export capabilities</li>
        </ul>
    </div>
</div>

<h2>📦 Installation</h2>

<pre><code># Build from source
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