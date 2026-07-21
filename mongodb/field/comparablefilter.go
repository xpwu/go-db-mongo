package field

import "github.com/xpwu/go-db-mongo/mongodb/filter"

func (b *baseField[T]) Eq(value T) filter.Filter {
	return filter.CompareByValue(b, filter.EQ, value)
}

func (b *baseField[T]) EqField(f filter.ComparableFilterField[T]) filter.Filter {
	return filter.CompareByField(b, filter.EQ, f)
}

func (b *baseField[T]) Ne(value T) filter.Filter {
	return filter.CompareByValue(b, filter.NE, value)
}

func (b *baseField[T]) NeField(f filter.ComparableFilterField[T]) filter.Filter {
	return filter.CompareByField(b, filter.NE, f)
}

func (b *baseField[T]) Gte(value T) filter.Filter {
	return filter.CompareByValue(b, filter.GTE, value)
}

func (b *baseField[T]) GteField(f filter.ComparableFilterField[T]) filter.Filter {
	return filter.CompareByValue(b, filter.GTE, f)
}

func (b *baseField[T]) Lte(value T) filter.Filter {
	return filter.CompareByValue(b, filter.LTE, value)
}

func (b *baseField[T]) LteField(f filter.ComparableFilterField[T]) filter.Filter {
	return filter.CompareByValue(b, filter.LTE, f)
}

func (b *baseField[T]) In(values []T) filter.Filter {
	return filter.New(b, "$in", values)
}

func (b *baseField[T]) Nin(values []T) filter.Filter {
	return filter.New(b, "$nin", values)
}

var (
	_ filter.ComparableFilterField[any] = &baseField[any]{}
)
