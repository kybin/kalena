package main

import "time"

func monthBefore(queryMonth int) int {
	switch queryMonth {
	case 1:
		return 12
	default:
		return queryMonth - 1
	}
}

func monthAfter(queryMonth int) int {
	switch queryMonth {
	case 12:
		return 1
	default:
		return queryMonth + 1
	}
}

func yearBefore(queryYear, queryMonth int) int {
	switch queryMonth {
	case 1:
		return queryYear - 1
	default:
		return queryYear
	}
}

func yearAfter(queryYear, queryMonth int) int {
	switch queryMonth {
	case 12:
		return queryYear + 1
	default:
		return queryYear
	}
}

func offset(year, month int) int {
	t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	return int(t.Weekday())
}

func onlyDate(str string) string {
	if regexWebdateTime.MatchString(str) {
		return str[len(str)-2 : len(str)]
	}
	return str
}
