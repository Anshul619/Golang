# Generics
- In [modern Go](https://www.geeksforgeeks.org/go-language/generics-in-golang/), generics offer a way to improve type safety while maintaining flexibility.

# Example 1

````
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

type List[T any] struct {
    head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}
````

# Example 2

````
// UserRequest is a request for user information.
type UserRequest struct {
    Login string
}

// GroupRequest is a request for group information.
type GroupRequest struct {
    ID string
}

// Request is the set of all possible request types.
type Request interface {
    UserRequest | GroupRequest
}

// UnmarshalJSON safely decodes JSON into a pointer to a specific type.
func UnmarshalJSON[T Request](data []byte, d *T) error {
    return json.Unmarshal(data, d)
}
````

# Read more
- [Go Tip #6: Using Generics for Better Type Safety in Go](https://medium.com/@lenonrodrigues/go-tip-6-using-generics-for-better-type-safety-in-go-e385a75233c0)