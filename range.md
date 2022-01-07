# Range function in Go
- range return two values `index` and `value`
- if you dont want to use index then you can use blank identifier in the go `_` yes underscore. 
- you cant define varaibles you dont want to use in the go proramming language. 
- so `_` will help you where you cant escaping the assigning but you dont need the value. best example is for loop

```go

package main

import (
    "fmt"
)

func main() {

    var a = [5]int{1,2,3,4,5};
    for _, value := range ( a ) {
        fmt.Println("Value : " , value);
    } 
}
```