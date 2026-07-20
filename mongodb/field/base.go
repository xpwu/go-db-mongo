package field

//type base struct {
//	name string
//}
//
//func (b *base) FullName() string {
//	return b.name
//}

//type BaseKey struct {
//	*base
//}
//
//func (b *BaseKey) AscIndex() index.Key {
//	return index.NewKey(b, 1)
//}
//
//func (b *BaseKey) DescIndex() index.Key {
//	return index.NewKey(b, -1)
//}
//
//type BaseUpdater[T any] struct {
//	*base
//}
//
//func (b *BaseUpdater[T]) Unset() updater.Updater {
//	return updater.New(b, `$unset`, "")
//}
//
//func (b *BaseUpdater[T]) Set(value T) updater.Updater {
//	return updater.New(b, `$set`, value)
//}
//
//func (b *BaseUpdater[T]) SetOnInsert(value T) updater.Updater {
//	return updater.New(b, `$setOnInsert`, value)
//}
//
//// SetMin finalValue = min(value, nowValue) or finalValue = value (if nowValue is Not exist)
//func (b *BaseUpdater[T]) SetMin(value T) updater.Updater {
//	return updater.New(b, "$min", value)
//}
//
//// SetMax finalValue = max(value, nowValue) or finalValue = value (if nowValue is Not exist)
//func (b *BaseUpdater[T]) SetMax(value T) updater.Updater {
//	return updater.New(b, "$max", value)
//}
//
//// Deprecated: Min using: SetMin
//func (b *BaseUpdater[T]) Min(value T) updater.Updater {
//	return b.SetMin(value)
//}
//
//// Deprecated: Max using: SetMax
//func (b *BaseUpdater[T]) Max(value T) updater.Updater {
//	return b.SetMax(value)
//}
//
//// Deprecated: SetOnIns using: SetOnInsert
//func (b *BaseUpdater[T]) SetOnIns(value T) updater.Updater {
//	return updater.New(b, "$setOnInsert", value)
//}
//
//type BaseFilter[T any] struct {
//	*base
//}
//
//func (b *BaseFilter[T]) Exist() filter.Filter {
//	return filter.Exist(b)
//}
//
//func (b *BaseFilter[T]) NotExist() filter.Filter {
//	return filter.NotExist(b)
//}
//
//func (b *BaseFilter[T]) Type(t bson.Type) filter.Filter {
//	return filter.Type(b, t)
//}
//
//func (b *BaseFilter[T]) Gt(value T) filter.Filter {
//	return filter.CompareByValue(b, filter.GT, value)
//}
//
//func (b *BaseFilter[T]) GtField(f *BaseFilter[T]) filter.Filter {
//	return filter.CompareByValue(b, filter.GT, f)
//}
//
//func (b *BaseFilter[T]) Lt(value T) filter.Filter {
//	return filter.CompareByValue(b, filter.LT, value)
//}
//
//func (b *BaseFilter[T]) LtField(f *BaseFilter[T]) filter.Filter {
//	return filter.CompareByValue(b, filter.LT, f)
//}
