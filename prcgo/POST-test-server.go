package main

import (
	"io"

	"net/http"
)

func hogeHandler(w http.ResponseWriter, req *http.Request) {
	// HTMLテキストをhttp.ResponseWriterへ書き込む
	io.WriteString(w, `
    <!DOCTYPE html>
    <html lang="ja">
    <head>
      <meta charset="UTF-8">
      <title>Go | net/httpパッケージ</title>
    </head>
    <body>
      <h1>Hello, World!</h1>
      <p>これはnet/httpで実装したlocalhostのWEBサーバです</p>
    </body>
    </html>
`)
}

func main() {
	// localhostへのリクエストをhogeHandler関数で処理する
	http.HandleFunc("/", hogeHandler)
	// localhost:8080でサーバー処理開始
	http.ListenAndServe(":8080", nil)
}
