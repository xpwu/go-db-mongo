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
	// 已经由 SetOnInsert(value T) Updater 代替
	// 暂时不考虑 interface{} 与 T 的兼容性，上层在调用旧版本时，正常使用都应该是传的 T
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
type Binary0F struct {
	BinaryField
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Binary0F) Min(value bson.Binary) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Binary0F) Max(value bson.Binary) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Binary0F) SetOnIns(value bson.Binary) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewBinary0F using NewBinaryField
func NewBinary0F(fName string) *Binary0F {
	return &Binary0F{NewBinaryField(fName)}
}

// Deprecated
type (
	Binary0FUpdaterF = Binary0F
	Binary0FFilterF  = Binary0F
)

// Deprecated
var (
	_ binary0F         = &Binary0F{}
	_ binary0FUpdaterF = &Binary0FUpdaterF{}
	_ binary0FFilterF  = &Binary0FFilterF{}
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

type bool0FUpdaterF interface {
	deprecatedBaseUpdater
	Min(value bool) updater.Updater
	Max(value bool) updater.Updater
	Set(value bool) updater.Updater
	SetOnIns(value bool) updater.Updater
}

type bool0FFilterF interface {
	deprecatedBaseFilter
	Eq(value bool) filter.Filter
	Ne(value bool) filter.Filter
	NeField(f filter.ComparableFilterField[bool]) filter.Filter
	EqField(f filter.ComparableFilterField[bool]) filter.Filter
	Gte(value bool) filter.Filter
	Lte(value bool) filter.Filter
	GteField(f filter.ComparableFilterField[bool]) filter.Filter
	LteField(f filter.ComparableFilterField[bool]) filter.Filter
	Gt(value bool) filter.Filter
	Lt(value bool) filter.Filter
	GtField(f filter.BaseFilterField[bool]) filter.Filter
	LtField(f filter.BaseFilterField[bool]) filter.Filter
	In(values []bool) filter.Filter
	Nin(values []bool) filter.Filter
}

type bool0F interface {
	bool0FUpdaterF
	bool0FFilterF
	deprecatedBaseKey
}

// Deprecated: Bool0F using BoolField
type Bool0F struct {
	ComparableField[bool]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Bool0F) Min(value bool) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Bool0F) Max(value bool) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Bool0F) SetOnIns(value bool) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewBool0F using NewBoolField
func NewBool0F(fName string) *Bool0F {
	return &Bool0F{NewBoolField(fName)}
}

// Deprecated:
type (
	Bool0FUpdaterF = Bool0F
	Bool0FFilterF  = Bool0F
)

var (
	_ bool0F         = &Bool0F{}
	_ bool0FUpdaterF = &Bool0FUpdaterF{}
	_ bool0FFilterF  = &Bool0FFilterF{}
)

type bool1Field struct {
	*depTypeArrayField[bool, BoolField]
}

// Deprecated: Push using ArrayField[].Push
func (da *bool1Field) Push(value bool) updater.Updater {
	return da.ArrayComparableField.Push([]bool{value})
}

// Deprecated: Bool1Field using ArrayComparableField[bool, BoolField]
type Bool1Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Bool0F
	EleOne() Bool0F
	EleThat() Bool0FUpdaterF
	EleAll() Bool0FUpdaterF
	EleByFid(identifier string) Bool0FUpdaterF
	DeclFid(identifier string) Bool0FFilterF
	Include(a []bool) filter.Filter
	Eq(a []bool) filter.Filter
	Set(a []bool) updater.Updater
	AddToSet(value bool) updater.Updater
	AddToSetValues(a []bool) updater.Updater
	Pull(value bool) updater.Updater
	PullAll(a []bool) updater.Updater
	Push(value bool) updater.Updater
	PushByModifier(m updater.PushModifier, each []bool) updater.Updater
}

// Deprecated: NewBool1Field using NewArrayEqualAbleField[bool, BoolField]
func NewBool1Field(fName string) Bool1Field {
	return &bool1Field{newDepCompAbleArrF[bool, BoolField](fName, NewBoolField)}
}

type decimal1280FUpdaterF interface {
	deprecatedBaseUpdater
	Inc(num bson.Decimal128) updater.Updater
	Mul(num bson.Decimal128) updater.Updater
	Min(value bson.Decimal128) updater.Updater
	Max(value bson.Decimal128) updater.Updater
	Set(value bson.Decimal128) updater.Updater
	SetOnIns(value bson.Decimal128) updater.Updater
}

type decimal1280FFilterF interface {
	deprecatedBaseFilter
	Gt(value bson.Decimal128) filter.Filter
	Lt(value bson.Decimal128) filter.Filter
	GtField(f filter.BaseFilterField[bson.Decimal128]) filter.Filter
	LtField(f filter.BaseFilterField[bson.Decimal128]) filter.Filter
	In(values []bson.Decimal128) filter.Filter
	Nin(values []bson.Decimal128) filter.Filter
}

type decimal1280F interface {
	decimal1280FUpdaterF
	decimal1280FFilterF
	deprecatedBaseKey
}

// Deprecated: Decimal1280F using Decimal128Field
type Decimal1280F struct {
	Decimal128Field
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Decimal1280F) Min(value bson.Decimal128) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Decimal1280F) Max(value bson.Decimal128) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Decimal1280F) SetOnIns(value bson.Decimal128) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewDecimal1280F using NewDecimal128Field
func NewDecimal1280F(fName string) *Decimal1280F {
	return &Decimal1280F{NewDecimal128Field(fName)}
}

// Deprecated:
type (
	Decimal1280FUpdaterF = Decimal1280F
	Decimal1280FFilterF  = Decimal1280F
)

var (
	_ decimal1280F         = &Decimal1280F{}
	_ decimal1280FUpdaterF = &Decimal1280FUpdaterF{}
	_ decimal1280FFilterF  = &Decimal1280FFilterF{}
)

type decimal1281Field struct {
	*depTypeArrayField[bson.Decimal128, Decimal128Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *decimal1281Field) Push(value bson.Decimal128) updater.Updater {
	return da.ArrayComparableField.Push([]bson.Decimal128{value})
}

// Deprecated: Decimal1281Field using ArrayComparableField[bson.Decimal128, Decimal128Field]
type Decimal1281Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Decimal1280F
	EleOne() Decimal1280F
	EleThat() Decimal1280FUpdaterF
	EleAll() Decimal1280FUpdaterF
	EleByFid(identifier string) Decimal1280FUpdaterF
	DeclFid(identifier string) Decimal1280FFilterF
	Include(a []bson.Decimal128) filter.Filter
	Eq(a []bson.Decimal128) filter.Filter
	Set(a []bson.Decimal128) updater.Updater
	AddToSet(value bson.Decimal128) updater.Updater
	AddToSetValues(a []bson.Decimal128) updater.Updater
	Pull(value bson.Decimal128) updater.Updater
	PullAll(a []bson.Decimal128) updater.Updater
	Push(value bson.Decimal128) updater.Updater
	PushByModifier(m updater.PushModifier, each []bson.Decimal128) updater.Updater
}

// Deprecated: NewDecimal1281Field using NewArrayEqualAbleField[bson.Decimal128, Decimal128Field]
func NewDecimal1281Field(fName string) Decimal1281Field {
	return &decimal1281Field{newDepCompAbleArrF[bson.Decimal128, Decimal128Field](fName, NewDecimal128Field)}
}
