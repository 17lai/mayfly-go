package persistence

import (
	"mayfly-go/internal/mq/kafka/domain/entity"
	"mayfly-go/internal/mq/kafka/domain/repository"
	"mayfly-go/pkg/base"
	"mayfly-go/pkg/model"
)

type kafkaRepoImpl struct {
	base.RepoImpl[*entity.Kafka]
}

func newKafkaRepo() repository.Kafka {
	return &kafkaRepoImpl{}
}

// 分页获取数据库信息列表
func (d *kafkaRepoImpl) GetList(condition *entity.KafkaQuery, orderBy ...string) (*model.PageResult[*entity.Kafka], error) {
	qd := model.NewCond().
		Like("name", condition.Name).
		Eq("code", condition.Code).
		In("code", condition.Codes)

	keyword := condition.Keyword
	if keyword != "" {
		keyword = "%" + keyword + "%"
		qd.And("name like ? or code like ? or hosts like ?", keyword, keyword, keyword)
	}
	return d.PageByCond(qd, condition.PageParam)
}
