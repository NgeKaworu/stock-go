package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NgeKaworu/stock/src/resultor"
	"github.com/NgeKaworu/stock/src/stock"
	"github.com/julienschmidt/httprouter"
)

func (d *DbEngine) FetchOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var s *stock.Stock
	for v, k := range stock.Stocks {
		s = stock.NewStock(v, k)
	}
	s, err := s.FetchCurrentInfor()
	if err != nil {
		resultor.RetFail(w, err)
	}
	s, err = s.FetchEnterPrise()
	if err != nil {
		resultor.RetFail(w, err)
	}

	bs, _ := json.Marshal(s)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("student=%v\n", out.String())

	resultor.RetOk(w, &s)
}
