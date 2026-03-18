package repository

import (
	"mayfly-go/internal/mq/kafka/domain/entity"
	"mayfly-go/pkg/base"
	"mayfly-go/pkg/model"
)

type Kafka interface {
	base.Repo[*entity.Kafka]

	// 分页获取列表
	GetList(condition *entity.KafkaQuery, orderBy ...string) (*model.PageResult[*entity.Kafka], error)
}
