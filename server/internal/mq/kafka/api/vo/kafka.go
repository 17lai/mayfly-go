package vo

import (
	tagentity "mayfly-go/internal/tag/domain/entity"
	"mayfly-go/pkg/model"
)

type Kafka struct {
	model.Model
	tagentity.ResourceTags

	Code     string `json:"code"`
	Name     string `json:"name"`
	Hosts    string `json:"hosts"`
	Username string `json:"username"`
	Password string `json:"password"`

	SshTunnelMachineId int    `json:"sshTunnelMachineId"` // ssh隧道机器id
	SaslMechanism      string `json:"saslMechanism"`      // sasl机制
}

func (m *Kafka) GetCode() string {
	return m.Code
}
