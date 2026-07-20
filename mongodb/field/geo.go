package field

type Geo struct {
	*base
	*BaseUpdater
	*BaseFilter
}

func NewGeo(name string) *Geo {
	ret := &Geo{
		base: &base{name: name},
	}
	ret.BaseFilter = &BaseFilter{ret.base}
	ret.BaseUpdater = &BaseUpdater{ret.base}

	return ret
}

func (g *Geo) FullName() string {
	return g.base.FullName()
}
