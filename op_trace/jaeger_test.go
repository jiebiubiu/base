package op_trace

import (
	"testing"

	"github.com/Ho-J/base/config"
	"github.com/Ho-J/base/viper_c"
)

func TestJaeger(t *testing.T) {
	viper_c.SetViper("../config/default.yaml")
	c := config.Config{}
	viper_c.LoadConfig(&c)
	InitJaeger(c.Jaeger)
}
