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
		switch x := stack[li].(type) {
		case []interface{}:
			if len(x) > 0 {
				stack[li] = x[1:]
				stack = append(stack, x[0])
			} else {
				//Remove the empty elements.
				stack = stack[:li]
			}
		case int:
			stack = stack[:li]
			result = append(result, x)
		default:
			return nil, fmt.Errorf("bad data format for type: %s, value: %v",
				reflect.TypeOf(stack[li]), stack[li])
		}
	}

	return result, nil
}
