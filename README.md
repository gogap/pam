pam
===

POST API Mux

### INSTALL

```bash
$ go get github.com/gogap/pam
```

### USE

```golang
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

```bash
$ curl -X POST -H "X-API: gogap.hello.test"  http://127.0.0.1:12345/hello
hello gogap.hello.test!
```