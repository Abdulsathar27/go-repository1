package main
import "fmt"
func main(){
	var num1,num2 int
	var Operator string

	//choosing frist number 
	
	fmt.Println("Enter Frist Number:")
	_,err1 := fmt.Scan(&num1)
	if err1 != nil {
		fmt.Println("Please Enter the Number")
		return
	}

	// Choosing the Operator
	fmt.Println("Choose Operator (+,-,*,/):")
	_,err2 := fmt.Scan(&Operator)
	if err2 !=nil {
		fmt.Println("Allow the Operator input.")
		return
	}

	//choosing second Number
	fmt.Println("Enter Second Number")
	_,err3 := fmt.Scan(&num2)
	if err3 !=nil{
		fmt.Println("Please Enter the Number.")
		return
	}

	//calculation
	switch Operator{
	case "+":
		fmt.Printf("%d + %d = %d\n" ,num1,num2,num1+num2)
	case"-":
		fmt.Printf("%d - %d = %d\n" ,num1,num2,num1-num2)
	case"*":
		fmt.Printf("%d * %d = %d\n" ,num1,num2,num1*num2)
	case"/":
		if num2 == 0{
			fmt.Println("Error:Division by zero is not allowed.")
		}else{
			fmt.Printf("%d/%d=%d\n",num1,num2,num1/num2)
		}
		default :
		fmt.Println("Operator are not use .Please use +,-,*,or/.")
	}




}
