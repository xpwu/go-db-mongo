package field

import (
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/geojson"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PointFlat struct {
	name string
}

func (f *PointFlat) FullName() string {
	return f.name
}

func NewFlat(name string) *PointFlat {
	return &PointFlat{name + ".coordinates"}
}

func (f *PointFlat) Index() index.Key {
	return index.NewKey(f, "2d")
}

func (f *PointFlat) GeoWithinBox(bottomLeft, upperRight geojson.Coordinate) filter.Filter {
	return filter.New(f, "$geoWithin",
		bson.M{"$box": bson.A{bottomLeft.ToSlice(), upperRight.ToSlice()}})
}

func (f *PointFlat) GeoWithinPolygon(c1, c2, c3 geojson.Coordinate, cs ...geojson.Coordinate) filter.Filter {
	all := []geojson.Coordinate{c1, c2, c3}
	all = append(all, cs...)

	ret := make(bson.A, len(all))

	for i, a := range all {
		ret[i] = a.ToSlice()
	}

	return filter.New(f, "$geoWithin", bson.M{"$polygon": ret})
}

func (f *PointFlat) GeoWithinCircle(center geojson.Coordinate, radiusNoUnit float64) filter.Filter {
	return filter.New(f, "$geoWithin",
		bson.M{"$center": bson.A{center.ToSlice(), radiusNoUnit}})
}

/**
 *
 * Mongodb 官网关于单位的说明，此处有很多疑问
 *
 * https://docs.mongodb.com/manual/geospatial-queries/ 中
 * "$near (legacy coordinates, 2d index)	PointFlat" 说明 使用传统坐标与2d索引，将是使用flat方式计算
 *  但是
 * https://docs.mongodb.com/manual/reference/operator/query/near/#op._S_near
 * 中说 $maxDistance  " $maxDistance: <distance in radians>"  直角坐标中怎么使用弧度 这个需要再次
 * 确认文章的意思
 *
 * 而在 https://docs.mongodb.com/manual/reference/operator/query/center/#op._S_center 中同样
 * 是半径却是使用 "The circle’s radius, as measured in the units used by the coordinate system."
 * 这个单位是符合传统坐标特性的
 *
 *
 * https://docs.mongodb.com/manual/reference/operator/query/maxDistance/
 *
 * The measuring units for the maximum distance are determined by
 * the coordinate system in use.
 * For GeoJSON point object,
 * specify the distance in meters, not radians.
 *
 * 但是根据此文中的说明，又应该是使用的坐标系自身的单位，也即是没有单位
 *
 * 最后综合分析：应该是文档中有部分错误，maxDistance这里应该没有单位
 *
 */
func (f *PointFlat) Near(point geojson.Coordinate, maxDistance float64) filter.Filter {
	return filter.And(filter.New(f, "$near", point.ToSlice()),
		filter.New(f, "$maxDistance", maxDistance))
}
