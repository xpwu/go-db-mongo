package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/geojson"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MultiPolygonField struct {
	*Geo
}

func NewMultiPolygonField(name string) *MultiPolygonField {
	return &MultiPolygonField{NewGeo(name)}
}

func (p *MultiPolygonField) Set(value *geojson.MultiPolygon) updater.Updater {
	return updater.New(p, `$set`, value)
}

func (p *MultiPolygonField) SetOnInsert(value *geojson.MultiPolygon) updater.Updater {
	return updater.New(p, `$setOnInsert`, value)
}

func (p *MultiPolygonField) Index() index.Key {
	return index.NewKey(p, "2dsphere")
}

func (p *MultiPolygonField) GeoIntersects(polygon *geojson.Polygon) filter.Filter {
	return filter.New(p, "$geoIntersects", bson.M{"$geometry": polygon})
}

func (p *MultiPolygonField) GeoIntersectsMul(polygon *geojson.MultiPolygon) filter.Filter {
	return filter.New(p, "$geoIntersects", bson.M{"$geometry": polygon})
}

func (p *MultiPolygonField) GeoIntersectsCcwCrs(polygon *geojson.Polygon) filter.Filter {
	return filter.New(p, "$geoIntersects", bson.M{"$geometry": NewCcwCrs(polygon)})
}
