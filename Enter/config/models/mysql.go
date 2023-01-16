package models

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址:端口
	DBName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`       // 数据库名
	DBUser   string `mapstructure:"dbuser" json:"dbuser" yaml:"dbuser"`       // 数据库用户名
	DBPasswd string `mapstructure:"dbpasswd" json:"dbpasswd" yaml:"dbpasswd"` // 数据库密码
}
