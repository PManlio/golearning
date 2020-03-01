package main

import "fmt"

func main() {
	fmt.Println("#### examples about map, slices and pointers ####")

	// tip: an array can be without name, like [dim_arr]type{}
	// you can use this nameless array inside function
	for i := 0; i < 3; i++ {
		showAr([3]int{i, i * i, i - 1})

		showArPointers(&[3]int{i, i * i, i - 1})
	}

	showSlices()

	//three different ways to create a map
	mappa1 := map[string]float64{"x": 1, "y": 2, "z": 3}
	var mappa2 = map[string]float64{"a": 1, "b": 2, "c": 3}
	var mappa3 = make(map[string]float64) // "make", not "new"

	fmt.Println(mappa1, mappa2, mappa3) // mappa3 will be empty

	fmt.Println(mappa1["y"], mappa2["c"])
	// fmt.Println(mappa1[2], mappa3[3]) this is an error
	// because you have to access to a map member via key, not map value nor index

	deleteMapMember(mappa2)
}

func deleteMapMember(mappa map[string]float64) {
	delete(mappa, "b")
	// didn't find a way to delete map element via,for example, map index
	fmt.Println(mappa)
}

func showAr(a [3]int) {
	fmt.Println(a)
}

func showArPointers(ap *[3]int) {
	fmt.Println(ap)
}

func showSlices() {
	array1 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceArr1 := array1[0:5] // you can write also sliceArr1 := array1[:5]
	sliceArr1 = append(sliceArr1, 5)
	fmt.Println("\n", sliceArr1)

	//	sliceArr1 := array1[:] // means [0:len(array1)] (see below)
	var slice2 []int // slices can be initially created like arrays with no dimension
	slice2 = array1[:]
	fmt.Println("\n", slice2)
}
