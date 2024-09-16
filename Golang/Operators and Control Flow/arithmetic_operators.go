package main
import "fmt"

func main(){
	var a,b string = "foo", "bar"
	fmt.Println(a + b)

	var c, d float64 = 79.02, 75.66
	fmt.Println(c - d)

	var e, f int = 12, 2
	fmt.Println(e * f)

	var g, h int = 24, 2
	fmt.Println(g / h)

	var i, j int = 24, 7
	fmt.Println(i % j)

	var k int = 1
	k++
	fmt.Println(k)

	var l int = 1
	l--
	fmt.Println(l)
}