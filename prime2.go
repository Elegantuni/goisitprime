package main

import (
	"fmt"
	"os"
	"math/big"
)

func isPrime(data1 string) int {
	var index1  = &big.Int{}
	var indextemp = &big.Int{}
	var tempzero = &big.Int{}

	index1.SetInt64(2)
	indextemp.SetInt64(1)
	tempzero.SetInt64(0)

	numberf,_ := new(big.Float).SetString(data1)
	var inttemp = &big.Int{}
	var floattemp = new(big.Float)
	thenumber,_ := new(big.Int).SetString(data1, 10)

	floattemp.Sqrt(numberf)
	floattemp.Int(inttemp)

	for (index1.Cmp(inttemp)) <= 0 {
		var theresult1 = &big.Int{}

		theresult1.Mod(thenumber, index1)

		if theresult1.Cmp(tempzero) == 0 {
			fmt.Println(data1, "isn't prime, because it is divisible by", index1)

			return 0
		}
		
		index1.Add(index1, indextemp)
	}

	return 1
}

func main() {
	var result1 int
	var data1temp string

	fmt.Println("Enter whole number: ")

	fmt.Scanln(&data1temp)

	result1 = isPrime(data1temp)

	if result1 == 0 {
		os.Exit(0)
	} else {
		fmt.Println(data1temp, "is prime")
	}
}



