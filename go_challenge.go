package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type StructData struct {
	Data []string `json:"Data"`
}

func main() {
	jsonString := `[
		["A", "B", "B"],
		["C", "B", "C"],
		["A", "B", "C"],
		["B", "B", "A"],
		["A", "B", "C"]
	]`

	var data [][]string
	//var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked []string
	// iterate through our results map to get the 'next level' of our JSON
	for _, i := range data {
		// marshal our next level as []byte
		stringified, err := json.Marshal(i)
		if err != nil {
			fmt.Println("error :", err)
		}
		// add item to our unpacked variable as a string
		unpacked = append(unpacked, string(stringified[:]))
	}
	//	var splitArrays []string
	var array1 []string
	//fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(splitArrays))
	for i, v := range unpacked {
		//	fmt.Println(v)
		v = strings.ReplaceAll(v, "[", "")
		v = strings.ReplaceAll(v, "]", "")
		array1[i] = v
		//		array1 = strings.Split(array1[1], ",")
		//		fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(splitArrays[i]))
		//		fmt.Println(array1)
		//		fmt.Println(splitArrays)
	}
	// for i, v := range splitArrays {
	// 	array1[i] = string(v)
	// 	array2 = strings.Split(array1[i], ",")
	// 	fmt.Println(array2)
	// 	fmt.Printf("unpacked[3]: %s\n", reflect.TypeOf(array2))
	// 	//		fmt.Println(splitArrays[i])
	// }
	// fmt.Println(splitArrays[0][0])
	// fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(splitArrays[1][0]))
	// for i := 0; i < len(splitArrays[i][0]); i++ {
	// 	fmt.Println(splitArrays[i][0])
	// }
}
