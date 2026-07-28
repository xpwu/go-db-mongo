package newfield

import "github.com/xpwu/go-db-mongo/mongodb/geojson"

type Wx struct {
	Age  geojson.Point
	Time *int
}

type UserInfo struct {
	Login int
	Pass  []int
	//Wx    Wx
	//Ws    []Wx
	Pass2 []int16
}
