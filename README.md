# concepts-go
🐇 Fundamental concepts about Golang.

[📚 Tutorial](https://gobyexample.com/variables)

# Values

Go has all basic types of values, and it can be used to define types from variables. They are `string`, `int`, `decimal` and `bool`.

# Variables

Go has static type inference. In other words, you don't need to define the type of variable, because lang get same type of first value. But, from there to behind, this type always will be followed.

```go
// So, if a define like this...
var text = "Some text"

// I can't do like this after...
text = 3;
```

In Go, variables are zero-valued, following their types. So, if a variable is `int`, if I do not declare initial value in variable definition, their value will be `0` (zero).
