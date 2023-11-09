package op_trace

import (
	"testing"

	"github.com/jiebiubiu/base/viper_c"
)

func TestJaeger(t *testing.T) {
	viperC := viper_c.NewViperC("../config/default.yaml")
	viperC.LoadConfig()
	c := viperC.GetConfig()
	InitJaeger(c.Jaeger)
}
