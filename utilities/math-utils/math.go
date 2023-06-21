package math_utils

import (
	"fmt"
	"math"
	"strconv"
)

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	limit := math.Sqrt(float64(n))
	max := math.Ceil(limit)
	for i := 3; i <= int(max); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// es1
func SommaMultipli3_5() {

	tot := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			tot += i
		}
	}
	fmt.Println(tot)
}

// es2
func SerieFibonacci() {

	tot := 0
	prec := 1
	for n := 1; n < 4000000; n = n + prec {
		prec = n - prec
		fmt.Println(n)
		if n%2 == 0 {
			tot += n
		}
	}
	fmt.Println("somma dei numeri pari: ", tot)
}

// es3
func MaxFattorePrimo(num int) {

	max := 0

	for i := num; i >= 2; i-- {
		fmt.Println(i)
		if num%i == 0 {
			if IsPrime(i) {
				max = i
			}
		}
	}
	fmt.Println("il fattore primo massimo è: ", max)
}

// es4
func LargestPalindromeProduct(c int) {

	max := 0
	for i := int(math.Pow(10, float64(c)) - 1); i > int(math.Pow(10, float64(c-1))-1); i-- {
		for j := int(math.Pow(10, float64(c)) - 1); j > int(math.Pow(10, float64(c-1))-1); j-- {
			if IsPalindrome(fmt.Sprintf(strconv.Itoa(i*j))) && i*j > max {
				max = i * j
				break
			} else {
				if IsPalindrome(fmt.Sprintf(strconv.Itoa(i * j))) {
					fmt.Println("Il prodotto palindromo è ", max)
					return
				}
			}
		}
	}
}

// es5
func SmallestMultiple(a, b int) {

	i := 0
	for i = 1; IsPrime(i); i++ {
	}
	fmt.Println(i)
}

// es6
func SumSquares(x, y int) {

	sum := 0
	for i := x; i <= y; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	sum2 := 0
	for i := y; i >= x; i-- {
		sum2 += i
	}
	sum2 = int(math.Pow(float64(sum2), 2))
	fmt.Println("La somma dei numeri quadrati è", sum, "invece la sommma quadrata dei numeri è", sum2, "la differenza è", sum2-sum)
}

// es7
func FindNPrime(n int) {

	i := 0
	j := 1

	for {
		j++

		if !IsPrime(j) {
			continue
		}
		i++
		if i == n {
			fmt.Println(j)
			return
		}
	}
}
