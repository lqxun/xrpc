package server

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

// XRpc 服务列表暂时用本地内存存储
type XRpc struct {
	services   map[string]reflect.Value
	success    chan struct{}
	middleware map[string]*Middleware
	listener   *net.Listener
	conn       *net.Conn
}

// NewXRpc 创建服务端，维护服务列表
func NewXRpc() *XRpc {
	a := &XRpc{
		services:   make(map[string]reflect.Value),
		success:    make(chan struct{}),
		middleware: make(map[string]*Middleware),
	}

	return a
}

// Register 注册服务
func (s *XRpc) Register(name string, service interface{}, middlewares ...Handler) {
	if _, ok := s.services[name]; ok {
		panic("XRpc already registered")
	}

	f := reflect.ValueOf(service)
	s.services[name] = f
	s.middleware[name] = NewMiddleware(middlewares...)
}

// ListenAndServe 启动服务，指定坚挺地址
func (s *XRpc) ListenAndServe(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	s.listener = &l
	fmt.Println("Listening on", addr)
	close(s.success)
	for {
		conn, err := l.Accept()
		s.conn = &conn
		if err != nil {
			panic(err)
		}
		go s.handler(conn)
	}
}

// Success 阻塞，等待链接成功
func (s *XRpc) Success() {
	<-s.success
}

func (s *XRpc) handler(conn net.Conn) {
	session := NewSession(conn)

	// 处理请求数据
	data, err := handleRequest(session)
	if err != nil {
		log.Println("request handle error:", err)
	}

	// 处理请求
	f, ok := s.services[data.Func]
	if !ok {
		log.Println("XRpc not found:", data.Func)
		return
	}

	args := make([]reflect.Value, len(data.Args))
	for i, arg := range data.Args {
		args[i] = reflect.ValueOf(arg)
	}
	request := func() {
		fmt.Println("handler...")
		// 处理响应
		responseRaw := f.Call(args)

		responseEncode, err := handleResponse(data, responseRaw)
		if err != nil {
			log.Println("response handle error:", err)
			return
		}

		// 发送响应
		err = session.Write(responseEncode)
		if err != nil {
			log.Println("response write error:", err)
			return
		}
	}

	s.middleware[data.Func].Handlers(handleMiddleware(request))
	s.middleware[data.Func].Run()
}

func handleMiddleware(f func()) Handler {
	return func(m *Middleware) {
		f()
		m.Next()
	}
}

// 处理请求，解析请求数据
func handleRequest(session *Session) (*RPCData, error) {
	raw, err := session.Read()
	if err != nil {
		log.Println("data read error:", err)
		return nil, err
	}

	data, err := decode(raw)
	if err != nil {
		log.Println("data decode error:", err)
		return nil, err
	}
	return &data, nil
}

// 处理响应数据
func handleResponse(req *RPCData, data []reflect.Value) ([]byte, error) {
	responseArgs := make([]interface{}, len(data))
	for i, arg := range data {
		responseArgs[i] = arg.Interface()
	}

	responseData := RPCData{
		Func: req.Func,
		Args: responseArgs,
	}

	responseEncode, err := encode(responseData)
	if err != nil {
		log.Println("response encode error:", err)
		return nil, err
	}

	return responseEncode, nil
}
