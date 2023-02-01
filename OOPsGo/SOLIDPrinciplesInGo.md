# SOLID principles in GoLang

| SOLID Principle                                                 | GoLang                                                                                                                                                                                                                                                     |
|-----------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Single Responsibility Principle                                 | Yes, through [struct](Structs.md)                                                                                                                                                                                                                                    |
| Open/Closed Principle                                           | No concept of generalization (class-based inheritance). Reliability is available through [embedding & interfaces](Interfaces.md).                                                                                                                          |
| Liskov substitution principle                                   | Instead of class-based inheritance, Golang provides a more powerful approach towards polymorphism via [Interfaces](Interfaces.md) and [Struct Embedding](Structs.md). <br/>- Go polymorphism involves creating many different data types that satisfy a common interface. |
| Interface Segregation                                           | [Interfaces](Interfaces.md) can be defined in GoLang.                                                                                                                                                                                                      |
| Dependency inversion principle                                  | Yes, through [interfaces](Interfaces.md) in GoLang.                                                                                                                                                                                                        |

# Single Responsibility
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
- Reusability is available as a form of [embedding](Interfaces.md).

Let’s take the example of the CommandExecutor, which is responsible for executing Commands.
- The Execute() and ValidateInput() methods need to handle each command separately.
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

## Liskov substitution principle

````go
type Command interface {
     Execute() ([]byte, error)
}

type CommandWithInput interface {
     Command
     ValidateInput() bool
}
````

## GoLang - Interface segregation principle
- In Golang [interfaces](Interfaces.md) are satisfied implicitly, rather than explicitly, which makes it easier to extend a class behaviour by implementing [multiple interface](Interfaces.md) based on needs.
- It also encourages to the design of small and reusable [interfaces](Interfaces.md).

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

# References
- [SOLID principle in GO](https://s8sg.medium.com/solid-principle-in-go-e1a624290346)
- [SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)