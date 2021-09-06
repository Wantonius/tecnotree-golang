package main

import "fmt"

func main() {

	//Introduce an array. It has a specific lenght and 
	//will be zero-valued unless initialized to something else

	var myArray [6]int

	fmt.Println("myArray:",myArray)
	fmt.Println("myArray length:",len(myArray))
	
	myArray[3] = 50
	fmt.Println("myArray again:",myArray)

	myInitializedArray := [3]int{0,1,2}
	
	fmt.Println("Initialized Array:",myInitializedArray)
	fmt.Println("---Slices---")
	
	//Introduce a slice. Slices are dynamic in length and support 
	//additional methods like append
	
	var mySlice []int //this has no memory allocated
	
	myAllocatedSlice := make([]int,10) //allocates a slice of initial size of ten. Values will be zero-valued

	fmt.Println("mySlice",mySlice)
	fmt.Println("mySlice size:",len(mySlice))

	fmt.Println("myAllocatedSlice",myAllocatedSlice)
	fmt.Println("myAllocatedSlice size:",len(myAllocatedSlice))
	
	//Next line will crash to program. mySlice has no memory allocated.
	//mySlice[0] = 0
	//You use append instead
	
	mySlice = append(mySlice,0)

	fmt.Println("mySlice again",mySlice)
	
	//When appending multiples or other slices use ... operator

	mySlice = append(mySlice,[]int{10,100}...)
	
	fmt.Println("mySlice again",mySlice)
	
	copiedSlice := make([]int,len(mySlice))

	copy(copiedSlice,mySlice)
	
	fmt.Println("Copied slice:",copiedSlice)
	
	partialSlice := mySlice[1:3]
	
	fmt.Println("Partial slice:",partialSlice)
	
	fmt.Println("---Maps---")

	//Maps are go's built-in associative data type, key-value pairs
	//or dictionaries. Use make to create empty maps

	intStringMap := make(map[int]string)
	stringIntMap := make(map[string]int)

	intStringMap[1] = "One"
	intStringMap[2] = "Two"
	
	stringIntMap["one"] = 1
	stringIntMap["two"] = 2

	fmt.Println("intStringMap",intStringMap)
	fmt.Println("stringIntMap",stringIntMap)
	fmt.Println("intStringMap value at 1",intStringMap[1])
	fmt.Println("stringIntMap value at one",stringIntMap["one"])

	//use delete to remove by key
	
	delete(stringIntMap,"two")
	fmt.Println("stringIntMap again",stringIntMap)
	
	//initializing map to init values
	
	initializedMap := map[int]string{1:"one",2:"two"}
	fmt.Println("initializedMap",initializedMap)
	
	//Checking if something exists within the map
	
	if val, ok := initializedMap[2]; ok {
		fmt.Printf("Initialized map contains %s\n",val)
	}
	if _, ok := initializedMap[3]; !ok {
		fmt.Println("Initialized map does not contain that key-value pair")
	}
}