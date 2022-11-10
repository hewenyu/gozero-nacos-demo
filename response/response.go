package response

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go.opentelemetry.io/otel/trace"
)

// Body 统一结构体返回
type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	TraceID string      `json:"trace"`
	Data    interface{} `json:"data,omitempty"`
}

// Response 统一结构体返回
func Response(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	var traceID string

	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		traceID = spanCtx.TraceID().String()
	} else {
		traceID = "-1"
	}

	body.TraceID = traceID

	if err == nil {
		body.Msg = "OK"
		body.Success = true
		body.Data = resp

		httpx.OkJson(w, body)
		return
	}

	// if errors.As(err, target interface{})
	switch e := err.(type) {
	case *DefaultCodeError:
		body.Code = e.Code
		body.Msg = e.Error()
		body.Data = e.Data()
	default:
		body.Code = http.StatusInternalServerError
		body.Msg = e.Error()
	}
	httpx.OkJson(w, body)
}
