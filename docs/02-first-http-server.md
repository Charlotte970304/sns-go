# First HTTP Server – Building the Initial HTTP Service

## 1. Modifying `main.go` into an HTTP Server

In the previous stage, the program only printed a message and exited
immediately. While this verified that the project was set up correctly,
it did not create a long-running backend service.

To build an SNS backend, the entry point must start an HTTP server
so that the program can remain running and respond to incoming requests.

The `main.go` file was modified as follows:

```go
func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        fmt.Fprint(w, `{"status":"ok"}`)
    })

    http.ListenAndServe(":8080", mux)
}
```

With this change, the program no longer exits immediately.
Instead, it starts an HTTP server and waits for requests.

---

## 2. Meaning of `ServeMux`

`ServeMux` is short for **HTTP request multiplexer**.

Its responsibility is to:
- Receive incoming HTTP requests
- Examine the request path (for example, `/health`)
- Route the request to the corresponding handler function

In practice, `ServeMux` acts as a routing table that connects
URL paths to handler functions.

Creating a new `ServeMux`:

```go
mux := http.NewServeMux()
```

At this point, the multiplexer is empty and contains no routing rules.
Routes are added by registering handlers.

---

## 3. Understanding `HandleFunc` and Handler Function Design

The `HandleFunc` method has the following signature:

```
func (mux *ServeMux) HandleFunc(
    pattern string,
    handler func(http.ResponseWriter, *http.Request),
)
```

### Anonymous Function Style

```go
mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"status":"ok"}`)
})
```

**Advantages**
- Compact and easy to read for small handlers

**Disadvantages**
- Harder to reuse and test

---

### Named Function Style

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"status":"ok"}`)
}

mux.HandleFunc("/health", healthHandler)
```

**Advantages**
- Clear intent and reusable
- Scales better as the project grows

**Disadvantages**
- Slightly more boilerplate

---

## 4. The Three Parts of an HTTP Response

### 1. Status Line

```
HTTP/1.1 200 OK
```

Set by:

```go
w.WriteHeader(http.StatusOK)
```

---

### 2. Headers

```
Content-Type: application/json
```

Set by:

```go
w.Header().Set("Content-Type", "application/json")
```

---

### 3. Body

```json
{"status":"ok"}
```

Written by:

```go
fmt.Fprint(w, `{"status":"ok"}`)
```

---

## 5. Running the Server

From the project root directory, run:

```bash
go run cmd/api/main.go
```
The program will continue running and start an HTTP server
listening on port 8080.

To verify the server is working, open a browser or run:
```bash
curl http://localhost:8080/health
```
Expected response:
```json
{"status":"ok"}
```
If this response is returned, the HTTP server is running correctly.