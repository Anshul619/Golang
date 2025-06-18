# What are Golang packages?
- In Go, the programs are built by using packages that help in managing the dependencies efficiently.
- The package is declared at the top of the Go source file as `package <package_name>`.
- The packages can be imported to our source file by writing: `import <package_name>`.
- :star: Files in the same package can freely call functions defined in other files.
- Struct, methods of the package are referenced using dot operator `.`. 

![img.png](../assets/gopackages_img.png)

# Exported Field/Method
This function takes a name parameter whose type is string.
- The function also returns a string.
- In Go, a function whose name starts with a capital letter can be called by a function not in the same package.

````go
func Hello(name string) string
````