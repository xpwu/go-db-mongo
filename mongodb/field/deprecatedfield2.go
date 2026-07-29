package field

import (
	"fmt"
	"github.com/xpwu/go-db-mongo/mongodb"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Deprecated: Array using ArrayField
type Array interface {
	ArrayField[any, mongodb.Field]
	// Deprecated: PullByF using ArrayField[].RemoveVirValue
	PullByF(f filter.Filter) updater.Updater
	// Deprecated: SameEleMatch using ArrayField[].SameElemMeet
	SameEleMatch(f filter.Filter) filter.Filter
}

type depArray struct {
	ArrayField[any, mongodb.Field]
}

// Deprecated: PullByF using ArrayField[].RemoveVirValue
func (a *depArray) PullByF(f filter.Filter) updater.Updater {
	return updater.PullByFilter(a, f)
}

// Deprecated: SameEleMatch using ArrayField[].SameElemMeet
func (a *depArray) SameEleMatch(f filter.Filter) filter.Filter {
	return filter.SameElemMatch(a, f)
}

// Deprecated: NewArray using NewArrayField
func NewArray(fName string) Array {
	return &depArray{NewArrayField[any, mongodb.Field](fName, func(name string) mongodb.Field {
		return &BaseField[any]{name: name}
	})}
}

type depTypeArrayFieldI[T any, ElemField mongodb.Field] interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) ElemField
	EleOne() ElemField
	EleThat() ElemField
	EleAll() ElemField
	EleByFid(identifier string) ElemField
	DeclFid(identifier string) ElemField
	Include(a []T) filter.Filter
	Eq(a []T) filter.Filter
	Set(a []T) updater.Updater
	AddToSet(value T) updater.Updater
	AddToSetValues(a []T) updater.Updater
	Pull(value T) updater.Updater
	PullAll(a []T) updater.Updater
	Push(value T) updater.Updater
	PushByModifier(m updater.PushModifier, each []T) updater.Updater
}

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
	*depTypeArrayField[bson.Binary, Binary0F]
}

// Deprecated: Push using ArrayField[].Push
func (da *binary1Field) Push(value bson.Binary) updater.Updater {
	return da.ArrayComparableField.Push([]bson.Binary{value})
}

// Deprecated: Binary1Field using ArrayComparableField[bson.Binary, BinaryField]
type Binary1Field = depTypeArrayFieldI[bson.Binary, Binary0F]

// Deprecated: NewBinary1Field using NewArrayEqualAbleField[bson.Binary, BinaryField]
func NewBinary1Field(fName string) Binary1Field {
	return &binary1Field{newDepEqualAbleArrF[bson.Binary, Binary0F](fName, NewBinary0F)}
}

type decimal1281Field struct {
	*depTypeArrayField[bson.Decimal128, Decimal1280F]
}

// Deprecated: Push using ArrayField[].Push
func (da *decimal1281Field) Push(value bson.Decimal128) updater.Updater {
	return da.ArrayComparableField.Push([]bson.Decimal128{value})
}

// Deprecated: Decimal1281Field using ArrayComparableField[bson.Decimal128, Decimal128Field]
type Decimal1281Field = depTypeArrayFieldI[bson.Decimal128, Decimal1280F]

// Deprecated: NewDecimal1281Field using NewArrayEqualAbleField[bson.Decimal128, Decimal128Field]
func NewDecimal1281Field(fName string) Decimal1281Field {
	return &decimal1281Field{newDepCompAbleArrF[bson.Decimal128, Decimal1280F](fName, NewDecimal1280F)}
}

type bool1Field struct {
	*depTypeArrayField[bool, Bool0F]
}

// Deprecated: Push using ArrayField[].Push
func (da *bool1Field) Push(value bool) updater.Updater {
	return da.ArrayComparableField.Push([]bool{value})
}

// Deprecated: Bool1Field using ArrayComparableField[bool, BoolField]
type Bool1Field = depTypeArrayFieldI[bool, Bool0F]

// Deprecated: NewBool1Field using NewArrayEqualAbleField[bool, BoolField]
func NewBool1Field(fName string) Bool1Field {
	return &bool1Field{newDepCompAbleArrF[bool, Bool0F](fName, NewBool0F)}
}

