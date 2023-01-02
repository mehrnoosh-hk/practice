package storage


type Storgae interface {
	Create(interface{}) (string, error)
	Read()
	Update()
	Delete()
}