package main

import "fmt"

func negative(arr []int, k int) {

	n := len(arr)
	i := 0

	j := 0

	temp := []int{}

	res := []int{}

	for j < n {
		if arr[j] < 0 {
			temp = append(temp, arr[j])
		}

		if j-i+1 < k {
			j++

		} else if j-i+1 == k {

			if len(temp) > 0 {
				negativeValue := temp[0]

				res = append(res, negativeValue)

				if arr[i] == negativeValue {
					temp = temp[1:]

				}

			}
			i++
			j++

		}

	}

	fmt.Println(res)
}
