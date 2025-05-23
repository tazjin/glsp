package server

import (
	"io"
	"log/slog"

	"github.com/gorilla/websocket"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (self *Server) ServeStream(stream io.ReadWriteCloser, log *slog.Logger) {
	if log == nil {
		log = self.Log
	}
	log.Info("new stream connection")
	<-self.newStreamConnection(stream).DisconnectNotify()
	log.Info("stream connection closed")
}

func (self *Server) ServeWebSocket(socket *websocket.Conn, log *slog.Logger) {
	if log == nil {
		log = self.Log
	}
	log.Info("new web socket connection")
	<-self.newWebSocketConnection(socket).DisconnectNotify()
	log.Info("web socket connection closed")
}
