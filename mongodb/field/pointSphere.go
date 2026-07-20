package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/geojson"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PointSphere struct {
	name string
}

func NewSphere(name string) *PointSphere {
	return &PointSphere{name}
}

func (s *PointSphere) FullName() string {
	return s.name
}

func (s *PointSphere) Index() index.Key {
	return index.NewKey(s, "2dsphere")
}

func (s *PointSphere) GeoWithinPyg(polygon *geojson.Polygon) filter.Filter {
	return filter.New(s, "$geoWithin", bson.M{"$geometry": polygon})
}

func (s *PointSphere) GeoWithinMulPyg(polygon *geojson.MultiPolygon) filter.Filter {
	return filter.New(s, "$geoWithin", bson.M{"$geometry": polygon})
}

func (s *PointSphere) GeoWithinCcwCrs(polygon *geojson.Polygon) filter.Filter {
	return filter.New(s, "$geoWithin", bson.M{"$geometry": NewCcwCrs(polygon)})
}

// 弧度
func (s *PointSphere) GeoWithinCircle(center geojson.Coordinate, radians float64) filter.Filter {
	return filter.New(s, "$geoWithin",
		bson.M{"$centerSphere": bson.A{center.ToSlice(), radians}})
}

/**
 *
 * 在球面坐标中，为什么长度又使用米？
 *
 * https://docs.mongodb.com/manual/reference/operator/query/maxDistance/
 *
 * The measuring units for the maximum distance are determined by
 * the coordinate system in use.
 * For GeoJSON point object,
 * specify the distance in meters, not radians.
 *
 * 根据此文中的说明，的确应该是使用的米
 *
 * 根据分析，GeoJson 是按照WGS84的标准来计算的，定死半径为地球半径，所以使用了单位"米"
 *
 */

func (s *PointSphere) Near(point geojson.Point, maxDistance float64, minDistance ...float64) filter.Filter {
	var value interface{}
	if len(minDistance) == 0 {
		value = bson.M{"$geometry": point, "$maxDistance": maxDistance}
	} else if len(minDistance) == 1 {
		value = bson.M{"$geometry": point, "$maxDistance": maxDistance, "$minDistance": minDistance[0]}
	} else {
		panic("args error")
	}

	return filter.New(s, "$nearSphere", value)
}
