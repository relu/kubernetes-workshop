use axum::{
    extract::Request,
    response::{Html, IntoResponse},
    Router,
};
use std::env;
use std::net::SocketAddr;
use std::sync::atomic::{AtomicU64, Ordering};
use tower_http::trace::TraceLayer;
use tracing::info;

static REQUEST_COUNT: AtomicU64 = AtomicU64::new(0);

const TEMPLATE: &str = r#"<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kubernetes Workshop - Rust App</title>
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
            color: #CE422B;
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
            <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/rust/rust-original.svg" alt="Rust" />
        </div>

        <h1>Hello from <span class="language">Rust</span>!</h1>
        <p class="subtitle">{subtitle}</p>

        <div class="info">
            <div class="info-item">
                <span class="info-label">Version:</span>
                <span class="info-value">Rust {rust_version}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Path:</span>
                <span class="info-value">{path}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Hostname:</span>
                <span class="info-value">{hostname}</span>
            </div>
            <div class="info-item">
                <span class="info-label">Request Count:</span>
                <span class="info-value">{request_count}</span>
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
</html>"#;

async fn handler(req: Request) -> impl IntoResponse {
    let count = REQUEST_COUNT.fetch_add(1, Ordering::SeqCst) + 1;

    let rust_version = env!("CARGO_PKG_RUST_VERSION", "unknown");
    let path = req.uri().path();
    let hostname = hostname::get()
        .unwrap_or_else(|_| std::ffi::OsString::from("unknown"))
        .to_string_lossy()
        .to_string();
    let subtitle = env::var("SUBTITLE")
        .unwrap_or_else(|_| "Kubernetes Workshop Example Application".to_string());

    let html = TEMPLATE
        .replace("{rust_version}", rust_version)
        .replace("{path}", path)
        .replace("{hostname}", &hostname)
        .replace("{request_count}", &count.to_string())
        .replace("{subtitle}", &subtitle);

    info!("{} - GET {}", req.uri().path(), count);

    Html(html)
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let app = Router::new()
        .fallback(handler)
        .layer(TraceLayer::new_for_http());

    let port = env::var("PORT")
        .unwrap_or_else(|_| "3000".to_string())
        .parse::<u16>()
        .unwrap_or(3000);

    let addr = SocketAddr::from(([0, 0, 0, 0], port));

    info!("Rust server listening on {}", addr);

    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();

    axum::serve(listener, app)
        .with_graceful_shutdown(shutdown_signal())
        .await
        .unwrap();
}

async fn shutdown_signal() {
    let ctrl_c = async {
        tokio::signal::ctrl_c()
            .await
            .expect("failed to install Ctrl+C handler");
    };

    #[cfg(unix)]
    let terminate = async {
        tokio::signal::unix::signal(tokio::signal::unix::SignalKind::terminate())
            .expect("failed to install signal handler")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {},
        _ = terminate => {},
    }

    info!("Shutdown signal received, starting graceful shutdown");
}
