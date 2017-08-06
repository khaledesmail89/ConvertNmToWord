package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
var arabicOnes = []string{
	"صفر","واحد", "اثنان", "ثلاثة", "أربعة", "خمسة", "ستة", "سبعة", "ثمانية", "تسعة",
	"عشرة", "أحد عشر", "اثنا عشر", "ثلاثة عشر", "أربعة عشر", "خمسة عشر", "ستة عشر", "سبعة عشر", "ثمانية عشر", "تسعة عشر",
}
var arabicTens = []string{
	"", "عشرون", "ثلاثون", "أربعون", "خمسون",
	"ستون", "سبعون", "ثمانون", "تسعون",
}
var arabicHundreds = []string{
	"", "مائة", "مئتان", "ثلاثمائة", "أربعمائة", "خمسمائة", "ستمائة", "سبعمائة", "ثمانمائة","تسعمائة",
}
var arabicTwos = []string{
	"مئتان", "ألفان", "مليونان", "ملياران", "تريليونان", "كوادريليونان", "كوينتليونان", "سكستيليونان",
}

var arabicGroup = []string{
	"مائة", "ألف", "مليون", "مليار", "تريليون", "كوادريليون", "كوينتليون", "سكستيليون",
}

var arabicPluralGroup = []string{
	"", "آلاف", "ملايين", "مليارات", "تريليونات", "كوادريليونات", "كوينتليونات", "سكستيليونات",
}

func main() {

	fmt.Print("Enter Number: ")
	var pounds, piasters, input string
	fmt.Scanf("%s",&input)
	var money[]string
	money=strings.Split(input, ".")
	pounds=money[0]
	output := Convert(pounds)
	output += " جنيهاً "
	if len(money) > 1 {
		piasters = money[1]
		digit := ConvertString2Digit(piasters)
		output += "و "
	if digit == 1 {
		output += Hundreds(digit)
	output += " قرش "
	}else if digit == 2 {
			output += " قرشان "
		}else {
		output += Hundreds(digit)
		if digit > 2 && digit <= 10 {
			output += " قروش "
		} else {
			output += " قرشاً "
		}
	}
	}
	fmt.Println(output)
}

func Convert(text string)string {
	var arr[]string
	var idx int = 0
	length := len(text)
	arrLen := math.Ceil(math.Floor(float64(length)) / 3)
	arr = make([]string, int(arrLen))
	for i := len(text); i >= 0; i -= 3 {
		if i >= 3 {
			arr[idx] = text[i - 3:i]
			idx++
		} else if i != 0 {
			arr[idx] = text[:i]
		}
	}
	// here i want to check numbers and matched words from arrays.
	var retVal string
	arrLength := len(arr)
	for j := arrLength; j > 0; j-- {
		digit := ConvertString2Digit(arr[j - 1])
		if digit != 0 {
			if j == arrLength {
				if digit == 1 {
					if arrLength > 1 {
						retVal += arabicGroup[j - 1]
					} else {
						retVal += Hundreds(digit)
					}
				} else if digit == 2 {
					if arrLength > 1 {
						retVal += arabicTwos[j - 1]
					} else {
						retVal += Hundreds(digit)
					}
				} else if digit > 2 && digit <= 10 {
					retVal += Hundreds(digit)
					retVal += arabicPluralGroup[j - 1]
				} else {
					retVal += Hundreds(digit)
					retVal += " "
					if arrLength > 1 {
						retVal += arabicGroup[j - 1]
					}
				}
			} else {
				retVal += " و "
				retVal += Hundreds(digit)
				retVal += " "
				if j > 1 {
					retVal += arabicGroup[j - 1]
				}
			}
		}
	}
return retVal
}
func ConvertString2Digit(word string)float64  {
	digit,_:= strconv.Atoi(word)
	return float64(digit)
}
func Tens (num float64)string{

	if num < 20{
		return  arabicOnes[int(num)]
	}else{
		mod := math.Mod(num,10)
		num/=10
		if mod == 0 {

			return arabicTens[int(math.Floor(num-1))]
		}else {
			return arabicOnes[int(mod)] + " و " + arabicTens[int(math.Floor(num - 1))]
		}
	}
}
func Hundreds (num float64)string{
	len:= math.Log10(num)
	//	check if tens
	if math.Floor(len+1) <= 2 {
		return Tens(num)
	}
	var retVal string
	mod := math.Mod(num,100)
	num/=100
	if mod == 0{
		retVal += arabicHundreds[int(math.Floor(num))]
	}else{
		retVal += arabicHundreds[int(math.Floor(num))]
		retVal += " و "
		retVal += Tens(mod)
	}
	return retVal
}