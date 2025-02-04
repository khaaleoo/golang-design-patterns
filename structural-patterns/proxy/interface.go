package proxy

type UserFinder interface {
	Find(ID int) (User, error)
}
