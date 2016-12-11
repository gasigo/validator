package Validator

import (
	"regexp"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func IsValidCPF(cpf string) bool {
	multFactors1 := [9]int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	multFactors2 := [10]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpf = strings.Trim(cpf, " ")
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	if len(cpf) != 11 {
		return false
	}

	smallCPF1, error := sliceAtoi(strings.Split(cpf[0:9], ""))
	sum := 0

	if error != nil {
		return false
	}

	for i := 0; i < 9; i++ {
		sum += smallCPF1[i] * multFactors1[i]
	}

	rest := sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}

	smallCPF2 := append(smallCPF1, rest)
	sum = 0

	for i := 0; i < 10; i++ {
		sum += smallCPF2[i] * multFactors2[i]
	}

	rest = sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}

	intCPF := append(smallCPF2, rest)
	finalCPF := []string{}

	for i := range intCPF {
		number := intCPF[i]
		text := strconv.Itoa(number)
		finalCPF = append(finalCPF, text)
	}

	result := strings.Join(finalCPF, "")
	return cpf == result
}

func IsValidCNPJ(cnpj string) bool {

	multFactors1 := [12]int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multFactors2 := [13]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpj = strings.Trim(cnpj, " ")
	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)

	if len(cnpj) != 14 {
		return false
	}

	smallCNPJ1, error := sliceAtoi(strings.Split(cnpj[0:12], ""))
	sum := 0
	if error != nil {
		return false
	}

	for i := 0; i < 12; i++ {
		sum += smallCNPJ1[i] * multFactors1[i]
	}

	rest := sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}

	smallCNPJ2 := append(smallCNPJ1, rest)
	sum = 0

	for i := 0; i < 13; i++ {
		sum += smallCNPJ2[i] * multFactors2[i]
	}

	rest = sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}

	intCNPJ := append(smallCNPJ2, rest)
	finalCNPJ := []string{}

	for i := range intCNPJ {
		number := intCNPJ[i]
		text := strconv.Itoa(number)
		finalCNPJ = append(finalCNPJ, text)
	}

	result := strings.Join(finalCNPJ, "")
	return cnpj == result
}
