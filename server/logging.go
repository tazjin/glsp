package server

import (
	"fmt"
	"log/slog"
	"strings"
)

type JSONRPCLogger struct {
	log *slog.Logger
}

// ([jsonrpc2.Logger] interface)
func (self *JSONRPCLogger) Printf(format string, v ...any) {
	self.log.Debug(fmt.Sprintf(strings.TrimSpace(format), v...), "source", "jsonrpc2")
}
