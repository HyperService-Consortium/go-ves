package server

import (
	"fmt"
	"github.com/HyperService-Consortium/go-ves/ves/config"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"go.uber.org/zap/zapcore"
)

func (srv *Server) instantiateLogger() error {
	var err error
	srv.Logger, err = logger.NewZapLogger(logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
	if err != nil {
		return err
	}
	srv.Module.Provide(config.ModulePath.Minimum.Global.Logger, srv.Logger)
	return nil
}

func (srv *Server) println(a ...interface{}) {
	_, err := fmt.Fprintln(srv.LoggerWriter, a...)
	if err != nil {
		fmt.Println(err)
	}
}

func (srv *Server) printf(format string, a ...interface{}) {
	_, err := fmt.Fprintf(srv.LoggerWriter, format, a...)
	if err != nil {
		fmt.Println(err)
	}
}
