package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	// "github.com/momotaro98/learn-go-with-goblueprints/trace"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	// "github.com/stretchr/gomniauth/providers/facebook"
	// "github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
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
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data) // リクエスト情報rをテンプレートへ渡す
}

func main() {
	var addr = flag.String("host", ":8080", "Address of the application")
	flag.Parse() // 入力された文字列の値を*addrへセットする。
	// Gomniauthのセットアップ
	gomniauth.SetSecurityKey("momotaro98") // クライアントサーバー間のデジタル署名のために必要。ランダムな値を設定しておくこと。
	gomniauth.WithProviders(
		// facebook.New(),
		// github.New(),
		google.New(os.Getenv("GOBLUEPRINT_GOOGLE_CLIENTID"),
			os.Getenv("GOBLUEPRINT_GOOGLE_SECRET"),
			"http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	// r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	// チャットルームを開始
	go r.run()
	// WEBサーバーを起動
	log.Println("Webサーバを開始する。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
