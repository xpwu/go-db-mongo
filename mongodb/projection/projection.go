package projection

import (
	"github.com/xpwu/go-db-mongo/mongodb"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type P bson.M

const (
	inc = 1
	exc = 0
)

func (p P) Inc(f mongodb.Field, fn ...mongodb.Field) P {
	p[f.FullName()] = inc
	for _, f := range fn {
		p[f.FullName()] = inc
	}
	return p
}

func (p P) Exc(f mongodb.Field, fn ...mongodb.Field) P {
	p[f.FullName()] = exc
	for _, f := range fn {
		p[f.FullName()] = exc
	}
	return p
}

// first five: [0,5]; last five:[-5, 5]; from 3rd to 7th: [3, 5]
func (p P) Slice(f mongodb.Field, pos, len int) P {
	var slice interface{}
	if pos == 0 {
		slice = len
	} else if -pos == len {
		slice = pos
	} else {
		slice = bson.A{pos, len}
	}

	p[f.FullName()] = bson.M{"$slice": slice}

	return p
}

func Inc(f mongodb.Field, fn ...mongodb.Field) P {
	p := P{}
	p.Inc(f, fn...)
	return p
}

func Exc(f mongodb.Field, fn ...mongodb.Field) P {
	p := P{}
	p.Exc(f, fn...)
	return p
}