type float321Field struct {
	*depTypeArrayField[float32, Float320F]
}

// Deprecated: Push using ArrayField[].Push
func (da *float321Field) Push(value float32) updater.Updater {
	return da.ArrayComparableField.Push([]float32{value})
}

// Deprecated: Float321Field using ArrayComparableField[float32, Float32Field]
type Float321Field = depTypeArrayFieldI[float32, Float320F]

// Deprecated: NewFloat321Field using ArrayComparableField[float32, Float32Field]
func NewFloat321Field(fName string) Float321Field {
	return &float321Field{newDepCompAbleArrF[float32, Float320F](fName, NewFloat320F)}
}

type int1Field struct {
	*depTypeArrayField[int, Int0F]
}

// Deprecated: Push using ArrayField[].Push
func (da *int1Field) Push(value int) updater.Updater {
	return da.ArrayComparableField.Push([]int{value})
}

// Deprecated: Int1Field using ArrayComparableField[int, IntField]
type Int1Field = depTypeArrayFieldI[int, Int0F]

// Deprecated: NewInt1Field using ArrayComparableField[int, IntField]
func NewInt1Field(fName string) Int1Field {
	return &int1Field{newDepCompAbleArrF[int, Int0F](fName, NewInt0F)}
}

type int81Field struct {
	*depTypeArrayField[int8, Int80F]
}

// Deprecated: Push using ArrayField[].Push
func (da *int81Field) Push(value int8) updater.Updater {
	return da.ArrayComparableField.Push([]int8{value})
}

// Deprecated: Int81Field using ArrayComparableField[int8, Int8Field]
type Int81Field = depTypeArrayFieldI[int8, Int80F]

// Deprecated: NewInt81Field using ArrayComparableField[int, UintField]
func NewInt81Field(fName string) Int81Field {
	return &int81Field{newDepCompAbleArrF[int8, Int80F](fName, NewInt80F)}
}

type int161Field struct {
	*depTypeArrayField[int16, Int160F]
}

// Deprecated: Push using ArrayField[].Push
func (da *int161Field) Push(value int16) updater.Updater {
	return da.ArrayComparableField.Push([]int16{value})
}

// Deprecated: Int161Field using ArrayComparableField[int16, Int16Field]
type Int161Field = depTypeArrayFieldI[int16, Int160F]

// Deprecated: NewInt161Field using ArrayComparableField[int16, Int16Field]
func NewInt161Field(fName string) Int161Field {
	return &int161Field{newDepCompAbleArrF[int16, Int160F](fName, NewInt160F)}
}

type int321Field struct {
	*depTypeArrayField[int32, Int320F]
}

// Deprecated: Push using ArrayField[].Push
func (da *int321Field) Push(value int32) updater.Updater {
	return da.ArrayComparableField.Push([]int32{value})
}

// Deprecated: Int321Field using ArrayComparableField[int32, Int32Field]
type Int321Field = depTypeArrayFieldI[int32, Int320F]

// Deprecated: NewInt321Field using ArrayComparableField[int32, Int32Field]
func NewInt321Field(fName string) Int321Field {
	return &int321Field{newDepCompAbleArrF[int32, Int320F](fName, NewInt320F)}
}

type int641Field struct {
	*depTypeArrayField[int64, Int640F]
}

// Deprecated: Push using ArrayField[].Push
func (da *int641Field) Push(value int64) updater.Updater {
	return da.ArrayComparableField.Push([]int64{value})
}

// Deprecated: Int641Field using ArrayComparableField[int64, Int64Field]
type Int641Field = depTypeArrayFieldI[int64, Int640F]

// Deprecated: NewInt641Field using ArrayComparableField[int64, Int64Field]
func NewInt641Field(fName string) Int641Field {
	return &int641Field{newDepCompAbleArrF[int64, Int640F](fName, NewInt640F)}
}

type uint1Field struct {
	*depTypeArrayField[uint, Uint0F]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint1Field) Push(value uint) updater.Updater {
	return da.ArrayComparableField.Push([]uint{value})
}

