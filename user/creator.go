package user

type Creator interface {
	Create(u User) (string, error)
}
