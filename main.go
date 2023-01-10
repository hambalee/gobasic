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
}

