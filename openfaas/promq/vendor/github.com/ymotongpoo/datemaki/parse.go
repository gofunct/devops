// Copyright 2015-2017 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datemaki

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	numericExp    = regexp.MustCompile(`^[0-9/\.\ \-:]+$`)
	hhmmExp       = regexp.MustCompile(`[0-9]{1,2}:[0-9]{1,2}`)
	hhmmssExp     = regexp.MustCompile(`[0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2}`)
	ordinalDayExp = regexp.MustCompile(`[0-9]{1,2}(th|st|nd|rd)`)
	timezoneExp   = regexp.MustCompile(`^[0-9]{1,4}-[0-9]{1,2}-[0-9]{1,2} [0-9]{1,2}:[0-9]{1,2}(:[0-9]{1,2})? (\+|-)([0-9]{1,2})?:?[0-9]{2}$`)
	unixZero      = time.Unix(0, 0)
)

var fullMonth = map[string]time.Month{
	"january":   time.January,
	"february":  time.February,
	"march":     time.March,
	"april":     time.April,
	"may":       time.May,
	"june":      time.June,
	"july":      time.July,
	"august":    time.August,
	"september": time.September,
	"october":   time.October,
	"november":  time.November,
	"december":  time.December,
}

var shortMonth = map[string]time.Month{
	"jan": time.January,
	"feb": time.February,
	"mar": time.March,
	"apr": time.April,
	"may": time.May,
	"jun": time.June,
	"jul": time.July,
	"aug": time.August,
	"sep": time.September,
	"oct": time.October,
	"nov": time.November,
	"dec": time.December,
}

// Parse accepts contextful date format and returns absolute time.Time value.
func Parse(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	switch {
	case strings.HasSuffix(value, "ago"):
		return ParseAgo(value)
	case hasRelative(value):
		return ParseRelative(value)
	default:
		return ParseAbsolute(value)
	}
	return time.Now().In(time.Local), nil // TODO(ymotongpoo): replace actual time.
}

// MustParse is like Parse but panics if the passed valuecannot be parsed.
// It simplifies safe initialization of global variables holding parsed time.
func MustParse(value string) time.Time {
	result, err := Parse(value)
	if err != nil {
		panic(err)
	}
	return result
}

// splitTokens splits value with commas, periods and spaces.
// Currently, it only expects single byte character tokenizer.
func splitTokens(value string) []string {
	f := func(c rune) bool {
		return c == rune(' ') || c == rune(',') || c == rune('.') || c == rune('/') || c == rune('-')
	}
	return strings.FieldsFunc(value, f)
}

// hasRelative confirms if value contains relative datatime words, such as
// "now", "today", "last xxx", "noon", "pm", "am"  and so on.
func hasRelative(value string) bool {
	keywords := []string{"now", "today", "yesterday", "last"}
	for _, k := range keywords {
		if strings.Contains(value, k) {
			return true
		}
	}
	return false
}

// ParseAgo parse "xxxx ago" format and returns corresponding absolute datetime.
func ParseAgo(value string) (time.Time, error) {
	tokens := splitTokens(value)
	now := time.Now().In(time.Local)
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t == "ago" {
			return now, nil
		}
		if i%2 == 0 {
			var err error
			n, err := strconv.Atoi(t)
			if err != nil {
				return unixZero, fmt.Errorf("Format error: %v", t)
			}
			now, err = subDate(now, n, tokens[i+1])
			if err != nil {
				return unixZero, err
			}
			i++
		}
	}
	return now, nil
}

// subDate subtracts n*unit duration from t and return the result.
// supportes units are "year", "month", "week", "day", "hour", "minute", "second", and those plurals.
func subDate(t time.Time, n int, unit string) (time.Time, error) {
	if strings.HasSuffix(unit, "s") {
		unit = string([]byte(unit)[:len(unit)-1])
	}
	switch unit {
	case "year":
		return t.AddDate(-1*n, 0, 0), nil
	case "month":
		return t.AddDate(0, -1*n, 0), nil
	case "week":
		return t.AddDate(0, 0, -7*n), nil
	case "day":
		return t.AddDate(0, 0, -1*n), nil
	case "hour":
		return t.Add(time.Duration(-1*n) * time.Hour), nil
	case "minute":
		return t.Add(time.Duration(-1*n) * time.Minute), nil
	case "second":
		return t.Add(time.Duration(-1*n) * time.Second), nil
	default:
		return t, fmt.Errorf("Unsupported time unit: %v", unit)
	}
}

