## About

The `typedenv` is a tiny library that provides functionality for working with typed environment variables.


## Usage

To get the library, use:

```bash
go get -u github.com/orewaee/typedenv
```

You can get env of a certain type as shown below. If the variable is missing, zero-value for the specified type will be returned.

```go
key := typedenv.String("KEY")
```

You can also specify a default value that will be returned if env is not present.

```go
number := typedenv.Int("NUMBER", 21)
```

If a default value is used many times, use global default values.

```go
typedenv.DefaultBool("BOOL", true)
```

You can also check the presence of a key in the environment or among default values.

```go
exists := typedenv.Lookup("KEY")
```
