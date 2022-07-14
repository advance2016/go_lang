package main

import fmt

func main()  {
	var a [3]int

	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])

	for i, v := range a {
		fmt.Println("%d %d", i, v)
	}
	
	for _, v := range a {
		fmt.Println("%d", v)
	}
}