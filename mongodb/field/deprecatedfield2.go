package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type uint0FUpdaterF interface {
	deprecatedBaseUpdater
	// Inc(num uint) is a bug.
	Inc(num int) updater.Updater
	Mul(num uint) updater.Updater
	Min(value uint) updater.Updater
	Max(value uint) updater.Updater
	Set(value uint) updater.Updater
	SetOnIns(value uint) updater.Updater
}

type uint0FFilterF interface {
	deprecatedBaseFilter
	Mod(divisor, remainder uint) filter.Filter
	Eq(value uint) filter.Filter
	Ne(value uint) filter.Filter
	NeField(f filter.ComparableFilterField[uint]) filter.Filter
	EqField(f filter.ComparableFilterField[uint]) filter.Filter
	Gte(value uint) filter.Filter
	Lte(value uint) filter.Filter
	GteField(f filter.ComparableFilterField[uint]) filter.Filter
	LteField(f filter.ComparableFilterField[uint]) filter.Filter
	Gt(value uint) filter.Filter
	Lt(value uint) filter.Filter
	GtField(f filter.BaseFilterField[uint]) filter.Filter
	LtField(f filter.BaseFilterField[uint]) filter.Filter
	In(values []uint) filter.Filter
	Nin(values []uint) filter.Filter
}

type uint0F interface {
	uint0FUpdaterF
	uint0FFilterF
	deprecatedBaseKey
}

// Deprecated: Uint0F using UintField
type Uint0F struct {
	UnIntegerField[uint, int]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Uint0F) Min(value uint) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Uint0F) Max(value uint) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Uint0F) SetOnIns(value uint) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewUint0F using NewUintField
func NewUint0F(fName string) *Uint0F {
	return &Uint0F{NewUintField(fName)}
}

// Deprecated:
type (
	Uint0FUpdaterF = Uint0F
	Uint0FFilterF  = Uint0F
)

var (
	_ uint0F         = &Uint0F{}
	_ uint0FUpdaterF = &Uint0F{}
	_ uint0FFilterF  = &Uint0F{}
)

type uint1Field struct {
	*depTypeArrayField[uint, UintField]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint1Field) Push(value uint) updater.Updater {
	return da.ArrayComparableField.Push([]uint{value})
}

// Deprecated: Uint1Field using ArrayComparableField[uint, UintField]
type Uint1Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Uint0F
	EleOne() Uint0F
	EleThat() Uint0FUpdaterF
	EleAll() Uint0FUpdaterF
	EleByFid(identifier string) Uint0FUpdaterF
	DeclFid(identifier string) Uint0FFilterF
	Include(a []uint) filter.Filter
	Eq(a []uint) filter.Filter
	Set(a []uint) updater.Updater
	AddToSet(value uint) updater.Updater
	AddToSetValues(a []uint) updater.Updater
	Pull(value uint) updater.Updater
	PullAll(a []uint) updater.Updater
	Push(value uint) updater.Updater
	PushByModifier(m updater.PushModifier, each []uint) updater.Updater
}

// Deprecated: NewInt1Field using ArrayComparableField[uint, UintField]
func NewUint1Field(fName string) Uint1Field {
	return &uint1Field{newDepCompAbleArrF[uint, UintField](fName, NewUintField)}
}

// Deprecated: use Uint8Field instead.
type Uint80F struct {
	UnIntegerField[uint8, int8]
}

// interfaces for compile-time check
type uint80F interface {
	uint80FUpdaterF
	uint80FFilterF
	deprecatedBaseKey
}

type uint80FUpdaterF interface {
	Inc(num int8) updater.Updater
	Mul(num uint8) updater.Updater
	Set(value uint8) updater.Updater
	Min(value uint8) updater.Updater
	Max(value uint8) updater.Updater
	SetOnIns(value uint8) updater.Updater
}