// Deprecated: Uint1Field using ArrayComparableField[uint, UintField]
type Uint1Field = depTypeArrayFieldI[uint, Uint0F]

// Deprecated: NewInt1Field using ArrayComparableField[uint, UintField]
func NewUint1Field(fName string) Uint1Field {
	return &uint1Field{newDepCompAbleArrF[uint, Uint0F](fName, NewUint0F)}
}

// todo
type uint81Field struct {
	*depTypeArrayField[uint8, Uint80F]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint81Field) Push(value uint8) updater.Updater {
	return da.ArrayComparableField.Push([]uint8{value})
}

// Deprecated: Uint81Field using ArrayComparableField[uint8, Uint8Field]
type Uint81Field = depTypeArrayFieldI[uint8, Uint80F]

// Deprecated: NewUint81Field using ArrayComparableField[uint8, Uint8Field]
func NewUint81Field(fName string) Uint81Field {
	return &uint81Field{newDepCompAbleArrF[uint8, Uint80F](fName, NewUint80F)}
}

type uint161Field struct {
	*depTypeArrayField[uint16, Uint160F]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint161Field) Push(value uint16) updater.Updater {
	return da.ArrayComparableField.Push([]uint16{value})
}

// Deprecated: Uint161Field using ArrayComparableField[uint16, Uint16Field]
type Uint161Field = depTypeArrayFieldI[uint16, Uint160F]

// Deprecated: NewUint161Field using ArrayComparableField[uint16, Uint16Field]
func NewUint161Field(fName string) Uint161Field {
	return &uint161Field{newDepCompAbleArrF[uint16, Uint160F](fName, NewUint160F)}
}

type uint321Field struct {
	*depTypeArrayField[uint32, Uint320F]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint321Field) Push(value uint32) updater.Updater {
	return da.ArrayComparableField.Push([]uint32{value})
}

// Deprecated: Uint321Field using ArrayComparableField[uint32, Uint32Field]
type Uint321Field = depTypeArrayFieldI[uint32, Uint320F]

// Deprecated: NewUint321Field using ArrayComparableField[uint32, Uint32Field]
func NewUint321Field(fName string) Uint321Field {
	return &uint321Field{newDepCompAbleArrF[uint32, Uint320F](fName, NewUint320F)}
}

type uint641Field struct {
	*depTypeArrayField[uint64, Uint640F]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint641Field) Push(value uint64) updater.Updater {
	return da.ArrayComparableField.Push([]uint64{value})
}

// Deprecated: Uint641Field using ArrayComparableField[uint64, Uint64Field]
type Uint641Field = depTypeArrayFieldI[uint64, Uint640F]

// Deprecated: NewUint641Field using ArrayComparableField[uint64, Uint64Field]
func NewUint641Field(fName string) Uint641Field {
	return &uint641Field{newDepCompAbleArrF[uint64, Uint640F](fName, NewUint640F)}
}

type float641Field struct {
	*depTypeArrayField[float64, Float640F]
}

// Deprecated: Push using ArrayField[].Push
func (da *float641Field) Push(value float64) updater.Updater {
	return da.ArrayComparableField.Push([]float64{value})
}

// Deprecated: Float641Field using ArrayComparableField[float64, Float64Field]
type Float641Field = depTypeArrayFieldI[float64, Float640F]

// Deprecated: NewFloat641Field using ArrayComparableField[float64, Float64Field]
func NewFloat641Field(fName string) Float641Field {
	return &float641Field{newDepCompAbleArrF[float64, Float640F](fName, NewFloat640F)}
}

type string1Field struct {
	*depTypeArrayField[string, String0F]
}

// Deprecated: Push using ArrayField[].Push
func (da *string1Field) Push(value string) updater.Updater {
	return da.ArrayComparableField.Push([]string{value})
}

// Deprecated: String1Field using ArrayComparableField[string, StringField]
type String1Field = depTypeArrayFieldI[string, String0F]

// Deprecated: NewString1Field using ArrayComparableField[string, StringField]
func NewString1Field(fName string) String1Field {
	return &string1Field{newDepCompAbleArrF[string, String0F](fName, NewString0F)}
}
