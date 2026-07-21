package mongodb

type Field interface {
	FullName() string
	InitName(name string)
}
