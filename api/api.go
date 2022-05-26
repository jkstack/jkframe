package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// Context api handler callback context
type Context struct {
	w      http.ResponseWriter
	r      *http.Request
	values context.Context
}

// NewContext create handler callback context
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:      w,
		r:      r,
		values: context.Background(),
	}
}

// OK response api handle ok
func (ctx *Context) OK(payload interface{}) {
	ctx.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.w).Encode(map[string]interface{}{
		"code":    0,
		"payload": payload,
	})
}

// ERR response api handle failed
func (ctx *Context) ERR(code int, msg string) {
	ctx.w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx.w).Encode(map[string]interface{}{
		"code": code,
		"msg":  msg,
	})
}

// Body response data from byte slice
func (ctx *Context) Body(data []byte) {
	ctx.w.Write(data)
}

// BodyFrom response data from reader
func (ctx *Context) BodyFrom(r io.Reader) {
	io.Copy(ctx.w, r)
}

// NotFound response not found error
func (ctx *Context) NotFound(what string) {
	panic(NotFound(what))
}

// Timeout response timeout error
func (ctx *Context) Timeout() {
	panic(Timeout{})
}

// URI get request uri
func (ctx *Context) URI() string {
	return ctx.r.URL.Path
}

// ServeFile response file data from dir
func (ctx *Context) ServeFile(dir string) {
	http.ServeFile(ctx.w, ctx.r, dir)
}

// HTTPNotFound response not found error with http_code=404
func (ctx *Context) HTTPNotFound(what string) {
	http.Error(ctx.w, what+" not found", http.StatusNotFound)
}

// HTTPServiceUnavailable response service unavailable error with http_code=503
func (ctx *Context) HTTPServiceUnavailable(msg string) {
	http.Error(ctx.w, msg, http.StatusServiceUnavailable)
}

// HTTPTimeout response timeout error with http_code=408
func (ctx *Context) HTTPTimeout() {
	http.Error(ctx.w, "timeout", http.StatusRequestTimeout)
}

// HTTPConflict response conflict data error with http_code=409
func (ctx *Context) HTTPConflict(msg string) {
	http.Error(ctx.w, msg, http.StatusConflict)
}

// HTTPForbidden response data forbidden error with http_code=403
func (ctx *Context) HTTPForbidden(msg string) {
	http.Error(ctx.w, msg, http.StatusForbidden)
}

// Token get token from request header by X-Token field
func (ctx *Context) Token() string {
	return ctx.r.Header.Get("X-Token")
}

// SetContentType set response Content-Type of data
func (ctx *Context) SetContentType(str string) {
	ctx.w.Header().Set("Content-Type", str)
}

// SetContentDisposition set response filename
func (ctx *Context) SetContentDisposition(name string) {
	ctx.w.Header().Set("Content-Disposition", `attachment; filename="`+name+`"`)
}

// AddValue add key/value to this context
func (ctx *Context) AddValue(k, v interface{}) {
	ctx.values = context.WithValue(ctx.values, k, v)
}

// Value get value from key in this context
func (ctx *Context) Value(k interface{}) interface{} {
	return ctx.values.Value(k)
}

// RawCallback callback function by this context
func (ctx *Context) RawCallback(cb func(http.ResponseWriter, *http.Request)) {
	cb(ctx.w, ctx.r)
}

// Method get request method
func (ctx *Context) Method() string {
	return ctx.r.Method
}
