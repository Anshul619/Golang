# Go modules
- A go module may contain any number of packages and sub-packages, or it may contain none at all.
- [Read more](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)

````go
go mod init example/project // Create module with name "example/project"
go get golang.org/x/text@v0.3.5 // To add, upgrade, or downgrade a dependency
````

# References
- [Go module proxy at Grab](https://engineering.grab.com/go-module-proxy)