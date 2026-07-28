package newfield

import (
  "fmt"
  "github.com/xpwu/go-db-mongo/mongodb/field"
)

func ExampleBuilder() {

  builder := field.NewBuilder()
  //builder.Build(reflect.TypeOf(UserInfo{}))
  field.BuildColl[UserInfo](builder)

  fmt.Println(true)
  // Output:
  // true
}
