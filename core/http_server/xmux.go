package http_server

import "net/http"

type Middleware func(http.Handler) http.Handler

type XMux struct {
	*http.ServeMux
	middlewares []Middleware
}

func NewXMux() *XMux {
	return &XMux{
		ServeMux: http.NewServeMux(),
	}
}

func handlerMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}

func (m *XMux) Use(middlewares ...Middleware) {
	m.middlewares = append(m.middlewares, middlewares...)
}

func (m *XMux) Handle(pattern string, handler http.Handler) {
	handler = handlerMiddlewares(handler, m.middlewares...)
	m.ServeMux.Handle(pattern, handler)
}

func (m *XMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	newHandler := handlerMiddlewares(handler, m.middlewares...)
	m.ServeMux.Handle(pattern, newHandler)
}

func (m *XMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h, _ := m.Handler(r)
	h = handlerMiddlewares(h, m.middlewares...)
	h.ServeHTTP(w, r)
}
