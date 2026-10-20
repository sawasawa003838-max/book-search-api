package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)


func main()  {
	fmt.Println("こんにちは！")
	http.HandleFunc("/api/books", homeHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

//r:お客さんからのお願いを読む　w:お客さんへの返事を書く
                 // 変数名 int stringとかの型と同じ
func homeHandler(w http.ResponseWriter, r *http.Request) {
		requestedTitle := r.URL.Query().Get("title")
		fmt.Println("注文された本:", requestedTitle)
		myBook := Book{
			Title: "ミステリー大冒険",
			Genre: "mystery",
			Rating: 4.5,
		}
		secondBook := Book{
			Title: "宇宙のひみつ",
			Genre: "science",
			Rating: 4.8,
		}
		matchedBooks := []Book{}
		for _, book := range []Book{myBook, secondBook} {
			fmt.Println("本棚チェック:", book.Title)

			if book.Title == requestedTitle {
				fmt.Println("見つけた！:", book.Title)
				matchedBooks = append(matchedBooks, book)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(matchedBooks)
}

type Book struct {
	Title string
	Genre string
	Rating float64
}