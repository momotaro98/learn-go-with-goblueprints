package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templは1つのテンプレートを表す
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServerHTTPはHTTPリクエストを処理する
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 一番はじめだけテンプレートを作成(コンパイル)
	// あとのリクエストに対してはすべてそのコンパイルされたテンプレートを利用
	t.once.Do(func() { // HTTPリクエストがあって初めて実行するか(lazy initialization)前もって実行するかの議論がある
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.templ.Execute(w, nil) // 注釈: 本当はエラーハンドリングが必要
}

func main() {
	// ルート
	http.Handle("/", &templateHandler{filename: "chat.html"}) // templateHandlerはhttpのHandlerインターフェースを実装しているのでHandleへ渡せる
	// WEBサーバーを開始
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
