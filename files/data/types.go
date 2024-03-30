package data

type distance float64
type distanceKm float64

func (miles distance) ToKM() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(4.5)
	km := d.ToKM()
	km.ToMiles()

	print(d)
}
