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
}
