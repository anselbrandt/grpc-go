
// Package glogger defines glog-based logging for grpc.
// Importing this package will install glog as the logger used by grpclog.
package glogger

import (
	"fmt"

	"github.com/golang/glog"
	"google.golang.org/grpc/grpclog"
)

const d = 2

func init() {
	grpclog.SetLoggerV2(&glogger{})
}

type glogger struct{}

func (g *glogger) Info(args ...interface{}) {
	glog.InfoDepth(d, args...)
}

func (g *glogger) Infoln(args ...interface{}) {
	glog.InfoDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Infof(format string, args ...interface{}) {
	glog.InfoDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) InfoDepth(depth int, args ...interface{}) {
	glog.InfoDepth(depth+d, args...)
}

func (g *glogger) Warning(args ...interface{}) {
	glog.WarningDepth(d, args...)
}

func (g *glogger) Warningln(args ...interface{}) {
	glog.WarningDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Warningf(format string, args ...interface{}) {
	glog.WarningDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) WarningDepth(depth int, args ...interface{}) {
	glog.WarningDepth(depth+d, args...)
}

func (g *glogger) Error(args ...interface{}) {
	glog.ErrorDepth(d, args...)
}

func (g *glogger) Errorln(args ...interface{}) {
	glog.ErrorDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Errorf(format string, args ...interface{}) {
	glog.ErrorDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) ErrorDepth(depth int, args ...interface{}) {
	glog.ErrorDepth(depth+d, args...)
}

func (g *glogger) Fatal(args ...interface{}) {
	glog.FatalDepth(d, args...)
}

func (g *glogger) Fatalln(args ...interface{}) {
	glog.FatalDepth(d, fmt.Sprintln(args...))
}

func (g *glogger) Fatalf(format string, args ...interface{}) {
	glog.FatalDepth(d, fmt.Sprintf(format, args...))
}

func (g *glogger) FatalDepth(depth int, args ...interface{}) {
	glog.FatalDepth(depth+d, args...)
}

func (g *glogger) V(l int) bool {
	return bool(glog.V(glog.Level(l)))
}
