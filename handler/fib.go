package handler

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

type Answer struct {
	Result *big.Int `json:"result"`
}

func Fib(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		v := r.URL.Query()
		if v == nil {
			return
		}
		// nが大きすぎるとオーバーフロー起こす気がするけどそこまで考慮する必要あるのか？
		n, err := strconv.Atoi(v.Get("n"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// エラー文は定義したほうがわかりやすいかも
		if n < 1 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ans := getFib(n)
		res := Answer{
			Result: ans,
		}
		output, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	// "Get" 意外だとerror
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func getFib(n int) *big.Int {
	// bigInt じゃないとデータ長が足りなかった
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}
