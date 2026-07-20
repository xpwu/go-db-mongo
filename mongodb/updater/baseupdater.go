package updater

import "github.com/xpwu/go-db-mongo/mongodb"

type BaseUpdater[T any] struct {
	mongodb.Field
}

func (b *BaseUpdater[T]) Unset() Updater {
	return New(b, `$unset`, "")
}

func (b *BaseUpdater[T]) Set(value T) Updater {
	return New(b, `$set`, value)
}

func (b *BaseUpdater[T]) SetOnInsert(value T) Updater {
	return New(b, `$setOnInsert`, value)
}

// SetMin finalValue = min(value, nowValue) or finalValue = value (if nowValue is Not exist)
func (b *BaseUpdater[T]) SetMin(value T) Updater {
	return New(b, "$min", value)
}

// SetMax finalValue = max(value, nowValue) or finalValue = value (if nowValue is Not exist)
func (b *BaseUpdater[T]) SetMax(value T) Updater {
	return New(b, "$max", value)
}

// Deprecated: Min using: SetMin
func (b *BaseUpdater[T]) Min(value T) Updater {
	return b.SetMin(value)
}

// Deprecated: Max using: SetMax
func (b *BaseUpdater[T]) Max(value T) Updater {
	return b.SetMax(value)
}

// Deprecated: SetOnIns using: SetOnInsert
func (b *BaseUpdater[T]) SetOnIns(value T) Updater {
	return New(b, "$setOnInsert", value)
}
