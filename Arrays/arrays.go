package main

import "fmt"

func main() {
	fmt.Println("Arrays in the Golang")

	// This is the Basic Declaration of the array
	var fruitlist [4]string
	// Add some data into the Array
	fruitlist[0] = "apple"
	fruitlist[1] = "mango"
	fruitlist[3] = "grapes"

	fmt.Println("The fruit list is:", fruitlist)
	fmt.Printf("Type of the fruit list is %T \n", fruitlist)
	//Type of the fruit list is [4]string

	//Another way of declaration of the array
	var veg_list = [3]string{"tomato", "beens"}
	fmt.Println("vegtable list is:", veg_list)
	fmt.Printf("Type of the veg list is %T \n", veg_list)
	//Type of the veg list is [4]string

	//Array declaration using the Make() function
	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&size)
	arr := make([]string, size)
	fmt.Println("Enter the string values into the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)

	//////////////////////////////////////////////////////
	// NEW PRACTICE
	/////////////////////////////////////////////////////

	fmt.Println("Arrays in the Golang")
	var size_1 int
	fmt.Println("Enter the size of the array:")
	fmt.Scan(&size_1)
	arr1 := make([]int, size_1)
	Elements_array := Enter_Elements_into_Array(arr1)
	total := Sum_Elemts(Elements_array)
	fmt.Println("Count is :", total)
	Reversed_array := Reversed(Elements_array)
	fmt.Println("Reversed array is :", Reversed_array)

	//////////////////////////////////////////////////
	// 2D array
	/////////////////////////////////////////////////
	///2d Array

	fmt.Println("2D array in the golang")
	fmt.Println("enter how amny rows:")
	var row int
	fmt.Scan(&row)
	fmt.Println("enter how many colomns")
	var col int
	fmt.Scan(&col)

	//Creating the 2d array
	arr_2d := make([][]int, row)
	for i := range arr_2d {
		arr_2d[i] = make([]int, col)
	}

	//Take elements into the 2d array
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scan(&arr_2d[i][j])
		}
	}
	fmt.Println(arr_2d)

	//Sum of the all elements in the 2D array
	var sum = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum = sum + arr_2d[i][j]
		}
	}
	fmt.Println(sum)

	//Sum diagonal elements
	sumdiagonal := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == j {
				sumdiagonal = sumdiagonal + arr_2d[i][j]
			}
		}
	}
	fmt.Println(sumdiagonal)

}

////////////////////////////////////////////////////////////

func Enter_Elements_into_Array(arr []int) []int {
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}
	return arr
}

func Sum_Elemts(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	return sum
}
func Reversed(arr []int) []int {
	n := 0
	for range arr {
		n++
	}
	res := make([]int, n)
	// for i, _ := range arr {
	// 	res[i] = arr[n-i-1]
	// }
	for i := 0; i < n; i++ {
		res[i] = arr[n-i-1]
	}
	return res
}
