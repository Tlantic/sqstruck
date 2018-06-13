package sqstruck

//Store ...
type Store interface {
	Get(key string) ([]byte, bool)
	Set(key string, value []byte) error
	GetName() string
}
