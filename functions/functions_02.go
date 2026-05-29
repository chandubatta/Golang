package functions

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

func [A,B any] (a A,b B) (A,B) {
	return a, b
}


////////////////////////////////////////////////////
// Custom Constraints
///////////////////////////////////////////////////
//func Add(x,y int/fliat64)int/fliat64{      // Don't do this like, for the reference declare this like.
//	....
//}

func Add[N Number] (x,y N) N {
	return x+y
}

type Number interface{
	~int | ~float64
}

type MyType int


func CallGenerics(){
	//Normal
	x:=Add(1,2)

	//Explicitly
	y:=Add[int](1,2)
}


////////////////////////////////////////////////////
// Methods
///////////////////////////////////////////////////
type MyType struct{}

	
func (t MyType) MyMethod(){
	...
}

func CallMethod(){
	m:= MyType{}
	m.MyMethod
}


////////////////////////////////////////////////////
// Deferring Function Execution
///////////////////////////////////////////////////
func Deferring(){
	defer fmt.Println("Printls Last")
	defer fmt.Println("Printls Second")
	fmt.Println("Printls First")
}



////////////////////////////////////////////////////
// Defer Gotcha
///////////////////////////////////////////////////
func DeferGGotcha(){
	var e error // nil
	// defer func (err error){      //Eanonimous Error
	// 	if err != nil{
	// 		fmt.Println(err)
	// 	}
	// }(e) //This is Pattern used a ton in go

	defer func (){      //This anotheway of writing Up side Eanonimous Error
		if e != nil{
			fmt.Println(e)
		}
	}() //This is Pattern used a ton in go

	val, ok := TrySomething()
	if  !ok {
		e = errors.New(Something Broken)
	}

}

////////////////////////////////////////////////////
// Modify Return in Defer
///////////////////////////////////////////////////

//Return 11
func DeferModifiedReturn()x int{
	defer func (){
		x ++
	}()

	return 10
}