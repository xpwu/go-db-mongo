package field

import (
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

type depTypeBaseUpdater[T any] interface {
	deprecatedBaseUpdater
	Set(value T) updater.Updater
	depMethods[T]
}

type depTypeBaseFilter[T any] interface {
	deprecatedBaseFilter
	In(values []T) filter.Filter
	Nin(values []T) filter.Filter
	Eq(value T) filter.Filter
	Ne(value T) filter.Filter
	NeField(f filter.ComparableFilterField[T]) filter.Filter
	EqField(f filter.ComparableFilterField[T]) filter.Filter
	Gte(value T) filter.Filter
	Lte(value T) filter.Filter
	GteField(f filter.ComparableFilterField[T]) filter.Filter
	LteField(f filter.ComparableFilterField[T]) filter.Filter
	Gt(value T) filter.Filter
	Lt(value T) filter.Filter
	GtField(f filter.BaseFilterField[T]) filter.Filter
	LtField(f filter.BaseFilterField[T]) filter.Filter
}

type depMethods[T any] interface {
	// Deprecated: Min using: updater.BaseUpdater[].SetMin
	Min(value T) updater.Updater
	// Deprecated: Max using: updater.BaseUpdater[].SetMax
	Max(value T) updater.Updater
	// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
	SetOnIns(value T) updater.Updater
}

type depMethodsImpl[T any] struct {
	updater.BaseUpdater[T]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *depMethodsImpl[T]) Min(value T) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *depMethodsImpl[T]) Max(value T) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *depMethodsImpl[T]) SetOnIns(value T) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: Binary0F using BinaryField
type Binary0F interface {
	BinaryField
	depMethods[bson.Binary]
}

// Deprecated
type (
	Binary0FUpdaterF = Binary0F
	Binary0FFilterF  = Binary0F
)

type binary0F struct {
	BinaryField
	depMethodsImpl[bson.Binary]
}

// Deprecated: NewBinary0F using NewBinaryField
func NewBinary0F(fName string) Binary0F {
	f := NewBinaryField(fName)
	return &binary0F{f, depMethodsImpl[bson.Binary]{f}}
}

// Deprecated: Bool0F using BoolField
type Bool0F interface {
	ComparableField[bool]
	depMethods[bool]
}

// Deprecated:
type (
	Bool0FUpdaterF = Bool0F
	Bool0FFilterF  = Bool0F
)

type bool0F struct {
	ComparableField[bool]
	depMethodsImpl[bool]
}

// Deprecated: NewBool0F using NewBoolField
func NewBool0F(fName string) Bool0F {
	f := NewBoolField(fName)
	return &bool0F{f, depMethodsImpl[bool]{f}}
}

// Deprecated: Decimal1280F using Decimal128Field
type Decimal1280F interface {
	Decimal128Field
	depMethods[bson.Decimal128]
}

// Deprecated:
type (
	Decimal1280FUpdaterF = Decimal1280F
	Decimal1280FFilterF  = Decimal1280F
)

type decimal1280F struct {
	Decimal128Field
	depMethodsImpl[bson.Decimal128]
}

// Deprecated: NewDecimal1280F using NewDecimal128Field
func NewDecimal1280F(fName string) Decimal1280F {
	f := NewDecimal128Field(fName)
	return &decimal1280F{f, depMethodsImpl[bson.Decimal128]{f}}
}

// Deprecated: Float320F using Float32Field
type Float320F interface {
	ComputableField[float32]
	depMethods[float32]
}

// Deprecated:
type (
	Float320FUpdaterF = Float320F
	Float320FFilterF  = Float320F
)

type float320F struct {
	ComputableField[float32]
	depMethodsImpl[float32]
}

// Deprecated: NewFloat320F using NewFloat32Field
func NewFloat320F(fName string) Float320F {
	f := NewFloat32Field(fName)
	return &float320F{f, depMethodsImpl[float32]{f}}
}

// Deprecated: Int0F using IntField
type Int0F interface {
	IntegerField[int]
	depMethods[int]
}

// Deprecated:
type (
	Int0FUpdaterF = Int0F
	Int0FFilterF  = Int0F
)

type int0F struct {
	IntegerField[int]
	depMethodsImpl[int]
}

// Deprecated: NewInt0F using NewIntField
func NewInt0F(fName string) Int0F {
	f := NewIntField(fName)
	return &int0F{f, depMethodsImpl[int]{f}}
}

// Deprecated: use Int8Field instead.
type Int80F interface {
	IntegerField[int8]
	depMethods[int8]
}

// Deprecated: use Int80F directly.
type (
	Int80FUpdaterF = Int80F
	Int80FFilterF  = Int80F
)

type int80F struct {
	IntegerField[int8]
	depMethodsImpl[int8]
}

// Deprecated: use NewInt8Field instead.
func NewInt80F(fieldName string) Int80F {
	f := NewInt8Field(fieldName)
	return &int80F{f, depMethodsImpl[int8]{f}}
}

// Deprecated: use Int16Field instead.
type Int160F interface {
	IntegerField[int16]
	depMethods[int16]
}

// Deprecated: use Int16Field directly.
type (
	Int160FUpdaterF = Int160F
	Int160FFilterF  = Int160F
)

type int160F struct {
	IntegerField[int16]
	depMethodsImpl[int16]
}

// Deprecated: use NewInt16Field instead.
func NewInt160F(fieldName string) Int160F {
	f := NewInt16Field(fieldName)
	return &int160F{f, depMethodsImpl[int16]{f}}
}

// Deprecated: use Int32Field instead.
type Int320F interface {
	IntegerField[int32]
	depMethods[int32]
}

