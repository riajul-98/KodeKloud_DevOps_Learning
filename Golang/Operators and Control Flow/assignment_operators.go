package main
import "fmt"

func main(){
	var x int = 10
	var y int
	y = x
	fmt.Println(y)

	var a, b int = 10, 20
	a += b
	fmt.Println(a)

	var c, d int = 10, 20
	c -= d
	fmt.Println(c)

	var e, f int = 10, 20
	e *= f
	fmt.Println(e)

	var g, h int = 200, 20
	g /= h
	fmt.Println(h)

	var i, j int = 210, 20
	i %= j
	fmt.Println(i)
}