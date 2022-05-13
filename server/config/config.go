package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var CONFIG Config

type Config struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`             // 服务器地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	Dbname   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // 数据库名
	Config   string `mapstructure:"config" json:"config" yaml:"config"`       // 高级配置
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 数据库用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func InitConfig() {
	viper.SetConfigFile("./configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&CONFIG); err != nil {
		fmt.Println(err)
	}

}
