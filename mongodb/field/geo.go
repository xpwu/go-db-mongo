package field

// todo Geo

type Geo struct {
	//*BaseField
	//*BaseUpdater
	//*BaseFilter
	name string
}

func NewGeo(name string) *Geo {
	//ret := &Geo{
	//	BaseField: &BaseField{name: name},
	//}
	//ret.BaseFilter = &BaseFilter{ret.BaseField}
	//ret.BaseUpdater = &BaseUpdater{ret.BaseField}
	//
	//return ret
	return &Geo{name: name}
}

func (g *Geo) FullName() string {
	//return g.BaseField.FullName()
	return g.name
}
