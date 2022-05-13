package region

import (
	"github.com/vela-security/vela-public/auxlib"
	"github.com/vela-security/vela-public/lua"
)

type config struct {
	name   string
	db     string
	method string
}

func newConfig(L *lua.LState) *config {
	tab := L.CheckTable(1)
	cfg := &config{}
	tab.ForEach(func(key lua.LValue, val lua.LValue) {
		if key.Type() != lua.LTString {
			L.RaiseError("invalid config table , got arr")
			return
		}

		switch key.String() {
		case "name":
			cfg.name = val.String()

		case "db":
			cfg.db = val.String()

		case "method":
			cfg.method = val.String()
		default:
			L.RaiseError("invalid %s config invalid", key.String())
			return
		}
	})

	if e := cfg.verify(); e != nil {
		L.RaiseError("%v", e)
		return nil
	}

	return cfg
}

func (cfg *config) verify() error {
	if e := auxlib.Name(cfg.name); e != nil {
		return e
	}

	return nil
}
