package customer

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Customer struct {
	User_id   int
	Latitude  string
	Longitude string
	Name      string
}

func (c *Customer) Lat() (float64, error) {
	return strconv.ParseFloat(c.Latitude, 64)
}

func (c *Customer) Long() (float64, error) {
	return strconv.ParseFloat(c.Longitude, 64)
}

func (c *Customer) String() string {
	return fmt.Sprintf("{id: %d name: %s latitude: %s longitude: %s}\n",
		c.User_id, c.Name, c.Latitude, c.Longitude)
}

func (c *Customer) Print_IdName() string {
	return fmt.Sprintf("{id: %d name: %s}", c.User_id, c.Name)
}

type Error struct {
	Level int
	Err   error
}

const (
	Fatal = iota
	Warning
)

func read(r io.Reader) ([]Customer, []Error) {
	var errs []Error
	scanner := bufio.NewScanner(r)
	var c struct {
		User_id   *int
		Latitude  *string
		Longitude *string
		Name      *string
	}

	var customers = make([]Customer, 0)

	for scanner.Scan() {
		c.User_id, c.Name, c.Latitude, c.Longitude = nil, nil, nil, nil
		// If can not unmashall skip the line.
		if err := json.Unmarshal(scanner.Bytes(), &c); err != nil {
			errs = append(errs, Error{Warning, err})
			continue
		}
		// If there are missing field skip the line.
		if c.User_id == nil || c.Latitude == nil || c.Longitude == nil || c.Name == nil {
			errs = append(errs, Error{Warning, errors.New("b")})
			continue
		}

		customers = append(customers, Customer{*c.User_id, *c.Latitude, *c.Longitude, *c.Name})
	}
	return customers, errs
}

// Return error only if can not open file.
func ReadFromFile(path string) ([]Customer, []Error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, []Error{Error{Fatal, err}}
	}
	defer file.Close()

	return read(file)
}

// Sort by id.
type ById []Customer

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].User_id < a[j].User_id }
