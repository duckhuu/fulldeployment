package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func isPrime(n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i*j == n {
				return true
			}
		}
	}
	return false
}

type StructData struct {
	Data []string `json:"Data"`
}

func GetData() {
	var myjson = `
	{
		"Learner-0001": [
			"Course-0001",
			"Course-0002",
			"Course-0003"
			],
		"Learner-0002": [
			"Course-0002",
			"Course-0003",
			"Course-0004"
			]
	}
	`
	var data map[string]interface{}
	json.Unmarshal([]byte(myjson), &data)
	//	empArray := "[{\"abc\":\"abc\"},{\"def\":\"def\"}]"

	var unpacked []string
	// iterate through our results map to get the 'next level' of our JSON
	for _, i := range data {
		// marshal our next level as []byte
		stringified, err := json.Marshal(i)
		if err != nil {
			fmt.Println("error :", err)
		}
		// add item to our unpacked variable as a string
		unpacked = append(unpacked, string(stringified))
	}
	//	array := unpacked
	mid := len(unpacked) / 2
	arr1 := unpacked[:mid]
	arr2 := unpacked[mid:]
	for i, r := range arr1 {
		r = strings.ReplaceAll(r, "[", "")
		r = strings.ReplaceAll(r, "]", "")
		arr1[i] = string(r)
		arr1 = strings.Split(arr1[i], ",")
		for j, v := range arr2 {
			v = strings.ReplaceAll(v, "[", "")
			v = strings.ReplaceAll(v, "]", "")
			arr2[j] = string(v)
			arr2 = strings.Split(arr2[j], ",")
		}
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("array1:", i, arr1[i])
		fmt.Println("array2:", i, arr2[i])
	}

	// for u := 0; u < len(arr1); u++ {
	// 	for y := 0; y < len(arr2); y++ {
	// 		if arr1[u] != arr2[y] {
	// 			fmt.Println("array1:", arr1[u])
	// 			// 			//			fmt.Println(arr2[u])
	// 		}
	// 	}
	// }

	// array := make([]string, len(unpacked))
	// //	array1 := make([]string, len(unpacked[1]))
	// for i, r := range unpacked {
	// 	array[i] = string(r)
	// 	fmt.Println(array[i])
	// 	array = strings.Split(array[i], ",")
	// 	for m := 0; m < len(array); m++ {
	// 		fmt.Println(array[m])
	// 	}
	// }
	// for j, v := range unpacked[1] {
	// 	array1[j] = string(v)
	// 	array1 = strings.Split(array1[j], ",")
	// }

}

func main() {
	GetData()
}
