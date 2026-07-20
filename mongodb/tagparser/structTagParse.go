package tagparser

import (
	"go.mongodb.org/mongo-driver/v2/bson/bsoncodec"
	"reflect"
	"strings"
)

// 建立struct field的转换，与 bsoncodec.DefaultStructTagParser 的区别是
// 不转换为小写，而是直接使用字段名。在其他CRUD函数中需要使用字符串的字段名时，如果默认
// 转换为小写，容易写错

func StructTagParser(sf reflect.StructField) (tags bsoncodec.StructTags, err error) {
	key := sf.Name
	tag, ok := sf.Tag.Lookup("bson")
	if !ok && !strings.Contains(string(sf.Tag), ":") && len(sf.Tag) > 0 {
		tag = string(sf.Tag)
	}
	var st bsoncodec.StructTags
	if tag == "-" {
		st.Skip = true
		return st, nil
	}

	for idx, str := range strings.Split(tag, ",") {
		if idx == 0 && str != "" {
			key = str
		}
		switch str {
		case "omitempty":
			st.OmitEmpty = true
		case "minsize":
			st.MinSize = true
		case "truncate":
			st.Truncate = true
		case "inline":
			st.Inline = true
		}
	}

	st.Name = key

	return st, nil
}
