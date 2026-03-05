package config

import (
	"cmp"
	sysapp "mayfly-go/internal/sys/application"
)

const (
	ConfigKeyFile string = "FileConfig" // 文件配置key
)

type FileConfig struct {
	BasePath string // 文件基础路径
}

func GetFileConfig() *FileConfig {
	c := sysapp.GetConfigApp().GetConfig(ConfigKeyFile)
	jm := c.GetJsonM()

	fc := new(FileConfig)
	fc.BasePath = cmp.Or(jm.GetStr("basePath"), "./file")
	return fc
}
