package main

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func getSliceFromInt(inp int) []int {
	str := strconv.Itoa(inp)
	var out []int

	for _, s := range str {
		val, _ := strconv.Atoi(string(s))
		out = append(out, val)
	}

	return out
}

func createStrBinValue(base string, cnt int) string {
	var strBuilder strings.Builder

	for i := 0; i < cnt; i++ {
		strBuilder.WriteString(base)
	}

	return strBuilder.String()
}

func paddingLeadZeros(str string, len int) string {
	var resultString strings.Builder

	strLen := utf8.RuneCountInString(str)

	for i := strLen + 1; i <= len; i++ {
		resultString.WriteString("0")
	}

	resultString.WriteString(str)

	return resultString.String()
}

func sumMask(ints []int, mask string) int {
	var resultSum int

	if len(ints) > 0 {
		resultSum = ints[0]

		for i, sign := range mask {
			if string(sign) == "0" {
				resultSum = resultSum - ints[i+1]
			} else {
				resultSum = resultSum + ints[i+1]
			}
		}
	}

	return resultSum
}

func convertMaskToSign(mask string) string {
	var resultString strings.Builder
	for _, s := range mask {
		if string(s) == "0" {
			resultString.WriteString("-")
		} else {
			resultString.WriteString("+")
		}
	}

	return resultString.String()
}

func PlusMinus(inp int) string {
	var resultMask string

	intValues := getSliceFromInt(inp)
	cnt := len(intValues)

	maxValueStr := createStrBinValue("1", cnt-1)
	maxValue, _ := strconv.ParseInt(maxValueStr, 2, 64)

	var resultExists bool
	var i int64

	for i = 0; i <= maxValue; i++ {
		strBin := strconv.FormatInt(i, 2)
		strBin = paddingLeadZeros(strBin, cnt-1)

		res := sumMask(intValues, strBin)
		if res == 0 {
			resultExists = true
			resultMask = convertMaskToSign(strBin)

			break
		}
	}

	if !resultExists {
		resultMask = "unreal"
	}

	return resultMask
}
