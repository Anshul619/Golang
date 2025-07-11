# Go is "Pass by Value"
- As in all languages in the C family, everything in Go is passed by value.
- That is, a function always gets a copy of the thing being passed, as if there were an assignment statement assigning the value to the parameter.
- For instance, passing an int value to a function makes a copy of the int, and passing a pointer value makes a copy of the pointer, but not the data it points to.

# :star: Map & Slices are passed by reference
- Map, [Channels](ConcurrencyGoRoutines) and [slice values](Slices.md) are passed by reference.
- They are descriptors that contain pointers to the underlying map or slice data.
- i.e. If you pass a slice into a function, the function can modify its contents (*) and the modifications will be visible to the caller once it returns.
- i.e. Copying a map or slice value doesn't copy the data it points to. Copying an interface value makes a copy of the thing stored in the interface value.

<code style="color:red">Important Note:</code>
- [Appending an element to the slice, will not reflect in the caller function](https://stackoverflow.com/questions/21035279/why-does-go-slice-append-not-take-a-reference).
- Hence, [either pointer needs to be used or new slice would have to be returned](https://stackoverflow.com/questions/52565597/cannot-append-to-slice-inside-a-function).

# Golang pointers

| Operator     | Description                                                                                                                                                                                                                                                                     |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `&` operator | This operator is called the [address operator](https://www.techopedia.com/definition/25580/address-of-operator--c) and is used for returning the address of the variable stored in the pointer.                                                                                 |
| `*` operator | This operator is called a [dereferencing operator](https://en.wikipedia.org/wiki/Dereference_operator) and is used for accessing the value in the address stored by the pointer.<br/>- When `*` is placed before a [type](https://github.com/Anshul619/LLD-OOPs-Design-Patterns-Go/OOPs/TypesGo.md), it means that we are working with a pointer to a [type](https://github.com/Anshul619/LLD-OOPs-Design-Patterns-Go/OOPs/TypesGo.md). |

```go
x := 100
y := &x // Pointer - stores address of x in y

var z *int // Means that we are working with a pointer to a int type

fmt.Println(*y) // prints x variable 
```

[Read more](https://www.geeksforgeeks.org/pointers-in-golang/)

# Should I define methods on values or pointers?
- :star: If method needs to modify the receiver, receiver must be [pointer](https://go.dev/tour/moretypes/1). (as if its pass by reference)

```
func (s *MyStruct) pointerMethod() { } // method on pointer. 
func (s MyStruct)  valueMethod()   { } // method on value

pointer := new(int) // new pointer

func x(output [][]*int){} //pointer passed to the method

(*h)[10] // access elements of a pointer array
```

# What's the difference between new and make?
- [Unlike new, make's return type is the same as the type of its argument](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new), not a pointer to it.
- new allocates memory, while make initializes the slice, map, and channel types.

```
themes := make([]*Template, 0)
theme := new(Theme) // Pointer to new Theme object
```

# When should I use a pointer to an interface?
- Almost never. 
- Pointers to interface values arise only in rare, tricky situations involving disguising an interface value's type for delayed evaluation.


