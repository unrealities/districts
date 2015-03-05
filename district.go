package main

type District struct {
	id       int
	name     string
	region   Region
	vertices []int
}

// func neighborDistrict(District d1, District d2) bool {
// 	neighbor := bool
//
// 	for key, value := range d1.vertices {
// 		//TODO
// 	}
// 	return neighbor
// }

// Given a slice of ints (vertices),
// return a slice of a slice of ints which represents the sides of the object
// func sides([]int vertices) [][]int{
//   numOfSides = len(vertices)
//   sides := [][]int
//   for key, vertex := range vertices {
//     sides
//   }
//   return sides
// }