type uint80FFilterF interface {
	Eq(value uint8) filter.Filter
	Ne(value uint8) filter.Filter
	Gt(value uint8) filter.Filter
	Lt(value uint8) filter.Filter
	Gte(value uint8) filter.Filter
	Lte(value uint8) filter.Filter
	In(values []uint8) filter.Filter
	Nin(values []uint8) filter.Filter
	Mod(divisor, remainder uint8) filter.Filter
	EqField(f filter.ComparableFilterField[uint8]) filter.Filter
	NeField(f filter.ComparableFilterField[uint8]) filter.Filter
	GtField(f filter.BaseFilterField[uint8]) filter.Filter
	LtField(f filter.BaseFilterField[uint8]) filter.Filter
	GteField(f filter.ComparableFilterField[uint8]) filter.Filter
	LteField(f filter.ComparableFilterField[uint8]) filter.Filter
}

// compile-time interface checks
var _ uint80F = &Uint80F{}
var _ uint80FUpdaterF = &Uint80F{}
var _ uint80FFilterF = &Uint80F{}

// Deprecated: use NewUint8Field instead.
func NewUint80F(fieldName string) *Uint80F {
	return &Uint80F{NewUint8Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Uint80F) Min(value uint8) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Uint80F) Max(value uint8) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Uint80F) SetOnIns(value uint8) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Uint80F directly.
type Uint80FUpdaterF = Uint80F

// Deprecated: use Uint80F directly.
type Uint80FFilterF = Uint80F

// Deprecated: use Uint16Field instead.
type Uint160F struct {
	UnIntegerField[uint16, int16]
}

// interfaces for compile-time check
type uint160F interface {
	uint160FUpdaterF
	uint160FFilterF
	deprecatedBaseKey
}

type uint160FUpdaterF interface {
	Inc(num int16) updater.Updater
	Mul(num uint16) updater.Updater
	Set(value uint16) updater.Updater
	Min(value uint16) updater.Updater
	Max(value uint16) updater.Updater
	SetOnIns(value uint16) updater.Updater
}

type uint160FFilterF interface {
	Eq(value uint16) filter.Filter
	Ne(value uint16) filter.Filter
	Gt(value uint16) filter.Filter
	Lt(value uint16) filter.Filter
	Gte(value uint16) filter.Filter
	Lte(value uint16) filter.Filter
	In(values []uint16) filter.Filter
	Nin(values []uint16) filter.Filter
	Mod(divisor, remainder uint16) filter.Filter
	EqField(f filter.ComparableFilterField[uint16]) filter.Filter
	NeField(f filter.ComparableFilterField[uint16]) filter.Filter
	GtField(f filter.BaseFilterField[uint16]) filter.Filter
	LtField(f filter.BaseFilterField[uint16]) filter.Filter
	GteField(f filter.ComparableFilterField[uint16]) filter.Filter
	LteField(f filter.ComparableFilterField[uint16]) filter.Filter
}

// compile-time interface checks
var _ uint160F = &Uint160F{}
var _ uint160FUpdaterF = &Uint160F{}
var _ uint160FFilterF = &Uint160F{}

// Deprecated: use NewUint16Field instead.
func NewUint160F(fieldName string) *Uint160F {
	return &Uint160F{NewUint16Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Uint160F) Min(value uint16) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Uint160F) Max(value uint16) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Uint160F) SetOnIns(value uint16) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Uint160F directly.
type Uint160FUpdaterF = Uint160F

// Deprecated: use Uint160F directly.
type Uint160FFilterF = Uint160F

// Deprecated: use Uint32Field instead.
type Uint320F struct {
	UnIntegerField[uint32, int32]
}

// interfaces for compile-time check
type uint320F interface {
	uint320FUpdaterF
	uint320FFilterF
	deprecatedBaseKey
}

type uint320FUpdaterF interface {
	Inc(num int32) updater.Updater
	Mul(num uint32) updater.Updater
	Set(value uint32) updater.Updater
	Min(value uint32) updater.Updater
	Max(value uint32) updater.Updater
	SetOnIns(value uint32) updater.Updater
}

type uint320FFilterF interface {
	Eq(value uint32) filter.Filter
	Ne(value uint32) filter.Filter
	Gt(value uint32) filter.Filter
	Lt(value uint32) filter.Filter
	Gte(value uint32) filter.Filter
	Lte(value uint32) filter.Filter
	In(values []uint32) filter.Filter
	Nin(values []uint32) filter.Filter
	Mod(divisor, remainder uint32) filter.Filter
	EqField(f filter.ComparableFilterField[uint32]) filter.Filter
	NeField(f filter.ComparableFilterField[uint32]) filter.Filter
	GtField(f filter.BaseFilterField[uint32]) filter.Filter
	LtField(f filter.BaseFilterField[uint32]) filter.Filter
	GteField(f filter.ComparableFilterField[uint32]) filter.Filter
	LteField(f filter.ComparableFilterField[uint32]) filter.Filter
}

