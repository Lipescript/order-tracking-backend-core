package repository

type userRepositoryImpl struct{}

func (u *userRepositoryImpl) GetUserByID(id int) (string, error) {
	panic("unimplemented")
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

type UserRepository interface {
	GetUserByID(id int) (string, error) // Exemplo de método
}
