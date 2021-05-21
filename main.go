package main

import (
	"time"

	"github.com/kpango/fastime"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

var hdrDate = []byte("Date")
var hdrContentType = []byte("Content-Type")
var hdrContentTypeJson = []byte("application/json; charset=UTF-8")

func main() {
	var t = fastime.SetFormat(time.UnixDate)
	var arenaPool = &fastjson.ArenaPool{}
	fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.SetBytesKV(hdrDate, t.FormattedNow())
		ctx.Response.Header.SetBytesKV(hdrContentType, contentTypeJson)
		arena := arenaPool.Get()
		arena.Reset()
		o, _ := arena.NewObject().Object()
		o.Set("message", arena.NewStringBytes([]byte("Hello, World!")))
		ctx.SetBody(o.MarshalTo(nil))
		arenaPool.Put(arena)
	})
}
