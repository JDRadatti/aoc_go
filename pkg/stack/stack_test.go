package stack

import "testing"

func TestPush(t *testing.T) {
    stack := Stack[int]{}
    inputs := []int{ 1, 10, -1 }
    expected := []int{}
    for i, input := range inputs {

        stack.Push(input)

        if stack[i] != input {
            t.Errorf("Stack test failed on input %d. Expected %d but got %d", i, expected[i], stack[i])
        }

    }
}

func TestPop(t *testing.T) {
    stack := Stack[int]{1, 10, -1}
    expected := []int{-1, 10, 1, 0}
    okExpected := []bool{true, true, true, false}
    for i := range expected {
        result, ok := stack.Pop()
        if expected[i] != result {
            t.Errorf("Stack test failed on input %d. Expected %d but got %d", i, expected[i], result)
        }

        if okExpected[i] != ok {
            t.Errorf("Stack test failed on input %d. Expected %t but got %t", i, okExpected[i], ok)
        }

    }
}
