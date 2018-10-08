// Copyright 2015-2017 Yoshi Yamaguchi, Yoshiki Shibukawa
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
	"time"
)

type formatType int

const (
	relative formatType = iota
	absolute
)

// FormatDuration is almost shortcut of the following function call:
//
//     FormatRelativeDurationFrom(time.Now(), time.Now().Add(duration))
//
func FormatDuration(duration time.Duration) string {
	now := time.Now()
	return formatDurationFrom(now, now.Add(duration), relative)
}

// FormatRelativeDurationFrom generates human readable string of time.Duration like GitHub.
//
// For example, duration between src and dst is less than 60 seconds,
// this function returns "now".
//
// It returns text like "1 minute ago", "5 minutes later", "2 days ago".
func FormatRelativeDurationFrom(src, dst time.Time) string {
	return formatDurationFrom(src, dst, relative)
}

// FormatDurationFrom generates human readable string of time.Duration like GitHub.
//
// Compare with FormatRelativeDurationFrom, it returns "1 day ago" instead of "yesterday".
func FormatDurationFrom(src, dst time.Time) string {
	return formatDurationFrom(src, dst, absolute)
}

func daysOfYear(t time.Time) int {
	year := time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
	return year.YearDay()
}

func formatDurationFrom(src, dst time.Time, ft formatType) string {
	suffix := "later"
	inverted := false
	if src.After(dst) {
		suffix = "ago"
		inverted = true
		src, dst = dst, src
	}
	dur := dst.Sub(src)

	if dur < time.Minute {
		return "now"
	}

	if dur < time.Hour {
		minutes := (dst.Minute() - src.Minute()) % 60
		if minutes < 0 {
			minutes += 60
		}
		if minutes == 1 {
			return "1 minute " + suffix
		}
		return fmt.Sprintf("%d minutes %s", minutes, suffix)
	}

	if dur < time.Hour*24 {
		hours := (dst.Hour() - src.Hour()) % 24
		if hours < 0 {
			hours += 24
		}
		if hours == 1 {
			return "1 hour " + suffix
		}
		return fmt.Sprintf("%d hours %s", hours, suffix)
	}

	if dur < time.Hour*24*31 {
		days := dst.YearDay() - src.YearDay()
		if days < 0 {
			days += daysOfYear(src)
		}
		if days == 1 {
			if ft == relative {
				if inverted {
					return "yesterday"
				}
				return "tomorrow"
			} else {
				return "1 day " + suffix
			}
		}
		return fmt.Sprintf("%d days %s", days, suffix)
	}

	if dur < time.Hour*24*365 {
		months := int(dst.Month()) - int(src.Month())
		if months < 0 {
			months += 12
		}
		if months == 1 {
			if ft == relative {
				if inverted {
					return "last month"
				}
				return "next month"
			} else {
				return "1 month " + suffix
			}
		}
		return fmt.Sprintf("%d months %s", months, suffix)
	}

	years := dst.Year() - src.Year()
	if years == 1 {
		if ft == relative {
			if inverted {
				return "last year"
			}
			return "next year"
		} else {
			return "1 year " + suffix
		}
	}
	return fmt.Sprintf("%d years %s", years, suffix)
}

func FormatRalative(date time.Time) string {
	return ""
}
