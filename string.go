package aefire

import (
	"github.com/dustin/go-humanize"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func AddDateHyphen(d, whenEmpty string) string {
	if len(d) == 0 {
		return whenEmpty
	} else if len(d) < 8 {
		return d
	} else {
		return d[:4] + "-" + d[4:6] + "-" + d[6:8]
	}
}

func HyphenWhenEmpty(s string) string {
	if len(s) == 0 {
		return "-"
	} else {
		return s
	}
}

func AddComma(s string) string {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return s
	} else {
		return humanize.Comma(i)
	}
}

func HyphenCorpNum(s string) string {
	if len(s) < 8 {
		return s
	} else if strings.Contains(s, "-") {
		return s
	} else {
		return s[:3] + "-" + s[3:5] + "-" + s[5:10]
	}
}

func HyphenPhoneNum(s string) string {
	if len(s) < 8 {
		return s
	}

	if s[0] == '+' {
		s = "0" + strings.TrimPrefix(s, "+82")
	}

	l := len(s)

	if l == 8 {
		return s[:4] + "-" + s[4:]
	} else {
		l0 := 0
		l1 := 0

		if strings.HasPrefix(s, "02") {
			l0 = 2
		} else {
			l0 = 3
		}

		l1 = l - l0 - 4

		return s[:l0] + "-" + s[l0:l0+l1] + "-" + s[l0+l1:]
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ToInt(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)

	PanicIfError(err)

	return i
}

//func IsJson(s string) bool{
//	if _, err := json.
//}
