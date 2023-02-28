# Go Struct
- A [struct (short for structure)](https://www.w3schools.com/go/go_struct.php) is used to create a collection of members of different data types, into a single variable.
- While arrays are used to store multiple values of the same data type into a single variable, [structs](https://www.w3schools.com/go/go_struct.php) are used to store multiple values of different data types into a single variable.
- A [struct](https://www.w3schools.com/go/go_struct.php) can be useful for grouping data together to create records.

````go
type Person struct {
  name string
  age int
  job string
  Salary int //exported
}

type Animal struct {
   //..
}

func (a *Animal) Eat() {...}
func (a *Animal) Sleep(){...}
func (a *Animal) Run(){...}

type Dog struct {
    Animal
    //..
}
````

# References
- [Structures in Golang](https://www.geeksforgeeks.org/structures-in-golang/)