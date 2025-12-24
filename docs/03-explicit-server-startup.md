# Explicit Server Startup – Making Server Lifecycle Observable

## 1. Why Server Startup Should Be Explicit

In earlier stages, the HTTP server was started silently.
If the server failed to start, the program would terminate
without clearly indicating what happened.

By explicitly logging server startup and handling startup errors,
the program makes its lifecycle visible.

This is important because:
- A backend service may fail to start due to configuration or environment issues
- Startup failures should be immediately observable
- The program should not continue running in an invalid state

Explicit startup logging helps distinguish between
"the server is running" and "the server failed to start".

---

## 2. Explicit Server Startup in Code
The explicit startup behavior is implemented in `main`:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"status":"ok"}`)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/health", healthHandler)

    log.Println("starting http server on :8080")

    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}

```
Key points:
- Startup is explicitly logged before the server begins listening
- ListenAndServe blocks on success and returns an error on failure
- A startup failure causes the program to exit immediately

---
## 3. How Startup Failure Was Tested

To verify that startup failures are handled correctly,
a failure scenario was intentionally created.

### Test Method

The HTTP server was started twice using the same port (`:8080`).

1. Start the server normally
2. Start the same server again in another terminal

Because only one process can bind to a port at a time,
the second startup attempt fails.

---

### Successful Startup Result

When the server starts successfully, the output shows:

```
2025/12/24 10:46:32 starting http server on :8080
```

The program continues running and waits for incoming requests.

---

### Failed Startup Result

When the server fails to start, the output shows an error similar to:

```
2025/12/24 10:48:11 starting http server on :8080
2025/12/24 10:48:11 listen tcp :8080: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
exit status 1
```

This confirms that:
- Startup failure is detected
- The failure reason is visible
- The program exits with a non-zero status code


