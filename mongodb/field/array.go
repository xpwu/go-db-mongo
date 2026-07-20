package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
)

type Array struct {
	*BaseFilter
	*BaseUpdater
	*BaseKey
}

func (a *Array) FullName() string {
	return a.BaseKey.FullName()
}

func (a *Array) PopFirst() updater.Updater {
	return updater.New(a, `$pop`, -1)
}

func (a *Array) PopLast() updater.Updater {
	return updater.New(a, `$pop`, 1)
}

func (a *Array) PullByF(f filter.Filter) updater.Updater {
	return updater.PullByFilter(a, f)
}

// 数组的同一个元素满足f
func (a *Array) SameEleMatch(f filter.Filter) filter.Filter {
	return filter.SameElemMatch(a, f)
}

func (a *Array) Size(sz int) filter.Filter {
	return filter.New(a, `$size`, sz)
}

func NewArray(fName string) *Array {
	b := &base{name: fName}

	return &Array{&BaseFilter{b},
		&BaseUpdater{b},
		&BaseKey{b},
	}
}
