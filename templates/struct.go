// build +ignore
// templates package comment
package templates

// StructTemplate comment
type StructTemplate struct {
	// String comment
	String string `json:"string"`
	// Int comment
	Int int `json:"integer"`
	// Float comment
	Float float64 `json:"float"`
	// Bool comment
	Bool bool `json:"bool"`
	// Struct comment
	Struct *StructTemplate `json:"struct"`
}
