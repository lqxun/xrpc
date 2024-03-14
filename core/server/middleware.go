package server

type Handler func(m *Middleware)

type Middleware struct {
	curr     int
	handlers []Handler
}

func NewMiddleware(handlres ...Handler) *Middleware {
	return &Middleware{
		curr:     -1,
		handlers: handlres,
	}
}

// Handlers 设置中间件
func (m *Middleware) Handlers(handlers ...Handler) {
	m.handlers = append(m.handlers, handlers...)
}

func (m *Middleware) Next() {
	// 这个++不能放到for中，
	m.curr++
	for m.curr < len(m.handlers) {
		m.handlers[m.curr](m)
	}
}

func (m *Middleware) Run() {
	m.Next()
}
