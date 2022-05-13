package region

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func newLuaRegionSearch(r *Region, L *lua.LState) int {
	ip := L.CheckString(1)

	info, err := r.Search(ip)
	if err != nil {
		L.Push(lua.S2L(err.Error()))
	} else {
		L.Push(L.NewAnyData(info))
	}
	return 1
}

func newLuaRegionDebug(r *Region, L *lua.LState) int {
	n := L.GetTop()
	if n <= 0 {
		return 0
	}

	tab := L.CreateTable(n, 0)

	for i := 1; i <= n; i++ {
		ip := L.CheckString(i)
		v, err := r.Search(ip)
		if err != nil {
			tab.RawSetInt(i, lua.S2L(err.Error()))
			continue
		}

		tab.RawSetInt(i, lua.B2L(v.Byte()))
		xEnv.Errorf("ip: %s , city_id: %d city: %s , Info: %s\n", ip, v.CityID(), v.City(), v.Byte())
	}

	L.Push(tab)
	return 1
}

func (r *Region) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "debug":
		return r.debug
	case "search":
		return r.search
	}
	return lua.LNil
}

func newLuaRegion(L *lua.LState) int {
	//取出对象并
	cfg := newConfig(L)
	proc := L.NewProc(cfg.name, regionTypeOf)

	if proc.IsNil() {
		proc.Set(newRegion(cfg))
	} else {
		proc.Data.(*Region).cfg = cfg
	}
	L.Push(proc)
	return 1
}

func regionL(L *lua.LState) int {
	ip := L.CheckString(1)
	v, err := xEnv.Region(ip)
	if err != nil {
		L.Push(lua.S2L("0|0|0|未知IP|未知IP"))
	} else {
		L.Push(lua.B2L(v.Byte()))
	}
	return 1
}

func WithEnv(env assert.Environment) {
	xEnv = env
	env.Set("ip2region", lua.NewFunction(newLuaRegion))
	env.Set("region", lua.NewFunction(regionL))
}
