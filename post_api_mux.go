package pam

import (
	"net/http"
	"strings"
)

const (
	REQ_X_API = "X-API"
)

const (
	methodPost = "POST"
)

type apiHandler struct {
	apiName string
	http.Handler
}

type PostAPIMux struct {
	appName  string
	handlers map[string][]*apiHandler
}

func New(appName string) *PostAPIMux {
	return &PostAPIMux{appName: appName, handlers: make(map[string][]*apiHandler)}
}

func (p *PostAPIMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	served := false
	defer func() {
		if !served {
			http.Error(w, "Method Not Allowed", 405)
		}
	}()

	if r.Method != methodPost {
		return
	}

	for _, h := range p.handlers[methodPost] {
		apiName := r.Header.Get(REQ_X_API)
		if apiName == h.apiName {
			if strings.HasPrefix(r.URL.Path, "/"+p.appName+"/") ||
				r.URL.Path == "/"+p.appName {
				h.Handler.ServeHTTP(w, r)
				served = true
				return
			}
		}
	}

	if r.Method == methodPost {
		http.NotFound(w, r)
		served = true
		return
	}
}

func (p *PostAPIMux) add(apiName string, h http.Handler) {
	p.handlers[methodPost] = append(p.handlers[methodPost], &apiHandler{apiName, h})
}

func (p *PostAPIMux) Post(apiName string, h http.Handler) {
	p.add(apiName, h)
}
