package context

// Re-export functions, types, and variables from package
// so it's easier to use for other packages.
var (
	Resolve = resolveContextInheritance
)
