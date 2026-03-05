package protocol

import (
	"context"
	"mayfly-go/internal/ai/config"
	"time"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

type Openai struct {
}

func (o *Openai) Name() string {
	return ProtocolOpenai
}

func (o *Openai) NewChatModel(ctx context.Context, modelConfig *config.ModelConfig) (model.ToolCallingChatModel, error) {
	return openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL:     modelConfig.BaseUrl,
		Model:       modelConfig.GetModelSpec().Model,
		APIKey:      modelConfig.ApiKey,
		Timeout:     time.Duration(modelConfig.TimeOut) * time.Second,
		MaxTokens:   &modelConfig.MaxTokens,
		Temperature: &modelConfig.Temperature,
	})
}
