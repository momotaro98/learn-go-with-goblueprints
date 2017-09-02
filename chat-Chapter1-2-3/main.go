package main

import (
	"flag"
	"log"
	"net/http"
	// "os"
	"path/filepath"
	"sync"
	"text/template"
	// "github.com/momotaro98/learn-go-with-goblueprints/trace"
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
			template.Must(template.ParseFiles(filepath.Join("../templates",
				t.filename)))
	})
	t.templ.Execute(w, r) // リクエスト情報rをテンプレートへ渡す
}

func main() {
	var addr = flag.String("host", ":8080", "Address of the application")
	flag.Parse() // 入力された文字列の値を*addrへセットする。
	r := newRoom()
	// r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// チャットルームを開始
	go r.run()
	// WEBサーバーを起動
	log.Println("Webサーバを開始する。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
