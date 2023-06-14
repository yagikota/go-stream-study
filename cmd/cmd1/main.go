package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// }

// func main() {
// 	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	var user User
// 	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("User: %+v\n", user)
// }

type Person struct {
	Name   string
	Parent *Person
}

func main() {
	// 大きなJSONデータを含むファイルを開く
	file, err := os.Open("../cmd2/large.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteFile, _ := ioutil.ReadAll(file)

	// デコードするJSONデータの型を指定
	var person Person

	err = json.Unmarshal(byteFile, &person)
	if err != nil {
		panic(err)
	}
	// personを使った処理を実行
}
