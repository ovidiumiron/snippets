package flatten

import "testing"

// Return true if a != b; else return false.
func notEq(a, b []int) bool {
	if a == nil && b == nil {
		return false
	}

	if a == nil || b == nil {
		return true
	}

	if len(a) != len(b) {
		return true
	}

	for i := range a {
		if a[i] != b[i] {
			return true
		}
	}

	return false
}

func TestFlatten(t *testing.T) {
	input := []interface{}{1, []interface{}{2}, 3, []interface{}{4}}

	want := []int{1, 2, 3, 4}

	if got, err := Flatten(input); err != nil || notEq(want, got) {
		t.Errorf("Flatten(%v) := %v '%v', want: %v", input, got, err, want)
	}
}

func TestFlattenOneLevelHuge(t *testing.T) {
	const size = 10000

	input := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		input = append(input, []interface{}{i})
	}

	want := make([]int, size)
	for i := 0; i < size; i++ {
		want[i] = i
	}

	if got, err := Flatten(input); err != nil || notEq(want, got) {
		t.Errorf("Flatten(%v) := %v '%v', want: %v", input, got, err, want)
	}
}

func TestFlattenHugeNested(t *testing.T) {
	const size = 100000

	input := make([]interface{}, 0)
	for i := 0; i < size; i++ {
		input = append(make([]interface{}, 0), append(input, i))
	}

	want := make([]int, size)
	for i := 0; i < size; i++ {
		want[i] = i
	}

	if got, err := Flatten(input); err != nil || notEq(want, got) {
		t.Errorf("Flatten(%v) := %v '%v', want: %v", input, got, err, want)
	}
}

func TestFlattenNegativeTests(t *testing.T) {
	input := []interface{}{1, []interface{}{"A"}, 3, []interface{}{4}}
	want := "bad data format for type: string, value: A"

	got, err := Flatten(input)
	if err == nil {
		t.Errorf("Flatten(%v) := %v '%v', want error: '%v'", input, got, err, want)
	} else if err.Error() != want {
		t.Errorf("Flatten(%v) := %v '%v', want error: '%v'", input, got, err, want)
	}
}
