package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type IntegerField[T Integer] struct {
	*baseKey
	*IntegerFieldFilter[T]
	*ComputableUpdater[T]
}

type IntegerFieldFilter[T Integer] struct {
	*ComparableFieldFilter[T]
}

func NewIntegerField[T Integer](fName string) *IntegerField[T] {
	b := &base{fName}
	flt := &IntegerFieldFilter[T]{
		&ComparableFieldFilter[T]{&baseFilter[T]{b}}}

	return &IntegerField[T]{&baseKey{b}, flt,
		&ComputableUpdater[T]{&baseUpdater[T]{b}},
	}
}

func (i *IntegerFieldFilter[T]) Mod(divisor, remainder T) filter.Filter {
	return filter.New(i, "$mod", bson.A{divisor, remainder})
}

type BoolField struct {
	*baseKey
	*ComparableFieldFilter[bool]
	*baseUpdater[bool]
}

func NewBoolField(fName string) *BoolField {
	b := &base{fName}

	return &BoolField{&baseKey{b},
		&ComparableFieldFilter[bool]{&baseFilter[bool]{b}},
		&baseUpdater[bool]{b},
	}
}

type StringField struct {
	*baseKey
	*StringFieldFilter
	*baseUpdater[string]
}

type StringFieldFilter struct {
	*ComparableFieldFilter[string]
}

func (s *StringFieldFilter) Regex(regex bson.Regex) filter.Filter {
	return filter.New(s, "$regex", regex)
}
