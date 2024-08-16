package repository

import (
	"context"
	"github.com/GoSimplicity/LinkMe/internal/domain"
	"github.com/GoSimplicity/LinkMe/internal/repository/dao"
)

// RelationRepository 定义了关注关系的存储库接口
type RelationRepository interface {
	ListFollowerRelations(ctx context.Context, followerID int64, pagination domain.Pagination) ([]domain.Relation, error)
	ListFolloweeRelations(ctx context.Context, followeeID int64, pagination domain.Pagination) ([]domain.Relation, error)
	FollowUser(ctx context.Context, followerID, followeeID int64) error
	CancelFollowUser(ctx context.Context, followerID, followeeID int64) error
	GetFolloweeCount(ctx context.Context, userID int64) (int64, error)
	GetFollowerCount(ctx context.Context, userID int64) (int64, error)
}

type relationRepository struct {
	dao dao.RelationDAO
}

func NewRelationRepository(dao dao.RelationDAO) RelationRepository {
	return &relationRepository{
		dao: dao,
	}
}

// FollowUser 关注用户
func (r *relationRepository) FollowUser(ctx context.Context, followerID, followeeID int64) error {
	return r.dao.FollowUser(ctx, followerID, followeeID)
}

func (r *relationRepository) CancelFollowUser(ctx context.Context, followerID, followeeID int64) error {
	return r.dao.CancelFollowUser(ctx, followerID, followeeID)
}

// ListFollowerRelations 列出关注列表
func (r *relationRepository) ListFollowerRelations(ctx context.Context, followerID int64, pagination domain.Pagination) ([]domain.Relation, error) {
	relations, err := r.dao.ListFollowerRelations(ctx, followerID, pagination)
	return r.toDomainRelationSlice(relations), err
}

// ListFolloweeRelations 获取特定的关注关系信息
func (r *relationRepository) ListFolloweeRelations(ctx context.Context, followeeID int64, pagination domain.Pagination) ([]domain.Relation, error) {
	relations, err := r.dao.ListFolloweeRelations(ctx, followeeID, pagination)
	return r.toDomainRelationSlice(relations), err
}

func (r *relationRepository) GetFolloweeCount(ctx context.Context, userID int64) (int64, error) {
	count, err := r.dao.FollowCount(ctx, userID)
	return count.FolloweeCount, err
}

func (r *relationRepository) GetFollowerCount(ctx context.Context, userID int64) (int64, error) {
	count, err := r.dao.FollowCount(ctx, userID)
	return count.FollowerCount, err
}

func (r *relationRepository) toDomainRelation(relation dao.Relation) domain.Relation {
	return domain.Relation{
		FolloweeId: relation.FolloweeID,
		FollowerId: relation.FollowerID,
	}
}

func (r *relationRepository) toDomainRelationSlice(relations []dao.Relation) []domain.Relation {
	relationSlice := make([]domain.Relation, len(relations))
	for i, relation := range relations {
		relationSlice[i] = r.toDomainRelation(relation)
	}
	return relationSlice
}

func (r *relationRepository) toDAORelation(relation dao.Relation) dao.Relation {
	return dao.Relation{
		FolloweeID: relation.FolloweeID,
		FollowerID: relation.FollowerID,
	}
}
