package init

import (
	"mayfly-go/internal/mq/kafka/api"
	"mayfly-go/internal/mq/kafka/application"
	"mayfly-go/internal/mq/kafka/infra/persistence"
	"mayfly-go/pkg/starter"
)

func init() {
	starter.AddInitIocFunc(func() {
		persistence.InitIoc()
		application.InitIoc()
		api.InitIoc()
	})
}