// ParseRelative returns absolute datetime corresponding to relative date expressed in value.
func ParseRelative(value string) (time.Time, error) {
	tokens := splitTokens(value)
	var t time.Time
	t = time.Now().In(time.Local)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "last":
			days, err := daysFromLast(tokens[i+1])
			if err != nil {
				return t, err
			}
			i++
			t = t.Add(time.Duration(-1*days) * time.Hour * 24)
		case "yesterday":
			t = t.Add(time.Duration(-24 * time.Hour))
		case "today":
			// pass
		case "noon", "tea", "midnight":
			var err error
			t, err = convertTimeWord(tokens[i])
			if err != nil {
				return t, err
			}
		case "now", "never":
			return convertTimeWord(tokens[i])
		default:
			var t1 time.Time
			var err error
			t1, err = parse12HourClock(tokens[i])
			if err != nil {
				t1, err = parseNumericTime(tokens[i])
				if err != nil {
					return unixZero, fmt.Errorf("Unexpected time value, %v: %v", value, err)
				}
			}
			t = time.Date(t.Year(), t.Month(), t.Day(), t1.Hour(), t1.Minute(), t1.Second(), 0, time.Local)
		}
	}
	return t, nil
}

// daysFromLast returns days from last weekday passed to
func daysFromLast(weekday string) (int, error) {
	now := time.Now().In(time.Local)
	var day int
	switch strings.ToLower(weekday) { // time.Weekday defines value.
	case "sunday":
		day = 0
	case "monday":
		day = 1
	case "tuesday":
		day = 2
	case "wednesday":
		day = 3
	case "thursday":
		day = 4
	case "friday":
		day = 5
	case "saturday":
		day = 6
	default:
		return 0, fmt.Errorf("%v is not weekday", weekday)
	}
	return int(now.Weekday()) - day + 7, nil
}

// convertTimeWord converts words of time of day to numerial expression.
func convertTimeWord(word string) (time.Time, error) {
	now := time.Now().In(time.Local)
	switch word {
	case "now":
		return now, nil
	case "noon":
		return time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.Local), nil
	case "tea":
		return time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, time.Local), nil
	case "midnight":
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local), nil
	case "never":
		return unixZero, nil
	}
	return now, fmt.Errorf("Unsupported time word: %v", word)
}

// parse12HourClock convers 12-hour clock time to 24-hour one.
func parse12HourClock(word string) (time.Time, error) {
	lower := strings.ToLower(word)
	now := time.Now().In(time.Local)

	start := 0
	hour := 0
	var err error
	for width := 0; start < len(lower); start += width {
		var r rune
		r, width = utf8.DecodeRuneInString(lower[start:])
		if !unicode.IsNumber(r) {
			hour, err = strconv.Atoi(lower[:start])
			if err != nil || hour > 12 || hour < 0 {
				return time.Now(), fmt.Errorf("Wrong hour: %v", word)
			}
			if string(lower[start:]) == "am" {
				break
			}
			if string(lower[start:]) == "pm" {
				hour += 12
				break
			}
			return time.Now(), fmt.Errorf("Unsupported 12 hour clock notation: %v", word)
		}
	}
	return time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local), nil
}

// ParseAbsolute converts absolute datetime into time.Time. Basic idea is same as time.Parse(),
// but this detects the format of value and convert it automatically.
func ParseAbsolute(value string) (time.Time, error) {
	if numericExp.MatchString(value) {
		return parseNumeric(value)
	}

	tokens := splitTokens(strings.ToLower(value))
	year, day := 0, 0
	var month time.Month
	monthParsed := false
	var t time.Time
	for _, token := range tokens {
		var ok bool
		var err error
		switch {
		case len(token) == 4 && numericExp.MatchString(token):
			year, err = strconv.Atoi(token)
			if err != nil {
				return unixZero, fmt.Errorf("%v, Unexpected year value: %v", value, err)
			}
		case ordinalDayExp.MatchString(token):
			day, err = strconv.Atoi(token[:len(token)-2])
			if err != nil {
				return unixZero, fmt.Errorf("%v, Unexpected day value: %v", value, err)
			}
		case strings.HasSuffix(token, "am") || strings.HasSuffix(token, "pm"):
			t, err = parse12HourClock(token)
			if err != nil {
				return unixZero, fmt.Errorf("%v, Unexpected 12-hour clock time: %v", value, err)
			}
		case strings.Index(token, ":") != -1:
			t, err = parseNumericTime(token)
			if err != nil {
				return unixZero, fmt.Errorf("%v, Unexpected numeric time: %v", value, err)
			}
		case monthParsed:
			day, err = strconv.Atoi(token)
			if err != nil || day > 31 || day < 0 {
				return unixZero, fmt.Errorf("%v, Unexpected day value: %v", value, err)
			}
		default:
			month, ok = fullMonth[token]
			if ok {
				monthParsed = true
				break
			}
			month, ok = shortMonth[token]
			if ok {
				monthParsed = true
				break
			}
			m, err := strconv.Atoi(token)
			if err != nil || m > 12 || m < 0 {
				return unixZero, fmt.Errorf("%v, Unexpected month value, %v, error: %v", value, m, err)
			}
			month = time.Month(m)
			monthParsed = true
		}
	}

	if year == 0 {
		year = time.Now().Year()
	}

	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), 0, time.Local), nil
}