// compile-time interface checks
var _ uint320F = &Uint320F{}
var _ uint320FUpdaterF = &Uint320F{}
var _ uint320FFilterF = &Uint320F{}

// Deprecated: use NewUint32Field instead.
func NewUint320F(fieldName string) *Uint320F {
	return &Uint320F{NewUint32Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Uint320F) Min(value uint32) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Uint320F) Max(value uint32) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Uint320F) SetOnIns(value uint32) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Uint320F directly.
type Uint320FUpdaterF = Uint320F

// Deprecated: use Uint320F directly.
type Uint320FFilterF = Uint320F

// Deprecated: use Uint64Field instead.
type Uint640F struct {
	UnIntegerField[uint64, int64]
}

// interfaces for compile-time check
type uint640F interface {
	uint640FUpdaterF
	uint640FFilterF
	deprecatedBaseKey
}

type uint640FUpdaterF interface {
	Inc(num int64) updater.Updater
	Mul(num uint64) updater.Updater
	Set(value uint64) updater.Updater
	Min(value uint64) updater.Updater
	Max(value uint64) updater.Updater
	SetOnIns(value uint64) updater.Updater
}

type uint640FFilterF interface {
	Eq(value uint64) filter.Filter
	Ne(value uint64) filter.Filter
	Gt(value uint64) filter.Filter
	Lt(value uint64) filter.Filter
	Gte(value uint64) filter.Filter
	Lte(value uint64) filter.Filter
	In(values []uint64) filter.Filter
	Nin(values []uint64) filter.Filter
	Mod(divisor, remainder uint64) filter.Filter
	EqField(f filter.ComparableFilterField[uint64]) filter.Filter
	NeField(f filter.ComparableFilterField[uint64]) filter.Filter
	GtField(f filter.BaseFilterField[uint64]) filter.Filter
	LtField(f filter.BaseFilterField[uint64]) filter.Filter
	GteField(f filter.ComparableFilterField[uint64]) filter.Filter
	LteField(f filter.ComparableFilterField[uint64]) filter.Filter
}

// compile-time interface checks
var _ uint640F = &Uint640F{}
var _ uint640FUpdaterF = &Uint640F{}
var _ uint640FFilterF = &Uint640F{}

// Deprecated: use NewUint64Field instead.
func NewUint640F(fieldName string) *Uint640F {
	return &Uint640F{NewUint64Field(fieldName)}
}

// Deprecated: use SetMin instead.
func (i *Uint640F) Min(value uint64) updater.Updater {
	return i.SetMin(value)
}

// Deprecated: use SetMax instead.
func (i *Uint640F) Max(value uint64) updater.Updater {
	return i.SetMax(value)
}

// Deprecated: use SetOnInsert instead.
func (i *Uint640F) SetOnIns(value uint64) updater.Updater {
	return i.SetOnInsert(value)
}

// Deprecated: use Uint640F directly.
type Uint640FUpdaterF = Uint640F

// Deprecated: use Uint640F directly.
type Uint640FFilterF = Uint640F

type uint81Field struct {
	*depTypeArrayField[uint8, Uint8Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint81Field) Push(value uint8) updater.Updater {
	return da.ArrayComparableField.Push([]uint8{value})
}

// Deprecated: Uint81Field using ArrayComparableField[uint8, Uint8Field]
type Uint81Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Uint80F
	EleOne() Uint80F
	EleThat() Uint80FUpdaterF
	EleAll() Uint80FUpdaterF
	EleByFid(identifier string) Uint80FUpdaterF
	DeclFid(identifier string) Uint80FFilterF
	Include(a []uint8) filter.Filter
	Eq(a []uint8) filter.Filter
	Set(a []uint8) updater.Updater
	AddToSet(value uint8) updater.Updater
	AddToSetValues(a []uint8) updater.Updater
	Pull(value uint8) updater.Updater
	PullAll(a []uint8) updater.Updater
	Push(value uint8) updater.Updater
	PushByModifier(m updater.PushModifier, each []uint8) updater.Updater
}

