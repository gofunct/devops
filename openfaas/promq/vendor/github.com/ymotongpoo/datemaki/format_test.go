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
	"testing"
	"time"
)

const (
	sec   time.Duration = time.Second
	min                 = time.Minute
	hour                = time.Hour
	day                 = time.Hour * 24
	month               = time.Hour * 31 * 24
	year                = time.Hour * 365 * 24
)

var (
	mid       = time.Date(2017, time.January, 15, 12, 30, 30, 0, time.Local)
	earlyEdge = time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)
	lateEdge  = time.Date(2017, time.December, 31, 23, 59, 59, 0, time.Local)
)

var formatAbsoluteDurationTests = []struct {
	label    string
	expect   string
	src      time.Time
	duration time.Duration
}{
	{"now", "now", mid, sec * 1},

	{"1 minute", "1 minute ago", mid, -min * 1},
	{"1 minute 2", "1 minute ago", mid, -min*1 - sec},
	{"1 minute 3", "1 minute later", mid, min * 1},
	{"1 minute 4", "1 minute later", mid, min*1 + sec},

	{"2 minutes 1", "2 minutes ago", mid, -min * 2},
	{"2 minutes 2", "2 minutes ago", earlyEdge, -min - sec*2},
	{"2 minutes 3", "2 minutes later", mid, min * 2},
	{"2 minutes 4", "2 minutes later", lateEdge, min + sec*2},

	{"1 hour", "1 hour ago", mid, -hour * 1},
	{"1 hour 2", "1 hour ago", mid, -hour*1 - min},
	{"1 hour 3", "1 hour later", mid, hour * 1},
	{"1 hour 4", "1 hour later", mid, hour*1 + min},

	{"2 hours 1", "2 hours ago", mid, -hour * 2},
	{"2 hours 2", "2 hours ago", earlyEdge, -hour - min*2},
	{"2 hours 3", "2 hours later", mid, hour * 2},
	{"2 hours 4", "2 hours later", lateEdge, hour + min*2},

	{"1 day", "1 day ago", mid, -day * 1},
	{"1 day 2", "1 day ago", mid, -day*1 - hour},
	{"1 day 3", "1 day later", mid, day * 1},
	{"1 day 4", "1 day later", mid, day*1 + hour},

	{"2 days 1", "2 days ago", mid, -day * 2},
	{"2 days 2", "2 days ago", earlyEdge, -day - hour*2},
	{"2 days 3", "2 days later", mid, day * 2},
	{"2 days 4", "2 days later", lateEdge, day + hour*2},

	{"1 month", "1 month ago", mid, -month * 1},
	{"1 month 2", "1 month ago", mid, -month*1 - day},
	{"1 month 3", "1 month later", mid, month * 1},
	{"1 month 4", "1 month later", mid, month*1 + day},

	{"2 months 1", "2 months ago", mid, -month * 2},
	{"2 months 2", "2 months ago", earlyEdge, -month - day*2},
	{"2 months 3", "2 months later", mid, month * 2},
	{"2 months 4", "2 months later", lateEdge, month + day*2},

	{"1 year", "1 year ago", mid, -year * 1},
	{"1 year 2", "1 year ago", mid, -year*1 - day},
	{"1 year 3", "1 year later", mid, year * 1},
	{"1 year 4", "1 year later", mid, year*1 + day},

	{"2 years 1", "2 years ago", mid, -year * 2},
	{"2 years 2", "2 years ago", earlyEdge, -year - day*2},
	{"2 years 3", "2 years later", mid, year * 2},
	{"2 years 4", "2 years later", lateEdge, year + day*2},
}

func TestFormatAbsoluteDuration(t *testing.T) {
	for _, tt := range formatAbsoluteDurationTests {
		dst := tt.src.Add(tt.duration)
		actual := FormatDurationFrom(tt.src, dst)
		if actual != tt.expect {
			t.Errorf("%s fails: expected='%s' actual='%s'", tt.label, tt.expect, actual)
		}
	}
}

var formatRelativeDurationTests = []struct {
	label    string
	expect   string
	src      time.Time
	duration time.Duration
}{
	{"1 day", "yesterday", mid, -day * 1},
	{"1 day 2", "yesterday", mid, -day*1 - hour},
	{"1 day 3", "tomorrow", mid, day * 1},
	{"1 day 4", "tomorrow", mid, day*1 + hour},

	{"2 days 1", "2 days ago", mid, -day * 2},
	{"2 days 2", "2 days ago", earlyEdge, -day - hour*2},
	{"2 days 3", "2 days later", mid, day * 2},
	{"2 days 4", "2 days later", lateEdge, day + hour*2},

	{"1 month", "last month", mid, -month * 1},
	{"1 month 2", "last month", mid, -month*1 - day},
	{"1 month 3", "next month", mid, month * 1},
	{"1 month 4", "next month", mid, month*1 + day},

	{"2 months 1", "2 months ago", mid, -month * 2},
	{"2 months 2", "2 months ago", earlyEdge, -month - day*2},
	{"2 months 3", "2 months later", mid, month * 2},
	{"2 months 4", "2 months later", lateEdge, month + day*2},

	{"1 year", "last year", mid, -year * 1},
	{"1 year 2", "last year", mid, -year*1 - day},
	{"1 year 3", "next year", mid, year * 1},
	{"1 year 4", "next year", mid, year*1 + day},

	{"2 years 1", "2 years ago", mid, -year * 2},
	{"2 years 2", "2 years ago", earlyEdge, -year - day*2},
	{"2 years 3", "2 years later", mid, year * 2},
	{"2 years 4", "2 years later", lateEdge, year + day*2},
}

func TestFormatRelativeDuration(t *testing.T) {
	for _, tt := range formatRelativeDurationTests {
		dst := tt.src.Add(tt.duration)
		actual := FormatRelativeDurationFrom(tt.src, dst)
		if actual != tt.expect {
			t.Errorf("%s fails: expected='%s' actual='%s'", tt.label, tt.expect, actual)
		}
	}
}
