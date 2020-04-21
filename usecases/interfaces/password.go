package interfaces

type Hasher interface {
	Hash(password string) (string, error)
	Verify(password, hash string) (bool, error)
}
