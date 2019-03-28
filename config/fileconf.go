package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

//配置文件信息
func InitFileConf() {
	log.Print("file conf init start...")

	fileConf := FileConf{
		Path: "config",
		Name: "props",
	}

	// 初始化配置文件
	if err := fileConf.initConfig(); err != nil {
		//配置文件错误，退出进程
		log.Fatalf("file conf init error ...")
		os.Exit(2)
	}

	// 监控配置文件变化并热加载程序
	fileConf.watchConfig()
	log.Print("file conf init end...")
}

//配置文件信息
type FileConf struct {
	Path string
	Name string
}

func (fileConf *FileConf) initConfig() error {
	viper.AddConfigPath(fileConf.Path)
	viper.SetConfigName(fileConf.Name)

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
func (fileConf *FileConf) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("file conf changed: %s", e.Name)
	})
}
