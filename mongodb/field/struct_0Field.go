package field

import "strings"

func StructNext(self, fName string) string {
	if self == "" {
		return fName
	}
	if fName == "" {
		return self
	}

	b := strings.Builder{}
	b.WriteString(self)
	b.WriteByte('.')
	b.WriteString(fName)

	return b.String()
}

type StructUpdaterF struct {
	*BaseUpdater
}

func NewStructUpdaterF(name string) *StructUpdaterF {
	return &StructUpdaterF{&BaseUpdater{&base{name: name}}}
}

type StructFilterF struct {
	*BaseFilter
}

func NewStructFilterF(name string) *StructFilterF {
	return &StructFilterF{&BaseFilter{&base{name: name}}}
}
