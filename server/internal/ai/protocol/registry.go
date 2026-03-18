package protocol

import (
	"context"
	"errors"
	"fmt"
	"mayfly-go/internal/ai/config"
	"mayfly-go/pkg/logx"
	"mayfly-go/pkg/utils/collx"

	"github.com/cloudwego/eino/components/model"
)

func init() {
	Register(new(Openai))
}

var (
	protocolMap collx.SM[string, Protocol]
	chatModels  collx.SM[string, model.ToolCallingChatModel]
)

// Register 注册协议
func Register(protocol Protocol) {
	protocolMap.Store(protocol.Name(), protocol)
}

// Get 获取协议
func Get(name string) Protocol {
	protocol, _ := protocolMap.Load(name)
	return protocol
}

// GetChatModel 获取Chat模型
func GetChatModel(ctx context.Context, modelConfig *config.ModelConfig) (model.ToolCallingChatModel, error) {
	modelSpec := modelConfig.GetModelSpec()
	modelProtocol := Get(modelSpec.Protocol)
	if modelProtocol == nil {
		return nil, errors.New("no supported AI model protocol found")
	}

	cacheKey := generateCacheKey(modelConfig)
	if chatModel, ok := chatModels.Load(cacheKey); ok {
		logx.Debugf("ai model [%s/%s] - get chat model from cache", modelSpec.Protocol, modelSpec.Model)
		return chatModel, nil
	}

	// 删除已存在的缓存
	chatModels.Clear()

	chatModel, err := modelProtocol.NewChatModel(ctx, modelConfig)
	if err != nil {
		return nil, err
	}
	logx.Debugf("ai model [%s/%s] - new chat model", modelSpec.Protocol, modelSpec.Model)
	chatModels.Store(cacheKey, chatModel)
	return chatModel, nil
}

// generateCacheKey 生成基于 modelConfig 关键字段的缓存键
func generateCacheKey(config *config.ModelConfig) string {
	return fmt.Sprintf("%s_%s_%s_%d_%f",
		config.Model,

		config.BaseUrl,
		config.ApiKey,
		config.MaxTokens,
		config.Temperature,
	)
}
