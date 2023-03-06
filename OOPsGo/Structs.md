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

// Receiver functions
func (a *Animal) Eat() {...}
func (a *Animal) Sleep(){...}
func (a *Animal) Run(){...}

type Dog struct {
    Animal
    //..
}

func main() {
    person1 := Person{name:"Anshul", age: 31, job: "Backend Developer", Salary: 100}
}
````

# Struct Embedding
- Go supports embedding of structs and interfaces to express a more seamless composition of types
- A container embeds a base. An embedding looks like a field without a name.

````go
package main

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base
    str string
}

func main() {

    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }

    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

    fmt.Println("also num:", co.base.num)

    fmt.Println("describe:", co.describe())

    type describer interface {
        describe() string
    }

    var d describer = co
    fmt.Println("describer:", d.describe())
}
````

# References
- [Structures in Golang](https://www.geeksforgeeks.org/structures-in-golang/)