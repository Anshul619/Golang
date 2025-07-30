# Mock Testing
- This is helpful to test a function by mocking internal function calls (external api, db calls)

# Best Practices

| Title                                                                           | Description                                                                                                                            |
|---------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Depending on **interfaces** is a best practice for testability and flexibility. | Interfaces help in dependency inversion, mocking methods of objects.                                                                   |
| Never Marshal Mocks (or structs embedding mock.Mock)                            | Mocks are for simulating behavior, not for data. You must replace mocks with real or fake data before marshalling.                     |
| Manual mocks (vs [gomock](https://github.com/uber-go/mock) framework)           | No external dependencies or code generation.<br/>- Full control over mocked behavior.<br/>-Perfect for small projects or simple tests. |

# Helpful Frameworks
- [Testify](https://github.com/stretchr/testify)
- [gomock](https://github.com/uber-go/mock)