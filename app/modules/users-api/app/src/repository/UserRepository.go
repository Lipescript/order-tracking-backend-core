package repository

// UserRepository define a interface para o repositório de usuários
type UserRepository interface {
	GetUserByID(id int) (string, error) // Exemplo de método
}
