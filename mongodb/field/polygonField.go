package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/geojson"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PolygonField struct {
	*Geo
}

func NewPolygonField(name string) *PolygonField {
	return &PolygonField{NewGeo(name)}
}

func (p *PolygonField) Set(value *geojson.Polygon) updater.Updater {
	return updater.New(p, `$set`, value)
}

func (p *PolygonField) SetOnInsert(value *geojson.Polygon) updater.Updater {
	return updater.New(p, `$setOnInsert`, value)
}

func (p *PolygonField) Index() index.Key {
	return index.NewKey(p, "2dsphere")
}

func (p *PolygonField) GeoIntersects(polygon *geojson.Polygon) filter.Filter {
	return filter.New(p, "$geoIntersects", bson.M{"$geometry": polygon})
}

func (p *PolygonField) GeoIntersectsMul(polygon *geojson.MultiPolygon) filter.Filter {
	return filter.New(p, "$geoIntersects", bson.M{"$geometry": polygon})
}

func (p *PolygonField) GeoIntersectsCcwCrs(polygon *geojson.Polygon) filter.Filter {
	return filter.New(p, "$geoIntersects", bson.M{"$geometry": NewCcwCrs(polygon)})
}
