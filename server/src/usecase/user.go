package usecase

import (
	"github.com/Masuda-1246/shares/domain/entity"
	"github.com/Masuda-1246/shares/domain/repository"
)

var _ IUserUsecase = &UserUsecase{}

type UserUsecase struct {
	userRepository repository.UserRepository
}

// UserUsecaseはユーザーに関するユースケースを担当します
type IUserUsecase interface {
	// GetByIDはIDを指定してユーザーを取得します
	GetByID(id string) (*entity.User, error)
	// GetByEmailはEmailを指定してユーザーを取得します
	GetByEmail(email string) (*entity.User, error)
	// CreateUserはユーザーを作成します
	CreateUser(user *entity.User) error
	// UpdateUserはユーザーを更新します
	UpdateUser(user *entity.User) error
	// DeleteUserはユーザーを削除します
	DeleteUser(id string) error
}

// NewUserUsecaseは新しいUserUsecaseを初期化し構造体のポインタを返します
func NewUserUsecase(ur repository.UserRepository) IUserUsecase {
	return &UserUsecase{
		userRepository: ur,
	}
}

// GetByIDはIDを指定してユーザーを取得します
func (uu *UserUsecase) GetByID(id string) (*entity.User, error) {
	return uu.userRepository.GetByID(id)
}

// GetByEmailはEmailを指定してユーザーを取得します
func (uu *UserUsecase) GetByEmail(email string) (*entity.User, error) {
	return uu.userRepository.GetByEmail(email)
}

// CreateUserはユーザーを作成します
func (uu *UserUsecase) CreateUser(user *entity.User) error {
	return uu.userRepository.CreateUser(user)
}

// UpdateUserはユーザーを更新します
func (uu *UserUsecase) UpdateUser(user *entity.User) error {
	return uu.userRepository.UpdateUser(user)
}

// DeleteUserはユーザーを削除します
func (uu *UserUsecase) DeleteUser(id string) error {
	return uu.userRepository.DeleteUser(id)
}
