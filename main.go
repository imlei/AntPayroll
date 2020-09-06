package main

//help to calculate the payroll, all the rate based on 2020
//Thanks for the CRA offer the useful tools help me to build this software
//https://www.canada.ca/en/revenue-agency/services/e-services/e-services-businesses/payroll-deductions-online-calculator.html
//last update 20200905

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Taxconfig struct {
	Abtax Abtax  `json:"abtax"`
	Date  string `json:"date"`
	Mbtax Mbtax `json:"mbtax"`
	Mysql Mysql  `json:"mysql"`
	Nbtax Nbtax `json:"nbtax"`
	Nltax Nltax `json:"nltax"`
	Nstax Nstax `json:"nstax"`
	Nttax Nttax `json:"nttax"`
	Nutax Nutax `json:"nutax"`
	Ontax Ontax `json:"ontax"`
	Petax Petax `json:"petax"`
	Sktax Sktax  `json:"sktax"`
	Yttax Yttax `json:"yttax"`
}
type Abtax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	L4 float64    `json:"l4"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	R5 float64 `json:"r5"`
	Pba float64 `json："pba"`
}
type Mbtax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	Pba float64 `json："pba"`
}
type Nbtax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	L4 float64    `json:"l4"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	R5 float64 `json:"r5"`
	Pba float64 `json："pba"`
}
type Nltax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	L4 float64    `json:"l4"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	R5 float64 `json:"r5"`
	Pba float64 `json："pba"`
}
type Nstax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	L4 float64    `json:"l4"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	R5 float64 `json:"r5"`
	Pba float64 `json："pba"`
}
type Nttax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	Pba float64 `json："pba"`
}
type Nutax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	Pba float64 `json："pba"`
}
type Ontax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	L4 float64    `json:"l4"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	R5 float64 `json:"r5"`
	Pba float64 `json："pba"`
}
type Petax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	Pba float64 `json："pba"`
}
type Sktax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	Pba float64 `json："pba"`
}
type Yttax struct {
	L1 float64    `json:"l1"`
	L2 float64     `json:"l2"`
	L3 float64    `json:"l3"`
	L4 float64    `json:"l4"`
	R1 float64 `json:"r1"`
	R2 float64 `json:"r2"`
	R3 float64 `json:"r3"`
	R4 float64 `json:"r4"`
	R5 float64 `json:"r5"`
	Pba float64 `json："pba"`
}
type Mysql struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	grossamount    float64
	pramt          float64 //prorate amount
	paytype int
	taxamount      float64
	payfre         float64
	fixamt1        float64
	fixamt2        float64
	fixamt3        float64
	fixamt4        float64
	cppamt         float64
	eiamt          float64
	cecamt         float64
	icamt, picamt  float64
	v,ptinfo             string
)

var incomeclass = map[string]float64{
	"level1": 48535.01,
	"level2": 97069.01,
	"level3": 150473.01,
	"level4": 214368.01,
	"rate1":  0.15,
	"rate2":  0.205,
	"rate3":  0.26,
	"rate4":  0.29,
	"rate5":  0.33,
	"fbpa":   13229.00,
	"fbpac":  1984.35,
	"kcpp":   0.0525,
	"kei":    0.0158,
}
var bctax = map[string]float64{
	"l1":  41725,
	"l2":  83451,
	"l3":  95812,
	"l4":  116344,
	"l5":  157748,
	"l6":  220000,
	"r1":  0.0506,
	"r2":  0.077,
	"r3":  0.105,
	"r4":  0.1229,
	"r5":  0.147,
	"r6":  0.168,
	"r7":  0.2420,
	"bpa": 10949,
	"rt1": 21185,
	"rt2": 34556,
}

func main() {


	fmt.Println("Please enter your payroll:")
	fmt.Scanln(&grossamount)
	fmt.Println("Please enter the year of income:/n 1 = monthly; 2 = Semi-monthly; 3 = bi-weekly; 4 = weekly; 5 = daily")
	fmt.Scanln(&paytype)
	fmt.Println("Please enter your provincial:/n 1 = AB; 2 = BC")
	fmt.Scanln(&ptinfo)
	//choose pay period frequency
	if paytype == 1 {
		payfre = 12
	} else if paytype == 2 {
		payfre = 24
	} else if paytype == 3 {
		payfre = 26
	} else if paytype == 4 {
		payfre = 52
	} else {
		payfre = 260
	}

	pramt = grossamount * payfre
switch {
case ptinfo == "AB":
	picamt = ptab(pramt)/payfre
    case ptinfo == "ON":
	picamt = pton(pramt)/payfre
	case ptinfo == "NU":
		picamt = ptnu(pramt)/payfre
case ptinfo == "BC":
	picamt = ptbc(pramt)/payfre
case ptinfo == "MB":
	picamt = ptmb(pramt)/payfre
case ptinfo == "NB":
	picamt = ptnb(pramt)/payfre
case ptinfo == "NL":
	picamt = ptnl(pramt)/payfre
case ptinfo == "NS":
	picamt = ptns(pramt)/payfre
case ptinfo == "NT":
	picamt = ptnt(pramt)/payfre
case ptinfo == "PE":
	picamt = ptpe(pramt)/payfre
case ptinfo == "SK":
	picamt = ptsk(pramt)/payfre
case ptinfo == "YT":
	picamt = ptyt(pramt)/payfre
}
	if ptinfo == "AB" {
		 picamt = ptab(pramt)/payfre
	} else if ptinfo == "BC"{
		picamt = ptbc(pramt)/payfre
	} else if ptinfo == "MB" {
		picamt = ptmb(pramt)/payfre
	} else {
		picamt = ptnb(pramt)/payfre
	}


	eiamt := ei(pramt) / payfre
	cppamt := cpp(pramt) / payfre
	cecamt := cec(pramt) / payfre
	taxamount := ft(grossamount)

	//less the federl personal amount T3 = (R × A) – K – K1 – K2 – K3 – K4

	if pramt > incomeclass["fbpa"] {
		icamt = (taxamount - incomeclass["fbpac"] - cecamt*payfre - cpp(pramt)*0.15 - ei(pramt)*0.15) / payfre
	} else {
		icamt = 0
	}

	fmt.Printf("your gross amount: %0.2f\n", grossamount)
	fmt.Printf(" tax payable: %0.2f\n", taxamount)
	fmt.Printf("Income tax payable: %0.2f\n", icamt)
	fmt.Printf("CPP tax payable: %0.2f\n", cppamt)
	fmt.Printf("EI tax payable: %0.2f\n", eiamt)
	fmt.Printf("provincial tax amount: %0.2f\n", picamt)
	fmt.Printf("AB provincial tax amount: %0.2f\n", picamt)

}

// to calculate the Federal income tax payable
func ft(f float64) float64 {
	//15% on the first $48,535 of taxable income
	fixamt1 = incomeclass["level1"] * incomeclass["rate1"]
	//20.5% on the portion of taxable income over $48,535 up to $97,069
	fixamt2 = (incomeclass["level2"]-incomeclass["level1"])*incomeclass["rate2"] + fixamt1
	//26% on the portion of taxable income over $97,069 up to $150,473
	fixamt3 = (incomeclass["level3"]-incomeclass["level2"])*incomeclass["rate3"] + fixamt2
	//29% on the portion of taxable income over $150,473 up to $214,368
	fixamt4 = (incomeclass["level4"]-incomeclass["level3"])*incomeclass["rate4"] + fixamt3
	//33% of taxable income over $214,368

	//calculate the Federal income tax payable
	if pramt < incomeclass["level1"] {
		f = pramt * incomeclass["rate1"]
	} else if pramt < incomeclass["level2"] {
		f = (pramt-incomeclass["level1"])*incomeclass["rate2"] + fixamt1
	} else if pramt < incomeclass["level3"] {
		f = (pramt-incomeclass["level2"])*incomeclass["rate3"] + fixamt2
	} else if pramt < incomeclass["level4"] {
		f = (pramt-incomeclass["level3"])*incomeclass["rate4"] + fixamt3
	} else {
		f = (pramt-incomeclass["level4"])*incomeclass["rate5"] + fixamt4
	}

	return f
}

func ReadJsonFile() {
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

//calculate the CPP and update for maximum amount.
func cpp(z float64) float64 {
	if pramt <= 3500 {
		z = 0
	} else {
		if pramt < 58700 {
			z = (pramt - 3500) * incomeclass["kcpp"]
		} else {
			z = 2898
		}
	}
	return z

}

//calculate the EI and update for maximum amount
func ei(x float64) float64 {
	if pramt > 54200 {
		x = 856.36
	} else {
		x = x * incomeclass["kei"]
	}

	return x

}

// to calculate CEC amount K4= The lesser of: (i)   0.15 × A; and (ii)  0.15 × $1,161.
func cec(y float64) float64 {
	if pramt > 1245 {
		y = 1245 * 0.15
	} else {
		y = pramt * 0.15
	}
	return y
}

// to calculate BC Tax
func ptbc(bc float64) float64 {
	var (
		bcamt1, bcamt2, bcamt3, bcamt4, bcamt5, bcamt6, bcrt float64
	)

	//first $41,725
	bcamt1 = bctax["l1"] * bctax["r1"]
	//over $41,725 up to $83,451
	bcamt2 = (bctax["l2"]-bctax["l1"])*bctax["r2"] + bcamt1
	//over $83,451 up to $95,812
	bcamt3 = (bctax["l3"]-bctax["l2"])*bctax["r3"] + bcamt2
	//over $95,812 up to $116,344
	bcamt4 = (bctax["l4"]-bctax["l3"])*bctax["r4"] + bcamt3
	//over $116,344 up to $157,748
	bcamt5 = (bctax["l5"]-bctax["l4"])*bctax["r5"] + bcamt4
	//over $157,748 up to $220,000
	bcamt6 = (bctax["l6"]-bctax["l5"])*bctax["r6"] + bcamt5
	//over $220,000

	//calculate the provincial tax reduction
	if pramt <= bctax["rt1"] {
		bcrt = 476
	} else if pramt <= bctax["rt2"] {
		bcrt = 476 - (pramt-bctax["rt1"])*0.0356
	} else {
		bcrt = 0
	}

	//calculate the Federal income tax payable

	if pramt < bctax["l1"] {
		bc = pramt * bctax["r1"]
	} else if pramt < bctax["l2"] {
		bc = (pramt-bctax["l1"])*bctax["r2"] + bcamt1
	} else if pramt < bctax["l3"] {
		bc = (pramt-bctax["l2"])*bctax["r3"] + bcamt2
	} else if pramt < bctax["l4"] {
		bc = (pramt-bctax["l3"])*bctax["r4"] + bcamt3
	} else if pramt < bctax["l5"] {
		bc = (pramt-bctax["l4"])*bctax["r5"] + bcamt4
	} else if pramt < bctax["l6"] {
		bc = (pramt-bctax["l5"])*bctax["r6"] + bcamt5
	} else if pramt < bctax["l7"] {
		bc = (pramt-bctax["l6"])*bctax["r7"] + bcamt6
	}

	bc = bc - bctax["bpa"]*bctax["r1"] - cpp(pramt)*bctax["r1"] - ei(pramt)*bctax["r1"] - bcrt
	if bc <= 0 {
		bc = 0
	} else {
	}
	return bc
}

// to calculate AB provincial tax
func ptab(ab float64) float64 {
	var (
		abamt1, abamt2, abamt3, abamt4 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $131,220
	abamt1 = v.Abtax.R1 * v.Abtax.L1
	//over $131,220 to $ 157,464
	abamt2 = (v.Abtax.L2-v.Abtax.L1)*v.Abtax.R2 + abamt1
	//over $157,464 to $209,950
	abamt3 = (v.Abtax.L3-v.Abtax.L2)*v.Abtax.R3 + abamt2
	//over $209,950 to $314,928
	abamt4 = (v.Abtax.L4-v.Abtax.L3)*v.Abtax.R4 + abamt3


	//calculate the Federal income tax payable

	if pramt < v.Abtax.L1 {
		ab = pramt * v.Abtax.R1
	} else if pramt < v.Abtax.L2 {
		ab = (pramt-v.Abtax.L1)*v.Abtax.R2 + abamt1
	} else if pramt < v.Abtax.L3 {
		ab = (pramt-v.Abtax.L2)*v.Abtax.R3 + abamt2
	} else if pramt < v.Abtax.L4 {
		ab = (pramt-v.Abtax.L3)*v.Abtax.R4 + abamt3
	} else {
		ab = (pramt-v.Abtax.L4)*v.Abtax.R5 + abamt4
	}

	ab = ab - v.Abtax.Pba * v.Abtax.R1 - cpp(pramt)*v.Abtax.R1 - ei(pramt)*v.Abtax.R1
	if ab <= 0 {
		ab = 0
	} else {
	}
	return ab
}

func ptmb(mb float64) float64 {
	var (
		mbamt1, mbamt2 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $33,389
	mbamt1 = v.Mbtax.R1 * v.Mbtax.L1
	//over $33,389 to $ 72,164
	mbamt2 = (v.Mbtax.L2-v.Mbtax.L1)*v.Mbtax.R2 + mbamt1


	//calculate the Federal income tax paymble

	if pramt < v.Mbtax.L1 {
		mb = pramt * v.Mbtax.R1
	} else if pramt < v.Mbtax.L2 {
		mb = (pramt-v.Mbtax.L1)*v.Mbtax.R2 + mbamt1
	} else {
		mb = (pramt-v.Mbtax.L2)*v.Mbtax.R3 + mbamt2
	}

	mb = mb - v.Mbtax.Pba * v.Mbtax.R1 - cpp(pramt)*v.Mbtax.R1 - ei(pramt)*v.Mbtax.R1
	if mb <= 0 {
		mb = 0
	} else {
	}
	return mb
}

func ptnb(nb float64) float64 {
	var (
		nbamt1, nbamt2, nbamt3, nbamt4 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $131,220
	nbamt1 = v.Nbtax.R1 * v.Nbtax.L1
	//over $131,220 to $ 157,464
	nbamt2 = (v.Nbtax.L2-v.Nbtax.L1)*v.Nbtax.R2 + nbamt1
	//over $157,464 to $209,950
	nbamt3 = (v.Nbtax.L3-v.Nbtax.L2)*v.Nbtax.R3 + nbamt2
	//over $209,950 to $314,928
	nbamt4 = (v.Nbtax.L4-v.Nbtax.L3)*v.Nbtax.R4 + nbamt3


	//calculate the Federal income tax payNble

	if pramt < v.Nbtax.L1 {
		nb = pramt * v.Nbtax.R1
	} else if pramt < v.Nbtax.L2 {
		nb = (pramt-v.Nbtax.L1)*v.Nbtax.R2 + nbamt1
	} else if pramt < v.Nbtax.L3 {
		nb = (pramt-v.Nbtax.L2)*v.Nbtax.R3 + nbamt2
	} else if pramt < v.Nbtax.L4 {
		nb = (pramt-v.Nbtax.L3)*v.Nbtax.R4 + nbamt3
	} else {
		nb = (pramt-v.Nbtax.L4)*v.Nbtax.R5 + nbamt4
	}

	nb = nb - v.Nbtax.Pba * v.Nbtax.R1 - cpp(pramt)*v.Nbtax.R1 - ei(pramt)*v.Nbtax.R1
	if nb <= 0 {
		nb = 0
	} else {
	}
	return nb
}
func ptnl(nl float64) float64 {
	var (
		nlamt1, nlamt2, nlamt3, nlamt4 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $131,220
	nlamt1 = v.Nltax.R1 * v.Nltax.L1
	//over $131,220 to $ 157,464
	nlamt2 = (v.Nltax.L2-v.Nltax.L1)*v.Nltax.R2 + nlamt1
	//over $157,464 to $209,950
	nlamt3 = (v.Nltax.L3-v.Nltax.L2)*v.Nltax.R3 + nlamt2
	//over $209,950 to $314,928
	nlamt4 = (v.Nltax.L4-v.Nltax.L3)*v.Nltax.R4 + nlamt3


	//calculate the Federal income tax paynlle

	if pramt < v.Nltax.L1 {
		nl = pramt * v.Nltax.R1
	} else if pramt < v.Nltax.L2 {
		nl = (pramt-v.Nltax.L1)*v.Nltax.R2 + nlamt1
	} else if pramt < v.Nltax.L3 {
		nl = (pramt-v.Nltax.L2)*v.Nltax.R3 + nlamt2
	} else if pramt < v.Nltax.L4 {
		nl = (pramt-v.Nltax.L3)*v.Nltax.R4 + nlamt3
	} else {
		nl = (pramt-v.Nltax.L4)*v.Nltax.R5 + nlamt4
	}

	nl = nl - v.Nltax.Pba * v.Nltax.R1 - cpp(pramt)*v.Nltax.R1 - ei(pramt)*v.Nltax.R1
	if nl <= 0 {
		nl = 0
	} else {
	}
	return nl
}
func ptns(ns float64) float64 {
	var (
		nsamt1, nsamt2, nsamt3, nsamt4 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $131,220
	nsamt1 = v.Nstax.R1 * v.Nstax.L1
	//over $131,220 to $ 157,464
	nsamt2 = (v.Nstax.L2-v.Nstax.L1)*v.Nstax.R2 + nsamt1
	//over $157,464 to $209,950
	nsamt3 = (v.Nstax.L3-v.Nstax.L2)*v.Nstax.R3 + nsamt2
	//over $209,950 to $314,928
	nsamt4 = (v.Nstax.L4-v.Nstax.L3)*v.Nstax.R4 + nsamt3


	//calculate the Federal income tax paynsle

	if pramt < v.Nstax.L1 {
		ns = pramt * v.Nstax.R1
	} else if pramt < v.Nstax.L2 {
		ns = (pramt-v.Nstax.L1)*v.Nstax.R2 + nsamt1
	} else if pramt < v.Nstax.L3 {
		ns = (pramt-v.Nstax.L2)*v.Nstax.R3 + nsamt2
	} else if pramt < v.Nstax.L4 {
		ns = (pramt-v.Nstax.L3)*v.Nstax.R4 + nsamt3
	} else {
		ns = (pramt-v.Nstax.L4)*v.Nstax.R5 + nsamt4
	}

	ns = ns - v.Nstax.Pba * v.Nstax.R1 - cpp(pramt)*v.Nstax.R1 - ei(pramt)*v.Nstax.R1
	if ns <= 0 {
		ns = 0
	} else {
	}
	return ns
}
func ptnt(nt float64) float64 {
	var (
		ntamt1, ntamt2, ntamt3 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./cntfig/cntfig.jsnt", &v)

	//read cntfig file

	//first $131,220
	ntamt1 = v.Nttax.R1 * v.Nttax.L1
	//over $131,220 to $ 157,464
	ntamt2 = (v.Nttax.L2-v.Nttax.L1)*v.Nttax.R2 + ntamt1
	//over $157,464 to $209,950
	ntamt3 = (v.Nttax.L3-v.Nttax.L2)*v.Nttax.R3 + ntamt2



	//calculate the Federal income tax payntle

	if pramt < v.Nttax.L1 {
		nt = pramt * v.Nttax.R1
	} else if pramt < v.Nttax.L2 {
		nt = (pramt-v.Nttax.L1)*v.Nttax.R2 + ntamt1
	} else if pramt < v.Nttax.L3 {
		nt = (pramt-v.Nttax.L2)*v.Nttax.R3 + ntamt2
	} else  {
		nt = (pramt-v.Nttax.L3)*v.Nttax.R4 + ntamt3
	}

	nt = nt - v.Nttax.Pba * v.Nttax.R1 - cpp(pramt)*v.Nttax.R1 - ei(pramt)*v.Nttax.R1
	if nt <= 0 {
		nt = 0
	} else {
	}
	return nt
}
func ptnu(nu float64) float64 {
	var (
		nuamt1, nuamt2, nuamt3 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./cnufig/cnufig.jsnu", &v)

	//read cnufig file

	//first $131,220
	nuamt1 = v.Nutax.R1 * v.Nutax.L1
	//over $131,220 to $ 157,464
	nuamt2 = (v.Nutax.L2-v.Nutax.L1)*v.Nutax.R2 + nuamt1
	//over $157,464 to $209,950
	nuamt3 = (v.Nutax.L3-v.Nutax.L2)*v.Nutax.R3 + nuamt2
	//over $209,950 to $314,928


	//calculate the Federal income tax paynule

	if pramt < v.Nutax.L1 {
		nu = pramt * v.Nutax.R1
	} else if pramt < v.Nutax.L2 {
		nu = (pramt-v.Nutax.L1)*v.Nutax.R2 + nuamt1
	} else if pramt < v.Nutax.L3 {
		nu = (pramt-v.Nutax.L2)*v.Nutax.R3 + nuamt2
	} else  {
		nu = (pramt-v.Nutax.L3)*v.Nutax.R4 + nuamt3
	}

	nu = nu - v.Nutax.Pba * v.Nutax.R1 - cpp(pramt)*v.Nutax.R1 - ei(pramt)*v.Nutax.R1
	if nu <= 0 {
		nu = 0
	} else {
	}
	return nu
}
func pton(on float64) float64 {
	var (
		onamt1, onamt2, onamt3, onamt4 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $131,220
	onamt1 = v.Ontax.R1 * v.Ontax.L1
	//over $131,220 to $ 157,464
	onamt2 = (v.Ontax.L2-v.Ontax.L1)*v.Ontax.R2 + onamt1
	//over $157,464 to $209,950
	onamt3 = (v.Ontax.L3-v.Ontax.L2)*v.Ontax.R3 + onamt2
	//over $209,950 to $314,928
	onamt4 = (v.Ontax.L4-v.Ontax.L3)*v.Ontax.R4 + onamt3


	//calculate the Federal income tax payonle

	if pramt < v.Ontax.L1 {
		on = pramt * v.Ontax.R1
	} else if pramt < v.Ontax.L2 {
		on = (pramt-v.Ontax.L1)*v.Ontax.R2 + onamt1
	} else if pramt < v.Ontax.L3 {
		on = (pramt-v.Ontax.L2)*v.Ontax.R3 + onamt2
	} else if pramt < v.Ontax.L4 {
		on = (pramt-v.Ontax.L3)*v.Ontax.R4 + onamt3
	} else {
		on = (pramt-v.Ontax.L4)*v.Ontax.R5 + onamt4
	}

	on = on - v.Ontax.Pba * v.Ontax.R1 - cpp(pramt)*v.Ontax.R1 - ei(pramt)*v.Ontax.R1
	if on <= 0 {
		on = 0
	} else {
	}
	return on
}
func ptpe(pe float64) float64 {
	var (
		peamt1, peamt2 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $33,389
	peamt1 = v.Petax.R1 * v.Petax.L1
	//over $33,389 to $ 72,164
	peamt2 = (v.Petax.L2-v.Petax.L1)*v.Petax.R2 + peamt1


	//calculate the Federal income tax paypele

	if pramt < v.Petax.L1 {
		pe = pramt * v.Petax.R1
	} else if pramt < v.Petax.L2 {
		pe = (pramt-v.Petax.L1)*v.Petax.R2 + peamt1
	} else {
		pe = (pramt-v.Petax.L2)*v.Petax.R3 + peamt2
	}

	pe = pe - v.Petax.Pba * v.Petax.R1 - cpp(pramt)*v.Petax.R1 - ei(pramt)*v.Petax.R1
	if pe <= 0 {
		pe = 0
	} else {
	}
	return pe
}
func ptsk(sk float64) float64 {
	var (
		skamt1, skamt2 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./config/config.json", &v)

	//read config file

	//first $33,389
	skamt1 = v.Sktax.R1 * v.Sktax.L1
	//over $33,389 to $ 72,164
	skamt2 = (v.Sktax.L2-v.Sktax.L1)*v.Sktax.R2 + skamt1


	//calculate the Federal income tax payskle

	if pramt < v.Sktax.L1 {
		sk = pramt * v.Sktax.R1
	} else if pramt < v.Sktax.L2 {
		sk = (pramt-v.Sktax.L1)*v.Sktax.R2 + skamt1
	} else {
		sk = (pramt-v.Sktax.L2)*v.Sktax.R3 + skamt2
	}

	sk = sk - v.Sktax.Pba * v.Sktax.R1 - cpp(pramt)*v.Sktax.R1 - ei(pramt)*v.Sktax.R1
	if sk <= 0 {
		sk = 0
	} else {
	}
	return sk
}
func ptyt(yt float64) float64 {
	var (
		ytamt1, ytamt2, ytamt3, ytamt4 float64
	)
	var JsonParse = NewJsonStruct()
	v := Taxconfig{}
	//下面使用的是相对路径
	JsonParse.Load("./cytfig/cytfig.jsyt", &v)

	//read cytfig file

	//first $131,220
	ytamt1 = v.Yttax.R1 * v.Yttax.L1
	//over $131,220 to $ 157,464
	ytamt2 = (v.Yttax.L2-v.Yttax.L1)*v.Yttax.R2 + ytamt1
	//over $157,464 to $209,950
	ytamt3 = (v.Yttax.L3-v.Yttax.L2)*v.Yttax.R3 + ytamt2
	//over $209,950 to $314,928
	ytamt4 = (v.Yttax.L4-v.Yttax.L3)*v.Yttax.R4 + ytamt3


	//calculate the Federal income tax payytle

	if pramt < v.Yttax.L1 {
		yt = pramt * v.Yttax.R1
	} else if pramt < v.Yttax.L2 {
		yt = (pramt-v.Yttax.L1)*v.Yttax.R2 + ytamt1
	} else if pramt < v.Yttax.L3 {
		yt = (pramt-v.Yttax.L2)*v.Yttax.R3 + ytamt2
	} else if pramt < v.Yttax.L4 {
		yt = (pramt-v.Yttax.L3)*v.Yttax.R4 + ytamt3
	} else {
		yt = (pramt-v.Yttax.L4)*v.Yttax.R5 + ytamt4
	}

	yt = yt - v.Yttax.Pba * v.Yttax.R1 - cpp(pramt)*v.Yttax.R1 - ei(pramt)*v.Yttax.R1
	if yt <= 0 {
		yt = 0
	} else {
	}
	return yt
}