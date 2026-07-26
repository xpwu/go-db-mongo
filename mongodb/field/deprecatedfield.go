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
			return &BaseField[any]{name: name}
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

type float320FUpdaterF interface {
	deprecatedBaseUpdater
	Min(value float32) updater.Updater
	Max(value float32) updater.Updater
	Set(value float32) updater.Updater
	SetOnIns(value float32) updater.Updater
}

type float320FFilterF interface {
	deprecatedBaseFilter
	Gt(value float32) filter.Filter
	Lt(value float32) filter.Filter
	GtField(f filter.BaseFilterField[float32]) filter.Filter
	LtField(f filter.BaseFilterField[float32]) filter.Filter
	In(values []float32) filter.Filter
	Nin(values []float32) filter.Filter
}

type float320F interface {
	float320FUpdaterF
	float320FFilterF
	deprecatedBaseKey
}

// Deprecated: Float320F using Float32Field
type Float320F struct {
	ComputableField[float32]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Float320F) Min(value float32) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Float320F) Max(value float32) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Float320F) SetOnIns(value float32) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: this is a bug
func (f *Float320F) In(values []float32) filter.Filter {
	return filter.New(f, "$in", values)
}

// Deprecated: this is a bug
func (f *Float320F) Nin(values []float32) filter.Filter {
	return filter.New(f, "$nin", values)
}

// Deprecated: NewFloat320F using NewFloat32Field
func NewFloat320F(fName string) *Float320F {
	return &Float320F{NewFloat32Field(fName)}
}

// Deprecated:
type (
	Float320FUpdaterF = Float320F
	Float320FFilterF  = Float320F
)

var (
	_ float320F         = &Float320F{}
	_ float320FUpdaterF = &Float320F{}
	_ float320FFilterF  = &Float320F{}
)

type float321Field struct {
	*depTypeArrayField[float32, Float32Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *float321Field) Push(value float32) updater.Updater {
	return da.ArrayComparableField.Push([]float32{value})
}

// Deprecated: Float321Field using ArrayComparableField[float32, Float32Field]
type Float321Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Float320F
	EleOne() Float320F
	EleThat() Float320FUpdaterF
	EleAll() Float320FUpdaterF
	EleByFid(identifier string) Float320FUpdaterF
	DeclFid(identifier string) Float320FFilterF
	Include(a []float32) filter.Filter
	Eq(a []float32) filter.Filter
	Set(a []float32) updater.Updater
	AddToSet(value float32) updater.Updater
	AddToSetValues(a []float32) updater.Updater
	Pull(value float32) updater.Updater
	PullAll(a []float32) updater.Updater
	Push(value float32) updater.Updater
	PushByModifier(m updater.PushModifier, each []float32) updater.Updater
}

// Deprecated: NewFloat321Field using ArrayComparableField[float32, Float32Field]
func NewFloat321Field(fName string) Float321Field {
	return &float321Field{newDepCompAbleArrF[float32, Float32Field](fName, NewFloat32Field)}
}

type int0FUpdaterF interface {
	deprecatedBaseUpdater
	Inc(num int) updater.Updater
	Mul(num int) updater.Updater
	Min(value int) updater.Updater
	Max(value int) updater.Updater
	Set(value int) updater.Updater
	SetOnIns(value int) updater.Updater
}

type int0FFilterF interface {
	deprecatedBaseFilter
	Mod(divisor, remainder int) filter.Filter
	Eq(value int) filter.Filter
	Ne(value int) filter.Filter
	NeField(f filter.ComparableFilterField[int]) filter.Filter
	EqField(f filter.ComparableFilterField[int]) filter.Filter
	Gte(value int) filter.Filter
	Lte(value int) filter.Filter
	GteField(f filter.ComparableFilterField[int]) filter.Filter
	LteField(f filter.ComparableFilterField[int]) filter.Filter
	Gt(value int) filter.Filter
	Lt(value int) filter.Filter
	GtField(f filter.BaseFilterField[int]) filter.Filter
	LtField(f filter.BaseFilterField[int]) filter.Filter
	In(values []int) filter.Filter
	Nin(values []int) filter.Filter
}

type int0F interface {
	int0FUpdaterF
	int0FFilterF
	deprecatedBaseKey
}

// Deprecated: Int0F using IntField
type Int0F struct {
	IntegerField[int]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Int0F) Min(value int) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Int0F) Max(value int) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Int0F) SetOnIns(value int) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewInt0F using NewIntField
func NewInt0F(fName string) *Int0F {
	return &Int0F{NewIntField(fName)}
}

