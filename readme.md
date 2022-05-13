# vela-region
## 说明 
`ipv4地址查询模块`

## 函数说明

### rock.ip2reion
- 函数: rock.ip2region(string)
- 语法: rock.ip2reiogn( "ip.db" )
```lua
    local r = http.client 
    r.output("resource/ip.db").GET("http://xxx.com/ip2region.db")

    local ip = vela.ip2region("resource/ip.db")
    ip.debug("122.112.221.2" , "114.114.114.114")
```