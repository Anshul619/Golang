# Generics
- In [modern Go](https://go.dev/tour/generics/1), generics offer a way to improve type safety while maintaining flexibility.

# When to Use Generics in Go?

|                                       | Description                                                                                                                                                                                                  | Example                                                                          |
|---------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------|
| Reusable Data Structures              | If you’re writing common data structures like stacks, queues, linked lists, or trees that work with multiple types, generics let you write them once and reuse them with any type.                           | A generic Stack[T any] that works with int, string, or any user-defined type.    |
| Generic Algorithms                    | Algorithms such as sorting, filtering, mapping, or searching often don’t depend on specific data types. Generics allow implementing these once and applying them across types.                               | A generic Filter function that filters a slice of any type based on a predicate. |
| Type Safety Across Similar Operations | Generics help catch type errors at compile time instead of using interface{} and then doing runtime type assertions. This makes code safer and easier to maintain.                                           |                                                                                  |
| Reducing Code Duplication             | Without generics, you might write nearly identical functions for different types (e.g., SumInts, SumFloats). <br/>- Generics allow writing a single function that works for all types that meet constraints. |                                                                                  |
| Improved Performance Over interface{} | Generics avoid boxing/unboxing and runtime type checks associated with interface{} types, potentially improving performance.                                                                                 |                                                                                  |

# When Not to Use Generics
- When the logic is tightly coupled to a specific type or requires type-specific methods not easily expressed with constraints.
- For very simple cases where generics might overcomplicate the code or reduce readability.
- When you only have one or two specific types to handle, simple functions might be better.

# Read more
- [Go Tip #6: Using Generics for Better Type Safety in Go](https://medium.com/@lenonrodrigues/go-tip-6-using-generics-for-better-type-safety-in-go-e385a75233c0)