package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (by ByAge) Len() int           { return len(by) }
func (by ByAge) Less(i, j int) bool { return by[i].Age < by[j].Age }
func (by ByAge) Swap(i, j int)      { by[i], by[j] = by[j], by[i] }

type ByLast []user

func (bL ByLast) Len() int           { return len(bL) }
func (bL ByLast) Less(i, j int) bool { return bL[i].Last < bL[j].Last }
func (bL ByLast) Swap(i, j int)      { bL[i], bL[j] = bL[j], bL[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("Sort by sayings ----")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("Sort by Age ----")
	fmt.Println(users)

	sort.Sort(ByAge(users))
	fmt.Println(users)

	fmt.Println("Sort by Last ----")
	fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println(users)

}