// Deprecated:
type (
	Int0FUpdaterF = Int0F
	Int0FFilterF  = Int0F
)

var (
	_ int0F         = &Int0F{}
	_ int0FUpdaterF = &Int0F{}
	_ int0FFilterF  = &Int0F{}
)

type int1Field struct {
	*depTypeArrayField[int, IntField]
}

// Deprecated: Push using ArrayField[].Push
func (da *int1Field) Push(value int) updater.Updater {
	return da.ArrayComparableField.Push([]int{value})
}

// Deprecated: Int1Field using ArrayComparableField[int, IntField]
type Int1Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Int0F
	EleOne() Int0F
	EleThat() Int0FUpdaterF
	EleAll() Int0FUpdaterF
	EleByFid(identifier string) Int0FUpdaterF
	DeclFid(identifier string) Int0FFilterF
	Include(a []int) filter.Filter
	Eq(a []int) filter.Filter
	Set(a []int) updater.Updater
	AddToSet(value int) updater.Updater
	AddToSetValues(a []int) updater.Updater
	Pull(value int) updater.Updater
	PullAll(a []int) updater.Updater
	Push(value int) updater.Updater
	PushByModifier(m updater.PushModifier, each []int) updater.Updater
}

// Deprecated: NewInt1Field using ArrayComparableField[int, IntField]
func NewInt1Field(fName string) Int1Field {
	return &int1Field{newDepCompAbleArrF[int, IntField](fName, NewIntField)}
}

// Deprecated: use Int8Field instead.
type Int80F struct {
	IntegerField[int8]
}

// interfaces for compile-time check
type int80F interface {
	int80FUpdaterF
	int80FFilterF
	deprecatedBaseKey
}

type int80FUpdaterF interface {
	Inc(num int8) updater.Updater
	Mul(num int8) updater.Updater
	Set(value int8) updater.Updater
	Min(value int8) updater.Updater
	Max(value int8) updater.Updater
	SetOnIns(value int8) updater.Updater
}

type int80FFilterF interface {
	Eq(value int8) filter.Filter
	Ne(value int8) filter.Filter
	Gt(value int8) filter.Filter
	Lt(value int8) filter.Filter
	Gte(value int8) filter.Filter
	Lte(value int8) filter.Filter
	In(values []int8) filter.Filter
	Nin(values []int8) filter.Filter
	Mod(divisor, remainder int8) filter.Filter
	EqField(f filter.ComparableFilterField[int8]) filter.Filter
	NeField(f filter.ComparableFilterField[int8]) filter.Filter
	GtField(f filter.BaseFilterField[int8]) filter.Filter
	LtField(f filter.BaseFilterField[int8]) filter.Filter
	GteField(f filter.ComparableFilterField[int8]) filter.Filter
	LteField(f filter.ComparableFilterField[int8]) filter.Filter
}

// compile-time interface checks
var _ int80F = &Int80F{}
var _ int80FUpdaterF = &Int80F{}
var _ int80FFilterF = &Int80F{}

// Deprecated: use NewInt8Field instead.
func NewInt80F(fieldName string) *Int80F {
	return &Int80F{NewInt8Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Int80F) Min(value int8) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Int80F) Max(value int8) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Int80F) SetOnIns(value int8) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Int80F directly.
type Int80FUpdaterF = Int80F

// Deprecated: use Int80F directly.
type Int80FFilterF = Int80F

type int81Field struct {
	*depTypeArrayField[int, Int8Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *int81Field) Push(value int) updater.Updater {
	return da.ArrayComparableField.Push([]int{value})
}

// Deprecated: Int81Field using ArrayComparableField[int, Int8Field]
type Int81Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Int80F
	EleOne() Int80F
	EleThat() Int80FUpdaterF
	EleAll() Int80FUpdaterF
	EleByFid(identifier string) Int80FUpdaterF
	DeclFid(identifier string) Int80FFilterF
	Include(a []int) filter.Filter
	Eq(a []int) filter.Filter
	Set(a []int) updater.Updater
	AddToSet(value int) updater.Updater
	AddToSetValues(a []int) updater.Updater
	Pull(value int) updater.Updater
	PullAll(a []int) updater.Updater
	Push(value int) updater.Updater
	PushByModifier(m updater.PushModifier, each []int) updater.Updater
}

// Deprecated: NewInt81Field using ArrayComparableField[int, UintField]
func NewInt81Field(fName string) Int81Field {
	return &int81Field{newDepCompAbleArrF[int, Int8Field](fName, NewInt8Field)}
}

// Deprecated: use Int16Field instead.
type Int160F struct {
	IntegerField[int16]
}

