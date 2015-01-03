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
