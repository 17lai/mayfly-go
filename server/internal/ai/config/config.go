package config

import (
	"cmp"
	sysapp "mayfly-go/internal/sys/application"
	"strings"
)

const (
	ConfigKeyModel string = "AiModelConfig"
)

type ModelConfig struct {
	Name        string  `json:"name"`  // 模型名称，主要用于展示
	Model       string  `json:"model"` // 模型标识，使用 协议/模型名 格式，如 openai/gpt-5.2
	BaseUrl     string  `json:"baseUrl"`
	ApiKey      string  `json:"apiKey"`  // api key
	TimeOut     int     `json:"timeOut"` // 请求超时时间，单位秒
	Temperature float32 `json:"temperature"`
	MaxTokens   int     `json:"maxTokens"`
}

func (c *ModelConfig) GetModelSpec() ModelSpec {
	return ParseModel(c.Model)
}

// ModelSpec 定义模型规范
type ModelSpec struct {
	Protocol string
	Model    string
}

// ParseModel 解析模型字符串，格式为 "protocol/model"，如 "openai/gpt-3.5-turbo"
func ParseModel(model string) ModelSpec {
	parts := strings.SplitN(model, "/", 2)

	if len(parts) != 2 {
		return ModelSpec{
			Protocol: "openai", // 默认协议
			Model:    model,
		}
	}

	return ModelSpec{
		Protocol: parts[0],
		Model:    parts[1],
	}
}

func GetModel() *ModelConfig {
	c := sysapp.GetConfigApp().GetConfig(ConfigKeyModel)
	jm := c.GetJsonM()

	conf := new(ModelConfig)
	conf.Name = jm.GetStr("name")
	conf.Model = jm.GetStr("model")
	conf.BaseUrl = jm.GetStr("baseUrl")
	conf.ApiKey = jm.GetStr("apiKey")
	conf.TimeOut = cmp.Or(jm.GetInt("timeOut"), 60)
	conf.Temperature = cmp.Or(jm.GetFloat32("temperature"), 0.7)
	conf.MaxTokens = cmp.Or(jm.GetInt("maxTokens"), 2048)
	return conf
}
