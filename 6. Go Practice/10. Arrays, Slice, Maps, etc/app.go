package main

import "fmt"

//example of array

func main() {
	var x [5]int
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50
	fmt.Println(x)
	fmt.Println(x[3])
	fmt.Println(x[1:3])
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0:5])
	fmt.Println(x[0:4])
	fmt.Println("---")
	fmt.Println(x[:4])
	fmt.Println(x[0:3])
	fmt.Println(x[0:2])
	fmt.Println(x[0:1])
	fmt.Println(x[0:0])

	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println(y)
	fmt.Println(y[3])
	fmt.Println(y[1:3])
	fmt.Println(len(y))
	fmt.Println(cap(y))
	fmt.Println(y[0:5])
	fmt.Println(y[0:4])

	// slices are like references to arrays so if we change the value of slice it will change the value of array

	// len is the number of elements in the slice
	// cap is the number of elements in the underlying array starting from the index from which the slice is created
	// alway can select more from the slice than the length of the slice but not more than the capacity of the slice and not  more to the right of the slice
	// arrays are fixed size but slices are dynamic

	fmt.Println("----")

	z := []int{10, 20, 30, 40, 50}
	fmt.Println(z)
	fmt.Println(z[3])
	fmt.Println(z[1:3])
	fmt.Println(len(z))
	fmt.Println(cap(z))
	fmt.Println(z[0:5])
	z = append(z, 60)
	fmt.Println(z)
	// to remove an element from the slice
	z = z[:2]
	fmt.Println(z)
	z = append(z, 40, 50, 60)
	fmt.Println(z)
	z = append(z, y[:]...)
	fmt.Println(z)

// maps in go

	fmt.Println("----")

websites := map[string]string{	// map[key]value
	"google": "https://www.google.com",
	"facebook": "https://www.facebook.com",
	"twitter": "https://www.twitter.com",

}

fmt.Println(websites)
fmt.Println(websites["google"])
fmt.Println(websites["facebook"])
fmt.Println(websites["twitter"])

websites["instagram"] = "https://www.instagram.com"
fmt.Println(websites)

delete(websites, "facebook")
fmt.Println(websites)

delete(websites, "facebook")
fmt.Println(websites)

// in maps any type can be used as key or value like int, string, float, bool, struct, array, slice, pointer, interface, etc.

// using make function to create a map to preallocate the space for the map

fmt.Println("----")

m := make([]string, 3)
m[0] = "google"
m[1] = "facebook"
m[2] = "twitter"

fmt.Println(m)

// using make function to create a map to preallocate the space for the map

fmt.Println("----")

n := make(map[string]string,3)
n["google"] = "https://www.google.com"
n["facebook"] = "https://www.facebook.com"
n["twitter"] = "https://www.twitter.com"

fmt.Println(n)

// we can also do aliases of more complex types like map[string]string
// type myMap map[string]string
// var myMap myMap
// myMap = make(myMap)
// myMap["google"] = "https://www.google.com"
// fmt.Println(myMap)

// using loops

fmt.Println("----")

for i := 0; i < 5; i++ {
	fmt.Println(i)

}

fmt.Println("----")

for index, value := range x {
	fmt.Println(index, value)
}


	
}