package conf

import (
	"ADDD_DOUYIN/util"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gopkg.in/ini.v1"
	"net/url"
	"strings"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.local.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}
	//初始化mysql
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	Database(path)

	//初始化cos
	u, _ := url.Parse(file.Section("cos").Key("url").String())
	b := &cos.BaseURL{BucketURL: u}
	id := file.Section("cos").Key("id").String()
	key := file.Section("cos").Key("key").String()
	util.InitCos(b, id, key)

	// 初始化redis
	conf := &redis.Options{
		Addr:     file.Section("redis").Key("url").String(),
		Password: file.Section("redis").Key("password").String(),
		DB:       0,
	}
	util.InitRedis(conf)
}
