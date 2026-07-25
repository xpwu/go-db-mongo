package field

import (
	"fmt"
	"github.com/xpwu/go-db-mongo/mongodb"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type deprecatedBase interface {
	FullName() string
}

type deprecatedBaseKey interface {
	deprecatedBase
	AscIndex() index.Key
	DescIndex() index.Key
}

type deprecatedBaseUpdater interface {
	deprecatedBase
	Unset() updater.Updater
	// 已经由 Set(value T) Updater 与 SetOnInsert(value T) Updater 代替
	// 暂时不考虑 interface{} 与 T 的兼容性，上层在调用旧版本时，正常使用都应该是传的 T
	//Set(value interface{}) updater.Updater
	//SetOnInsert(value interface{}) updater.Updater
}

type deprecatedBaseFilter interface {
	deprecatedBase
	Exist() filter.Filter
	NotExist() filter.Filter
	Type(t bson.Type) filter.Filter
}

type binary0FUpdaterF interface {
	deprecatedBaseUpdater
	Set(value bson.Binary) updater.Updater
	SetOnIns(value bson.Binary) updater.Updater
}

type binary0FFilterF interface {
	deprecatedBaseFilter
	In(values []bson.Binary) filter.Filter
	Nin(values []bson.Binary) filter.Filter
}

type binary0F interface {
	binary0FUpdaterF
	binary0FFilterF
	deprecatedBaseKey
}

// Deprecated: Binary0F using BinaryField
type (
	Binary0F         = BinaryField
	Binary0FUpdaterF = BinaryField
	Binary0FFilterF  = BinaryField
)

// Deprecated: NewBinary0F using NewBinaryField
var (
	_           binary0F              = Binary0F(nil)
	_           binary0FUpdaterF      = Binary0FUpdaterF(nil)
	_           binary0FFilterF       = Binary0FFilterF(nil)
	NewBinary0F func(string) Binary0F = NewBinaryField
)

type arrayField interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	PopFirst() updater.Updater
	PopLast() updater.Updater
	// Deprecated:
	PullByF(f filter.Filter) updater.Updater
	// Deprecated:
	SameEleMatch(f filter.Filter) filter.Filter
	Size(sz int) filter.Filter
}

// Deprecated: Array using ArrayField
type Array struct {
	ArrayField[any, mongodb.Field]
}

// Deprecated: PullByF using ArrayField[].RemoveVirValue
func (a *Array) PullByF(f filter.Filter) updater.Updater {
	return updater.PullByFilter(a, f)
}

// Deprecated: SameEleMatch using ArrayField[].SameElemMeet
func (a *Array) SameEleMatch(f filter.Filter) filter.Filter {
	return filter.SameElemMatch(a, f)
}

// Deprecated: NewArray using NewArrayField
var (
	NewArray func(fName string) *Array = func(fName string) *Array {
		return &Array{NewArrayField[any, mongodb.Field](fName, func(name string) mongodb.Field {
			return &baseField[any]{name: name}
		})}
	}

	_ arrayField = NewArray("")
)

type depTypeArrayField[T any, ElemField mongodb.Field] struct {
	ArrayComparableField[T, ElemField]
}

// Deprecated: EleAt using ArrayField[].AtPos
func (da *depTypeArrayField[T, ElemField]) EleAt(index int) ElemField {
	return da.AtPos(index)
}

// Deprecated: PullByF using ArrayField[].RemoveVirValue
func (da *depTypeArrayField[T, ElemField]) PullByF(f filter.Filter) updater.Updater {
	return updater.PullByFilter(da, f)
}

// Deprecated: SameEleMatch using ArrayField[].SameElemMeet
func (da *depTypeArrayField[T, ElemField]) SameEleMatch(f filter.Filter) filter.Filter {
	return filter.SameElemMatch(da, f)
}

// Deprecated: EleOne using ArrayField[].Elems
func (da *depTypeArrayField[T, ElemField]) EleOne() ElemField {
	return da.Elems()
}

// Deprecated: EleThat using ArrayField[].FirstMatched
func (da *depTypeArrayField[T, ElemField]) EleThat() ElemField {
	return da.FirstMatched()
}

