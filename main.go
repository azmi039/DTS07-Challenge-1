package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%d\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")
	j := true
	fmt.Printf("%t\n", j)
	fmt.Printf("\n")

	fmt.Printf("%b \n", i)
	fmt.Printf("%c\n", '\u042F')
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	f := 15
	fmt.Printf("%x\n", f)
	fmt.Printf("%X\n", f)

	r := int32('Я')
	u := fmt.Sprintf("%U", r)
	fmt.Println(u)
	fmt.Printf("\n")

	k := 123.456
	fmt.Printf("%f\n", k)
	fmt.Printf("%e\n", k)
}
