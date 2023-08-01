package usecase

import (
	"github.com/Masuda-1246/shares/domain/entity"
	"github.com/Masuda-1246/shares/domain/repository"
)

var _ IFavoriteUseCase = &FavoriteUseCase{}

// FavoriteUseCaseは投稿に関するユースケースを担当します
type FavoriteUseCase struct {
	fr repository.FavoriteRepository
}

// IFavoriteUseCaseは投稿に関するユースケースを担当します
type IFavoriteUseCase interface {
	// GetByIDはIDを指定して投稿を取得します
	GetByID(id string) (*entity.Favorite, error)
	// GetByUserIDはUserIDを指定して投稿を取得します
	GetByUserID(userID string) (*entity.Favorites, error)
	// GetByPostIDはPostIDを指定して投稿を取得します
	GetByPostID(postID string) (*entity.Favorites, error)
	// GetAllは全ての投稿を取得します
	GetAll() (*entity.Favorites, error)
	// Createは投稿を作成します
	CreateFavorite(favorite *entity.Favorite) error
	// DeleteFavoriteは投稿を削除します
	DeleteFavorite(id string) error
}

// NewFavoriteUseCaseは新しいFavoriteUseCaseを初期化し構造体のポインタを返します
func NewFavoriteUseCase(fr repository.FavoriteRepository) IFavoriteUseCase {
	return &FavoriteUseCase{
		fr: fr,
	}
}

// GetByIDはIDを指定して投稿を取得します
func (fu *FavoriteUseCase) GetByID(id string) (*entity.Favorite, error) {
	return fu.fr.GetByID(id)
}

// GetByUserIDはUserIDを指定して投稿を取得します
func (fu *FavoriteUseCase) GetByUserID(userID string) (*entity.Favorites, error) {
	return fu.fr.GetByUserID(userID)
}

// GetByPostIDはPostIDを指定して投稿を取得します
func (fu *FavoriteUseCase) GetByPostID(postID string) (*entity.Favorites, error) {
	return fu.fr.GetByPostID(postID)
}

// GetAllは全ての投稿を取得します
func (fu *FavoriteUseCase) GetAll() (*entity.Favorites, error) {
	return fu.fr.GetAll()
}

// Createは投稿を作成します
func (fu *FavoriteUseCase) CreateFavorite(favorite *entity.Favorite) error {
	return fu.fr.CreateFavorite(favorite)
}

// DeleteFavoriteは投稿を削除します
func (fu *FavoriteUseCase) DeleteFavorite(id string) error {
	return fu.fr.DeleteFavorite(id)
}