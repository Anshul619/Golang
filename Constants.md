# Constants
- The keyword const is used for declaring constants followed by the desired name and the type of value the constant will hold. You must assign a value at the time of the constant declaration, you can't assign a value later as with variables.
- You can also omit the type at the time the constant is declared. The type of the value assigned to the constant will be used as the type of that variable.
- By convention, constant names are usually written in uppercase letters. This is for their easy identification and differentiation from variables in the source code.

````go
package main

import "fmt"

const PRODUCT string = "Canada"
const PRICE = 500

func main() {
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)
}
````

## Multiple Constants

````go
package main

import "fmt"

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK  = true
)

func main() {
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(STOCK)
}
````