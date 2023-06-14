package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

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
// 	bodyByte, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := json.Unmarshal(bodyByte, &user); err != nil {
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
	file, err := os.Open("large.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ファイルを読み込むbufio.Readerを作成
	reader := bufio.NewReader(file)

	// JSONデコーダーを作成
	decoder := json.NewDecoder(reader)

	// デコードするJSONデータの型を指定
	var person Person

	// JSONデータをストリームとして処理
	for {
		err := decoder.Decode(&person)
		if err != nil {
			if err == io.EOF {
				// ストリームの終了を検知した場合は処理を終了
				break
			} else {
				panic(err)
			}
		}
		// personを使った処理を実行
	}
}
