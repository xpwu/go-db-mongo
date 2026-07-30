package newfield

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Wx struct {
	Age  bson.Decimal128
	Time *int
}

type UserInfo struct {
	Login int
	Pass  []int
	Wx    Wx
	Ws    []Wx
	Pass2 [][]int16
	InWx  Wx `bson:"inWx,inline"`
}

//var filter2 = UserInfoDoc.AgeF().
