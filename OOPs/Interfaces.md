# What are interfaces?
- [Interfaces](https://gobyexample.com/interfaces) are a special type in Go that define a set of method signatures but do not provide implementations.
- A [Go type](TypesGo.md) satisfies an interface by implementing the methods of that interface, nothing more. 

> Note - We can't have fields, or methods implementation in interfaces.

# Example

````go
type geometry interface {
 area() float64
 perim() float64
}

// rectangle struct/type would implement geometry interface 
type rectangle struct {
}

func (r *rectangle) area() float64 { 
} 

func (r *rectangle) perim() float64 {
} 
````

# Interface names
- By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

# Embedding interface
- In [embedding](https://www.geeksforgeeks.org/embedding-interfaces-in-golang/?ref=lbp), an interface can embed other interfaces or an interface can embed other interface’s method signatures in it.
- You are allowed to embed any number of interfaces in a single interface. 
- And when an interface, embed other interfaces in it if we made any changes in the methods of the interfaces, then it will reflect in the embedded interface also.

````go
// Interface 1
type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
}

// Interface 3 embedded with interface 1 and 2
type FinalDetails interface {
	AuthorDetails
	AuthorArticles
}
````