package stock

import (
	"log"
	"strconv"
)

// Calc 计算全部
func (s *Stock) Calc() {
	s.CalcPB()
	s.CalcPE()
	s.CalcAAGR()
	s.CalcPEG()
	s.CalcROE()
}

// Discount 估值
func (s *Stock) Discount(r float64) {
	s.CalcDCE(r)
	s.CalcDCER()
	s.CalcDPE(r)
	s.CalcDPER()
}

// CalcPB 计算市净率
func (s *Stock) CalcPB() {
	// 股价
	cp, err := strconv.ParseFloat(*s.CurrentInfo.CurrentPrice, 64)
	if err != nil {
		log.Println(err, s.Code, "cp")
		return
	}
	var bps float64
	// 每股净资产

	if len(s.Enterprise) > 0 {
		bps, err = strconv.ParseFloat(*s.Enterprise[0].Bps, 64)
		if err != nil || bps == 0 {
			log.Println(err, s.Code, "bps")
			return
		}
		s.PB = cp / bps
	}

}

// CalcPE 计算市盈率
func (s *Stock) CalcPE() {
	// 股价
	cp, err := strconv.ParseFloat(*s.CurrentInfo.CurrentPrice, 64)
	if err != nil || cp == 0 {
		log.Println(err, s.Code, "cp")
		return
	}
	// 每股未分配利润
	var mgwfplr float64
	if len(s.Enterprise) > 0 {
		mgwfplr, err = strconv.ParseFloat(*s.Enterprise[0].Mgwfplr, 64)
		if err != nil {
			log.Println(err, s.Code, "mgfplr")
			return
		}
		s.PE = mgwfplr / cp
	}

}

// CalcAAGR 计算平均年增长率
func (s *Stock) CalcAAGR() {
	enterpriseList := s.Enterprise
	len := len(enterpriseList)
	var sum float64

	for k, v := range enterpriseList {
		n := k + 1
		if n >= len {
			break
		}
		lastBps, err := strconv.ParseFloat(*enterpriseList[n].Bps, 64)
		if err != nil || lastBps == 0 {
			log.Println(err, s.Code)
			continue
		}
		Bps, err := strconv.ParseFloat(*v.Bps, 64)

		if err != nil {
			log.Println(err, s.Code)
			continue
		}

		curAAGR := (Bps - lastBps) / lastBps

		sum += curAAGR

	}
	s.AAGR = sum / float64((len - 1))

}

// CalcPEG 计算市盈增长比
func (s *Stock) CalcPEG() {
	s.PEG = s.PE / s.AAGR

}

// CalcROE 计算净资产收益率
func (s *Stock) CalcROE() {
	// 每股净值
	var mgwfplr, bps float64
	var err error
	if len(s.Enterprise) > 0 {
		mgwfplr, err = strconv.ParseFloat(*s.Enterprise[0].Mgwfplr, 64)
		if err != nil {
			log.Println(err, s.Code, "mgwfplr")
			return
		}
		// 每股未分配利润
		bps, err = strconv.ParseFloat(*s.Enterprise[0].Bps, 64)
		if err != nil || bps == 0 {
			log.Println(err, s.Code, "bps")
			return
		}
		s.ROE = mgwfplr / bps
	}

}

// CalcDPE 计算动态利润估值
func (s *Stock) CalcDPE(r float64) {
	if len(s.Enterprise) > 0 {

		bps, err := strconv.ParseFloat(*s.Enterprise[0].Bps, 64)
		if err != nil {
			log.Println(err, s.Code, "bps")
			return
		}
		s.DPE = bps / (r - s.AAGR)
	}
}

// CalcDPER 估值 现值比
func (s *Stock) CalcDPER() {
	cp, err := strconv.ParseFloat(*s.CurrentInfo.CurrentPrice, 64)
	if err != nil || cp == 0 {
		log.Println(err, s.Code, "dper")
		return
	}

	s.DPER = s.DPE / cp
}

// CalcDCE 计算动态现金估值
func (s *Stock) CalcDCE(r float64) {
	// 每股经营现金流(元)
	if len(s.Enterprise) > 0 {
		mgjyxjje, err := strconv.ParseFloat(*s.Enterprise[0].Mgjyxjje, 64)
		if err != nil {
			log.Println(err, s.Code, "mgjyxjje")
			return
		}
		s.DCE = mgjyxjje / (r - s.AAGR)
	}

}

// CalcDCER 估值 现值比
func (s *Stock) CalcDCER() {
	cp, err := strconv.ParseFloat(*s.CurrentInfo.CurrentPrice, 64)
	if err != nil || cp == 0 {
		log.Println(err, s.Code, "dcer")
		return
	}

	s.DCER = s.DCE / cp
}
