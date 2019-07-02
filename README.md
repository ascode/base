# base

### ding
发送@所有人的消息:
func SendMsgAll(msg, token string) error

发送@特定人员的消息:
func SendMsg(msg, token string, at ...string) error

钉钉自定义机器人官方文档：https://open-doc.dingtalk.com/microapp/serverapi2/qf2nxq

### sn

产生序列号，可用在分布式环境下产生在不同设备上不会重复的序列号。

```$xslt
/**
生成分布式场景下18位的唯一序列号(可以使用到2022年5月21日)
terminal: 终端识别号
bizCode: 业务编号
machineCode: 机器编号
*/
func GenCode(terminal string, bizCode string, machineCode string) (string, error)

/**
生成分布式场景下15位的唯一序列号(可以使用到2022年5月21日)
terminal: 终端识别号
bizCode: 业务编号
machineCode: 机器编号
 */
func GenShortCode(terminal string, bizCode string, machineCode string) (string, error)

```

### Restful Api 返回值的处理

#### 注意： 使用此节中的方法，会对当前服务的所有api生效，所以目前要使用返回值处理，为了不影响其他api,需要新建一个服务。时间有限，欢迎大家可以改成指定api生效。

可以支持对返回值的装饰效果进行处理，例如我们默认的返回值为
```$xslt
{
    "x": {
        "ok": true
    },
    "data": {
        "echostr": "aksjkldfs"
    }
}
```

要对返回值进行加工需要实现Server结构体的ParseFlag方法，目前支持三种返回值风格
```$xslt
type Server struct {
	msvc.Service
}
```

默认情况
```$xslt
// 不需要实现
func (s *Server) ParseFlag(p *pflag.FlagSet)
```

去掉固定模式返回值的x和data
```$xslt
//需要在service里面实现
func (s *Server) ParseFlag(p *pflag.FlagSet) {
	s.SetRestfulResultNoWrap(true)
}
```

单一键值对返回值的时候只取值作为字符串返回
```$xslt
//需要在service里面实现
func (s *Server) ParseFlag(p *pflag.FlagSet) {
	s.SetRestfulResultNoWrap(true)
	s.SetSingleKVtoVString(true)
}
```
