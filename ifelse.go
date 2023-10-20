// ifelse provides equivalent of the ternary operator.
package ifelse

// If returns t if cond is true, otherwise f.
//
// Example:
//
//	If(true, 1, 2) // 1
//	If(false, 1, 2) // 2
func If[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}

// LazyIf returns t() if cond is true, otherwise f().
// LazyIf is useful when t and f are expensive to compute.
//
// Example:
//
//	LazyIf(true, func() int { return 1 }, func() int { return 2 }) // 1
//	LazyIf(false, func() int { return 1 }, func() int { return 2 }) // 2
func LazyIf[T any](cond bool, t, f func() T) T {
	if cond {
		return t()
	}
	return f()
}
