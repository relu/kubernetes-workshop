import os
import sys
import signal
import threading
from flask import Flask, request, render_template_string

app = Flask(__name__)

# Signal handler for graceful shutdown
def signal_handler(sig, frame):
    signame = 'SIGINT' if sig == signal.SIGINT else 'SIGTERM'
    print(f'\n{signame} received, shutting down gracefully...', flush=True)
    sys.exit(0)

signal.signal(signal.SIGINT, signal_handler)
signal.signal(signal.SIGTERM, signal_handler)

# Thread-safe request counter
request_count = 0
request_lock = threading.Lock()

TEMPLATE = """<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kubernetes Workshop - Python App</title>
    <link rel="icon" type="image/svg+xml" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/python/python-original.svg">
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
            color: #3776AB;
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
            <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/python/python-original.svg" alt="Python" />
        </div>

        <h1>Hello from <span class="language">Python</span>!</h1>
        <p class="subtitle">{{ subtitle }}</p>

        <div class="info">
            <div class="info-item">
                <span class="info-label">Version:</span>
                <span class="info-value">{{ python_version }}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Path:</span>
                <span class="info-value">{{ path }}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Hostname:</span>
                <span class="info-value">{{ hostname }}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Count:</span>
                <span class="info-value">{{ request_count }}</span>
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
</html>"""

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def index(path):
    global request_count

    with request_lock:
        request_count += 1
        current_count = request_count

    pod_name = os.environ.get('NAME', 'unknown')
    python_version = f"{sys.version_info.major}.{sys.version_info.minor}.{sys.version_info.micro}"
    hostname = os.uname().nodename
    request_path = request.path
    subtitle = os.environ.get('SUBTITLE', 'Kubernetes Workshop Example Application')

    return render_template_string(
        TEMPLATE,
        pod_name=pod_name,
        python_version=python_version,
        path=request_path,
        hostname=hostname,
        subtitle=subtitle,
        request_count=current_count
    )

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 3000))
    app.run(host='0.0.0.0', port=port)
