
package field

// ---- auto generated by builder.go, NOT modify this file ----

import (
  fmt "fmt"
  filter "github.com/xpwu/go-db-mongo/mongodb/filter"
  updater "github.com/xpwu/go-db-mongo/mongodb/updater"
  bson "go.mongodb.org/mongo-driver/bson"
)

type Uint1Field struct {
  *Array
}

func NewUint1Field(fName string) *Uint1Field {
  return &Uint1Field { NewArray(fName)}
}

func (i *Uint1Field) EleAt(index int) *Uint0F {
  return NewUint0F(fmt.Sprintf("%s.%d", i.FullName(), index))
}

// 数组的某一个元素，有时也可以理解为 数组的任何一个元素
func (i *Uint1Field) EleOne() *Uint0F {
  return NewUint0F(i.FullName())
}

// update 操作中被filter匹配的那第一个元素
func (i *Uint1Field) EleThat() *Uint0FUpdaterF {
  return NewUint0F(i.FullName() + ".$").Uint0FUpdaterF
}

func (i *Uint1Field) EleAll() *Uint0FUpdaterF {
  return NewUint0F(i.FullName() + ".$[]").Uint0FUpdaterF
}

func (i *Uint1Field) EleByFid(identifier string) *Uint0FUpdaterF {
  return NewUint0F(fmt.Sprintf("%s.$[%s]", i.FullName(), i.FullName()+identifier)).Uint0FUpdaterF
}

func (i *Uint1Field) DeclFid(identifier string) *Uint0FFilterF {
  return NewUint0F(i.FullName()+identifier).Uint0FFilterF
}

func (i *Uint1Field) Include(a []uint) filter.Filter {
  return filter.New(i, "$all", a)
}

func (i *Uint1Field) Eq(a []uint) filter.Filter {
  return filter.CompareByValue(i, filter.EQ, a)
}

func (i *Uint1Field) Set(a []uint) updater.Updater {
  return updater.New(i, "$set", a)
}

func (i *Uint1Field) AddToSet(value uint) updater.Updater {
  return updater.New(i, "$addToSet", value)
}

func (i *Uint1Field) AddToSetValues(a []uint) updater.Updater {
  return updater.New(i, "$addToSet", bson.M{"$each":a})
}

func (i *Uint1Field) Pull(value uint) updater.Updater {
  return updater.New(i, "$pull", value)
}

func (i *Uint1Field) PullAll(a []uint) updater.Updater {
  return updater.New(i, "$pullAll", a)
}

func (i *Uint1Field) Push (value uint) updater.Updater {
  return updater.New(i, "$push", value)
}

func (i *Uint1Field) PushByModifier(m updater.PushModifier, each []uint) updater.Updater {
  return updater.PushByModifier(i, m, each)
}
