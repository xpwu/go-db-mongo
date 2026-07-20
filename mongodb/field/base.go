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

type baseUpdater struct {
	*base
}

func (b *baseUpdater) Unset() updater.Updater {
	return updater.New(b, `$unset`, "")
}

func (b *baseUpdater) Set(value interface{}) updater.Updater {
	return updater.New(b, `$set`, value)
}

func (b *baseUpdater) SetOnInsert(value interface{}) updater.Updater {
	return updater.New(b, `$setOnInsert`, value)
}

type baseFilter struct {
	*base
}

func (b *baseFilter) Exist() filter.Filter {
	return filter.Exist(b)
}

func (b *baseFilter) NotExist() filter.Filter {
	return filter.NotExist(b)
}

func (b *baseFilter) Type(t bson.Type) filter.Filter {
	return filter.Type(b, t)
}
