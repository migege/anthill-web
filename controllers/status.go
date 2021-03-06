package controllers

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	ahlog "github.com/migege/anthill/proto/log"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	cli = ahlog.NewLoggerClient("", client.NewClient(client.RequestTimeout(120*time.Second), client.Registry(registry.NewRegistry(registry.Addrs("ah.mayibot.com:8500")))))
)

type StatusController struct {
	baseController
}

func (this *StatusController) Get() {
	tag := this.GetString("tag")

	this.TplName = "status.tpl"
	this.Data["Tag"] = tag
}

func (this *StatusController) Stream() {
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup websocket connection:", err)
		return
	}
	defer ws.Close()

	var req ahlog.Info
	if err = ws.ReadJSON(&req); err != nil {
		beego.Error(err)
		return
	}

	infos := strings.Split(req.Info, ",")
	if len(infos) < 3 {
		beego.Error("invalid request")
		return
	}

	go func() {
		for {
			if _, _, err := ws.NextReader(); err != nil {
				break
			}
		}
	}()

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id":    infos[0],
		"X-Host":       infos[1],
		"X-Process-Id": infos[2],
	})
	stream, err := cli.Status(ctx, &req)
	if err != nil {
		beego.Error(err)
		return
	}
	defer stream.Close()

	for {
		rsp, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				beego.Error(err)
				return
			}
			break
		}

		err = ws.WriteJSON(rsp)
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				break
			} else {
				beego.Error(err)
				return
			}
		}
	}
}
