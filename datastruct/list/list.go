package list

type FactoryList interface {
	Get(key int) (interface{}, bool)

	Set(key int, v interface{})

	Search(key int) (interface{}, bool)

	Delete(key int)

	IsEmpty()

	DeleteTail() (interface{}, bool)

	DeleteHead() (interface{}, bool)
}
