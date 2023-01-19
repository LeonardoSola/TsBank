package tools

func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	var sum, rest int

	for i := 1; i <= 9; i++ {
		peso := 11 - i
		num := int(cpf[i-1] - '0')

		sum += peso * num
	}

	rest = (sum * 10) % 11

	if rest == 10 || rest == 11 {
		rest = 0
	}

	if rest != int(cpf[9]-'0') {
		return false
	}

	sum = 0

	for i := 1; i <= 10; i++ {
		peso := 12 - i
		num := int(cpf[i-1] - '0')

		sum += peso * num
	}

	rest = (sum * 10) % 11

	if rest == 10 || rest == 11 {
		rest = 0
	}

	if rest != int(cpf[10]-'0') {
		return false
	}

	return true
}

func ValidateCNPJ(str string) bool {
	size := len(str) - 2
	numbers := str[:size]
	digits := str[size:]

	sum := 0
	pos := size - 7

	for i := size; i >= 1; i-- {
		num := int(numbers[size-i] - '0')
		sum += num * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}

	result := sum % 11
	if result < 2 {
		result = 0
	} else {
		result = 11 - result
	}

	if result != int(digits[0]-'0') {
		return false
	}

	size++
	numbers = str[:size]
	sum = 0
	pos = size - 7

	for i := size; i >= 1; i-- {
		sum += int(numbers[size-i]-'0') * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}

	result = sum % 11
	if result < 2 {
		result = 0
	} else {
		result = 11 - result
	}

	if result != int(digits[1]-'0') {
		return false
	}

	return true
}
