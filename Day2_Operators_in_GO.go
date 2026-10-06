package main
import "fmt"

func main() {
	fmt.Printf("Operators\n")
	var num int
	var num2 int 
	fmt.Printf("Input for num: ")
    fmt.Scan(&num)
	fmt.Printf("Input for num2: ")
	fmt.Scan(&num2)
	// Arithmetic operators
	fmt.Println("\n--- Arithmetic Operators ---")
	fmt.Printf("num + num2 = %d\n", num+num2)
	fmt.Printf("num - num2 = %d\n", num-num2)
	fmt.Printf("num * num2 = %d\n", num*num2)
	fmt.Printf("num / num2 = %d\n", num/num2)
	fmt.Printf("num %% num2 = %d\n", num%num2)

	// Comparison operators
	fmt.Println("\n--- Comparison Operators ---")
	fmt.Printf("num == num2 : %t\n", num == num2)
	fmt.Printf("num != num2 : %t\n", num != num2)
	fmt.Printf("num > num2  : %t\n", num > num2)
	fmt.Printf("num < num2  : %t\n", num < num2)
	fmt.Printf("num >= num2 : %t\n", num >= num2)
	fmt.Printf("num <= num2 : %t\n", num <= num2)

	// Logical operators
	fmt.Println("\n--- Logical Operators ---")
	fmt.Printf("(num > 0) && (num2 > 0) : %t\n", (num > 0) && (num2 > 0))
	fmt.Printf("(num > 0) || (num2 > 0) : %t\n", (num > 0) || (num2 > 0))
	fmt.Printf("!(num > 0)              : %t\n", !(num > 0))

	// Bitwise operators
	fmt.Println("\n--- Bitwise Operators ---")
	fmt.Printf("num & num2 = %d\n", num&num2)
	fmt.Printf("num | num2 = %d\n", num|num2)
	fmt.Printf("num ^ num2 = %d\n", num^num2)
	fmt.Printf("num << 1   = %d\n", num<<1)
	fmt.Printf("num >> 1   = %d\n", num>>1)
}
