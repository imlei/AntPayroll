package main

//help to calculate the payroll
//Thanks for the CRA offer the usefull tools help me to build this software
//https://www.canada.ca/en/revenue-agency/services/e-services/e-services-businesses/payroll-deductions-online-calculator.html
//last update 20200901

import "fmt"

var (
	grossamount float64
	paytype     int
	taxamount   float64
	payfre      float64
	fixamt1     float64
	fixamt2     float64
	fixamt3     float64
	fixamt4     float64
	cppamt      float64
	eiamt       float64
	cecamt      float64
	icamt       float64
)

func main() {

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
		"fbpac":  1810.35,
		"kcpp":   0.0525,
		"kei":    0.0158,
	}
	fmt.Println("Please enter your payroll:")
	fmt.Scanln(&grossamount)
	fmt.Println("Please enter the year of income:/n 1 = monthly; 2 = Semi-monthly; 3 = bi-weekly; 4 = weekly; 5 = daily")
	fmt.Scanln(&paytype)
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

	//Calculate the basic tax
	grossamount = grossamount * payfre
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
	if grossamount < incomeclass["level1"] {
		taxamount = grossamount * incomeclass["rate1"]
	} else if grossamount < incomeclass["level2"] {
		taxamount = (grossamount-incomeclass["level1"])*incomeclass["rate2"] + fixamt1
	} else if grossamount < incomeclass["level3"] {
		taxamount = (grossamount-incomeclass["level2"])*incomeclass["rate3"] + fixamt2
	} else if grossamount < incomeclass["level4"] {
		taxamount = (grossamount-incomeclass["level3"])*incomeclass["rate4"] + fixamt3
	} else {
		taxamount = (grossamount-incomeclass["level4"])*incomeclass["rate5"] + fixamt4
	}
	//calculate the CPP and update for maximum amount.
	
	if grossamount <= 3500 {
		cppamt = 0 
	} else {
		if grossamount < 58700 {
			cppamt = (grossamount-3500) * incomeclass["kcpp"] 
		} else {
			cppamt = 2898
		}
	}
	//calculate the EI and update for maximum amount
	eiamt = grossamount * incomeclass["kei"] / payfre
	// K4= The lesser of: (i)   0.15 × A; and (ii)  0.15 × $1,161.
	if grossamount > 1245 {
		cecamt = 1245 * 0.15
	} else {
		cecamt = grossamount * 0.15
	}

	//less the federl personal amount T3 = (R × A) – K – K1 – K2 – K3 – K4

	if grossamount > incomeclass["fbpa"] {
		icamt = (taxamount - incomeclass["fbpac"] - cecamt - payfre*cppamt*0.15 - payfre*eiamt*0.15) / payfre
	} else {
		icamt = 0
	}

	fmt.Println("your gross amount:", grossamount)
	fmt.Println("Income tax payable:", taxamount)
	fmt.Println("Income tax payable:", icamt)
	fmt.Println("CPP tax payable:", cppamt)
	fmt.Println("EI tax payable:", eiamt)
	fmt.Println("EI tax payable:", cecamt)
}
