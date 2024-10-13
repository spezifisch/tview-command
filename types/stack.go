package types

import (
	"fmt"
	"io"
	"os"
)

// ContextStack is a stack to manage the UI contexts.
type ContextStack struct {
	stack []string
}

// NewContextStack creates and initializes a new ContextStack.
func NewContextStack() *ContextStack {
	return &ContextStack{
		stack: []string{"Global"}, // Start with the Global context
	}
}

// Push adds a new context to the stack.
func (cs *ContextStack) Push(context string) {
	cs.stack = append(cs.stack, context)
}

// Pop removes the current context from the stack, unless it's the last one.
func (cs *ContextStack) Pop() {
	if len(cs.stack) > 1 {
		cs.stack = cs.stack[:len(cs.stack)-1]
	}
}

// Current returns the currently active context.
func (cs *ContextStack) Current() string {
	if len(cs.stack) > 0 {
		return cs.stack[len(cs.stack)-1]
	}
	return "Global"
}

// Reset clears the stack and resets to the Global context.
func (cs *ContextStack) Reset() {
	cs.stack = []string{"Global"}
}

// PrintStack prints the current context stack (for debugging).
func (cs *ContextStack) PrintStack() {
	cs.PrintStackTo(os.Stdout)
}

// PrintStackTo allows printing the stack to a specified writer (useful for testing).
func (cs *ContextStack) PrintStackTo(w io.Writer) {
	fmt.Fprintf(w, "Current Context Stack: %v\n", cs.stack)
}
