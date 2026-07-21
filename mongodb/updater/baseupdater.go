package updater

type BaseUpdater[T any] interface {
	Unset() Updater
	Set(value T) Updater
	SetOnInsert(value T) Updater
	// SetMin finalValue = min(value, nowValue) or finalValue = value (if nowValue is Not exist)
	SetMin(value T) Updater
	// SetMax finalValue = max(value, nowValue) or finalValue = value (if nowValue is Not exist)
	SetMax(value T) Updater

	// Deprecated: Min using: SetMin
	Min(value T) Updater
	// Deprecated: Max using: SetMax
	Max(value T) Updater
	// Deprecated: SetOnIns using: SetOnInsert
	SetOnIns(value T) Updater
}
