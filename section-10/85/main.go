package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "Miss Moneypenny")
	fmt.Println(m)
}
