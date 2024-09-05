package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func subslices(slice []string) [][]string {
	var ss [][]string
	for _, e := range slice {
		if strings.HasPrefix(e, "101") || len(ss) == 0 {
			ss = append(ss, make([]string, 0, 3))
		}
		end := len(ss) - 1
		ss[end] = append(ss[end], e)
	}
	return ss
}

func splitSlice(slice []string, size int) [][]string {
	var result [][]string
	for size < len(slice) {
		slice, result = slice[size:], append(result, slice[0:size:size])
	}
	result = append(result, slice)
	return result
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

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err)
	}
	var unpacked []string
	for _, i := range data {
		// marshal our next level as []byte
		stringified, err := json.Marshal(i)
		if err != nil {
			fmt.Println("error :", err)
		}
		// add item to our unpacked variable as a string
		unpacked = append(unpacked, string(stringified[:]))
	}
	split := splitSlice(unpacked, 1)
	//	fmt.Println(split)
	//	fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(unpacked))
	//	var array []string
	//	for i, r := range split {
	var array [3]string
	for j := 0; j < len(split); j++ {
		//	fmt.Println("mang split", j, split[j])
		for j, r := range split[j] {
			r = strings.ReplaceAll(r, "[", "")
			r = strings.ReplaceAll(r, "]", "")
			array[j] = string(r)
			fmt.Println(array[j])
			fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(array[j]))
			array = [3]string(strings.Split(array[j], ","))
			//	fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(r))
			fmt.Printf("unpacked[3]: %s\n", reflect.TypeOf(array))
			fmt.Println(array)
		}
	}
	//	fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(r))
	//	fmt.Println(r)
	//	fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(r))
	//	fmt.Println(split[i])
}

//	}
//	fmt.Println(len(unpacked))
// fmt.Println(unpacked[0])
// result := make([]string, 4)
// fmt.Println(len(result[2]))
// mid := len(unpacked) / len(unpacked)
// arr1 := unpacked[:mid]
// fmt.Println(arr1)
// arr2 := unpacked[mid::,:]
// fmt.Println(arr2)
//	var array []string
// for i, r := range unpacked {
// 	r = strings.ReplaceAll(r, "[", "")
// 	r = strings.ReplaceAll(r, "]", "")
// 	// 	//	mid := len(unpacked) / len(unpacked)
// 	array[i] = string(r)
// 	array = strings.Split(array[i], ",")
// 	fmt.Println(array[i])
// 	fmt.Println(unpacked)
//	array = strings.Split(array[i], ",")
// 	//		fmt.Println(array)
//	}
// arr1 := unpacked[:mid]
// arr2 := unpacked[mid:]
// for i, r := range arr1 {
// 	r = strings.ReplaceAll(r, "[", "")
// 	r = strings.ReplaceAll(r, "]", "")
// 	arr1[i] = string(r)
// 	arr1 = strings.Split(arr1[i], ",")
// 	for j, v := range arr2 {
// 		v = strings.ReplaceAll(v, "[", "")
// 		v = strings.ReplaceAll(v, "]", "")
// 		arr2[j] = string(v)
// 		arr2 = strings.Split(arr2[j], ",")
// 	}
// }
// for i := 0; i < len(arr1); i++ {
// 	fmt.Println("array1:", i, arr1[i])
// 	fmt.Println("array2:", i, arr2[i])
// }
