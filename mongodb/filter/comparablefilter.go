package filter

type ComparableFilter[T any] struct {
	*BaseFilter[T]
}

func (c *ComparableFilter[T]) Eq(value T) Filter {
	return CompareByValue(c, EQ, value)
}

func (c *ComparableFilter[T]) EqField(f *ComparableFilter[T]) Filter {
	return CompareByField(c, EQ, f)
}

func (c *ComparableFilter[T]) Ne(value T) Filter {
	return CompareByValue(c, NE, value)
}

func (c *ComparableFilter[T]) NeField(f *ComparableFilter[T]) Filter {
	return CompareByField(c, NE, f)
}

func (c *ComparableFilter[T]) Gte(value T) Filter {
	return CompareByValue(c, GTE, value)
}

func (c *ComparableFilter[T]) GteField(f *ComparableFilter[T]) Filter {
	return CompareByValue(c, GTE, f)
}

func (c *ComparableFilter[T]) Lte(value T) Filter {
	return CompareByValue(c, LTE, value)
}

func (c *ComparableFilter[T]) LteField(f *ComparableFilter[T]) Filter {
	return CompareByValue(c, LTE, f)
}

func (c *ComparableFilter[T]) In(values []T) Filter {
	return New(c, "$in", values)
}

func (c *ComparableFilter[T]) Nin(values []T) Filter {
	return New(c, "$nin", values)
}

func NewComparableFilter[T comparable](base *BaseFilter[T]) *ComparableFilter[T] {
	return &ComparableFilter[T]{base}
}

type EqualAble[T EqualAble[T]] interface {
	Equal(T) bool
}

func NewEqualAbleFilter[T EqualAble[T]](base *BaseFilter[T]) *ComparableFilter[T] {
	return &ComparableFilter[T]{base}
}
