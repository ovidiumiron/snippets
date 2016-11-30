package flatten

import (
	"fmt"
	"reflect"
)

// Flatten an array of arbitrarily nested arrays of integers into a flat array
// of integers. e.g. [[1,2,[3]],4] → [1,2,3,4].
func Flatten(a []interface{}) ([]int, error) {
	result := make([]int, 0)
	stack := []interface{}{a}

	for len(stack) > 0 {
		li := len(stack) - 1
		if e, ok := stack[li].([]interface{}); ok {
			if len(e) > 0 {
				stack[li] = e[1:]
				stack = append(stack, e[0])
			} else {
				//Remove the empty elements.
				stack = stack[:li]
			}
		} else if i, ok := stack[li].(int); ok {
			stack = stack[:li]
			result = append(result, i)
		} else {
			return nil, fmt.Errorf(
				fmt.Sprintf("bad data format for type: %s, value: %v",
					reflect.TypeOf(stack[li]), stack[li]))
		}
	}

	return result, nil
}
