pam
===

POST API Mux

### INSTALL

```bash
$ go get github.com/gogap/pam
```

### USE

#### classic http server
```go
package main

import (
    "io"
    "log"
    "net/http"

    "github.com/gogap/pam"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello "+req.Header.Get(pam.REQ_X_API)+"!\n")
}

func main() {
    m := pam.New("hello")

    m.Post("gogap.hello.test", http.HandlerFunc(HelloServer))

    http.Handle("/", m)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

#### martini http server

```go
package main

import (
    "io"
    "net/http"

    "github.com/go-martini/martini"
    "github.com/gogap/pam"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello "+req.Header.Get(pam.REQ_X_API)+"!\n")
}

func main() {
    m := martini.Classic()

    apiMux := pam.New("hello")
    apiMux.Post("gogap.hello.test", http.HandlerFunc(HelloServer))

    m.Post("/hello", apiMux.APIMatcher())
    m.Get("/hello", func() string { return "GET from func" })

    m.Run()
}

```

```bash
$ curl -X POST -H "X-API: gogap.hello.test"  http://127.0.0.1:12345/hello
hello gogap.hello.test!
```