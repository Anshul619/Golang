# Mock Testing
- This is helpful to test a function by mocking internal function calls (external api, db calls)

# Best Practices

| Title                                                                       | Remarks                                                                                                            |
|-----------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| Depending on interfaces is a best practice for testability and flexibility. | Interfaces help in dependency inversion, mocking methods of objects.                                               |
| Never Marshal Mocks (or structs embedding mock.Mock)                        | Mocks are for simulating behavior, not for data. You must replace mocks with real or fake data before marshalling. |

# Helpful Packages
- [Testify](https://github.com/stretchr/testify)