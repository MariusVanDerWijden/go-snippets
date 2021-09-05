package terragen

import "testing"

func TestTerragen(t *testing.T) {
	field := createMap(20, 30, EARTH)
	field = makeCave(field, 15, 15, 12)
	field = makeCave(field, 10, 25, 12)
	field = addWater(field, 5, 20)
	field = addWater(field, 10, 8)
	printMap(field)
	panic("adsf")
}
