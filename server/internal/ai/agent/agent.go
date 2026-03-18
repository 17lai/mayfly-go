package agent

import (
	"context"
	"mayfly-go/pkg/logx"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/adk/prebuilt/deep"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

// newOpsExpertAgent 创建一个新的运维专家Agent
func newOpsExpertAgent(ctx context.Context, tools ...tool.BaseTool) (adk.Agent, error) {
	toolableChatModel, err := GetChatModel(ctx)
	if err != nil {
		return nil, err
	}

	return adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "OpsExpertAgent",
		Description: "an agent for ops expert task",
		Model:       toolableChatModel,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: tools,
			},
		},
	})
}

// GetOpsExpertAgent 获取运维专家agent
func GetOpsExpertAgent(ctx context.Context, toolTypes ...ToolType) (*AiAgent, error) {
	tools := make([]tool.BaseTool, 0)
	for _, toolType := range toolTypes {
		if t, exists := GetTools(toolType); exists {
			tools = append(tools, t...)
		}
	}

	agent, err := newOpsExpertAgent(ctx, tools...)
	if err != nil {
		return nil, err
	}
	return &AiAgent{
		Agent: agent,
	}, nil
}

type AiAgent struct {
	adk.Agent
}

// Run 运行，并返回最终结果
func (aiAgent *AiAgent) Run(ctx context.Context, sysPrompt string, question string) (string, error) {
	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		EnableStreaming: true,
		Agent:           aiAgent.Agent,
		CheckPointStore: NewInMemoryStore(),
	})

	deep.New(ctx, &deep.Config{})

	iter := runner.Run(ctx, []adk.Message{
		{
			Role:    schema.System,
			Content: sysPrompt,
		},
		{
			Role:    schema.User,
			Content: question,
		},
	})

	var lastMessage adk.Message
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}

		err := event.Err
		if err != nil {
			logx.Error(err.Error())
			return "", err
		}

		if lastMessage, _, err = adk.GetMessage(event); err != nil {
			return "", err
		}

		LogEvent(event)
	}

	if lastMessage != nil {
		return lastMessage.Content, nil
	}

	return "finished without output message", nil
}

func NewDeepAgent(ctx context.Context, tools ...tool.BaseTool) (adk.Agent, error) {
	toolableChatModel, err := GetChatModel(ctx)
	if err != nil {
		return nil, err
	}
	return deep.New(ctx, &deep.Config{
		Name:        "OpsExpertAgent",
		Description: "an agent for ops expert task",
		ChatModel:   toolableChatModel,
		// SubAgents:   []adk.Agent{ca, wa},
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: tools,
			},
		},
		MaxIteration: 100,
	})
}
