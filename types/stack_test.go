package types

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextStack_PushPop(t *testing.T) {
	stack := NewContextStack()

	// Check the initial state
	assert.Equal(t, "Global", stack.Current(), "Initial context should be Global")

	// Push new contexts
	stack.Push("QueuePage")
	stack.Push("QueueList")
	assert.Equal(t, "QueueList", stack.Current(), "Current context should be QueueList")

	// Pop contexts
	stack.Pop()
	assert.Equal(t, "QueuePage", stack.Current(), "Current context should be QueuePage")

	stack.Pop()
	assert.Equal(t, "Global", stack.Current(), "Should return to Global context")

	// Ensure Pop does not remove Global context
	stack.Pop()
	assert.Equal(t, "Global", stack.Current(), "Global context should remain after all pops")
}

func TestContextStack_Reset(t *testing.T) {
	stack := NewContextStack()

	// Push some contexts
	stack.Push("QueuePage")
	stack.Push("QueueList")

	// Reset the stack
	stack.Reset()
	assert.Equal(t, "Global", stack.Current(), "After reset, context should be Global")
}

func TestContextStack_MultiplePushes(t *testing.T) {
	stack := NewContextStack()

	// Push several contexts
	stack.Push("QueuePage")
	stack.Push("QueueSidebar")
	stack.Push("QueueDetails")

	// Check that the current context is the last pushed
	assert.Equal(t, "QueueDetails", stack.Current(), "Current context should be QueueDetails")
}

func TestContextStack_PrintStack(t *testing.T) {
	stack := NewContextStack()

	// Capture the output of PrintStack
	var buf bytes.Buffer
	stack.PrintStackTo(&buf)
	output := buf.String()

	// Verify the initial state (should only have "Global")
	assert.Equal(t, "Current Context Stack: [Global]\n", output, "Stack output should start with Global context")

	// Push new context and verify
	stack.Push("QueuePage")
	stack.Push("QueueList")
	buf.Reset()
	stack.PrintStackTo(&buf)
	output = buf.String()
	assert.Equal(t, "Current Context Stack: [Global QueuePage QueueList]\n", output, "Stack output should reflect the current context stack")
}

func TestContextStack_Current_EmptyStack(t *testing.T) {
	// Create an empty stack (directly modify for testing purposes)
	stack := &ContextStack{
		stack: []string{},
	}

	// Test Current() when stack is empty
	assert.Equal(t, "Global", stack.Current(), "Current should return 'Global' even if stack is empty")
}
