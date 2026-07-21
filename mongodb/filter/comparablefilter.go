package filter

import "github.com/xpwu/go-db-mongo/mongodb"

type ComparableFilterField[T any] interface {
	mongodb.Field
	ComparableFilter[T]
}

// ComparableFilter T ~ comparable | EqualAble
type ComparableFilter[T any] interface {
	BaseFilter[T]
	Eq(value T) Filter
	EqField(f ComparableFilterField[T]) Filter
	Ne(value T) Filter
	NeField(f ComparableFilterField[T]) Filter
	Gte(value T) Filter
	GteField(f ComparableFilterField[T]) Filter
	Lte(value T) Filter
	LteField(f ComparableFilterField[T]) Filter
	In(values []T) Filter
	Nin(values []T) Filter
}

type EqualAble[T EqualAble[T]] interface {
	Equal(T) bool
}
