package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Computable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | bson.Decimal128
}

type ComputableUpdater[T Computable] struct {
	*baseUpdater[T]
}

// todo -num?

// Inc finalValue = T(nowValue + num) or finalValue = value (if nowValue is Not exist)
func (c *ComputableUpdater[T]) Inc(num T) updater.Updater {
	return updater.New(c, "$inc", num)
}

// Dec finalValue = T(nowValue - num) or finalValue = T(-value) (if nowValue is Not exist)
func (c *ComputableUpdater[T]) Dec(num T) updater.Updater {
	return updater.New(c, "$inc", -num)
}

// Mul finalValue = nowValue * num or finalValue = 0 (if nowValue is Not exist)
func (c *ComputableUpdater[T]) Mul(num T) updater.Updater {
	return updater.New(c, "$mul", num)
}
