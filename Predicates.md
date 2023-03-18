# Predicates
- In Go, functions are first-class values. 
- This means that you can do with functions the same things you can do with all other values - [assign functions to variables](https://exercism.org/tracks/go/concepts/first-class-functions), pass them as arguments to other functions or even return functions from other functions.

````go
import "fmt"

func engGreeting(name string) string {
	return fmt.Sprintf("Hello %s, nice to meet you!", name)
}

func espGreeting(name string) string {
	return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := engGreeting			// greeting is a variable of type func(string) string
fmt.Println(greeting("Alice"))	// Hello Alice, nice to meet you!

greeting = espGreeting
fmt.Println(greeting("Alice")) 	// ¡Hola Alice, mucho gusto!
````