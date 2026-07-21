package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Deprecated:
type (
	Binary0F         = BinaryField
	Binary0FUpdaterF = updater.BaseUpdater[bson.Binary]
	Binary0FFilterF  = filter.ComparableFilter[bson.Binary]
)

// Deprecated:
var (
	_           binary0F               = &Binary0F{}
	_           binary0FUpdaterF       = &Binary0FUpdaterF{}
	_           binary0FFilterF        = &Binary0FFilterF{}
	NewBinary0F func(string) *Binary0F = NewBinaryField
)

// Binary0F
type (
	binary0FUpdaterF interface {
		Set(value bson.Binary) updater.Updater
		SetOnIns(value bson.Binary) updater.Updater
	}

	binary0FFilterF interface {
		In(values []bson.Binary) filter.Filter
		Nin(values []bson.Binary) filter.Filter
	}

	binary0F interface {
		binary0FUpdaterF
		binary0FFilterF
		FullName() string
	}
)
