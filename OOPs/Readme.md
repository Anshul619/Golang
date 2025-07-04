# Object-Oriented Programming in GoLang
- Although Go has types and methods and allows an [object-oriented](https://github.com/Anshul619/LLD-OOPs-Design-Patterns/blob/main/OOP.md) style of programming, there is [no type hierarchy](https://staff.fnwi.uva.nl/a.j.p.heck/Courses/JAVAcourse/ch3/s1.html). 
- The concept of [interface](https://medium.com/@ubale.vikas9/interface-in-oops-6eae3731c242) in Go provides a different approach that is believed to be easy to use and in some ways more general. 
- There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing.
- Also, the lack of a type hierarchy makes [objects](https://www.techtarget.com/searchapparchitecture/definition/object-oriented-programming-OOP) in Go feel much more lightweight than in languages such as C++ or Java.

# :star: OOPs concepts in GoLang

| OOPs Concept                             | Supported in GoLang? | Implementation                                                                                                                                                                                                                                                                                        | Example Code                                         |
|------------------------------------------|----------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------|
| Inheritance - Extend a class behaviour   | :white_check_mark:   | - By embedding another [type](TypesGo.md) or [interface](Interfaces.md), a type can extend<br/>- There is no `extends` keyword in GoLang.                                                                                                                                                               | ```go type I3 interface {I1}``` |
| Inheritance - Implement a class behavior | :white_check_mark:   | - By implementing [all methods of the interface](https://go.dev/doc/faq#implements_interface), a type can implement an interface.<br/>- There is no `implements` keyword in GoLang.                                                                                                                     | -                                                    |
| Abstraction                              | :white_check_mark:   | - [Interfaces](https://www.golangprograms.com/example-of-abstraction-using-interfaces-in-golang.html) can be used to create common abstraction that can be used by multiple types. <br/>- Interfaces define one or more method declarations that must be satisfied to be compatible with the interface. | -                                                    |
| Encapsulation                            | :white_check_mark:   | Using [exported identifier](https://www.geeksforgeeks.org/encapsulation-in-golang/) (i.e. first letter in capital), we can expose a method publicly in the package.                                                                                                                                   | -                                                    |
| Polymorphism - Overloading               | :x:                  | Go [doesn't](https://go.dev/doc/faq#overloading) support method overloading.                                                                                                                                                                                                                          | -                                                    |
| Polymorphism - Overriding                | :x:                  | Go [doesn't](https://stackoverflow.com/questions/38123911/golang-method-override) support method overriding                                                                                                                                                                                           | -                                                    |

# Why is there no type inheritance in Go?
- Object-oriented programming involves too much discussion of the relationships between types, relationships that often could be derived automatically.
- Go takes a different approach.
- [Go types](TypesGo.md) can satisfy many [interfaces](Interfaces.md) at once, without the complexities of traditional multiple inheritance.
- [Interfaces](Interfaces.md) can be very lightweight with one or even zero methods can express a useful concept.
- Because there are no explicit relationships between [types](TypesGo.md) and [interfaces](Interfaces.md), there is no type hierarchy to manage or discuss.

# References
- [Frequently Asked Questions (FAQ)](https://go.dev/doc/faq)
- [Head First Design Patterns using Go](https://faun.pub/head-first-design-patterns-using-go-1-welcome-to-design-patterns-the-strategy-pattern-6cbd940e113a)
