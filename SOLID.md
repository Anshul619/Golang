# Single Responsibility
- Single responsibility principal aims to maintain a good level of Coupling that also maintains a good level of Cohesion.

> “Do one thing and do it well” — McIlroy (Unix philosophy)

# :star: SOLID principles in GoLang

| Principle                                                   | Supported in GoLang?   | Implementation                                                                                                                                                                                                                                                     |
|-------------------------------------------------------------|------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Single Responsibility Principle                             | :white_check_mark:     | Yes, through [struct](Structs.md)                                                                                                                                                                                                                                        |
| Open/Closed Principle                                       | :x:                    | No concept of generalization (class-based inheritance). <br/>- Reliability is available through [embedding & interfaces](Interfaces.md).                                                                                                                                 |
| Liskov substitution principle                               | :x:                    | Instead of class-based inheritance, Golang provides a more powerful approach towards polymorphism via [Interfaces](Interfaces.md) and [Struct Embedding](Structs.md). <br/>- Go polymorphism involves creating many different data types that satisfy a common interface. |
| Interface Segregation                                       | :white_check_mark:     | [Interfaces](Interfaces.md) can be defined in GoLang.                                                                                                                                                                                                                    |
| Dependency inversion principle                              | :white_check_mark:     | Yes, through [interfaces](Interfaces.md) in GoLang.                                                                                                                                                                                                                      |

[Read more about SOLID Principles in GoLang](../SOLIDPrinciples.md)


# Example
- CommandFactory and CommandExecutor are loosely coupled via Command module.

````go
type Command struct {
    commandType string 
    args []string
}

type CommandFactory struct {
    ...
}

// Create decode and validate the command
func (cf CommandFactory) Create(data []byte) (*Command, error) {
    // decode command
    command, err := cf.Decode(data)
    if err != nil {
        return nil, err
    }
    // validate type
    switch cf.Type { 
       case Foo:
       case Bar:
       default:
          return nil, InvalidCommandType    
    }
    return command, nil
}

type CommandExecutor struct {
}

// Execute executes the command 
func (ce CommandExecutor) Execute(command *Command) ([]byte, error) {
   // validate input and execute 
   switch command.Type {
   case Foo: 
       if len(args) == 0 || len(args[0]) == 0 {
           return nil, InvalidInput
       }
       return ExecuteFoo(command)   case Bar: // Bar doesn't take any input
       return ExecuteBar(command)
   }
}
````

# Open/Closed Principle
- In Golang there is no concept of generalization.
- Reusability is available as a form of [embedding](OOPs/Interfaces.md).

> “A module should be open for extensions, but closed for modification” — Robert C. Martin

## Example
- Let’s take the example of the CommandExecutor, which is responsible for executing Commands.
- Execute() and ValidateInput() methods need to handle each command separately.
- So every time a new command is added Execute() implementation needs to change.

````go
type Command interface {
     Execute() ([]byte, error)
     ValidateInput() bool
}
type CommandExecutor struct {
}

func (c CommandExecutor) Execute(command Command) {
     if command.ValidateInput() {
          command.Execute()
     }
}

type FooCommand struct {
     args []string // need args
}

func (c FooCommand) ValidateInput() {
    // validate args 
    if len(args) >= 1 && len(args[0]) > 0 {
       return true
    }
    return false
}

func (c FooCommand) Execute() ([]byte, error) {
    ...
}

type BarCommand struct {
}

func (c BarCommand) ValidateInput() {
    // does nothing 
    return false
}

func (c BarCommand) Execute() ([]byte, error) {
    ...
}
````

# Liskov substitution principle

> “Derived methods should expect no more and provide no less” — Robert C. Martin

````go
type Command interface {
     Execute() ([]byte, error)
}

type CommandWithInput interface {
     Command
     ValidateInput() bool
}
````

# Interface segregation principle
- In Golang [interfaces](OOPs/Interfaces.md) are satisfied implicitly, rather than explicitly, which makes it easier to extend a class behaviour by implementing [multiple interface](OOPs/Interfaces.md) based on needs.
- It also encourages to the design of small and reusable [interfaces](OOPs/Interfaces.md).

> “Many client specific interfaces are better than one general purpose interface” — Robert C. Martin

````go
type I1 interface { // consumed by C1
    M1()
    M2()
    M3()       // also defined in I2
}
type I2 interface { // consumed by C2 and C3
    M3()       // here M3 not in a separate interface as no client, use an interface with only M3()
    M4()
}
type I3 interface { // consumed by C4
    M5()       // similarly M5() only used along with I1 and I2, thus not needed to have it in a separate interface
    I1
    I2
}
````

# Dependency inversion principle

> “Depend upon Abstractions. Do not depend upon concretions” — Robert C. Martin

````go
type CommandFactory struct {
     decoder JsonDecoder // decoder decodes the command
}// Create decode and validate the command
func (cf CommandFactory) Create(encoded String) (Command, error) {
    // decode command
    command, err := cf.decoder.Decode(data)
    if err != nil {
        return nil, err
    }
    ...
}
````

# References
- [SOLID principle in GO](https://s8sg.medium.com/solid-principle-in-go-e1a624290346)
- [SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)