// Deprecated: NewUint81Field using ArrayComparableField[uint8, Uint8Field]
func NewUint81Field(fName string) Uint81Field {
	return &uint81Field{newDepCompAbleArrF[uint8, Uint8Field](fName, NewUint8Field)}
}

type uint161Field struct {
	*depTypeArrayField[uint16, Uint16Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint161Field) Push(value uint16) updater.Updater {
	return da.ArrayComparableField.Push([]uint16{value})
}

// Deprecated: Uint161Field using ArrayComparableField[uint16, Uint16Field]
type Uint161Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Uint160F
	EleOne() Uint160F
	EleThat() Uint160FUpdaterF
	EleAll() Uint160FUpdaterF
	EleByFid(identifier string) Uint160FUpdaterF
	DeclFid(identifier string) Uint160FFilterF
	Include(a []uint16) filter.Filter
	Eq(a []uint16) filter.Filter
	Set(a []uint16) updater.Updater
	AddToSet(value uint16) updater.Updater
	AddToSetValues(a []uint16) updater.Updater
	Pull(value uint16) updater.Updater
	PullAll(a []uint16) updater.Updater
	Push(value uint16) updater.Updater
	PushByModifier(m updater.PushModifier, each []uint16) updater.Updater
}

// Deprecated: NewUint161Field using ArrayComparableField[uint16, Uint16Field]
func NewUint161Field(fName string) Uint161Field {
	return &uint161Field{newDepCompAbleArrF[uint16, Uint16Field](fName, NewUint16Field)}
}

type uint321Field struct {
	*depTypeArrayField[uint32, Uint32Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint321Field) Push(value uint32) updater.Updater {
	return da.ArrayComparableField.Push([]uint32{value})
}

// Deprecated: Uint321Field using ArrayComparableField[uint32, Uint32Field]
type Uint321Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Uint320F
	EleOne() Uint320F
	EleThat() Uint320FUpdaterF
	EleAll() Uint320FUpdaterF
	EleByFid(identifier string) Uint320FUpdaterF
	DeclFid(identifier string) Uint320FFilterF
	Include(a []uint32) filter.Filter
	Eq(a []uint32) filter.Filter
	Set(a []uint32) updater.Updater
	AddToSet(value uint32) updater.Updater
	AddToSetValues(a []uint32) updater.Updater
	Pull(value uint32) updater.Updater
	PullAll(a []uint32) updater.Updater
	Push(value uint32) updater.Updater
	PushByModifier(m updater.PushModifier, each []uint32) updater.Updater
}

// Deprecated: NewUint321Field using ArrayComparableField[uint32, Uint32Field]
func NewUint321Field(fName string) Uint321Field {
	return &uint321Field{newDepCompAbleArrF[uint32, Uint32Field](fName, NewUint32Field)}
}

type uint641Field struct {
	*depTypeArrayField[uint64, Uint64Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *uint641Field) Push(value uint64) updater.Updater {
	return da.ArrayComparableField.Push([]uint64{value})
}

// Deprecated: Uint641Field using ArrayComparableField[uint64, Uint64Field]
type Uint641Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Uint640F
	EleOne() Uint640F
	EleThat() Uint640FUpdaterF
	EleAll() Uint640FUpdaterF
	EleByFid(identifier string) Uint640FUpdaterF
	DeclFid(identifier string) Uint640FFilterF
	Include(a []uint64) filter.Filter
	Eq(a []uint64) filter.Filter
	Set(a []uint64) updater.Updater
	AddToSet(value uint64) updater.Updater
	AddToSetValues(a []uint64) updater.Updater
	Pull(value uint64) updater.Updater
	PullAll(a []uint64) updater.Updater
	Push(value uint64) updater.Updater
	PushByModifier(m updater.PushModifier, each []uint64) updater.Updater
}

// Deprecated: NewUint641Field using ArrayComparableField[uint64, Uint64Field]
func NewUint641Field(fName string) Uint641Field {
	return &uint641Field{newDepCompAbleArrF[uint64, Uint64Field](fName, NewUint64Field)}
}

type float640FUpdaterF interface {
	deprecatedBaseUpdater
	Min(value float64) updater.Updater
	Max(value float64) updater.Updater
	Set(value float64) updater.Updater
	SetOnIns(value float64) updater.Updater
}

