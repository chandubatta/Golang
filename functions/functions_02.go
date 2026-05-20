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


//Name result parameters- Mostly used for the improving the readability of your function definition
func NameResultParam()(x,y int){
	//x=0,y=0
	return // 0, 0

	x= 10
	return //10,0

	return 1, 2
}

//Name result parameters- Mostly used for the improving the readability of your function definition
type Gogo interface{
	MyFunc()(x,y int)
}

////////////////////////////////////////////////////
// Function as Types
///////////////////////////////////////////////////
//We are passing the function as a 'Input Parameter' and 'Result Parameter'
func FunctionParam(f func(i int) int) func(s string) string{
	...
}

double_01:= func(i int)int{
			return i*2
		}

func CallFuncWithFunctionParam(){
//we can call a 'function' with 'function parameter' 3 ways
        //1st way is an 'inline' anonimous function
		f1:= FunctionParam(func(i int) int{
			return i++
		})

		//2nd way
		double:= func(i int)int{
			return i*2
		}
		f2:= FunctionParam(double)

		//3rd way
		f3:= FunctionParam(double_01)

}


////////////////////////////////////////////////////
// Generics in Functions
///////////////////////////////////////////////////
//Go Supports Generics in the Function Definitions, We specified in the [] Square brackets

func Min[T cmp.Ordered] (a,b T) T{
	if a<b{
		return a
	} 
	return b
}

func [a,b any] (a A,b B) (A,B) {
	return a, b
}


////////////////////////////////////////////////////
// Custom Cconstraints
///////////////////////////////////////////////////
func Add[N Number] (x,y N) N {
	return x+y
}

type Number interface{
	~int | ~float64
}

type MyType int