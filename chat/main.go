package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// templは1つのテンプレート
type templateHandler struct {
	once     sync.Once // テンプレートを一回だけコンパイル
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

/**
main関数が呼ばれたらサーバがスタートするだけ。
"/"にアクセスすると、chat.htmlを返す。
*/
func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始します。バックグラウンドで。
	go r.run()
	// webサーバの開始
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
