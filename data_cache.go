package mjölnir

type dataCache interface {
	Get(key []byte) ([]byte, error)
	Set(key, value []byte) error
}
