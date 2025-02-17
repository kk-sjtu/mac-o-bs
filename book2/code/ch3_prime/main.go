package main

import (
	"fmt"
	"github.com/otiai10/primes"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage:", os.Args[0], "<number>")
		os.Exit(1)
	}
	number, err := strconv.Atoi(args[0])
	/*
		strconv.Atoi(s string) (int, error)：将字符串转换为整数。
		strconv.Itoa(i int) string：将整数转换为字符串。
		strconv.ParseBool(str string) (bool, error)：将字符串转换为布尔值。
		strconv.ParseFloat(s string, bitSize int) (float64, error)：将字符串转换为浮点数。
		strconv.ParseInt(s string, base int, bitSize int) (int64, error)：将字符串转换为指定进制和位数的整数。
		strconv.ParseUint(s string, base int, bitSize int) (uint64, error)：将字符串转换为无符号整数
	*/
	if err != nil {
		panic(err)
	}
	f := primes.Factorize(int64(number))
	fmt.Println("primes:", len(f.Powers()) == 1)
}
