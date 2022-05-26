package api

import (
	"mime/multipart"
	"strconv"
	"strings"
)

/* XInt get required int argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not int, panic with BadParam error
*/
func (ctx *Context) XInt(name string) int {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return int(n)
}

/* OInt get unrequired int argument by name
   if not found argument from request, return default value by def argument
   if value by name is not int, return zero value of int value
*/
func (ctx *Context) OInt(name string, def int) int {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseInt(v, 10, 64)
	return int(n)
}

/* XInt32 get required int32 argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not int32, panic with BadParam error
*/
func (ctx *Context) XInt32(name string) int32 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return int32(n)
}

/* OInt32 get unrequired int32 argument by name
   if not found argument from request, return default value by def argument
   if value by name is not int32, return zero value of int32 value
*/
func (ctx *Context) OInt32(name string, def int32) int32 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseInt(v, 10, 32)
	return int32(n)
}

/* XInt get required int64 argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not int64, panic with BadParam error
*/
func (ctx *Context) XInt64(name string) int64 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return int64(n)
}

/* OInt64 get unrequired int64 argument by name
   if not found argument from request, return default value by def argument
   if value by name is not int64, return zero value of int64 value
*/
func (ctx *Context) OInt64(name string, def int64) int64 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseInt(v, 10, 64)
	return int64(n)
}

/* XUInt get required uint argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not uint, panic with BadParam error
*/
func (ctx *Context) XUInt(name string) uint {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return uint(n)
}

/* OUint get unrequired uint argument by name
   if not found argument from request, return default value by def argument
   if value by name is not uint, return zero value of uint value
*/
func (ctx *Context) OUInt(name string, def uint) uint {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseUint(v, 10, 64)
	return uint(n)
}

/* XUint32 get required uint32 argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not uint32, panic with BadParam error
*/
func (ctx *Context) XUInt32(name string) uint32 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return uint32(n)
}

/* OUint32 get unrequired uint32 argument by name
   if not found argument from request, return default value by def argument
   if value by name is not uint32, return zero value of uint32 value
*/
func (ctx *Context) OUInt32(name string, def uint32) uint32 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseUint(v, 10, 64)
	return uint32(n)
}

/* XUint64 get required uint64 argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not uint64, panic with BadParam error
*/
func (ctx *Context) XUInt64(name string) uint64 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	n, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return uint64(n)
}

/* OUint64 get unrequired uint64 argument by name
   if not found argument from request, return default value by def argument
   if value by name is not uint64, return zero value of uint64 value
*/
func (ctx *Context) OUInt64(name string, def uint64) uint64 {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	n, _ := strconv.ParseUint(v, 10, 64)
	return uint64(n)
}

/* XStr get required string argument by name
   if not found argument from request, panic with MissingParam error
*/
func (ctx *Context) XStr(name string) string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	return v
}

/* OStr get unrequired string argument by name
   if not found argument from request, return default value by def argument
*/
func (ctx *Context) OStr(name, def string) string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	return v
}

func transCsv(data []string) []string {
	for i, v := range data {
		data[i] = strings.ReplaceAll(v, "%2c%", ",")
	}
	return data
}

/* XCsv get required csv argument by name
   if not found argument from request, panic with MissingParam error

   in each string value of "%2c%" will change to symbol ","
*/
func (ctx *Context) XCsv(name string) []string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	return transCsv(strings.Split(v, ","))
}

/* OCsv get unrequired csv argument by name
   if not found argument from request, return default value by def argument

   in each string value of "%2c%" will change to symbol ","
*/
func (ctx *Context) OCsv(name string, def []string) []string {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	return transCsv(strings.Split(v, ","))
}

/* XBool get required bool argument by name
   if not found argument from request, panic with MissingParam error
   if value by name is not bool, panic with BadParam error
*/
func (ctx *Context) XBool(name string) bool {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		panic(MissingParam(name))
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(BadParam(name + ":" + err.Error()))
	}
	return b
}

/* OBool get unrequired bool argument by name
   if not found argument from request, return default value by def argument
   if value by name is not bool, return zero value of bool value
*/
func (ctx *Context) OBool(name string, def bool) bool {
	v := ctx.r.FormValue(name)
	if len(v) == 0 {
		return def
	}
	b, _ := strconv.ParseBool(v)
	return b
}

// File get upload file argument by name
func (ctx *Context) File(name string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.r.FormFile(name)
}
