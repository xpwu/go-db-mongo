package field

import (
	"github.com/xpwu/go-db-mongo/mongodb"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type IntegerFilter[T Integer] interface {
	filter.ComparableFilter[T]
	Mod(divisor, remainder T) filter.Filter
}

// IntegerField todo *index.BaseKey
type IntegerField[T Integer] interface {
	mongodb.Field
	IntegerFilter[T]
	updater.ComputableUpdater[T]
}

func (b *baseField[T]) Mod(divisor, remainder T) filter.Filter {
	return filter.New(b, "$mod", bson.A{divisor, remainder})
}

func NewIntegerField[T Integer](name string) IntegerField[T] {
	return &baseField[T]{name}
}

type IntField = IntegerField[int]
type Int8Field = IntegerField[int8]
type Int16Field = IntegerField[int16]
type Int32Field = IntegerField[int32]
type Int64Field = IntegerField[int64]

type UnInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// UnIntegerField todo *index.BaseKey
type UnIntegerField[T UnInteger] interface {
	mongodb.Field
	IntegerFilter[T]
	updater.UnsignedComputableUpdater[T]
}

type unIntegerField[T UnInteger] struct {
	baseField[T]
}

// Dec finalValue = T(nowValue - num) or finalValue = T(-value) (if nowValue is Not exist)
func (u *unIntegerField[T]) Dec(num T) updater.Updater {
	return updater.New(u, "$inc", -num)
}

func NewUnIntegerField[T UnInteger](name string) UnIntegerField[T] {
	return &unIntegerField[T]{baseField[T]{name}}
}

type UintField = UnIntegerField[uint]
type ByteField = UnIntegerField[byte]
type Uint8Field = UnIntegerField[uint8]
type Uint16Field = UnIntegerField[uint16]
type Uint32Field = UnIntegerField[uint32]
type Uint64Field = UnIntegerField[uint64]

type StringFilter interface {
	Regex(regex bson.Regex) filter.Filter
	filter.ComparableFilter[string]
}

// StringField todo *index.BaseKey
type StringField interface {
	mongodb.Field
	StringFilter
	updater.BaseUpdater[string]
}

type stringField struct {
	baseField[string]
}

func (s *stringField) Regex(regex bson.Regex) filter.Filter {
	return filter.New(s, "$regex", regex)
}

func NewStringField(name string) StringField {
	return &stringField{baseField[string]{name}}
}

// ComparableField todo *index.BaseKey
type ComparableField[T ~bool | bson.ObjectID] interface {
	mongodb.Field
	filter.ComparableFilter[T]
	updater.BaseUpdater[T]
}

func NewComparableField[T ~bool | bson.ObjectID](name string) ComparableField[T] {
	return &baseField[T]{name}
}

type BoolField = ComparableField[bool]
type ObjectIDField = ComparableField[bson.ObjectID]

// ComputableField todo *index.BaseKey
type ComputableField[T ~float32 | ~float64] interface {
	mongodb.Field
	filter.BaseFilter[T]
	updater.ComputableUpdater[T]
}

func NewComputableField[T ~float32 | ~float64](name string) ComputableField[T] {
	return &baseField[T]{name}
}

type Float32Field = ComputableField[float32]
type Float64Field = ComputableField[float64]

// Decimal128Field todo *index.BaseKey
type Decimal128Field interface {
	mongodb.Field
	filter.ComparableFilter[bson.Decimal128]
	updater.ComputableUpdater[bson.Decimal128]
}

func NewDecimal128Field(name string) Decimal128Field {
	return &baseField[bson.Decimal128]{name}
}

// BinaryField todo *index.BaseKey
type BinaryField interface {
	mongodb.Field
	filter.ComparableFilter[bson.Binary]
	updater.BaseUpdater[bson.Binary]
}

func NewBinaryField(name string) BinaryField {
	return &baseField[bson.Binary]{name}
}
