////////////////////////////////////////////////////
// Basic Function Use
///////////////////////////////////////////////////

//No input parameters and No Return parameters
func Basic(){
	fmt.Println("Do Something")
}

//Function with Parameters. You must pass the parameters (No optinal)
func WithParams(i int, x,y string){ ... }

func CallingFunction(){
	Basic()
	WithParams(10,"1","2")

	//Not allow
	WithParams(i: 10, x: "1",y: "2") //invalid
}


// Variadic Parameters- ...
func Sum(nums ...int, x int)int{
	total:=0
	for index, num:= range nums{
		total=total+num
	}
	return total
}

func CallVariadic(){
	sum := Sum(1,2,3,4,5,6)

	nums:=[]int{1,2,3}
	sum:= Sum(nums...)
}

// Single ResultParameter

func SingleResultParameter(i int)int{
	return i+2
}

func MultipuleResultParameters(i int)(int,int){
	return i+2, i-2
}

func CallFuncWithResultParameter(){
	x := SingleResultParameter(10)
	y, z:= MultipuleResultParameters(10)
	_, k:= MultipuleResultParameters(10) //skipping the first valuess
}



