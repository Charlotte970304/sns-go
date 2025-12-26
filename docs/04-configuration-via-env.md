# Configuration via Environment Variables – Making Server Port Configurable

## 1. Why the Server Port Should Be Configurable

In previous stages, the HTTP server port was hard-coded in the program.
While this works for local development, it is not suitable for real-world usage.

In practice, the port a service listens on is often determined by:
- The execution environment (local, staging, production)
- Deployment platforms or container runtimes
- Port conflicts with other services

Because of this, the port should not be treated as application logic.
Instead, it should be treated as configuration provided at runtime.

Using environment variables allows the same program binary
to run in different environments without code changes.

---

## 2. Complete Program with Configurable Port

The server was updated to read the port number from an environment variable.
If the variable is not set, a default value is used.

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/health", healthHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    addr := ":" + port
    log.Println("starting http server on", addr)

    err := http.ListenAndServe(addr, mux)
    if err != nil {
        log.Fatal(err)
    }
}
```

This approach separates configuration from code,
while keeping a sensible default for local development.

---

## 3. Usage Examples

### Running Without Setting `PORT`

If no environment variable is provided,
the server uses the default port.

```bash
go run cmd/api/main.go
```

Expected output:

```
starting http server on :8080
```

---

### Running With a Custom Port (Windows PowerShell)

```powershell
$env:PORT="3000"
go run cmd/api/main.go
```

Expected output:

```
starting http server on :3000
```

---

### Running With a Custom Port (Unix / macOS)

```bash
PORT=3000 go run cmd/api/main.go
```

Expected output:

```
starting http server on :3000
```
