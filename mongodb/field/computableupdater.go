package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/updater"
)

// Inc finalValue = T(nowValue + num) or finalValue = value (if nowValue is Not exist)
func (b *baseField[T]) Inc(num T) updater.Updater {
	return updater.New(b, "$inc", num)
}

// Mul finalValue = nowValue * num or finalValue = 0 (if nowValue is Not exist)
func (b *baseField[T]) Mul(num T) updater.Updater {
	return updater.New(b, "$mul", num)
}
