package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type base struct {
	name string
}

func (b *base) FullName() string {
	return b.name
}

type baseKey struct {
	*base
}

func (b *baseKey) AscIndex() index.Key {
	return index.NewKey(b, 1)
}

func (b *baseKey) DescIndex() index.Key {
	return index.NewKey(b, -1)
}

type baseUpdater[T any] struct {
	*base
}

func (b *baseUpdater[T]) Unset() updater.Updater {
	return updater.New(b, `$unset`, "")
}

func (b *baseUpdater[T]) Set(value T) updater.Updater {
	return updater.New(b, `$set`, value)
}

func (b *baseUpdater[T]) SetOnInsert(value T) updater.Updater {
	return updater.New(b, `$setOnInsert`, value)
}

// Min finalValue = min(value, nowValue) or finalValue = value (if nowValue is Not exist)
func (b *baseUpdater[T]) Min(value T) updater.Updater {
	return updater.New(b, "$min", value)
}

// Max finalValue = max(value, nowValue) or finalValue = value (if nowValue is Not exist)
func (b *baseUpdater[T]) Max(value T) updater.Updater {
	return updater.New(b, "$max", value)
}

// Deprecated: SetOnIns using: SetOnInsert
func (b *baseUpdater[T]) SetOnIns(value T) updater.Updater {
	return updater.New(b, "$setOnInsert", value)
}

type baseFilter[T any] struct {
	*base
}

func (b *baseFilter[T]) Exist() filter.Filter {
	return filter.Exist(b)
}

func (b *baseFilter[T]) NotExist() filter.Filter {
	return filter.NotExist(b)
}

func (b *baseFilter[T]) Type(t bson.Type) filter.Filter {
	return filter.Type(b, t)
}

func (b *baseFilter[T]) Gt(value T) filter.Filter {
	return filter.CompareByValue(b, filter.GT, value)
}

func (b *baseFilter[T]) GtField(f *baseFilter[T]) filter.Filter {
	return filter.CompareByValue(b, filter.GT, f)
}

func (b *baseFilter[T]) Lt(value T) filter.Filter {
	return filter.CompareByValue(b, filter.LT, value)
}

func (b *baseFilter[T]) LtField(f *baseFilter[T]) filter.Filter {
	return filter.CompareByValue(b, filter.LT, f)
}
