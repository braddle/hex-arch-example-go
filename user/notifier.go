package user

type CreationNotifier interface {
	UserCreated(id string, u User) error
}
