package main

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no explicit return value
}
