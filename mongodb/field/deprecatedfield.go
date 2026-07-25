package field

import (
	"github.com/xpwu/go-db-mongo/mongodb"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type deprecatedBase interface {
	FullName() string
}

type deprecatedBaseKey interface {
	deprecatedBase
	AscIndex() index.Key
	DescIndex() index.Key
}

type deprecatedBaseUpdater interface {
	deprecatedBase
	Unset() updater.Updater
	// 已经由 Set(value T) Updater 与 SetOnInsert(value T) Updater 代替
	// 暂时不考虑 interface{} 与 T 的兼容性，上层在调用旧版本时，正常使用都应该是传的 T
	//Set(value interface{}) updater.Updater
	//SetOnInsert(value interface{}) updater.Updater
}

type deprecatedBaseFilter interface {
	deprecatedBase
	Exist() filter.Filter
	NotExist() filter.Filter
	Type(t bson.Type) filter.Filter
}

type binary0FUpdaterF interface {
	deprecatedBaseUpdater
	Set(value bson.Binary) updater.Updater
	SetOnIns(value bson.Binary) updater.Updater
}

type binary0FFilterF interface {
	deprecatedBaseFilter
	In(values []bson.Binary) filter.Filter
	Nin(values []bson.Binary) filter.Filter
}

type binary0F interface {
	binary0FUpdaterF
	binary0FFilterF
	deprecatedBaseKey
}

// Deprecated:
type (
	Binary0F         = BinaryField
	Binary0FUpdaterF = BinaryField
	Binary0FFilterF  = BinaryField
)

// Deprecated:
var (
	_           binary0F              = Binary0F(nil)
	_           binary0FUpdaterF      = Binary0FUpdaterF(nil)
	_           binary0FFilterF       = Binary0FFilterF(nil)
	NewBinary0F func(string) Binary0F = NewBinaryField
)

type array interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	PopFirst() updater.Updater
	PopLast() updater.Updater
	PullByF(f filter.Filter) updater.Updater
	SameEleMatch(f filter.Filter) filter.Filter
	Size(sz int) filter.Filter
}

// Deprecated:
type Array struct {
	ArrayField[any, mongodb.Field]
}

func (a *Array) PullByF(f filter.Filter) updater.Updater {
	return updater.PullByFilter(a, f)
}

func (a *Array) SameEleMatch(f filter.Filter) filter.Filter {
	return filter.SameElemMatch(a, f)
}

// Deprecated:
var (
	NewArray func(fName string) *Array = func(fName string) *Array {
		return &Array{NewArrayField[any, mongodb.Field](fName, func(name string) mongodb.Field {
			return &baseField[any]{name: name}
		})}
	}

	_ array = NewArray("")
)
