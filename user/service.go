package user

func NewService(r Repository, n CreationNotifier) Service {
	return Service{ r, n}
}

type Service struct {
	r Repository
	n CreationNotifier
}

func (b Service) Create(u User) (string, error) {
	if u.Age < 18 {
		return "", InvalidUser{"Must be over 18"}
	}

	id, err := b.r.Save(u)
	if (err != nil) {
		return "", StorageError{"Failed to save user"}
	}

	return id, b.n.UserCreated(id, u)
}

func (b Service) RetrieveUser(id string) (User, error) {
	// TODO actual error handling and remapping!
	return b.r.GetUserById(id)
}

type InvalidUser struct {
	msg string
}

func (i InvalidUser) Error() string {
	return i.msg
}

type StorageError struct {
	msg string
}

func (s StorageError) Error() string {
	return s.msg
}
