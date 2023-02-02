# Object-Oriented Programming in GoLang
- Although Go has types and methods and allows an object-oriented style of programming, [there is no type hierarchy](https://staff.fnwi.uva.nl/a.j.p.heck/Courses/JAVAcourse/ch3/s1.html). 
- The concept of [interface](https://medium.com/@ubale.vikas9/interface-in-oops-6eae3731c242) in Go provides a different approach that is believed to be easy to use and in some ways more general. 
- There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing.
- Also, the lack of a type hierarchy makes [objects](https://www.techtarget.com/searchapparchitecture/definition/object-oriented-programming-OOP) in Go feel much more lightweight than in languages such as C++ or Java.

# Why is there no type inheritance in Go?
- Object-oriented programming involves too much discussion of the relationships between types, relationships that often could be derived automatically. Go takes a different approach.
- [Go types](../TypesGo.md) can satisfy many [interfaces](Interfaces.md) at once, without the complexities of traditional multiple inheritance. 
- [Interfaces](Interfaces.md) can be very lightweight with one or even zero methods can express a useful concept. 
- Because there are no explicit relationships between [types](../TypesGo.md) and [interfaces](Interfaces.md), there is no type hierarchy to manage or discuss.

# :star: Key Features

| Feature                                              | Remarks |
|------------------------------------------------------|---------|
| [Interfaces](Interfaces.md)                          | -       |
| [Struct](Structs.md)                                 | -       |
| [SOLID Principles in GoLang](SOLIDPrinciplesInGo.md) | -       |

# How do I get dynamic dispatch of methods?
- The only way to have dynamically dispatched methods is through an [interface](Interfaces.md).
- Methods on a struct or any other concrete type are always resolved statically.

# References
- [Frequently Asked Questions (FAQ)](https://go.dev/doc/faq)
- [Head First Design Patterns using Go](https://faun.pub/head-first-design-patterns-using-go-1-welcome-to-design-patterns-the-strategy-pattern-6cbd940e113a)
