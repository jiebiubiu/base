package logs

import "go.uber.org/zap"

type ZapJaegerLog struct{}

// jaeger 接入日志
func (ZapJaegerLog) Error(msg string) {
	zap.S().Error(msg)
}

func (ZapJaegerLog) Infof(msg string, args ...interface{}) {
	zap.S().Infof(msg, args...)
}

// gorm 接入日志
type ZapGormLog struct {
}

func (l ZapGormLog) Printf(s string, params ...interface{}) {
	zap.S().Infof(s, params...)
}
