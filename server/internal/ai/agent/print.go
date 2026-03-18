package agent

import (
	"fmt"
	"mayfly-go/pkg/logx"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func LogEvent(event *adk.AgentEvent) {
	agentTag := fmt.Sprintf("Agent - [%s|%s]", event.AgentName, event.RunPath)

	msg, _, err := adk.GetMessage(event)
	if err != nil {
		logx.Debugf("%s error: get msg error: %v", agentTag, err)
		return
	}

	// 消息内容
	if len(msg.Content) > 0 {
		content := msg.Content
		if len(content) > 500 {
			content = content[:500] + "..."
		}
		if msg.Role == schema.Tool {
			logx.Infof("%s tool[%s] response: %s", agentTag, msg.ToolName, content)
		} else {
			logx.Infof("%s answer: %s", agentTag, content)
		}
	}

	// 工具调用
	if len(msg.ToolCalls) > 0 {
		for _, tc := range msg.ToolCalls {
			args := tc.Function.Arguments
			if len(args) > 300 {
				args = args[:300] + "..."
			}
			logx.Infof("%s call %s(%s)", agentTag, tc.Function.Name, args)
		}
	}

	// 动作
	if event.Action != nil {
		if event.Action.TransferToAgent != nil {
			logx.Debugf("%s ->  transfer: %s", agentTag, event.Action.TransferToAgent.DestAgentName)
		}
		if event.Action.Interrupted != nil {
			for _, ic := range event.Action.Interrupted.InterruptContexts {
				if str, ok := ic.Info.(fmt.Stringer); ok {
					logx.Infof("%s ⏸  %s", agentTag, str.String())
				} else {
					logx.Infof("%s ⏸  %v", agentTag, ic.Info)
				}
			}
		}
		if event.Action.Exit {
			logx.Infof("%s done", agentTag)
		}
	}

	// 错误
	if event.Err != nil {
		logx.Errorf("%s error: %v", agentTag, event.Err)
	}
}
