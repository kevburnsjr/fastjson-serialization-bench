package main

import (
	"time"

	"github.com/kpango/fastime"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

var contentTypeJson = []byte("application/json; charset=UTF-8")

func main() {
	var t = fastime.SetFormat(time.UnixDate)
	var arenaPool = &fastjson.ArenaPool{}
	fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.SetBytesV("Date", t.FormattedNow())
		ctx.Response.Header.SetBytesV("Content-Type", contentTypeJson)
		arena := arenaPool.Get()
		s := arena.NewString("Hello, World!")
		o, _ := arena.NewObject().Object()
		o.Set("message", s)
		ctx.SetBody(o.MarshalTo(nil))
		arenaPool.Put(arena)
	})
}