// interfaces for compile-time check
type int160F interface {
	int160FUpdaterF
	int160FFilterF
	deprecatedBaseKey
}

type int160FUpdaterF interface {
	Inc(num int16) updater.Updater
	Mul(num int16) updater.Updater
	Set(value int16) updater.Updater
	Min(value int16) updater.Updater
	Max(value int16) updater.Updater
	SetOnIns(value int16) updater.Updater
}

type int160FFilterF interface {
	Eq(value int16) filter.Filter
	Ne(value int16) filter.Filter
	Gt(value int16) filter.Filter
	Lt(value int16) filter.Filter
	Gte(value int16) filter.Filter
	Lte(value int16) filter.Filter
	In(values []int16) filter.Filter
	Nin(values []int16) filter.Filter
	Mod(divisor, remainder int16) filter.Filter
	EqField(f filter.ComparableFilterField[int16]) filter.Filter
	NeField(f filter.ComparableFilterField[int16]) filter.Filter
	GtField(f filter.BaseFilterField[int16]) filter.Filter
	LtField(f filter.BaseFilterField[int16]) filter.Filter
	GteField(f filter.ComparableFilterField[int16]) filter.Filter
	LteField(f filter.ComparableFilterField[int16]) filter.Filter
}

// compile-time interface checks
var _ int160F = &Int160F{}
var _ int160FUpdaterF = &Int160F{}
var _ int160FFilterF = &Int160F{}

// Deprecated: use NewInt16Field instead.
func NewInt160F(fieldName string) *Int160F {
	return &Int160F{NewInt16Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Int160F) Min(value int16) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Int160F) Max(value int16) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Int160F) SetOnIns(value int16) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Int16Field directly.
type Int160FUpdaterF = Int160F

// Deprecated: use Int16Field directly.
type Int160FFilterF = Int160F

type int161Field struct {
	*depTypeArrayField[int16, Int16Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *int161Field) Push(value int16) updater.Updater {
	return da.ArrayComparableField.Push([]int16{value})
}

// Deprecated: Int161Field using ArrayComparableField[int16, Int16Field]
type Int161Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Int160F
	EleOne() Int160F
	EleThat() Int160FUpdaterF
	EleAll() Int160FUpdaterF
	EleByFid(identifier string) Int160FUpdaterF
	DeclFid(identifier string) Int160FFilterF
	Include(a []int16) filter.Filter
	Eq(a []int16) filter.Filter
	Set(a []int16) updater.Updater
	AddToSet(value int16) updater.Updater
	AddToSetValues(a []int16) updater.Updater
	Pull(value int16) updater.Updater
	PullAll(a []int16) updater.Updater
	Push(value int16) updater.Updater
	PushByModifier(m updater.PushModifier, each []int16) updater.Updater
}

// Deprecated: NewInt161Field using ArrayComparableField[int16, Int16Field]
func NewInt161Field(fName string) Int161Field {
	return &int161Field{newDepCompAbleArrF[int16, Int16Field](fName, NewInt16Field)}
}

// Deprecated: use Int32Field instead.
type Int320F struct {
	IntegerField[int32]
}

// interfaces for compile-time check
type int320F interface {
	int320FUpdaterF
	int320FFilterF
	deprecatedBaseKey
}

type int320FUpdaterF interface {
	Inc(num int32) updater.Updater
	Mul(num int32) updater.Updater
	Set(value int32) updater.Updater
	Min(value int32) updater.Updater
	Max(value int32) updater.Updater
	SetOnIns(value int32) updater.Updater
}

type int320FFilterF interface {
	Eq(value int32) filter.Filter
	Ne(value int32) filter.Filter
	Gt(value int32) filter.Filter
	Lt(value int32) filter.Filter
	Gte(value int32) filter.Filter
	Lte(value int32) filter.Filter
	In(values []int32) filter.Filter
	Nin(values []int32) filter.Filter
	Mod(divisor, remainder int32) filter.Filter
	EqField(f filter.ComparableFilterField[int32]) filter.Filter
	NeField(f filter.ComparableFilterField[int32]) filter.Filter
	GtField(f filter.BaseFilterField[int32]) filter.Filter
	LtField(f filter.BaseFilterField[int32]) filter.Filter
	GteField(f filter.ComparableFilterField[int32]) filter.Filter
	LteField(f filter.ComparableFilterField[int32]) filter.Filter
}

// compile-time interface checks
var _ int320F = &Int320F{}
var _ int320FUpdaterF = &Int320F{}
var _ int320FFilterF = &Int320F{}

