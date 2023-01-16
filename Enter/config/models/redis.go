package models

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 地址
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}
