package gojson

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestSetValue(t *testing.T)  {
	s := `
	{
		"name": "gopher",
		"age": {
			"name": "gopher",
			"age": {
				"name": "gopher",
				"age": 7
			}
		},
		"array": [1, 2, 3, 4, 5]
	}
	`

	ss := `
	{
		"name":"test"
	}
	`

	var u interface{}
	json.Unmarshal([]byte(s), &u)
	json.Unmarshal([]byte(ss), &u)
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}
	JsonSetValue(u, "data", "add data")
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	JsonSetValue(u, "age.age.name", "change data")
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	JsonSetValue(u, "age.age", JsonNewObject() )
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	JsonSetValue(u, "age.age.name", "change data")
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	result, _ := json.Marshal(u)
	fmt.Printf("%+v\n", string(result))
}


func TestGetValue(t *testing.T)  {
	s := `
	{
		"name": "gopher",
		"age": {
			"name": "gopher",
			"age": {
				"name": "gopher",
				"age": 7
			},
			"array": [
				{
					"name": "test1"
				},
				{
					"name": "test2"
				},
				{
					"name": "test3"
				}
			]
		},
		"array": [1, 2, 3, 4, 5]
	}
	`

	var u interface{}
	json.Unmarshal([]byte(s), &u)

	// 키 호출 테스트
	fmt.Printf("%+v\n", JsonGetValue(u, "age"))
	fmt.Printf("%+v\n", JsonGetValue(u, "name"))

	// 키 조합 호출테스트
	fmt.Printf("%+v\n", JsonGetValue(u, "age.name"))
	fmt.Printf("%+v\n", JsonGetValue(u, "age.age.age"))

	// 배열 인덱스 호출테스트
	fmt.Printf("%+v\n", JsonGetValue(u, "array[1]"))
	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[1]"))
	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[0].name"))
	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[2].name"))

	// 배열 인덱스를 벗어난 호출
	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[3].name"))
	fmt.Printf("%+v\n", JsonGetValue(u, "array[5]"))

	// 데이터가 없을경우 기본값을 반환하는 펑션 테스트
	fmt.Printf("%+v\n", JsonGetValueDefault(u, "array[5]", "test"))
}

func TestCheckNulls(t *testing.T)  {
	s := `
	{
		"name": "gopher",
		"age": {
			"name": "gopher",
			"age": {
				"name": "gopher",
				"age": 7
			},
			"array": [
				{
					"name": "test1"
				},
				{
					"name": "test2"
				},
				{
					"name": "test3"
				}
			]
		},
		"array": [1, 2, 3, 4, 5]
	}
	`
	var u interface{}
	json.Unmarshal([]byte(s), &u)

	fmt.Printf("Normal Case : %+v\n", CheckNulls(u, []string{"name", "age", "age.name", "array[3]"}))
	fmt.Printf("Abnormal Case : %+v\n", CheckNulls(u, []string{"name.age", "age", "age.name", "array[7]"}))
}