// Deprecated: use NewInt32Field instead.
func NewInt320F(fieldName string) *Int320F {
	return &Int320F{NewInt32Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Int320F) Min(value int32) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Int320F) Max(value int32) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Int320F) SetOnIns(value int32) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Int320F directly.
type Int320FUpdaterF = Int320F

// Deprecated: use Int320F directly.
type Int320FFilterF = Int320F

type int321Field struct {
	*depTypeArrayField[int32, Int32Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *int321Field) Push(value int32) updater.Updater {
	return da.ArrayComparableField.Push([]int32{value})
}

// Deprecated: Int321Field using ArrayComparableField[int32, Int32Field]
type Int321Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Int320F
	EleOne() Int320F
	EleThat() Int320FUpdaterF
	EleAll() Int320FUpdaterF
	EleByFid(identifier string) Int320FUpdaterF
	DeclFid(identifier string) Int320FFilterF
	Include(a []int32) filter.Filter
	Eq(a []int32) filter.Filter
	Set(a []int32) updater.Updater
	AddToSet(value int32) updater.Updater
	AddToSetValues(a []int32) updater.Updater
	Pull(value int32) updater.Updater
	PullAll(a []int32) updater.Updater
	Push(value int32) updater.Updater
	PushByModifier(m updater.PushModifier, each []int32) updater.Updater
}

// Deprecated: NewInt321Field using ArrayComparableField[int32, Int32Field]
func NewInt321Field(fName string) Int321Field {
	return &int321Field{newDepCompAbleArrF[int32, Int32Field](fName, NewInt32Field)}
}

// Deprecated: use Int64Field instead.
type Int640F struct {
	IntegerField[int64]
}

// interfaces for compile-time check
type int640F interface {
	int640FUpdaterF
	int640FFilterF
	deprecatedBaseKey
}

type int640FUpdaterF interface {
	Inc(num int64) updater.Updater
	Mul(num int64) updater.Updater
	Set(value int64) updater.Updater
	Min(value int64) updater.Updater
	Max(value int64) updater.Updater
	SetOnIns(value int64) updater.Updater
}

type int640FFilterF interface {
	Eq(value int64) filter.Filter
	Ne(value int64) filter.Filter
	Gt(value int64) filter.Filter
	Lt(value int64) filter.Filter
	Gte(value int64) filter.Filter
	Lte(value int64) filter.Filter
	In(values []int64) filter.Filter
	Nin(values []int64) filter.Filter
	Mod(divisor, remainder int64) filter.Filter
	EqField(f filter.ComparableFilterField[int64]) filter.Filter
	NeField(f filter.ComparableFilterField[int64]) filter.Filter
	GtField(f filter.BaseFilterField[int64]) filter.Filter
	LtField(f filter.BaseFilterField[int64]) filter.Filter
	GteField(f filter.ComparableFilterField[int64]) filter.Filter
	LteField(f filter.ComparableFilterField[int64]) filter.Filter
}

// compile-time interface checks
var _ int640F = &Int640F{}
var _ int640FUpdaterF = &Int640F{}
var _ int640FFilterF = &Int640F{}

// Deprecated: use NewInt64Field instead.
func NewInt640F(fieldName string) *Int640F {
	return &Int640F{NewInt64Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Int640F) Min(value int64) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Int640F) Max(value int64) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Int640F) SetOnIns(value int64) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Int640F directly.
type Int640FUpdaterF = Int640F

// Deprecated: use Int640F directly.
type Int640FFilterF = Int640F

type int641Field struct {
	*depTypeArrayField[int64, Int64Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *int641Field) Push(value int64) updater.Updater {
	return da.ArrayComparableField.Push([]int64{value})
}

// Deprecated: Int641Field using ArrayComparableField[int64, Int64Field]
type Int641Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Int640F
	EleOne() Int640F
	EleThat() Int640FUpdaterF
	EleAll() Int640FUpdaterF
	EleByFid(identifier string) Int640FUpdaterF
	DeclFid(identifier string) Int640FFilterF
	Include(a []int64) filter.Filter
	Eq(a []int64) filter.Filter
	Set(a []int64) updater.Updater
	AddToSet(value int64) updater.Updater
	AddToSetValues(a []int64) updater.Updater
	Pull(value int64) updater.Updater
	PullAll(a []int64) updater.Updater
	Push(value int64) updater.Updater
	PushByModifier(m updater.PushModifier, each []int64) updater.Updater
}

// Deprecated: NewInt641Field using ArrayComparableField[int64, Int64Field]
func NewInt641Field(fName string) Int641Field {
	return &int641Field{newDepCompAbleArrF[int64, Int64Field](fName, NewInt64Field)}
}
