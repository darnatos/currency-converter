package objects

import (
	"errors"
	"strconv"
)

type Currencies struct {
	Currencies map[string]map[string]float64 `json:"currencies"`
}

type ICurrency interface {
	Rate(from string, to string) (float64, error)
	Convert(amount string, from string, to string) (string, error)
}

func (c Currencies) Rate(from string, to string) (float64, error) {
	if _, ok := c.Currencies[from]; !ok {
		return 0, errors.New("no fromCurrency: " + from)
	}
	if rate, ok := c.Currencies[from][to]; !ok {
		return 0, errors.New("no toCurrency: " + to)
	} else {
		return rate, nil
	}
}

// Convert 轉換
func (c Currencies) Convert(amount string, from string, to string) (string, error) {
	rate, err := c.Rate(from, to)
	if err != nil {
		return "", err
	}
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}
	toAmt := amt * rate

	sign := false
	if toAmt < 0 {
		sign = true
		toAmt = -toAmt
	}

	var res string
	if sign {
		res = "-" + formatAmount(strconv.FormatFloat(toAmt, 'f', 2, 64))
	} else {
		res = formatAmount(strconv.FormatFloat(toAmt, 'f', 2, 64))
	}

	return res, nil
}

func formatAmount(num string) string {
	res := make([]byte, 0, (len(num)/3+1)*4)
	pointIndex := len(num) - 3
	for i := pointIndex - 1; i >= 0; i-- {
		if pointIndex-i-1 > 0 && (pointIndex-i-1)%3 == 0 {
			res = append(res, ',')
		}
		res = append(res, num[i])
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res) + num[pointIndex:]
}
