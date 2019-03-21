package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

//配置信息
func init() {
	log.Print("config init ...")

	c := Config{
		FilePath: "config",
		FileName: "props",
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		//配置文件错误，退出进程
		log.Print("config file error ...")
		os.Exit(2)
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()
}

type Config struct {
	FilePath string
	FileName string
}

func (c *Config) initConfig() error {
	viper.AddConfigPath(c.FilePath)
	viper.SetConfigName(c.FileName)

	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
