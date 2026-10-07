package main
import "fmt"

func main() {
	fmt.Printf("Calculator\n")
	var num int 
    var into string 
	var num2 int 
	fmt.Printf("Enter Number : ")
	fmt.Scanf("%d",&num)
	fmt.Printf("Enter Number 2 : ")
	fmt.Scanf("%d",&num)
	fmt.Printf("Enter the Operation : ")
	fmt.Scanf("%s",&into)

	if into == "+"{
		fmt.Printf("The addition is of numbers is : %d",num + num2)
	} else if into == "-"{
		fmt.Printf("The subtraction is of numbers is : %d",num - num2)
	} else if into == "*"{
		fmt.Printf("The Multiplication is of numbers is : %d",num*num2)
	} else if into == "/"{
		fmt.Printf("The Divition is of numbers is : %d",num/num2)
	}
}
