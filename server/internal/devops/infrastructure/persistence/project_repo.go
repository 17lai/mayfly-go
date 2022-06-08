package persistence

import (
	"mayfly-go/internal/devops/domain/entity"
	"mayfly-go/internal/devops/domain/repository"
	"mayfly-go/pkg/biz"
	"mayfly-go/pkg/model"
)

type projectRepo struct{}

var ProjectRepo repository.Project = &projectRepo{}

func (p *projectRepo) GetPageList(condition *entity.Project, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult {
	return model.GetPage(pageParam, condition, toEntity, orderBy...)
}

func (p *projectRepo) Count(condition *entity.Project) int64 {
	return model.CountBy(condition)
}

func (p *projectRepo) GetByIdIn(ids []uint64, toEntity interface{}, orderBy ...string) {
	model.GetByIdIn(new(entity.Project), toEntity, ids, orderBy...)
}

func (p *projectRepo) Save(project *entity.Project) {
	biz.ErrIsNil(model.Insert(project), "保存项目失败")
}

func (p *projectRepo) Update(project *entity.Project) {
	biz.ErrIsNil(model.UpdateById(project), "更新项目信息")
}

func (p *projectRepo) Delete(id uint64) {
	model.DeleteById(new(entity.Project), id)
}
