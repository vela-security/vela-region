package region

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

type Region struct {
	lua.ProcEx

	//服务标注
	cfg *config

	//服务对象
	i2r *ip2region

	//lua 接口方法
	search *lua.LFunction
	debug  *lua.LFunction
}

func newRegion(cfg *config) *Region {
	r := &Region{cfg: cfg}
	r.V(lua.PTInit, regionTypeOf)

	r.search = lua.NewFunction(func(L *lua.LState) int {
		return newLuaRegionSearch(r, L)
	})

	r.debug = lua.NewFunction(func(L *lua.LState) int {
		return newLuaRegionDebug(r, L)
	})

	return r
}

func (r *Region) Name() string {
	return r.cfg.name
}

func (r *Region) Start() error {
	i2, err := newIp2Region(r.cfg.db)
	if err != nil {
		return err
	}
	r.i2r = i2
	return nil
}

func (r *Region) Close() error {
	r.i2r.Close()
	xEnv.Errorf("%s region proc close", r.Name())
	return nil
}

func (r *Region) Search(v string) (*assert.IPv4Info, error) {
	switch r.cfg.method {
	case "Binary":
		return r.i2r.BinarySearch(v)
	case "BtreeSearch":
		return r.i2r.BtreeSearch(v)
	default:
		return r.i2r.MemorySearch(v)
	}
}
