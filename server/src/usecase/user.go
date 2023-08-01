package usecase

import (
	"github.com/Masuda-1246/shares/domain/entity"
	"github.com/Masuda-1246/shares/domain/repository"
)

var _ IUserUseCase = &UserUseCase{}

type UserUseCase struct {
	userRepository repository.UserRepository
}

// UserUseCaseはユーザーに関するユースケースを担当します
type IUserUseCase interface {
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

// NewUserUseCaseは新しいUserUseCaseを初期化し構造体のポインタを返します
func NewUserUseCase(ur repository.UserRepository) IUserUseCase {
	return &UserUseCase{
		userRepository: ur,
	}
}

// GetByIDはIDを指定してユーザーを取得します
func (uu *UserUseCase) GetByID(id string) (*entity.User, error) {
	return uu.userRepository.GetByID(id)
}

// GetByEmailはEmailを指定してユーザーを取得します
func (uu *UserUseCase) GetByEmail(email string) (*entity.User, error) {
	return uu.userRepository.GetByEmail(email)
}

// CreateUserはユーザーを作成します
func (uu *UserUseCase) CreateUser(user *entity.User) error {
	return uu.userRepository.CreateUser(user)
}

// UpdateUserはユーザーを更新します
func (uu *UserUseCase) UpdateUser(user *entity.User) error {
	return uu.userRepository.UpdateUser(user)
}

// DeleteUserはユーザーを削除します
func (uu *UserUseCase) DeleteUser(id string) error {
	return uu.userRepository.DeleteUser(id)
}

