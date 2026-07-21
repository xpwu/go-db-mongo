package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type baseField[T any] struct {
	name string
}

func (b *baseField[T]) FullName() string {
	return b.name
}

func (b *baseField[T]) Exist() filter.Filter {
	return filter.Exist(b)
}

func (b *baseField[T]) NotExist() filter.Filter {
	return filter.NotExist(b)
}

func (b *baseField[T]) Type(t bson.Type) filter.Filter {
	return filter.Type(b, t)
}

func (b *baseField[T]) Gt(value T) filter.Filter {
	return filter.CompareByValue(b, filter.GT, value)
}

func (b *baseField[T]) GtField(f filter.BaseFilterField[T]) filter.Filter {
	return filter.CompareByValue(b, filter.GT, f)
}

func (b *baseField[T]) Lt(value T) filter.Filter {
	return filter.CompareByValue(b, filter.LT, value)
}

func (b *baseField[T]) LtField(f filter.BaseFilterField[T]) filter.Filter {
	return filter.CompareByValue(b, filter.LT, f)
}

var (
	_ filter.BaseFilterField[any] = &baseField[any]{}
)

func (b *baseField[T]) Unset() updater.Updater {
	return updater.New(b, `$unset`, "")
}

func (b *baseField[T]) Set(value T) updater.Updater {
	return updater.New(b, `$set`, value)
}

func (b *baseField[T]) SetOnInsert(value T) updater.Updater {
	return updater.New(b, `$setOnInsert`, value)
}

// SetMin finalValue = min(value, nowValue) or finalValue = value (if nowValue is Not exist)
func (b *baseField[T]) SetMin(value T) updater.Updater {
	return updater.New(b, "$min", value)
}

// SetMax finalValue = max(value, nowValue) or finalValue = value (if nowValue is Not exist)
func (b *baseField[T]) SetMax(value T) updater.Updater {
	return updater.New(b, "$max", value)
}

// Deprecated: Min using: SetMin
func (b *baseField[T]) Min(value T) updater.Updater {
	return b.SetMin(value)
}

// Deprecated: Max using: SetMax
func (b *baseField[T]) Max(value T) updater.Updater {
	return b.SetMax(value)
}

// Deprecated: SetOnIns using: SetOnInsert
func (b *baseField[T]) SetOnIns(value T) updater.Updater {
	return updater.New(b, "$setOnInsert", value)
}

var (
	_ updater.BaseUpdater[any] = &baseField[any]{}
)
