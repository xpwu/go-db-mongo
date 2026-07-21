package updater

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Computable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 | bson.Decimal128
}

type ComputableUpdater[T Computable] interface {
	// Inc finalValue = T(nowValue + num) or finalValue = value (if nowValue is Not exist)
	Inc(num T) Updater
	// Mul finalValue = nowValue * num or finalValue = 0 (if nowValue is Not exist)
	Mul(num T) Updater
}

type UnsignedComputable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type UnsignedComputableUpdater[T UnsignedComputable] interface {
	ComputableUpdater[T]
	// Dec finalValue = T(nowValue - num) or finalValue = T(-value) (if nowValue is Not exist)
	Dec(num T) Updater
}
