package main

import "net/http"

// handleFunc is a function that handles a request.
// It writes reply headers and data to the ResponseWriter and then return.
// Returning signals that the request is finished; it is not valid to use
// the ResponseWriter or read from the Request.Body after or concurrently
// with the completion of the handleFunc call.
// 最主要的是为了控制Context由框架创建，而不是由业务代码创建
type handleFunc func(ctx *Context)

type Server interface {
	// method to start the server
	Route(method string, pattern string, handle func(ctx Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

// Route 注册路由
// handleFunc是一个函数，它处理一个请求
// 它将回复头和数据写入ResponseWriter，然后返回
func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	// HandleFunc registers the handleFunc function for the given pattern in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	// HandleFunc is a wrapper around DefaultServeMux.HandleFunc.
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		// 创建Context
		// 由框架创建，而不是由业务代码创建
		// 使用闭包去handle这个context
		ctx := NewContext(w, r)
		//handleFunc(&Context{
		//	W: w,
		//	R: r,
		//})
		handleFunc(ctx)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}
