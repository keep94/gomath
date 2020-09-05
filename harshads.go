package gomath

// Harshads generates the harshad numbers in order that are greater than or
// equal to start.
func Harshads(start int64) IntStream {
	if start < 1 {
		start = 1
	}
	return &harshadStream{start: start, sum: sumDigits(start)}
}

type harshadStream struct {
	start int64
	sum   int64
}

func (h *harshadStream) Next() (result int64, ok bool) {
	for {
		if h.start < 0 {
			return
		}
		if h.start%h.sum == 0 {
			result = h.start
			ok = true
			h.increment()
			return
		}
		h.increment()
	}
}

func (h *harshadStream) increment() {
	h.start++
	h.sum++
	temp := h.start
	for temp%10 == 0 {
		temp /= 10
		h.sum -= 9
	}
}

func sumDigits(n int64) int64 {
	result := int64(0)
	for n > 0 {
		result += n % 10
		n /= 10
	}
	return result
}
