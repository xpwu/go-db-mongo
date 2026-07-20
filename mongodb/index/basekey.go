package index

import "github.com/xpwu/go-db-mongo/mongodb"

type BaseKey struct {
	mongodb.Field
}

func (b *BaseKey) AscIndex() Key {
	return NewKey(b, 1)
}

func (b *BaseKey) DescIndex() Key {
	return NewKey(b, -1)
}
