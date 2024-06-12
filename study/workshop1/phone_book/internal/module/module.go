package module

type Storage interface {
	// CreateContact(name string)
}

type Deps struct {
	Storage Storage
}


type Module struct {
	Deps
}

// Module .. TODO сделать описание функции
func NewModule() Module {
	return Module{}
}

func (m Module) AddContact(telephone models.Telephone) error {
	// Поход в сторонний сервис
	// Запись в kafka

	// Запись в хранилище
	return m.Storage.AddContact(telephone)
}