package customer

import "testing"

func notEq(x, y []Customer) bool {
	if len(x) != len(y) {
		return true
	}
	for i := range x {
		if x[i] != y[i] {
			return true
		}
	}
	return false
}

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
		/*Skip lines can not be unmarshall.*/
		{"not_unmarshall.txt", []Customer{
			Customer{User_id: 1, Name: "Alice Cahill", Latitude: "51.92893", Longitude: "-10.27699"}}},
		/*Skip lines can not have all require fiels.*/
		{"missing_fields.txt", []Customer{
			Customer{User_id: 4, Name: "Alice Cahill", Latitude: "51.92893", Longitude: "-10.27699"}}},
	}
	for _, test := range tests {
		if got, err := ReadFromFile(test.input); err != nil || notEq(got, test.want) {
			t.Errorf("ReadFromFile(%v) = %v %v, want %v", test.input, got, err, test.want)
		}
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
