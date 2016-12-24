package customer

import (
	"errors"
	"reflect"
	"testing"
)

// Example how to use an io.Reader. Kept just for further references.
/*
type Reader struct {
	read string
	done bool
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
	}

	for i, b := range []byte(r.read) {
		p[i] = b
	}
	r.done = true
	return len(r.read), nil
}

func NewReader(toRead string) *Reader {
	return &Reader{toRead, false}
}

func TestReadFunction(t *testing.T) {
	var tests = []struct {
		input string
		want  []Customer
	}{
		{
			"{\"latitude\": \"52.986375\", \"user_id\": 12, \"name\": \"Christina McArdle\", \"longitude\": \"-6.043701\"}",
			[]Customer{
				Customer{User_id: 12, Name: "Christina McArdle", Latitude: "52.986375", Longitude: "-6.043701"}}},
	}
	for _, test := range tests {
		if got, err := read(NewReader(test.input)); err != nil || !reflect.DeepEqual(got, test.want) {
			t.Errorf("read(%v) = %v %v, want %v", test.input, got, err, test.want)
		}
	}
}
*/

func TestReadFromFilePositiveTests(t *testing.T) {
	var tests = []struct {
		input string
		want  []Customer
	}{
		/*Clean input.*/
		{"noerrors.txt", []Customer{
			Customer{User_id: 12, Name: "Christina McArdle", Latitude: "52.986375", Longitude: "-6.043701"},
			Customer{User_id: 1, Name: "Alice Cahill", Latitude: "51.92893", Longitude: "-10.27699"}}},
		/*Can handle unicode.*/
		{"unicode_name.txt", []Customer{
			Customer{User_id: 3, Name: " 世界", Latitude: "33", Longitude: "333"}}},
	}
	for _, test := range tests {
		if got, err := ReadFromFile(test.input); err != nil || !reflect.DeepEqual(got, test.want) {
			t.Errorf("ReadFromFile(%v) = %v %v, want %v", test.input, got, err, test.want)
		}
	}
}

// Skip lines can not be unmarshall.
func TestReadFromFileSkipeUnmarshall(t *testing.T) {
	input := "not_unmarshall.txt"
	customers := []Customer{Customer{User_id: 1, Name: "Alice Cahill", Latitude: "51.92893", Longitude: "-10.27699"}}

	errors := []Error{Error{2, errors.New("customer: missing at least one field")}}

	got, err := ReadFromFile(input)
	if !reflect.DeepEqual(got, customers) && reflect.DeepEqual(err, errors) {
		t.Errorf("ReadFromFile(%v) = %v %v, want errors %v", input, got, err, errors)
	}
}

// Skip lines can not have all require fiels.
func TestReadFromFileSkipMissingFields(t *testing.T) {
	input := "missing_fields.txt"
	customers := []Customer{Customer{User_id: 4, Name: "Alice Cahill", Latitude: "51.92893", Longitude: "-10.27699"}}
	errors := []Error{
		Error{1, errors.New("customer: missing at least one field")}, Error{1, errors.New("customer: missing at least one field")},
		Error{1, errors.New("customer: missing at least one field")}, Error{1, errors.New("customer: missing at least one field")}}

	got, err := ReadFromFile(input)

	if !reflect.DeepEqual(got, customers) || !reflect.DeepEqual(err, errors) {
		t.Errorf("ReadFromFile(%v) = %v %v, want errors %v", input, got, err, errors)
	}
}

func TestReadFromFileNegativeTests(t *testing.T) {
	var tests = []string{"file_not_exist.txt"}

	for _, test := range tests {
		if got, err := ReadFromFile(test); err == nil {
			t.Errorf("ReadFromFile(%v) = %v %v, want error", test, got, err)
		}
	}
}
