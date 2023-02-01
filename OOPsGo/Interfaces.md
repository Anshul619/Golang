# What are interfaces?
- [Interfaces](https://gobyexample.com/interfaces) are a special type in Go that define a set of method signatures but do not provide implementations.
- A [Go type](../TypesGo.md) satisfies an interface by implementing the methods of that interface, nothing more.
- [Interfaces]() essentially act as placeholders for methods that will have multiple implementations based on what object is using it.
- In Go language, the [interface]() is a collection of method signatures and it is also a type means you can create a variable of an interface type.

````go
type geometry interface {
 area() float64
 perim() float64
}
````

# Embedding interface
- In embedding, [an interface can embed other interfaces or an interface can embed other interface’s method signatures](https://www.geeksforgeeks.org/embedding-interfaces-in-golang/?ref=lbp) in it, the result of both is the same as shown in Example 1 and 2.
- You are allowed to embed any number of interfaces in a single interface. And when an interface, embed other interfaces in it if we made any changes in the methods of the interfaces, then it will reflect in the embedded interface also.

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