package user

type IDRetriever interface {
	RetrieveUser(id string) (User, error)
}
