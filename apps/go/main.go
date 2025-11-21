package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"
)

const templateStr = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kubernetes Workshop - Go App</title>
    <link rel="icon" type="image/svg+xml" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }

        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
            padding: 60px;
            max-width: 600px;
            text-align: center;
        }

        .logo {
            width: 120px;
            height: 120px;
            margin: 0 auto 30px;
        }

        h1 {
            color: #2d3748;
            font-size: 2.5em;
            margin-bottom: 10px;
        }

        .language {
            color: #00ADD8;
            font-weight: bold;
        }

        .subtitle {
            color: #718096;
            font-size: 1.2em;
            margin-bottom: 30px;
        }

        .info {
            background: #f7fafc;
            border-radius: 10px;
            padding: 20px;
            margin-top: 30px;
        }

        .info-item {
            display: flex;
            justify-content: space-between;
            padding: 10px 0;
            border-bottom: 1px solid #e2e8f0;
        }

        .info-item:last-child {
            border-bottom: none;
        }

        .info-label {
            color: #718096;
            font-weight: 500;
        }

        .info-value {
            color: #2d3748;
            font-family: 'Monaco', 'Courier New', monospace;
            font-size: 0.9em;
        }

        .footer {
            margin-top: 30px;
            padding-top: 20px;
            border-top: 1px solid #e2e8f0;
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 10px;
        }

        .kubernetes-badge {
            display: inline-flex;
            align-items: center;
            gap: 8px;
            background: #326ce5;
            color: white;
            padding: 8px 16px;
            border-radius: 20px;
            font-size: 0.9em;
        }

        .k8s-logo {
            width: 20px;
            height: 20px;
            filter: brightness(0) invert(1);
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo">
            <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" alt="Go" />
        </div>

        <h1>Hello from <span class="language">Go</span>!</h1>
        <p class="subtitle">{{.Subtitle}}</p>

        <div class="info">
            <div class="info-item">
                <span class="info-label">Version:</span>
                <span class="info-value">{{.GoVersion}}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Path:</span>
                <span class="info-value">{{.Path}}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Hostname:</span>
                <span class="info-value">{{.Hostname}}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Count:</span>
                <span class="info-value">{{.RequestCount}}</span>
            </div>
        </div>

        <div class="footer">
            <div class="kubernetes-badge">
                <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/kubernetes/kubernetes-plain.svg" alt="Kubernetes" class="k8s-logo" />
                <span>Running on Kubernetes</span>
            </div>
        </div>
    </div>
</body>
</html>`

var tmpl = template.Must(template.New("index").Parse(templateStr))
var requestCount uint64

type PageData struct {
	PodName      string
	GoVersion    string
	Path         string
	Hostname     string
	Subtitle     string
	RequestCount uint64
}

func handler(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddUint64(&requestCount, 1)

	podName := os.Getenv("NAME")
	if podName == "" {
		podName = "unknown"
	}

	subtitle := os.Getenv("SUBTITLE")
	if subtitle == "" {
		subtitle = "Kubernetes Workshop Example Application"
	}

	hostname, _ := os.Hostname()

	data := PageData{
		PodName:      podName,
		GoVersion:    runtime.Version(),
		Path:         r.URL.Path,
		Hostname:     hostname,
		Subtitle:     subtitle,
		RequestCount: count,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)

	log.Printf("%s - %s %s", r.RemoteAddr, r.Method, r.URL.Path)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", handler)

	server := &http.Server{
		Addr: ":" + port,
	}

	// Channel to listen for interrupt or terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		log.Printf("Go server listening on port %s...", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	log.Println("Shutdown signal received, shutting down gracefully...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
