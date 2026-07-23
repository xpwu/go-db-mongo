package field

import (
	"fmt"
	"github.com/xpwu/go-db-mongo/mongodb"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ArrayField[T any, ElemField mongodb.Field] interface {
	mongodb.Field

	// At Element field at index pos.
	// https://www.mongodb.com/docs/manual/core/document/#arrays
	// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-for-an-element-by-the-array-index-position
	// https://www.mongodb.com/docs/manual/reference/operator/update/set/#set-elements-in-arrays
	At(pos int) *ElemField
}

// VirValue VirValue is a virtual value described by filters.
//
//	eg: { "$elemMatch" : { size: "M", num: { $gt: 50} }
//	    { votes: { $gte: 6 } }
//	    { results: { score: 8 , item: "B" } }
//	    { answers: { $elemMatch: { q: 2, a: { $gte: 8 } } } }
type VirValue filter.Filter

type ArrayBaseFilter[T any, ElemField mongodb.Field] interface {
	filter.BaseFilter[[]T]
	Size(sz int) filter.Filter

	// AnyElemMeet Different elements can satisfy different conditions,
	// or a single element can satisfy all conditions.
	//
	// NOTE THAT: If using a Negation operator, such as $ne, $not, or $nin, 'anyElem' is 'AllElem'.
	// In other words, none of elem meets the positive operator (the counterpart of the negation operator).
	//
	//	op: { dim_cm: { $gt: 15, $lt: 20 } }
	//
	// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-an-array-with-compound-filter-conditions-on-the-array-elements
	AnyElemMeet(f func(anyElem *ElemField) filter.Filter) filter.Filter

	// SameElemMeet Must be the same element satisfying all filters.
	//
	// NOTE THAT: If using a Negation operator, such as $ne, $not, or $nin, 'theOne' is 'someOne'.
	// So that means if some element of the array meets the Negation operator, the document is selected.
	// In other words, not all elem meets the positive operator (the counterpart of the Negation operator).
	//
	//	op: { dim_cm: { $elemMatch: { $gt: 22, $lt: 30 } } }
	//
	// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-for-an-array-element-that-meets-multiple-criteria
	SameElemMeet(f func(theOne *ElemField) filter.Filter) filter.Filter

	// PosElemMeet Element at a fixed index must satisfy all filters.
	//
	//	op: { 'dim_cm.1': { $gt: 25 } }
	//
	// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-for-an-element-by-the-array-index-position
	//PosElemMeet(pos int, f func(atPosElem *ElemField) filter.Filter) filter.Filter

	// CoverVirValues The array must cover the given VirValues — either by a single element covering all,
	// or by multiple elements covering them collectively (order doesn't matter).
	//
	//		op: { tags {
	//	         $all: [
	//	            { "$elemMatch" : { size: "M", num: { $gt: 50} } },
	//	            { "$elemMatch" : { num : 100, color: "green" } }
	//	         ]}
	//	     }
	//
	// Both document
	//
	//	{ tags: [
	//	   { size: "M", num: 100, color: "green" }
	//	] }
	//
	// and document
	//
	//	{ tags: [
	//	   { size: "S", num: 10, color: "blue" },
	//	   { size: "M", num: 100, color: "blue" },
	//	   { size: "L", num: 100, color: "green" }
	//	] }
	//
	// can cover these VirValues
	//
	//	[
	//	  { "$elemMatch" : { size: "M", num: { $gt: 50} } },
	//	  { "$elemMatch" : { num : 100, color: "green" } }
	//	]
	//
	// https://www.mongodb.com/docs/manual/reference/operator/query/all/#use--all-with--elemmatch
	CoverVirValues(f func(sameElem *ElemField) []VirValue) filter.Filter
}

type ArrayComparableFilter[T comparable, ElemField mongodb.Field] interface {
	ArrayBaseFilter[T, ElemField]

	// CoverValues The array must cover the given Values (order doesn't matter).
	//
	// op: {tags: { $all: ['red', 'blank'] }}
	//
	// https://www.mongodb.com/docs/manual/reference/operator/query/all/#use--all-to-match-values
	CoverValues(values []T) filter.Filter
}

type ArrayBaseUpdater[T any, ElemField mongodb.Field] interface {
	updater.BaseUpdater[[]T]

	PopFirst() updater.Updater
	PopLast() updater.Updater

	// AddEach Adds each value of values to an array unless the value is already present,
	// in which case the value isn't added to that array.
	//
	// op: { $addToSet: { tags: { $each: [ "camera", "electronics", "accessories" ] } }
	//
	// document:
	//
	//	{ _id: 2, item: "cable",
	//	  tags: [ "electronics", "supplies" ]
	//	}
	//
	// only adds "camera" and "accessories" to the tags array. "electronics" was already in the array.
	//
	// result:
	//
	//	{
	//	  _id: 2,
	//	  item: "cable",
	//	  tags: [ "electronics", "supplies", "camera", "accessories" ]
	//	}
	//
	// https://www.mongodb.com/docs/manual/reference/operator/update/addToSet/#examples
	AddEach(values []T) updater.Updater

	// RemoveVirValue Removes all instances that match the specified VirValue from an existing array.
	//
	// op:
	//
	//		 { $pull: { votes: { $gte: 6 } } }
	//	  { $pull: { results: { score: 8 , item: "B" } } }
	//	  { $pull: { results: { answers: { $elemMatch: { q: 2, a: { $gte: 8 } } } } } }
	//
	// https://www.mongodb.com/docs/manual/reference/operator/update/pull/#examples
	RemoveVirValue(func(elem *ElemField) VirValue) updater.Updater

	// Push Appends multiple values to the array field
	//
	// op: { $push: { genres: { $each: [ "Modern Classic", "Award-Winning" ] } } }
	// https://www.mongodb.com/docs/manual/reference/operator/update/push/#append-multiple-values-to-an-array
	Push(values []T) updater.Updater

	// PushWith Appends multiple values to the array field with the PushModifier
	//
	//   op: { $push: {
	//         quizzes: {
	//           $each: [ { wk: 5, score: 8 }, { wk: 6, score: 7 }, { wk: 7, score: 6 } ],
	//           $sort: { score: -1 },
	//           $slice: 3
	//         }
	//       } }
	//
	// Operation with modifiers occur in the following order, regardless of the order in which the modifiers appear:
	//
	// 1. Update array to add elements in the correct position.
	//
	// 2. Apply sort, if specified.
	//
	// 3. Slice the array, if specified.
	//
	// 4. Store the array.
	//
	// https://www.mongodb.com/docs/manual/reference/operator/update/push/#use--push-operator-with-multiple-modifiers
	PushWith(values []T, f func(elem *ElemField) updater.PushModifier) updater.Updater
}

type ArrayComparableUpdater[T comparable, ElemField mongodb.Field] interface {
	ArrayBaseUpdater[T, ElemField]

	// RemoveValues Removes all instances of the specified values from an existing array.
	//
	// op: { $pullAll: { scores: [ 0, 5 ] } }
	//
	// https://www.mongodb.com/docs/manual/reference/operator/update/pullAll/#examples
	RemoveValues(values []T) updater.Updater
}

type arrayBaseField[T any, ElemField mongodb.Field] struct {
	baseField[[]T]
	newElemField func(name string) *ElemField
}

func (a *arrayBaseField[T, ElemField]) At(pos int) *ElemField {
	return a.newElemField(fmt.Sprintf("%s.%d", a.FullName(), pos))
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

// AddEach Adds each value of values to an array unless the value is already present,
// in which case the value isn't added to that array.
//
// op: { $addToSet: { tags: { $each: [ "camera", "electronics", "accessories" ] } }
//
// document:
//
//	{ _id: 2, item: "cable",
//	  tags: [ "electronics", "supplies" ]
//	}
//
// only adds "camera" and "accessories" to the tags array. "electronics" was already in the array.
//
// result:
//
//	{
//	  _id: 2,
//	  item: "cable",
//	  tags: [ "electronics", "supplies", "camera", "accessories" ]
//	}
//
// https://www.mongodb.com/docs/manual/reference/operator/update/addToSet/#examples
func (a *arrayBaseField[T, ElemField]) AddEach(values []T) updater.Updater {
	return updater.New(a, "$addToSet", bson.M{"$each": values})
}

// RemoveValues Removes all instances of the specified values from an existing array.
//
// op: { $pullAll: { scores: [ 0, 5 ] } }
//
// https://www.mongodb.com/docs/manual/reference/operator/update/pullAll/#examples
func (a *arrayBaseField[T, ElemField]) RemoveValues(values []T) updater.Updater {
	return updater.New(a, "$pullAll", values)
}

// RemoveVirValue Removes all instances that match the specified VirValue from an existing array.
//
// op:
//
//		 { $pull: { votes: { $gte: 6 } } }
//	  { $pull: { results: { score: 8 , item: "B" } } }
//	  { $pull: { results: { answers: { $elemMatch: { q: 2, a: { $gte: 8 } } } } } }
//
// https://www.mongodb.com/docs/manual/reference/operator/update/pull/#examples
func (a *arrayBaseField[T, ElemField]) RemoveVirValue(f func(sameElem *ElemField) VirValue) updater.Updater {
	fil := f(a.newElemField(""))
	return updater.PullByFilter(a, fil)
}

// Push Appends multiple values to the array field
//
// op: { $push: { genres: { $each: [ "Modern Classic", "Award-Winning" ] } } }
// https://www.mongodb.com/docs/manual/reference/operator/update/push/#append-multiple-values-to-an-array
func (a *arrayBaseField[T, ElemField]) Push(values []T) updater.Updater {
	return updater.New(a, "$push", values)
}

// PushWith Appends multiple values to the array field with the PushModifier
//
//   op: { $push: {
//         quizzes: {
//           $each: [ { wk: 5, score: 8 }, { wk: 6, score: 7 }, { wk: 7, score: 6 } ],
//           $sort: { score: -1 },
//           $slice: 3
//         }
//       } }
//
// Operation with modifiers occur in the following order, regardless of the order in which the modifiers appear:
//
// 1. Update array to add elements in the correct position.
//
// 2. Apply sort, if specified.
//
// 3. Slice the array, if specified.
//
// 4. Store the array.
//
// https://www.mongodb.com/docs/manual/reference/operator/update/push/#use--push-operator-with-multiple-modifiers
func (a *arrayBaseField[T, ElemField]) PushWith(values []T,
	f func(elem *ElemField) updater.PushModifier) updater.Updater {

	return updater.PushByModifier(a, f(a.newElemField("")), values)
}

// AnyElemMeet Different elements can satisfy different conditions,
// or a single element can satisfy all conditions.
//
// NOTE THAT: If using a Negation operator, such as $ne, $not, or $nin, 'anyElem' is 'AllElem'.
// In other words, none of elem meets the positive operator (the counterpart of the negation operator).
//
//	op: { dim_cm: { $gt: 15, $lt: 20 } }
//
// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-an-array-with-compound-filter-conditions-on-the-array-elements
func (a *arrayBaseField[T, ElemField]) AnyElemMeet(f func(anyElem *ElemField) filter.Filter) filter.Filter {
	return f(a.newElemField(a.FullName()))
}

// SameElemMeet Must be the same element satisfying all filters.
//
// NOTE THAT: If using a Negation operator, such as $ne, $not, or $nin, 'theOne' is 'someOne'.
// So that means if some element of the array meets the Negation operator, the document is selected.
// In other words, not all elem meets the positive operator (the counterpart of the Negation operator).
//
//	op: { dim_cm: { $elemMatch: { $gt: 22, $lt: 30 } } }
//
// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-for-an-array-element-that-meets-multiple-criteria
func (a *arrayBaseField[T, ElemField]) SameElemMeet(f func(theOne *ElemField) filter.Filter) filter.Filter {
	fil := f(a.newElemField(""))
	return filter.SameElemMatch(a, fil)
}

// PosElemMeet Element at a fixed index must satisfy all filters.
//
//	op: { 'dim_cm.1': { $gt: 25 } }
//
// https://www.mongodb.com/docs/manual/tutorial/query-arrays/#query-for-an-element-by-the-array-index-position
//func (a *arrayBaseField[T, ElemField]) PosElemMeet(pos int,
//	f func(atPosElem *ElemField) filter.Filter) filter.Filter {
//
//	return f(a.newElemField(fmt.Sprintf("%s.%d", a.FullName(), pos)))
//}

// CoverValues The array must cover the given Values (order doesn't matter).
//
// op: {tags: { $all: ['red', 'blank'] }}
//
// https://www.mongodb.com/docs/manual/reference/operator/query/all/#use--all-to-match-values
func (a *arrayBaseField[T, ElemField]) CoverValues(values []T) filter.Filter {
	return filter.New(a, "$all", values)
}

// CoverVirValues The array must cover the given VirValues — either by a single element covering all,
// or by multiple elements covering them collectively (order doesn't matter).
//
//		op: { tags {
//	         $all: [
//	            { "$elemMatch" : { size: "M", num: { $gt: 50} } },
//	            { "$elemMatch" : { num : 100, color: "green" } }
//	         ]}
//	     }
//
// Both document
//
//	{ tags: [
//	   { size: "M", num: 100, color: "green" }
//	] }
//
// and document
//
//	{ tags: [
//	   { size: "S", num: 10, color: "blue" },
//	   { size: "M", num: 100, color: "blue" },
//	   { size: "L", num: 100, color: "green" }
//	] }
//
// can cover these VirValues
//
//	[
//	  { "$elemMatch" : { size: "M", num: { $gt: 50} } },
//	  { "$elemMatch" : { num : 100, color: "green" } }
//	]
//
// https://www.mongodb.com/docs/manual/reference/operator/query/all/#use--all-with--elemmatch
func (a *arrayBaseField[T, ElemField]) CoverVirValues(f func(sameElem *ElemField) []VirValue) filter.Filter {
	virValues := f(a.newElemField(""))
	return filter.New(a, "$all", virValues)
}

/**
更新：

一、普通域

二、元素级
$  就更新前面匹配第一个的那个值为新的值

$[]  所有的值  update every element of an array

$[<identifier>]
*/
