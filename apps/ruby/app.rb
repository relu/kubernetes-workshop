require 'sinatra'
require 'socket'

set :port, ENV['PORT'] || 3000
set :bind, '0.0.0.0'

# Signal handlers for graceful shutdown
trap('INT') do
  puts "\nSIGINT received, shutting down gracefully..."
  exit(0)
end

trap('TERM') do
  puts "\nSIGTERM received, shutting down gracefully..."
  exit(0)
end

# Thread-safe request counter
$request_count = 0
$request_mutex = Mutex.new

TEMPLATE = <<-HTML
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kubernetes Workshop - Ruby App</title>
    <link rel="icon" type="image/svg+xml" href="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/ruby/ruby-original.svg">
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
            color: #CC342D;
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
            <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/ruby/ruby-original.svg" alt="Ruby" />
        </div>

        <h1>Hello from <span class="language">Ruby</span>!</h1>
        <p class="subtitle"><%= @subtitle %></p>

        <div class="info">
            <div class="info-item">
                <span class="info-label">Version:</span>
                <span class="info-value"><%= @ruby_version %></span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Path:</span>
                <span class="info-value"><%= @path %></span>
            </div>
            <div class="info-item">
                <span class="info-label">Hostname:</span>
                <span class="info-value"><%= @hostname %></span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Count:</span>
                <span class="info-value"><%= @request_count %></span>
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
</html>
HTML

configure do
  set :host_authorization, permitted_hosts: []
end

get '/*' do
  $request_mutex.synchronize do
    $request_count += 1
  end

  @pod_name = ENV.fetch('NAME', 'unknown')
  @ruby_version = RUBY_VERSION
  @path = request.path
  @hostname = Socket.gethostname
  @subtitle = ENV.fetch('SUBTITLE', 'Kubernetes Workshop Example Application')
  @request_count = $request_count

  erb TEMPLATE
end
