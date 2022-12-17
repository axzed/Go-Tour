package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// ReadJson 处理Json数据
func (c *Context) ReadJson(data any) error {
	// 读出 body
	body, err := ioutil.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	// 反序列化
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

// WriteJson 返回Json数据
// core method
func (c *Context) WriteJson(code int, resp any) error {
	c.W.WriteHeader(code)
	// 序列化
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// 写入响应
	_, err = c.W.Write(respJson)
	return err
}

// OkJson 返回成功Json数据
// extension method from WriteJson
func (c *Context) OkJson(resp any) error {
	return c.WriteJson(http.StatusOK, resp)
}

// BadRequest 返回失败Json数据
func (c *Context) BadRequest(resp any) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

// NewContext 创建Context
// 这是进一步的封装，将Context的创建和使用分离开
// 不需在server暴露Context的创建
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}
