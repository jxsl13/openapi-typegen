package literals

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

type stringLit interface {
	string |
		int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128 |
		bool
}

// String returns a string literal for the given value.
// Supported types are string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, and bool.
func String[T stringLit](value T) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: strconv.Quote(fmt.Sprint(value)),
	}
}

type integerLit interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	// TODO: do we need a string here?
}

// Int returns an integer literal for the given value.
// Supported types are int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, and uintptr.
func Int[T integerLit](value T) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.INT,
		Value: fmt.Sprint(value),
	}
}

type floatLit interface {
	float32 | float64
}

// Float returns a floating-point literal for the given value.
// Supported types are float32 and float64.
func Float[T floatLit](value T) *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.FLOAT,
		Value: fmt.Sprint(value),
	}
}
