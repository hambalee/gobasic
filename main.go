package main

func main() {
	countries := map[string]string{}
	countries["th"] = "Thailland"

	country, ok := countries["th"]
	if ok {
		println(country)
	} else {
		println("no value")
	}

	//loop
	values := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(values); i++ {
		println(values[i])
	}
	//for range
	for i, v := range values {
		println(i, v)
	}
	for _, v := range values {
		println(v)
	}

	//function
	c := sum(10, 20)
	println(c)
	c, x := sum2(10, 20)
	println(c, x)

	//anonymouse function
	z := func(a, b int) int {
		return a + b
	}
	y := z(1, 2)
	println(y)

	//function as parameter
	cal(sum)
	cal(add)
	cal(sub)
	cal(z)
	cal(func(a, b int) int {
		return a + b
	})

	//variodic
	v := []int{1, 2, 3}
	s := sum3(v)
	println("variodic")
	println(s)
	println("parameter array")
	s2 := sum4(1, 2, 3, 4, 5)
	println(s2)
}
func sum3(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
func sum4(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func sum(a, b int) int {
	return a + b
}
func sum2(a, b int) (int, string) {
	return a + b, "Hello"
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func cal(f func(int, int) int) {
	sum := f(50, 10)
	println(sum)
}
