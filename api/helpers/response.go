package helpers

import (
	"github.com/valyala/fasthttp"
)

type sResponse struct {
	statusCode  int
	data        []byte
	ctx         *fasthttp.RequestCtx
	contentType string
}

func Response(ctx *fasthttp.RequestCtx) *sResponse {
	return &sResponse{
		ctx:         ctx,
		statusCode:  fasthttp.StatusOK,
		contentType: "text/plain; charset=utf-8",
	}
}

func (resp *sResponse) Text(text string) *sResponse {
	resp.data = []byte(text)
	return resp
}

func (resp *sResponse) Json(data any) *sResponse {
	resp.contentType = "application/json"
	resp.data = ToJson(data)
	return resp
}

func (resp *sResponse) StatusCode(code int) *sResponse {
	resp.statusCode = code
	return resp
}

func (resp *sResponse) Send() {
	resp.ctx.Response.SetStatusCode(resp.statusCode)
	resp.ctx.Response.Header.Set("Content-Type", resp.contentType)
	resp.ctx.SetBody(resp.data)
}
