package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type base struct {
	name string
}

func (b *base) FullName() string {
	return b.name
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type IntegerFieldFilter[T Integer] struct {
	*filter.ComparableFilter[T]
}

func (i *IntegerFieldFilter[T]) Mod(divisor, remainder T) filter.Filter {
	return filter.New(i, "$mod", bson.A{divisor, remainder})
}

type IntegerField[T Integer] struct {
	*index.BaseKey
	*IntegerFieldFilter[T]
	*updater.ComputableUpdater[T]
}

func NewIntegerField[T Integer](fName string) *IntegerField[T] {
	b := &base{fName}
	flt := &IntegerFieldFilter[T]{
		ComparableFilter: filter.NewComparableFilter[T](&filter.BaseFilter[T]{Field: b})}

	return &IntegerField[T]{
		BaseKey:            &index.BaseKey{Field: b},
		IntegerFieldFilter: flt,
		ComputableUpdater:  &updater.ComputableUpdater[T]{BaseUpdater: &updater.BaseUpdater[T]{Field: b}},
	}
}

type IntField = IntegerField[int]
type Int8Field = IntegerField[int8]
type Int16Field = IntegerField[int16]
type Int32Field = IntegerField[int32]
type Int64Field = IntegerField[int64]

type UnInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type UnIntegerField[T UnInteger] struct {
	*index.BaseKey
	*IntegerFieldFilter[T]
	*updater.UnsignedComputableUpdater[T]
}

func NewUnIntegerField[T UnInteger](fName string) *UnIntegerField[T] {
	b := &base{fName}
	flt := &IntegerFieldFilter[T]{
		ComparableFilter: filter.NewComparableFilter[T](&filter.BaseFilter[T]{Field: b})}
	c := &updater.ComputableUpdater[T]{BaseUpdater: &updater.BaseUpdater[T]{Field: b}}

	return &UnIntegerField[T]{
		BaseKey:                   &index.BaseKey{Field: b},
		IntegerFieldFilter:        flt,
		UnsignedComputableUpdater: &updater.UnsignedComputableUpdater[T]{ComputableUpdater: c},
	}
}

type UintField = UnIntegerField[uint]
type ByteField = UnIntegerField[byte]
type Uint8Field = UnIntegerField[uint8]
type Uint16Field = UnIntegerField[uint16]
type Uint32Field = UnIntegerField[uint32]
type Uint64Field = UnIntegerField[uint64]

type StringField struct {
	*index.BaseKey
	*StringFieldFilter
	*updater.BaseUpdater[string]
}

type StringFieldFilter struct {
	*filter.ComparableFilter[string]
}

func (s *StringFieldFilter) Regex(regex bson.Regex) filter.Filter {
	return filter.New(s, "$regex", regex)
}

func NewStringField(fName string) *StringField {
	b := &base{fName}
	flt := &StringFieldFilter{
		ComparableFilter: filter.NewComparableFilter[string](&filter.BaseFilter[string]{Field: b})}

	return &StringField{
		BaseKey:           &index.BaseKey{Field: b},
		StringFieldFilter: flt,
		BaseUpdater:       &updater.BaseUpdater[string]{Field: b},
	}
}

type ComparableField[T ~bool | bson.ObjectID] struct {
	*index.BaseKey
	*filter.ComparableFilter[T]
	*updater.BaseUpdater[T]
}

func NewComparableField[T ~bool | bson.ObjectID](fName string) *ComparableField[T] {
	b := &base{fName}

	return &ComparableField[T]{
		BaseKey:          &index.BaseKey{Field: b},
		ComparableFilter: filter.NewComparableFilter[T](&filter.BaseFilter[T]{Field: b}),
		BaseUpdater:      &updater.BaseUpdater[T]{Field: b},
	}
}

type BoolField = ComparableField[bool]
type ObjectIDField = ComparableField[bson.ObjectID]

type ComputableField[T ~float32 | ~float64] struct {
	*index.BaseKey
	*filter.BaseFilter[T]
	*updater.ComputableUpdater[T]
}

func NewComputableField[T ~float32 | ~float64](fName string) *ComputableField[T] {
	b := &base{fName}

	return &ComputableField[T]{
		BaseKey:           &index.BaseKey{Field: b},
		BaseFilter:        &filter.BaseFilter[T]{Field: b},
		ComputableUpdater: &updater.ComputableUpdater[T]{BaseUpdater: &updater.BaseUpdater[T]{Field: b}},
	}
}

type Float32Field = ComputableField[float32]
type Float64Field = ComputableField[float64]

type Decimal128Field struct {
	*index.BaseKey
	*filter.ComparableFilter[bson.Decimal128]
	*updater.ComputableUpdater[bson.Decimal128]
}

func NewDecimal128Field(fName string) *Decimal128Field {
	b := &base{fName}

	return &Decimal128Field{
		BaseKey:          &index.BaseKey{Field: b},
		ComparableFilter: filter.NewComparableFilter[bson.Decimal128](&filter.BaseFilter[bson.Decimal128]{Field: b}),
		ComputableUpdater: &updater.ComputableUpdater[bson.Decimal128]{
			BaseUpdater: &updater.BaseUpdater[bson.Decimal128]{Field: b}},
	}
}

type BinaryField struct {
	*index.BaseKey
	*filter.ComparableFilter[bson.Binary]
	*updater.BaseUpdater[bson.Binary]
}

func NewBinaryField(fName string) *BinaryField {
	b := &base{fName}

	return &BinaryField{
		BaseKey:          &index.BaseKey{Field: b},
		ComparableFilter: filter.NewEqualAbleFilter[bson.Binary](&filter.BaseFilter[bson.Binary]{Field: b}),
		BaseUpdater:      &updater.BaseUpdater[bson.Binary]{Field: b},
	}
}
