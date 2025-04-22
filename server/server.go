package server

import (
	"log/slog"
	"time"

	"github.com/tazjin/glsp"
)

var DefaultTimeout = time.Minute

//
// Server
//

type Server struct {
	Handler glsp.Handler
	Debug   bool

	Log              *slog.Logger
	Timeout          time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	StreamTimeout    time.Duration
	WebSocketTimeout time.Duration
}

func NewServer(handler glsp.Handler, logger *slog.Logger, debug bool) *Server {
	return &Server{
		Handler:          handler,
		Debug:            debug,
		Log:              logger,
		Timeout:          DefaultTimeout,
		ReadTimeout:      DefaultTimeout,
		WriteTimeout:     DefaultTimeout,
		StreamTimeout:    DefaultTimeout,
		WebSocketTimeout: DefaultTimeout,
	}
}
