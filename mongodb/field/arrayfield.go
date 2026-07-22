package field

import (
	"fmt"
	"github.com/xpwu/go-db-mongo/mongodb"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ArrayBaseFilter[T any, ElemField mongodb.Field] interface {
	filter.BaseFilter[[]T]
	Size(sz int) filter.Filter
	// AnyElemMeet Different elements can satisfy different conditions, or a single element can satisfy all conditions.
	// NOTE THAT: If using a negation operator, such as $ne, $not, or $nin, 'anyElem' is 'AllElem'.
	// In other words, none of elem meets the positive operator (the counterpart of the negation operator).
	AnyElemMeet(f func(anyElem *ElemField) filter.Filter) filter.Filter
	// SameElemMeet Must be the same element satisfying all conditions.
	// NOTE THAT: If using a negation operator, such as $ne, $not, or $nin, 'theOne' is 'someOne'.
	// So that means if some element of the array meets the negation operator, the document is selected.
	// In other words, not all elem meets the positive operator (the counterpart of the negation operator).
	SameElemMeet(f func(theOne *ElemField) filter.Filter) filter.Filter
	// PosElemMeet Element at a fixed index must satisfy all conditions
	PosElemMeet(pos int, f func(atPosElem *ElemField) filter.Filter) filter.Filter
}

type ArrayComparableFilter[T comparable, ElemField mongodb.Field] interface {
	ArrayBaseFilter[T, ElemField]

	IncludeAll(values []T) filter.Filter
}

type ArrayBaseUpdater[T any, ElemField mongodb.Field] interface {
	updater.BaseUpdater[[]T]

	PopFirst() updater.Updater
	PopLast() updater.Updater
	// $addToSet
	AddIfNotExist(value T) updater.Updater
	AddEachIfNotExist(values []T) updater.Updater
	RemoveBy(func(elem *ElemField) filter.Filter) updater.Updater
	Push(values []T, f func(elem *ElemField) updater.PushModifier) updater.Updater
}

type ArrayComparableUpdater[T comparable, ElemField mongodb.Field] interface {
	ArrayBaseUpdater[T, ElemField]

	// RemoveAll: remove all db value in the array, where the values meet the condition
	RemoveAll(value []T) updater.Updater
	Remove(value T) updater.Updater
}

type arrayBaseField[T any, ElemField mongodb.Field] struct {
	baseField[[]T]
	newElemField func(name string) *ElemField
}

func (a *arrayBaseField[T, ElemField]) Size(sz int) filter.Filter {
	return filter.New(a, `$size`, sz)
}

func (a *arrayBaseField[T, ElemField]) PopFirst() updater.Updater {
	return updater.New(a, `$pop`, -1)
}

func (a *arrayBaseField[T, ElemField]) PopLast() updater.Updater {
	return updater.New(a, `$pop`, 1)
}

func (a *arrayBaseField[T, ElemField]) AddIfNotExist(value T) updater.Updater {
	return updater.New(a, "$addToSet", value)
}

func (a *arrayBaseField[T, ElemField]) AddEachIfNotExist(values []T) updater.Updater {
	return updater.New(a, "$addToSet", bson.M{"$each": values})
}

// Remove: remove all db value in the array, where the values meet the condition

func (a *arrayBaseField[T, ElemField]) RemoveAll(value []T) updater.Updater {
	return updater.New(a, "$pullAll", value)
}

func (a *arrayBaseField[T, ElemField]) Remove(value T) updater.Updater {
	return updater.New(a, "$pull", value)
}

func (a *arrayBaseField[T, ElemField]) RemoveBy(f func(elem *ElemField) filter.Filter) updater.Updater {
	fil := f(a.newElemField(""))
	return updater.PullByFilter(a, fil)
}

func (a *arrayBaseField[T, ElemField]) Push(values []T,
	f func(elem *ElemField) updater.PushModifier) updater.Updater {

	return updater.PushByModifier(a, f(a.newElemField("")), values)
}

func (a *arrayBaseField[T, ElemField]) IncludeAll(values []T) filter.Filter {
	return filter.New(a, "$all", values)
}

// AnyElemMeet Different elements can satisfy different conditions, or a single element can satisfy all conditions.
// NOTE THAT: If using a negation operator, such as $ne, $not, or $nin, 'anyElem' is 'AllElem'.
// In other words, none of elem meets the positive operator (the counterpart of the negation operator).
func (a *arrayBaseField[T, ElemField]) AnyElemMeet(f func(anyElem *ElemField) filter.Filter) filter.Filter {
	return f(a.newElemField(a.FullName()))
}

// SameElemMeet Must be the same element satisfying all conditions.
// NOTE THAT: If using a negation operator, such as $ne, $not, or $nin, 'theOne' is 'someOne'.
// So that means if some element of the array meets the negation operator, the document is selected.
// In other words, not all elem meets the positive operator (the counterpart of the negation operator).
func (a *arrayBaseField[T, ElemField]) SameElemMeet(f func(theOne *ElemField) filter.Filter) filter.Filter {
	fil := f(a.newElemField(""))
	return filter.SameElemMatch(a, fil)
}

// PosElemMeet Element at a fixed index must satisfy all conditions
func (a *arrayBaseField[T, ElemField]) PosElemMeet(pos int, f func(atPosElem *ElemField) filter.Filter) filter.Filter {
	return f(a.newElemField(fmt.Sprintf("%s.%d", a.FullName(), pos)))
}

/**
查找

一、作为普通域
1、 整个数组严格相等或者某种大小关系   -----  这是 BaseFilter 或者  ComparableFilter 的能力

二、元素级别
1、元素包括这些值 [a, b]  ---- $all  where  element==
{tags: { $all: ['red', 'blank'] }}
contains both "red" and "blank" regardless of order or other elements in the array, use the $all operator

2、任意多个元素满足一个或多个条件
dim_cm: { $gt: 15, $lt: 20 }
tags: 'red'
dim_cm: { $gt: 25 }
任意一个元素 == value 或者其他田间。  如果是多个条件，数组中不同的元素满足不同的条件，但是合起来整个数组满足所有条件就行，不要求是
同一个元素满足所有条件
One element can satisfy the greater than 15 condition
	and another element can satisfy the less than 20 condition,
or a single element can satisfy both:

3、任意同一个元素满足所有条件  $elemMatch
{
  dim_cm: { $elemMatch: { $gt: 22, $lt: 30 } }
}
on array elements so that at least one array element satisfies all the specified criteria

4、指定位置的特定元素满足条件
'dim_cm.1': { $gt: 25 }

5、{
       qty: { $all: [
					{ "$elemMatch" : { size: "M", num: { $gt: 50} } },
           { "$elemMatch" : { num : 100, color: "green" } }
        ] }
   }

		====>

	qty: [
      { size: "S", num: 10, color: "blue" },
      { size: "M", num: 100, color: "blue" },
      { size: "L", num: 100, color: "green" }
   ]

	qty: [
      { size: "M", num: 100, color: "green" }
   ]


*/

/**
更新：

一、普通域

二、元素级
$  就更新前面匹配第一个的那个值为新的值

$[]  所有的值  update every element of an array

$[<identifier>]
*/

// extract
