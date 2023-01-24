package models

type QiNiu struct {
	AccessKey   string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey   string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	Bucket      string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	QiniuServer string `mapstructure:"qiniu-server" json:"qiniu-server" yaml:"qiniu-server"`
	Path        string `mapstructure:"path" json:"path" yaml:"path"`
}
