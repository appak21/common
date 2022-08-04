package common

import (
	"strconv"
	"time"
)

//When formats and returns time like "5 seconds ago", "6 months ago",
//past time format must be based on RFC3339 or "2006-01-02 15:04:05".
//If an error occurs, it returns past time without formatting
func When(past string) string {
	dictionary := map[int][]string{
		0: {" year", " years"},
		1: {" month", " months"},
		2: {" day", " days"},
		3: {" hour", " hours"},
		4: {" minute", " minutes"},
		5: {" second", " seconds"},
	}
	now := time.Now().Format(time.RFC3339)
	k, till := 0, 4
	for i := 0; i < 6; i++ {
		if i != 0 {
			till = 2
		}
		pastTime, err1 := strconv.Atoi(past[k : k+till])
		currTime, err2 := strconv.Atoi(now[k : k+till])
		if err1 != nil || err2 != nil {
			return past
		}
		if t := currTime - pastTime; t == 1 {
			return strconv.Itoa(t) + dictionary[i][0] + " ago"
		} else if t > 1 {
			return strconv.Itoa(t) + dictionary[i][1] + " ago"
		}
		k = k + till + 1
	}
	return "just now"
}
