# Is Go an object-oriented language?
- Yes and no. 
- Although Go has types and methods and allows an object-oriented style of programming, [there is no type hierarchy](https://staff.fnwi.uva.nl/a.j.p.heck/Courses/JAVAcourse/ch3/s1.html). 
- The concept of [interface](https://medium.com/@ubale.vikas9/interface-in-oops-6eae3731c242) in Go provides a different approach that we believe is easy to use and in some ways more general. 
- There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing.
- Also, the lack of a type hierarchy makes [objects](https://www.techtarget.com/searchapparchitecture/definition/object-oriented-programming-OOP) in Go feel much more lightweight than in languages such as C++ or Java.
- [SOLID Principles in GoLang](SOLIDPrinciplesInGo.md)

# How do I get dynamic dispatch of methods?
- The only way to have dynamically dispatched methods is through an [interface](Interfaces.md). 
- Methods on a struct or any other concrete type are always resolved statically.

# Why is there no type inheritance in Go?
- Object-oriented programming, at least in the best-known languages, involves too much discussion of the relationships between types, relationships that often could be derived automatically. Go takes a different approach.
- Rather than requiring the programmer to declare ahead of time that two types are related, in Go a type automatically satisfies any interface that specifies a subset of its methods. 
- Besides, reducing the bookkeeping, this approach has real advantages. 
- [Go types](../TypesGo.md) can satisfy many [interfaces](Interfaces.md) at once, without the complexities of traditional multiple inheritance. 
- [Interfaces](Interfaces.md) can be very lightweight—an [interface](Interfaces.md) with one or even zero methods can express a useful concept. 
- [Interfaces](Interfaces.md) can be added after the fact if a new idea comes along or for testing—without annotating the original types. 
- Because there are no explicit relationships between [types](../TypesGo.md) and [interfaces](Interfaces.md), there is no type hierarchy to manage or discuss.
- It takes some getting used to but this implicit style of type dependency is one of the most productive things about Go.

# Class Example

````go
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
- [Frequently Asked Questions (FAQ)](https://go.dev/doc/faq)
- [Head First Design Patterns using Go](https://faun.pub/head-first-design-patterns-using-go-1-welcome-to-design-patterns-the-strategy-pattern-6cbd940e113a)
