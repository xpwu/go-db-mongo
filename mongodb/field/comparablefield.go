package field

import "github.com/xpwu/go-db-mongo/mongodb/filter"

type ComparableFieldFilter[T comparable] struct {
	*baseFilter[T]
}

func (c *ComparableFieldFilter[T]) Eq(value T) filter.Filter {
	return filter.CompareByValue(c, filter.EQ, value)
}

func (c *ComparableFieldFilter[T]) EqField(f *ComparableFieldFilter[T]) filter.Filter {
	return filter.CompareByField(c, filter.EQ, f)
}

func (c *ComparableFieldFilter[T]) Ne(value T) filter.Filter {
	return filter.CompareByValue(c, filter.NE, value)
}

func (c *ComparableFieldFilter[T]) NeField(f *ComparableFieldFilter[T]) filter.Filter {
	return filter.CompareByField(c, filter.NE, f)
}

func (c *ComparableFieldFilter[T]) Gte(value T) filter.Filter {
	return filter.CompareByValue(c, filter.GTE, value)
}

func (c *ComparableFieldFilter[T]) GteField(f *ComparableFieldFilter[T]) filter.Filter {
	return filter.CompareByValue(c, filter.GTE, f)
}

func (c *ComparableFieldFilter[T]) Lte(value T) filter.Filter {
	return filter.CompareByValue(c, filter.LTE, value)
}

func (c *ComparableFieldFilter[T]) LteField(f *ComparableFieldFilter[T]) filter.Filter {
	return filter.CompareByValue(c, filter.LTE, f)
}

func (c *ComparableFieldFilter[T]) In(values []T) filter.Filter {
	return filter.New(c, "$in", values)
}

func (c *ComparableFieldFilter[T]) Nin(values []T) filter.Filter {
	return filter.New(c, "$nin", values)
}