type float640FFilterF interface {
	deprecatedBaseFilter
	Gt(value float64) filter.Filter
	Lt(value float64) filter.Filter
	GtField(f filter.BaseFilterField[float64]) filter.Filter
	LtField(f filter.BaseFilterField[float64]) filter.Filter
	In(values []float64) filter.Filter
	Nin(values []float64) filter.Filter
}

type float640F interface {
	float640FUpdaterF
	float640FFilterF
	deprecatedBaseKey
}

// Deprecated: Float640F using Float64Field
type Float640F struct {
	ComputableField[float64]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *Float640F) Min(value float64) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *Float640F) Max(value float64) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *Float640F) SetOnIns(value float64) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: this is a bug
func (f *Float640F) In(values []float64) filter.Filter {
	return filter.New(f, "$in", values)
}

// Deprecated: this is a bug
func (f *Float640F) Nin(values []float64) filter.Filter {
	return filter.New(f, "$nin", values)
}

// Deprecated: NewFloat640F using NewFloat64Field
func NewFloat640F(fName string) *Float640F {
	return &Float640F{NewFloat64Field(fName)}
}

// Deprecated:
type (
	Float640FUpdaterF = Float640F
	Float640FFilterF  = Float640F
)

var (
	_ float640F         = &Float640F{}
	_ float640FUpdaterF = &Float640F{}
	_ float640FFilterF  = &Float640F{}
)

type float641Field struct {
	*depTypeArrayField[float64, Float64Field]
}

// Deprecated: Push using ArrayField[].Push
func (da *float641Field) Push(value float64) updater.Updater {
	return da.ArrayComparableField.Push([]float64{value})
}

// Deprecated: Float641Field using ArrayComparableField[float64, Float64Field]
type Float641Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) Float640F
	EleOne() Float640F
	EleThat() Float640FUpdaterF
	EleAll() Float640FUpdaterF
	EleByFid(identifier string) Float640FUpdaterF
	DeclFid(identifier string) Float640FFilterF
	Include(a []float64) filter.Filter
	Eq(a []float64) filter.Filter
	Set(a []float64) updater.Updater
	AddToSet(value float64) updater.Updater
	AddToSetValues(a []float64) updater.Updater
	Pull(value float64) updater.Updater
	PullAll(a []float64) updater.Updater
	Push(value float64) updater.Updater
	PushByModifier(m updater.PushModifier, each []float64) updater.Updater
}

// Deprecated: NewFloat641Field using ArrayComparableField[float64, Float64Field]
func NewFloat641Field(fName string) Float641Field {
	return &float641Field{newDepCompAbleArrF[float64, Float64Field](fName, NewFloat64Field)}
}

type string0FUpdaterF interface {
	deprecatedBaseUpdater
	Min(value string) updater.Updater
	Max(value string) updater.Updater
	Set(value string) updater.Updater
	SetOnIns(value string) updater.Updater
	Regex(regex bson.Regex) filter.Filter
}

type string0FFilterF interface {
	deprecatedBaseFilter
	Eq(value string) filter.Filter
	Ne(value string) filter.Filter
	Gt(value string) filter.Filter
	Lt(value string) filter.Filter
	Gte(value string) filter.Filter
	Lte(value string) filter.Filter
	In(values []string) filter.Filter
	Nin(values []string) filter.Filter
	EqField(f filter.ComparableFilterField[string]) filter.Filter
	NeField(f filter.ComparableFilterField[string]) filter.Filter
	GtField(f filter.BaseFilterField[string]) filter.Filter
	LtField(f filter.BaseFilterField[string]) filter.Filter
	GteField(f filter.ComparableFilterField[string]) filter.Filter
	LteField(f filter.ComparableFilterField[string]) filter.Filter
}

type string0F interface {
	string0FUpdaterF
	string0FFilterF
	deprecatedBaseKey
}

// Deprecated: String0F using StringField
type String0F struct {
	StringField
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *String0F) Min(value string) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *String0F) Max(value string) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *String0F) SetOnIns(value string) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewString0F using NewStringField
func NewString0F(fName string) *String0F {
	return &String0F{NewStringField(fName)}
}

