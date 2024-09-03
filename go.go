package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

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
		unpacked = append(unpacked, string(stringified[:]))
		//		var newarray []interface{} = []interface{}{string(stringified)}
		//		fmt.Println(newarray[0])
	}
	fmt.Println(unpacked)
	arrayofstring := strings.Split(unpacked[0], ",")
	fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(arrayofstring[0]))
	fmt.Println(arrayofstring[1])
	array1ofstring := strings.Split(unpacked[1], ",")
	fmt.Println(array1ofstring[2])
	fmt.Printf("unpacked[2]: %s\n", reflect.TypeOf(array1ofstring[0]))
	if arrayofstring[0] != array1ofstring[0] {
		fmt.Println(arrayofstring[0])
	}
	arr := make([][]string, len(unpacked))
	fmt.Println(unpacked[0])
	fmt.Println(unpacked[1])
	//	arr1 := make([][]string, len(unpacked[1]))
	for i, r := range unpacked {
		//		for j, v := range unpacked[1] {
		arr[i] = string(r)
		//			arr[j] = string(v)
		arr = strings.Split(arr[i], ",")
		//			arr1 = strings.Split(arr[j], "")
		fmt.Println("the new arr:", arr)
		//			fmt.Println("the new arr1:", arr1)
		fmt.Println(arr[i])
		//			fmt.Println(arr1[j])
		//			if arr[i] !=
		//		}
	}
	// fmt.Println("phan tu mang:", arr[0])
	// fmt.Println("phan tu mang:", arr[1])
	pr := &StructData{unpacked}
	prAsBytes, err := json.Marshal(pr)
	if err != nil {
		fmt.Println("error :", err)
	}
	fmt.Println(string(prAsBytes[:]))
	//	var array3 []string
	// for _, u := range unpacked {
	// 	if str, ok := u([]string); ok {
	// 		array3 = append(array3, str)
	// 	}
	// }
	//	fmt.Println(array3)
}

func main() {
	GetData()
	//	fmt.Printf("t1: %s\n", reflect.TypeOf(prAsBytes))
	//	fmt.Println(unpacked)
	var myInterfaceSlice []interface{} = []interface{}{"Course-0001", "Course-0002", "Course-0003"}
	fmt.Printf("t4: %s\n", reflect.TypeOf(myInterfaceSlice))
	var array1 []string
	fmt.Println(array1)
	fmt.Printf("t5: %s\n", reflect.TypeOf(array1))
	for _, v := range myInterfaceSlice {
		if str, ok := v.(string); ok {
			array1 = append(array1, str)
		}
	}
	fmt.Println(array1)
	// Tạo một slice chứa các giá trị string
	stringSlice := []string{"hello", "world", "golang"}
	// Tạo một slice để chứa các giá trị interface{}
	interfaceSlice := make([]interface{}, len(stringSlice))

	// Duyệt qua từng phần tử trong stringSlice và thêm vào interfaceSlice
	for i, v := range stringSlice {
		interfaceSlice[i] = v
	}

	fmt.Println(interfaceSlice) // Output: [hello world golang]
	// for i := 0; i < len(unpacked); i++ {
	// 	var array1 []string
	// 	if unpacked[i] == "Course-0001,Course-0002,Course-0003" {
	// 		fmt.Println("Hello World")
	// 	}
	// }
	//	if (array[0][1]) = (array[0][2]) {
	//		fmt.Println("Hellw World")
	//	}
	//	array := [][]unpacked
	// array := unpacked[0]
	// fmt.Println(unpacked)
	// indicated := []string{"User_001", "User_002", "User_003", "User_004"}
	// var result [][]string
	// for _, index := range indicated {
	//  	start := index[0]
	//  	end := index[3]
	//  	if start < reflect.Array.String() && end <= len(array) {
	//  		result = append(result, arr[start:end])
	//  	}
	//  }

	// fmt.Println(result)
	//	array := []string(prAsBytes)
	//	fmt.Println(array)

}
