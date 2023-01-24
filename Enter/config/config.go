package config

import (
	models2 "DouYin/Enter/config/models"
)

type Server struct {
	JWT   models2.JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	DB    models2.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis models2.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	ZAP   models2.Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	QiNiu models2.QiNiu `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
}
