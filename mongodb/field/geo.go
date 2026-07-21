package field

type Geo struct {
	*baseField
	*BaseUpdater
	*BaseFilter
}

func NewGeo(name string) *Geo {
	ret := &Geo{
		baseField: &baseField{name: name},
	}
	ret.BaseFilter = &BaseFilter{ret.baseField}
	ret.BaseUpdater = &BaseUpdater{ret.baseField}

	return ret
}

func (g *Geo) FullName() string {
	return g.baseField.FullName()
}
