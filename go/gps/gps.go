package gps

import "math"

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Return distance in meters between point A(lat1, long1) and point B(lat2, long2)
// using the radius of the Earth 6371 km.
func Distance(lat1, long1, lat2, long2 float64) int32 {
	const r float64 = 6371 * 1000

	// Convert to radians
	la1 := lat1 * math.Pi / 180
	lo1 := long1 * math.Pi / 180
	la2 := lat2 * math.Pi / 180
	lo2 := long2 * math.Pi / 180

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return int32(2 * r * math.Asin(math.Sqrt(h)))
}
