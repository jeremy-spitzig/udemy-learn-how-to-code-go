package main

import "fmt"

func main() {
	info := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(info)
	for key, record := range info {
		fmt.Println(key)
		for _, thing := range record {
			fmt.Println("\t", thing)
		}
	}
}
