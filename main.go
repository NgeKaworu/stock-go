package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 结构体定义
type test struct {
	Value   float64
	str     string
	Another float64
}

func main() {

	s := make([]test, 5)
	s[0] = test{Value: 2.21423, Another: 2.21423, str: "test2"}
	s[1] = test{Value: 4.21423, Another: 3.21423, str: "test4"}
	s[2] = test{Value: 1.21423, Another: 5.21423, str: "test1"}
	s[3] = test{Value: 5.21423, Another: 7.21423, str: "test5"}
	s[4] = test{Value: 3.21423, Another: 0.21423, str: "test3"}
	fmt.Println("初始化结果:")
	fmt.Println(s)

	CusSort(s, true, "Value")

	// sort.Slice(s, func(i, j int) bool {
	// 	s1 := reflect.ValueOf(s[i])
	// 	s2 := reflect.ValueOf(s[j])
	// 	if s1.FieldByName("Value").Interface().(float64) < s2.FieldByName("Value").Interface().(float64) {
	// 		return true
	// 	}
	// 	return false
	// })

	fmt.Println("\n从小到大排序结果:")
	fmt.Println(s)

	// var (
	// 	dbinit = flag.Bool("i", false, "init database flag")
	// 	mongo  = flag.String("m", "mongodb://localhost:27017", "mongod addr flag")
	// 	db     = flag.String("db", "stock", "database name")
	// )
	// flag.Parse()

	// log.SetOutput(os.Stdout)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// eng := dbengin.NewDbEngine()
	// err := eng.Open(*mongo, *db, *dbinit)

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// // 风险收益率(Rate of Risked Return)
	// // 假设10年内 > 80% 30年内 < 20%
	// RRR := 0.086
	// // 通货
	// CPI := 0.052
	// // 无风险利率 (The risk-free rate of interest)
	// RFR := 0.0285
	// discount := RRR + CPI + RFR
	// log.Println(discount)
	// stocks := utils.Merge(constants.Ss50, constants.Hs300)
	// for k, v := range stocks {
	// 	s := &stock.Stock{
	// 		Code:       k,
	// 		BourseCode: v,
	// 	}
	// 	switch v {
	// 	case "01":
	// 		s.Bourse = "sh"
	// 	case "02":
	// 		s.Bourse = "sz"
	// 	default:
	// 		break
	// 	}
	// 	s.FetchCurrentInfo()
	// 	s.FetchMainIndicator()
	// 	s.FetchClassify()
	// 	s.Calc()

	// 	log.Printf("%+v\n", s)
	// 	break
	// }
	// for k, v := range constants.Ss50 {
	// 	log.Println(k, v)
	//
	// }

}

// CusSort 自定义排序
func CusSort(s interface{}, gt bool, key string) {
	// var isGt bool
	sort.Slice(s, func(i, j int) bool {
		val := reflect.ValueOf(s)
		s1 := val.Index(i)
		s2 := val.Index(j)
		if s1.FieldByName(key).Interface().(float64) > s2.FieldByName(key).Interface().(float64) {
			return true
		}
		return false

	})
}
