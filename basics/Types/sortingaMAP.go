package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"alice": 24,
		"marco": 32,
		"james": 27,
		"xiago": 24,
		"zeno":  22,
		"jenny": 28,
		"anna":  29,
	}

	fmt.Println(m) // unordered

	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, m[name])
	}

}
