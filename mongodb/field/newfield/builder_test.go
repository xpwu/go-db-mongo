package newfield

import (
	"fmt"
	"github.com/xpwu/go-db-mongo/mongodb/field"
	"github.com/xpwu/go-db-mongo/mongodb/field/elsejson"
)

func ExampleBuilder() {

	builder := field.NewBuilder()

	field.BuildColl[UserInfo](builder)
	field.BuildColl[elsejson.ThirdParty](builder)

	fmt.Println(true)
	// Output:
	// true
}
