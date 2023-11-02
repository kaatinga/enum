[![GitHub release](https://img.shields.io/github/release/kaatinga/enum.svg)](https://github.com/kaatinga/enum/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/enum/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/enum/branch/main/graph/badge.svg)](https://codecov.io/gh/kaatinga/enum)
[![lint workflow](https://github.com/kaatinga/enum/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/enum/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/enum/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# Enum Package

Some databases do not have enums and we ofter just store strings. This package provides functionality to encode and
decode strings into Enum values. This package is designed for situations where you need to represent strings as unique
integer values, often used for enumeration-like functionality. It goes with improved performance and memory usage.

## Installation

To install the `enum` package, simply run:

```bash
go get github.com//enum
```

## Usage

Here's a quick guide on how to use the package:

### Import the Package

```go
import "github.com/kaatinga/enum"
```

### Encoding a String

To encode a string into an `Enum` value, use the `Encode` function. It will convert a string of characters into an
integer value:

```go
s := "Hello"
encoded, err := enum.Encode(s)
if err != nil {
fmt.Println("Error:", err)
} else {
fmt.Println("Encoded:", encoded)
}
```

### Decoding an Enum

To decode an `Enum` back into its original string representation, you can use the `String` method of the `Enum` type:

```go
s := "Hello" // Replace with your actual encoded value
encoded, err := enum.Encode(s)
decoded := encoded.String()
fmt.Println("Decoded:", decoded)
```

## Supported Characters

The package supports encoding and decoding using the following characters:

- Digits: 0-9
- Uppercase Letters: A-Z
- Lowercase Letters: a-z
- Underscore: _

## String Length Limitation

The package enforces a maximum string length for encoding, supporting strings with lengths between 1 and 10 characters.
If the input string length does not meet this criterion, the `errInvalidLength` error is returned.

## Error Handling

In cases where an error occurs during encoding, the `Encode` function returns an error, allowing you to handle it
gracefully.

## License

This package is provided under the [MIT License](LICENSE.md). Feel free to use and modify it as needed.

## Contribution

Contributions to the package are welcome. If you encounter any issues or have ideas for improvements, please open an
issue or a pull request on the [GitHub repository](https://github.com/your-repository/enum).
