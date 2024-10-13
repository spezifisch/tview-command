package types

import (
	"bytes"
	"strings"
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

func TestPopExpect(t *testing.T) {
	// Test case where PopExpect succeeds
	t.Run("PopExpect_Success", func(t *testing.T) {
		cs := &ContextStack{
			stack: []string{"Global", "QueuePage", "QueueList"},
		}

		cs.PopExpect("QueueList") // This should succeed without panicking

		if len(cs.stack) != 2 {
			t.Fatalf("Expected stack length 2 after PopExpect, got %d", len(cs.stack))
		}
		if cs.stack[len(cs.stack)-1] != "QueuePage" {
			t.Fatalf("Expected 'QueuePage' at the top of the stack, got '%s'", cs.stack[len(cs.stack)-1])
		}
	})

	// Test case where PopExpect fails (mismatched value)
	t.Run("PopExpect_Fail", func(t *testing.T) {
		cs := &ContextStack{
			stack: []string{"Global", "QueuePage", "QueueList"},
		}

		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Expected panic but got none")
			} else {
				// Ensure panic message is as expected
				expectedMessage := "PopExpect: expected 'QueueSidebar' but got 'QueueList'"
				if !strings.Contains(r.(string), expectedMessage) {
					t.Fatalf("Unexpected panic message, got: '%v'", r)
				}
			}
		}()

		cs.PopExpect("QueueSidebar") // This should panic with the error message
	})

	// Test case where PopExpect is called on an empty stack
	t.Run("PopExpect_EmptyStack", func(t *testing.T) {
		cs := &ContextStack{
			stack: []string{},
		}

		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Expected panic for empty stack but got none")
			} else {
				// Ensure panic message is as expected
				expectedMessage := "PopExpect called on an empty stack"
				if !strings.Contains(r.(string), expectedMessage) {
					t.Fatalf("Unexpected panic message, got: '%v'", r)
				}
			}
		}()

		cs.PopExpect("AnyValue") // This should panic with the empty stack message
	})
}
