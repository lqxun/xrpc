package server

import (
	"context"
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) Call(ctx context.Context, name string, funcPtr interface{}) {
	// 反射初始化 funcPtr 函数原型
	fn := reflect.ValueOf(funcPtr).Elem()

	// RPC 调用远程的函数
	f := func(args []reflect.Value) []reflect.Value {
		// 参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}

		// 请求
		requestRPCData := &RPCData{
			Func: name,
			Args: inArgs,
		}

		select {
		case <-ctx.Done():
			panic(ctx.Err())
			return nil
		default:
			return handlerRequest(c.conn, &fn, requestRPCData)
		}
	}

	// 将 RPC 调用函数赋给 fn
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}

// 处理请求
func handlerRequest(conn net.Conn, fn *reflect.Value, requestRPCData *RPCData) []reflect.Value {
	cliSession := NewSession(conn)

	requestEncoded, err := encode(*requestRPCData)
	if err != nil {
		panic(err)
	}

	if err := cliSession.Write(requestEncoded); err != nil {
		panic(err)
	}

	// 响应
	response, err := cliSession.Read()
	if err != nil {
		panic(err)
	}
	respRPCData, err := decode(response)

	if err != nil {
		panic(err)
	}
	outArgs := make([]reflect.Value, 0, len(respRPCData.Args))
	for i, arg := range respRPCData.Args {
		if arg == nil {
			outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
		} else {
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
	}

	// 返回远程函数的返回值
	return outArgs
}