// Deprecated:
type (
	String0FUpdaterF = String0F
	String0FFilterF  = String0F
)

var (
	_ string0F         = &String0F{}
	_ string0FUpdaterF = &String0F{}
	_ string0FFilterF  = &String0F{}
)

type string1Field struct {
	*depTypeArrayField[string, StringField]
}

// Deprecated: Push using ArrayField[].Push
func (da *string1Field) Push(value string) updater.Updater {
	return da.ArrayComparableField.Push([]string{value})
}

// Deprecated: String1Field using ArrayComparableField[string, StringField]
type String1Field interface {
	deprecatedBaseFilter
	deprecatedBaseUpdater
	deprecatedBaseKey
	EleAt(index int) String0F
	EleOne() String0F
	EleThat() String0FUpdaterF
	EleAll() String0FUpdaterF
	EleByFid(identifier string) String0FUpdaterF
	DeclFid(identifier string) String0FFilterF
	Include(a []string) filter.Filter
	Eq(a []string) filter.Filter
	Set(a []string) updater.Updater
	AddToSet(value string) updater.Updater
	AddToSetValues(a []string) updater.Updater
	Pull(value string) updater.Updater
	PullAll(a []string) updater.Updater
	Push(value string) updater.Updater
	PushByModifier(m updater.PushModifier, each []string) updater.Updater
}

// Deprecated: NewString1Field using ArrayComparableField[string, StringField]
func NewString1Field(fName string) String1Field {
	return &string1Field{newDepCompAbleArrF[string, StringField](fName, NewStringField)}
}

type objectID0FUpdaterF interface {
	deprecatedBaseUpdater
	Min(value bson.ObjectID) updater.Updater
	Max(value bson.ObjectID) updater.Updater
	Set(value bson.ObjectID) updater.Updater
	SetOnIns(value bson.ObjectID) updater.Updater
}

type objectID0FFilterF interface {
	deprecatedBaseFilter
	Eq(value bson.ObjectID) filter.Filter
	Ne(value bson.ObjectID) filter.Filter
	Gt(value bson.ObjectID) filter.Filter
	Lt(value bson.ObjectID) filter.Filter
	Gte(value bson.ObjectID) filter.Filter
	Lte(value bson.ObjectID) filter.Filter
	In(values []bson.ObjectID) filter.Filter
	Nin(values []bson.ObjectID) filter.Filter
	EqField(f filter.ComparableFilterField[bson.ObjectID]) filter.Filter
	NeField(f filter.ComparableFilterField[bson.ObjectID]) filter.Filter
	GtField(f filter.BaseFilterField[bson.ObjectID]) filter.Filter
	LtField(f filter.BaseFilterField[bson.ObjectID]) filter.Filter
	GteField(f filter.ComparableFilterField[bson.ObjectID]) filter.Filter
	LteField(f filter.ComparableFilterField[bson.ObjectID]) filter.Filter
}

type objectID0F interface {
	objectID0FUpdaterF
	objectID0FFilterF
	deprecatedBaseKey
}

// Deprecated: ObjectID0F using ObjectIDField
type ObjectID0F struct {
	ComparableField[bson.ObjectID]
}

// Deprecated: Min using: updater.BaseUpdater[].SetMin
func (f *ObjectID0F) Min(value bson.ObjectID) updater.Updater {
	return f.SetMin(value)
}

// Deprecated: Max using: updater.BaseUpdater[].SetMax
func (f *ObjectID0F) Max(value bson.ObjectID) updater.Updater {
	return f.SetMax(value)
}

// Deprecated: SetOnIns using: updater.BaseUpdater[].SetOnInsert
func (f *ObjectID0F) SetOnIns(value bson.ObjectID) updater.Updater {
	return f.SetOnInsert(value)
}

// Deprecated: NewObjectID0F using NewObjectIDField
func NewObjectID0F(fName string) *ObjectID0F {
	return &ObjectID0F{NewObjectIDField(fName)}
}

// Deprecated:
type (
	ObjectID0FUpdaterF = ObjectID0F
	ObjectID0FFilterF  = ObjectID0F
)

var (
	_ objectID0F         = &ObjectID0F{}
	_ objectID0FUpdaterF = &ObjectID0F{}
	_ objectID0FFilterF  = &ObjectID0F{}
)