// Deprecated: use Int320F directly.
type (
	Int320FUpdaterF = Int320F
	Int320FFilterF  = Int320F
)

type int320F struct {
	IntegerField[int32]
	depMethodsImpl[int32]
}

// Deprecated: use NewInt32Field instead.
func NewInt320F(fieldName string) Int320F {
	f := NewInt32Field(fieldName)
	return &int320F{f, depMethodsImpl[int32]{f}}
}

// Deprecated: use Int64Field instead.
type Int640F interface {
	IntegerField[int64]
	depMethods[int64]
}

// Deprecated: use Int640F directly.
type (
	Int640FUpdaterF = Int640F
	Int640FFilterF  = Int640F
)

type int640F struct {
	IntegerField[int64]
	depMethodsImpl[int64]
}

// Deprecated: use NewInt64Field instead.
func NewInt640F(fieldName string) Int640F {
	f := NewInt64Field(fieldName)
	return &int640F{f, depMethodsImpl[int64]{f}}
}

// Deprecated: Uint0F using UintField
type Uint0F interface {
	UnIntegerField[uint, int]
	depMethods[uint]
}

// Deprecated:
type (
	Uint0FUpdaterF = Uint0F
	Uint0FFilterF  = Uint0F
)

type uint0F struct {
	UnIntegerField[uint, int]
	depMethodsImpl[uint]
}

// Deprecated: NewUint0F using NewUintField
func NewUint0F(fName string) Uint0F {
	f := NewUintField(fName)
	return &uint0F{f, depMethodsImpl[uint]{f}}
}

// Deprecated: use Uint8Field instead.
type Uint80F interface {
	UnIntegerField[uint8, int8]
	depMethods[uint8]
}

// Deprecated: use Uint80F directly.
type (
	Uint80FUpdaterF = Uint80F
	Uint80FFilterF  = Uint80F
)

type uint80F struct {
	UnIntegerField[uint8, int8]
	depMethodsImpl[uint8]
}

// Deprecated: use NewUint8Field instead.
func NewUint80F(fieldName string) Uint80F {
	f := NewUint8Field(fieldName)
	return &uint80F{f, depMethodsImpl[uint8]{f}}
}

// Deprecated: use Uint16Field instead.
type Uint160F interface {
	UnIntegerField[uint16, int16]
	depMethods[uint16]
}

// Deprecated: use Uint160F directly.
type (
	Uint160FUpdaterF = Uint160F
	Uint160FFilterF  = Uint160F
)

type uint160F struct {
	UnIntegerField[uint16, int16]
	depMethodsImpl[uint16]
}

// Deprecated: use NewUint16Field instead.
func NewUint160F(fieldName string) Uint160F {
	f := NewUint16Field(fieldName)
	return &uint160F{f, depMethodsImpl[uint16]{f}}
}

// Deprecated: use Uint32Field instead.
type Uint320F interface {
	UnIntegerField[uint32, int32]
	depMethods[uint32]
}

// Deprecated: use Uint320F directly.
type (
	Uint320FUpdaterF = Uint320F
	Uint320FFilterF  = Uint320F
)

type uint320F struct {
	UnIntegerField[uint32, int32]
	depMethodsImpl[uint32]
}

// Deprecated: use NewUint32Field instead.
func NewUint320F(fieldName string) Uint320F {
	f := NewUint32Field(fieldName)
	return &uint320F{f, depMethodsImpl[uint32]{f}}
}

// Deprecated: use Uint64Field instead.
type Uint640F interface {
	UnIntegerField[uint64, int64]
	depMethods[uint64]
}

// Deprecated: use Uint640F directly.
type (
	Uint640FUpdaterF = Uint640F
	Uint640FFilterF  = Uint640F
)

type uint640F struct {
	UnIntegerField[uint64, int64]
	depMethodsImpl[uint64]
}

// Deprecated: use NewUint640F instead.
func NewUint640F(fieldName string) Uint640F {
	f := NewUint64Field(fieldName)
	return &uint640F{f, depMethodsImpl[uint64]{f}}
}

// Deprecated: Float640F using Float64Field
type Float640F interface {
	ComputableField[float64]
	depMethods[float64]
}

// Deprecated:
type (
	Float640FUpdaterF = Float640F
	Float640FFilterF  = Float640F
)

type float640F struct {
	ComputableField[float64]
	depMethodsImpl[float64]
}

// Deprecated: NewFloat640F using NewFloat64Field
func NewFloat640F(fName string) Float640F {
	f := NewFloat64Field(fName)
	return &float640F{f, depMethodsImpl[float64]{f}}
}

// Deprecated: String0F using StringField
type String0F interface {
	StringField
	depMethods[string]
}

type string0F struct {
	StringField
	depMethodsImpl[string]
}

// Deprecated:
type (
	String0FUpdaterF = String0F
	String0FFilterF  = String0F
)

// Deprecated: NewString0F using NewStringField
func NewString0F(fName string) String0F {
	f := NewStringField(fName)
	return &string0F{f, depMethodsImpl[string]{f}}
}

// Deprecated: ObjectID0F using ObjectIDField
type ObjectID0F interface {
	ComparableField[bson.ObjectID]
	depMethods[bson.ObjectID]
}

// Deprecated:
type (
	ObjectID0FUpdaterF = ObjectID0F
	ObjectID0FFilterF  = ObjectID0F
)

type objectID0F struct {
	ComparableField[bson.ObjectID]
	depMethodsImpl[bson.ObjectID]
}

// Deprecated: NewObjectID0F using NewObjectIDField
func NewObjectID0F(fName string) ObjectID0F {
	f := NewObjectIDField(fName)
	return &objectID0F{f, depMethodsImpl[bson.ObjectID]{f}}
}
