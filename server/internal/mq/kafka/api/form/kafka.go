package form

import "mayfly-go/pkg/model"

type Kafka struct {
	model.IdModel

	Code          string  `json:"code"`
	Name          string  `json:"name" binding:"required"`
	Hosts         string  `json:"hosts" binding:"required"` // Kafka 连接地址，格式: host1:port1,host2:port2 或单个 broker
	Username      *string `json:"username"`
	Password      *string `json:"password"`
	SaslMechanism *string `json:"saslMechanism"`

	SshTunnelMachineId int `json:"sshTunnelMachineId"`

	TagCodePaths []string `json:"tagCodePaths" binding:"required"`
}
