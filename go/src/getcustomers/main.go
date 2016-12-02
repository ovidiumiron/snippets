package main

import (
	"customer"
	"fmt"
	"gps"
	"log"
	"sort"
)

func main() {
	// Path to the file with customers.
	const path = "customers.txt"
	// Accepted distance in meters.
	const distance = 100 * 1000

	intercom := struct {
		lat  float64
		long float64
	}{53.3393, -6.2576841}

	customers, err := customer.ReadFromFile(path)
	for _, e := range err {
		if e.Level == customer.Fatal {
			log.Fatal(e.Err.Error())
		}
		log.Println(e)
	}

	sorted := make(customer.ById, 0)
	for _, c := range customers {
		lat, err := c.Lat()
		if err != nil {
			// Can not use latitude skip over the client.
			log.Println(err)
			continue
		}
		long, err := c.Long()
		if err != nil {
			// Can not use longitude skip over the client.
			log.Println(err)
			continue
		}

		if gps.Distance(intercom.lat, intercom.long, lat, long) <= distance {
			sorted = append(sorted, c)
		}
	}
	sort.Sort(customer.ById(sorted))
	for _, customer := range sorted {
		fmt.Println(customer.Print_IdName())
	}
}
