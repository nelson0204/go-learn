package main

import "fmt"

func main() {
	age := 32 //regular variable
	var agePointer *int
	agePointer = &age

	// fmt.Println("Age", age)
	// fmt.Println("Age", agePointer)
	fmt.Println("Age", *agePointer) //dereferencing this pointer

	// adultYears := getAdultYears(agePointer) // variable when function return value and not change value in memory
	// fmt.Println(adultYears)

	editAgeAdultYears(agePointer) // don't need variable because se value of age changed in memory
	fmt.Println(age)              // age was changed when getAdultYears called
}

// Age is duplicated in memory
//	func getAdultYears(age int) int {
//		return age - 18
//	}

// Age is a address of memory and function return a new value of calculation
// func getAdultYears(age *int) int {
// 	return *age - 18
// }

// Age is a address of memory and function change the value in memory of pointer
func editAgeAdultYears(age *int) {
	*age = *age - 18 //dereferencing pointer age and change your value in memory
}
