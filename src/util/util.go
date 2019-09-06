package util

import "strconv"

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func GetCoinValue(value string) (uint64, error) {
	coin, err := strconv.ParseFloat(value, 64)
	return uint64(coin * 1000000), err
}

func Format(n uint64) string {
    in := strconv.FormatUint(n, 10)
    out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
    if in[0] == '-' {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}