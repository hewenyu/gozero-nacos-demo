package handler

import (
	"net/http"

	"git-devops.zeekrlife.com/dr/response"
	"github.com/hewenyu/gozero-nacos-demo/api/internal/logic"
	"github.com/hewenyu/gozero-nacos-demo/api/internal/svc"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		response.Response(r.Context(), w, resp, err)
	}
}
