package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type branchService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewBranchService(storage storage.IStorage, log logger.ILogger) branchService {
	return branchService{storage: storage, log: log}
}

func (b branchService) Create(ctx context.Context, branch models.CreateBranch) (models.Branch, error) {
	b.log.Info("branch create service layer", logger.Any("branch", branch))

	id, err := b.storage.Branch().Create(ctx, branch)
	if err != nil {
		b.log.Error("error in service layer while creating branch", logger.Error(err))
		return models.Branch{}, err
	}

	createdBranch, err := b.storage.Branch().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting branch by id", logger.Error(err))
		return models.Branch{}, err
	}

	return createdBranch, nil
}

func (b branchService) Get(ctx context.Context, id string) (models.Branch, error) {
	branch, err := b.storage.Branch().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting branch by id", logger.Error(err))
		return models.Branch{}, err
	}

	return branch, nil
}

func (b branchService) GetList(ctx context.Context, request models.GetListRequest) (models.BranchResponse, error) {
	b.log.Info("branch get list service layer", logger.Any("branch", request))

	branches, err := b.storage.Branch().GetList(ctx, request)
	if err != nil {
		b.log.Error("error in service layer while getting list", logger.Error(err))
		return models.BranchResponse{}, err
	}

	return branches, nil
}

func (b branchService) Update(ctx context.Context, branch models.UpdateBranch) (models.Branch, error) {
	id, err := b.storage.Branch().Update(ctx, branch)
	if err != nil {
		b.log.Error("error in service layer while updating branch", logger.Error(err))
		return models.Branch{}, err
	}

	updatedBranch, err := b.storage.Branch().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting  branch by id", logger.Error(err))
		return models.Branch{}, err
	}

	return updatedBranch, nil
}

func (b branchService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.Branch().Delete(ctx, key)

	return err
}
