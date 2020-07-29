package user

type Repository interface {
	GetUserById(id string) (User, error)
	Save(u User) (string, error)
}
