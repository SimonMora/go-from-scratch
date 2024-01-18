package funciones

var originalNumber int
var poweredNumber int

func powCall(number int, pow int, isRecursive bool) {
	if !isRecursive {
		originalNumber = number
	}

	if pow == 0 && !isRecursive {
		poweredNumber = 1
		return
	}

	if pow == 1 && isRecursive {
		poweredNumber = number
		return
	}

	newpow := pow - 1
	powCall(number*originalNumber, newpow, true)
}

func Pow(number int, pow int) int {
	powCall(number, pow, false)
	return poweredNumber
}
