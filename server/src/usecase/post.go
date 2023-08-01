package usecase

import (
	"github.com/Masuda-1246/shares/domain/entity"
	"github.com/Masuda-1246/shares/domain/repository"
)

var _ IPostUseCase = &PostUseCase{}

// PostUseCaseは投稿に関するユースケースを担当します
type PostUseCase struct {
	pr repository.PostRepository
}

// IPostUseCaseは投稿に関するユースケースを担当します
type IPostUseCase interface {
	// GetByIDはIDを指定して投稿を取得します
	GetByID(id string) (*entity.Post, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(userID string) (*entity.Posts, error)
	// GetAllは全ての投稿を取得します
	GetAll() (*entity.Posts, error)
	// Createは投稿を作成します
	CreatePost(post *entity.Post) error
	// UpdatePostは投稿を更新します
	UpdatePost(post *entity.Post) error
	// DeletePostは投稿を削除します
	DeletePost(id string) error
}

// NewPostUseCaseは新しいPostUseCaseを初期化し構造体のポインタを返します
func NewPostUseCase(pr repository.PostRepository) IPostUseCase {
	return &PostUseCase{
		pr: pr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (pu *PostUseCase) GetByID(id string) (*entity.Post, error) {
	return pu.pr.GetByID(id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (pu *PostUseCase) GetByUserID(userID string) (*entity.Posts, error) {
	return pu.pr.GetByUserID(userID)
}

// GetAllは全ての投稿を取得します
func (pu *PostUseCase) GetAll() (*entity.Posts, error) {
	return pu.pr.GetAll()
}

// Createは投稿を作成します
func (pu *PostUseCase) CreatePost(post *entity.Post) error {
	return pu.pr.CreatePost(post)
}

// UpdatePostは投稿を更新します
func (pu *PostUseCase) UpdatePost(post *entity.Post) error {
	return pu.pr.UpdatePost(post)
}

// DeletePostは投稿を削除します
func (pu *PostUseCase) DeletePost(id string) error {
	return pu.pr.DeletePost(id)
}

