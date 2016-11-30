package gps

import "testing"

func TestDistance(t *testing.T) {
	type Point struct {
		lat  float64
		long float64
	}

	var tests = []struct {
		p1   Point
		p2   Point
		want int32
	}{
		{Point{38.898556, -77.037852}, Point{38.897147, -77.043934}, 549},
		{Point{52.986375, -6.043701}, Point{53.3393, -6.2576841}, 41755}}

	for _, test := range tests {
		got := Distance(test.p1.lat, test.p1.long, test.p2.lat, test.p2.long)
		if got != test.want {
			t.Errorf("Distance(%f, %f, %f, %f) rounded to %v, want %v",
				test.p1.lat, test.p1.long, test.p2.lat, test.p2.long, got, test.want)
		}
	}
}
