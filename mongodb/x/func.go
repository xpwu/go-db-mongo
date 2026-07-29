package x

import (
	"reflect"
	"strings"
)

func TypeFor[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func BaseTypeName(t reflect.Type) string {
	name := t.Name()
	// 去掉泛型参数部分（如果有）
	if idx := strings.Index(name, "["); idx != -1 {
		return name[:idx]
	}
	return name
}