// Deprecated: EleAll using ArrayField[].UpdateAll
func (da *depTypeArrayField[T, ElemField]) EleAll() ElemField {
	return da.UpdateAll()
}

func (da *depTypeArrayField[T, ElemField]) arrayBaseField() *arrayBaseField[T, ElemField] {
	return da.ArrayComparableField.(*arrayBaseField[T, ElemField])
}

// Deprecated: EleByFid using ArrayField[].AtVirPos
func (da *depTypeArrayField[T, ElemField]) EleByFid(identifier string) ElemField {
	return da.arrayBaseField().newElemField(fmt.Sprintf("%s.$[%s]", da.FullName(), da.FullName()+identifier))
}

// Deprecated: DeclFid using ArrayField[].AtVirPos
func (da *depTypeArrayField[T, ElemField]) DeclFid(identifier string) ElemField {
	return da.arrayBaseField().newElemField(da.FullName() + identifier)
}

// Deprecated: Include using ArrayComparableUpdater[].CoverValues
func (da *depTypeArrayField[T, ElemField]) Include(a []T) filter.Filter {
	return da.CoverValues(a)
}

// Deprecated: AddToSet using ArrayField[].AddEach
func (da *depTypeArrayField[T, ElemField]) AddToSet(value T) updater.Updater {
	return da.AddEach([]T{value})
}

// Deprecated: AddToSetValues using ArrayField[].AddEach
func (da *depTypeArrayField[T, ElemField]) AddToSetValues(a []T) updater.Updater {
	return da.AddEach(a)
}

// Deprecated: Pull using ArrayComparableField[].RemoveValues
func (da *depTypeArrayField[T, ElemField]) Pull(value T) updater.Updater {
	return da.RemoveValues([]T{value})
}

// Deprecated: PullAll using ArrayComparableField[].RemoveValues
func (da *depTypeArrayField[T, ElemField]) PullAll(a []T) updater.Updater {
	return da.RemoveValues(a)
}

// Deprecated: PushByModifier using ArrayComparableField[].PushWith
func (da *depTypeArrayField[T, ElemField]) PushByModifier(m updater.PushModifier, each []T) updater.Updater {
	return updater.PushByModifier(da, m, each)
}

func newDepEqualAbleArrF[T filter.EqualAble[T], ElemField mongodb.Field](fName string,
	newElem func(name string) ElemField) *depTypeArrayField[T, ElemField] {

	arr := NewArrayEqualAbleField[T, ElemField](fName, newElem)
	return &depTypeArrayField[T, ElemField]{arr}
}

func newDepCompAbleArrF[T comparable, ElemField mongodb.Field](fName string,
	newElem func(name string) ElemField) *depTypeArrayField[T, ElemField] {

	arr := NewArrayComparableField[T, ElemField](fName, newElem)
	return &depTypeArrayField[T, ElemField]{arr}
}

type binary1Field struct {
	*depTypeArrayField[bson.Binary, BinaryField]
}

// Deprecated: Push using ArrayField[].Push
func (da *binary1Field) Push(value bson.Binary) updater.Updater {
	return da.ArrayComparableField.Push([]bson.Binary{value})
}

// Deprecated: Binary1Field using ArrayComparableField[bson.Binary, BinaryField]
type Binary1Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Binary0F
	EleOne() Binary0F
	EleThat() Binary0FUpdaterF
	EleAll() Binary0FUpdaterF
	EleByFid(identifier string) Binary0FUpdaterF
	DeclFid(identifier string) Binary0FFilterF
	Include(a []bson.Binary) filter.Filter
	Eq(a []bson.Binary) filter.Filter
	Set(a []bson.Binary) updater.Updater
	AddToSet(value bson.Binary) updater.Updater
	AddToSetValues(a []bson.Binary) updater.Updater
	Pull(value bson.Binary) updater.Updater
	PullAll(a []bson.Binary) updater.Updater
	Push(value bson.Binary) updater.Updater
	PushByModifier(m updater.PushModifier, each []bson.Binary) updater.Updater
}

// Deprecated: NewBinary1Field using NewArrayEqualAbleField[bson.Binary, BinaryField]
func NewBinary1Field(fName string) Binary1Field {
	return &binary1Field{newDepEqualAbleArrF[bson.Binary, BinaryField](fName, NewBinaryField)}
}
