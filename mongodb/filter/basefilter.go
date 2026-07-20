package filter

import (
	"github.com/xpwu/go-db-mongo/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type BaseFilter[T any] struct {
	mongodb.Field
}

func (b *BaseFilter[T]) Exist() Filter {
	return Exist(b)
}

func (b *BaseFilter[T]) NotExist() Filter {
	return NotExist(b)
}

func (b *BaseFilter[T]) Type(t bson.Type) Filter {
	return Type(b, t)
}

func (b *BaseFilter[T]) Gt(value T) Filter {
	return CompareByValue(b, GT, value)
}

func (b *BaseFilter[T]) GtField(f *BaseFilter[T]) Filter {
	return CompareByValue(b, GT, f)
}

func (b *BaseFilter[T]) Lt(value T) Filter {
	return CompareByValue(b, LT, value)
}

func (b *BaseFilter[T]) LtField(f *BaseFilter[T]) Filter {
	return CompareByValue(b, LT, f)
}
