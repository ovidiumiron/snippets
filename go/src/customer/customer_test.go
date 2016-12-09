package customer

import (
	"errors"
	"reflect"
	"testing"
)

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

	errors := []Error{Error{2, errors.New("b")}}

	got, err := ReadFromFile(input)
	if !reflect.DeepEqual(got, customers) && reflect.DeepEqual(err, errors) {
		t.Errorf("aaa")
	}
}

// Skip lines can not have all require fiels.
func TestReadFromFileSkipMissingFields(t *testing.T) {
	input := "missing_fields.txt"
	customers := []Customer{Customer{User_id: 4, Name: "Alice Cahill", Latitude: "51.92893", Longitude: "-10.27699"}}
	errors := []Error{
		Error{1, errors.New("b")}, Error{1, errors.New("b")},
		Error{1, errors.New("b")}, Error{1, errors.New("b")}}

	got, err := ReadFromFile(input)

	if !reflect.DeepEqual(got, customers) || !reflect.DeepEqual(err, errors) {
		t.Errorf("aaa")
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
