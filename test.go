//go:build ignore

package main

import (
    f "fmt"
	m "math"
	_ "reflect"
	hu "local/main/humain"
)

func main() {
	f.Println("done")
}

func main1() {
	num := 1234
	check := 23
	e := 1

	for counter := 0; counter < 2; {
		num2 := num%int(m.Pow10(e))
		if num2 == num { counter++ }
		f.Printf("%v - %v\n", e, num2)
		f.Println("------------")
		e++
	}

	f.Println()
	e = e-2
	e2 := e
	num_list := make([]int, e)
	for i := 0; e > 0; i++ {
		num_list[i] = num%int(m.Pow10(e))
		f.Printf("%v - %v\n", e, num_list[i])
		f.Println("++++++++++++")
		e--
	}

	f.Println()
	subt_list := make([]int, e2)
	slli := len(subt_list) - 1
	nlli := len(num_list) - 1
	for i := 0; i < nlli; i++ {
		if i == 0 { subt_list[slli] = num_list[nlli] }

		subt_list[i] = num_list[i] - num_list[i+1]
		f.Printf("%v - %v\n", i, subt_list[i])
		f.Println("============")
	}


	ind_num_list := make([]int, e2)
	for i := 0; i < e2; i++ {
		ind_num_list[i] = subt_list[i]/int(m.Pow10(e2 - (i+1)))
		f.Println(ind_num_list)
		f.Println("''''''''''''")
	}

	e3 := e2-1
	i := 0
	k := 0
	j := e3
	for ; i < e3; i++ {
		for l := 0; l < e3; l++ {
			f.Println(ind_num_list[l]*int(m.Pow10(k+1)) + ind_num_list[l+1])
		}
		k++ 
		j--
	}
	return

	f.Println(ind_num_list[0]*10 + ind_num_list[1])
	f.Println(ind_num_list[1]*10 + ind_num_list[2])
	f.Println(ind_num_list[2]*10 + ind_num_list[3])

	f.Println(ind_num_list[0]*100 + ind_num_list[1]*10 + ind_num_list[2])
	f.Println(ind_num_list[1]*100 + ind_num_list[2]*10 + ind_num_list[3])
	
	num1 := num%1000
	num2 := num%100
	num3 := num%10
	f.Printf("%d - %d - %d - %d\n", num, num1, num2, num3)

	num11 := num-num1
	num22 := num1-num2
	num33 := num2-num3
	f.Printf("%d - %d - %d - %d - %d\n", num, num11, num22, num33, num3)

	value, type_ := hu.TrueType(float64(num22 + num33)/10)
	val, _ := value.(int);

	if type_ == "int" && val == check {
		f.Println("true")
	}
}