// parseNumericTime converts a time expressed in digits to time.Time.
func parseNumericTime(value string) (time.Time, error) {
	now := time.Now()
	var t time.Time
	var err error
	if hhmmssExp.MatchString(value) {
		t, err = time.Parse("15:04:05", value)
		if err != nil {
			return parseBroadcasterTime(value)
		}
		return t, nil
	} else if hhmmExp.MatchString(value) {
		t, err = time.Parse("15:04", value)
		if err != nil {
			return parseBroadcasterTime(value)
		}
		return t, nil
	}
	return time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local), nil
}

// parseBroadcasterTime converts a non-standard time expression used in TV schedules into normal one.
// eg. 30:00 (= 6am tomorrow)
//
// TODO(ymotongpoo): tentative implementation for Go Advent Calendar 2015
func parseBroadcasterTime(value string) (time.Time, error) {
	day := time.Now()
	var err error
	var hh, mm, ss int
	if hhmmssExp.MatchString(value) {
		tokens := strings.Split(value, ":")
		hh, err = strconv.Atoi(tokens[0])
		if err != nil {
			return day, fmt.Errorf("HHMMSS Unexpected format: %v", value)
		}
		mm, err = strconv.Atoi(tokens[1])
		if err != nil {
			return day, fmt.Errorf("HHMMSS Unexpected format: %v", value)
		}
		ss, err = strconv.Atoi(tokens[2])
		if err != nil {
			return day, fmt.Errorf("HHMMSS Unexpected format: %v", value)
		}
		if ss > 60 {
			mm += ss / 60
			ss = ss % 60
		}
		if mm > 60 {
			hh += mm / 60
			mm = mm % 60
		}
		if hh > 24 {
			day = day.AddDate(0, 0, hh/24)
			hh = hh % 24
		}
		return time.Date(day.Year(), day.Month(), day.Day(), hh, mm, ss, 0, time.Local), nil
	} else if hhmmExp.MatchString(value) { // TODO(ymotongpoo): merge into the condition above by treating this as HH:MM:00.
		tokens := strings.Split(value, ":")
		hh, err = strconv.Atoi(tokens[0])
		if err != nil {
			return day, fmt.Errorf("HHMM Unexpected format: %v", value)
		}
		mm, err = strconv.Atoi(tokens[1])
		if err != nil {
			return day, fmt.Errorf("HHMM Unexpected format: %v", value)
		}
		ss = 0
		if mm > 60 {
			hh += mm / 60
			mm = mm % 60
		}
		if hh > 24 {
			day = day.AddDate(0, 0, hh/24)
			hh = hh % 24
		}
	}
	return time.Date(day.Year(), day.Month(), day.Day(), hh, mm, ss, 0, time.Local), nil
}

// parseNumeric convers a datetime expressed all in digits to time.Time.
func parseNumeric(value string) (time.Time, error) {
	tokens := splitTokens(value)
	now := time.Now()
	year, month, day := 0, 0, 0
	var t time.Time
	for _, token := range tokens {
		var err error
		switch {
		case len(token) == 4:
			year, err = strconv.Atoi(token)
			if err != nil { // time package can handle days before unixtime 0.
				return now, fmt.Errorf("Error on parsing year: %v", value)
			}
		case len(token) == 2 || len(token) == 1:
			if month == 0 {
				month, err = strconv.Atoi(token)
				if err != nil || month < 0 || month > 12 {
					return now, fmt.Errorf("Error on parsing month: %v", value)
				}
			} else if day == 0 {
				day, err = strconv.Atoi(token)
				if err != nil || day < 0 || day > 31 {
					return now, fmt.Errorf("Error on parsing day: %v", value)
				}
			}
		case hhmmssExp.MatchString(token) || hhmmExp.MatchString(token):
			t, err = parseNumericTime(token)
			if err != nil {
				return now, fmt.Errorf("HHMMSS Unexpected format: %v, error: %v", value, err)
			}
		}
	}
	if year == 0 {
		year = now.Year()
	}
	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), 0, time.Local), nil
}
