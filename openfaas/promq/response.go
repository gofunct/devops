package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

func formatRespose(resp *QueryRangeResponse, format string) (string, error) {
	switch format {
	case "json":
		return responseToJSON(resp)
	case "table":
		return responseToTable(resp)
	}

	return "", fmt.Errorf("unknown format: %s", format)
}

func responseToJSON(resp *QueryRangeResponse) (string, error) {
	type valueEntry struct {
		Metric map[string]string `json:"metric"`
		Value  float64           `json:"value"`
	}
	type timeEntry struct {
		Time   int64         `json:"time"`
		Values []*valueEntry `json:"values"`
	}
	entryByTime := map[int64]*timeEntry{}

	for _, r := range resp.Data.Result {
		for _, v := range r.Values {
			t := v.Time()
			u := t.Unix()
			e, ok := entryByTime[u]
			if !ok {
				e = &timeEntry{
					Time:   u,
					Values: []*valueEntry{},
				}
				entryByTime[u] = e
			}

			val, err := v.Value()
			if err != nil {
				return "", err
			}
			e.Values = append(e.Values, &valueEntry{
				Metric: r.Metric,
				Value:  val,
			})
		}
	}

	s := make([]*timeEntry, len(entryByTime))
	i := 0
	for _, e := range entryByTime {
		s[i] = e
		i++
	}

	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func responseToTable(resp *QueryRangeResponse) (string, error) {
	type valueByMetric map[string]float64

	valuesByTime := map[time.Time]valueByMetric{}
	metrics := []string{}
	delimiter := "\t"

	for _, r := range resp.Data.Result {
		metric := stringMapToString(r.Metric, "|")
		for _, v := range r.Values {
			t := v.Time()
			d, ok := valuesByTime[t]
			if !ok {
				d = valueByMetric{}
				valuesByTime[t] = d
			}
			var err error
			d[metric], err = v.Value()
			if err != nil {
				return "", err
			}
		}

		found := false
		for _, m := range metrics {
			if m == metric {
				found = true
			}
		}
		if !found {
			metrics = append(metrics, metric)
		}
	}

	type st struct {
		time time.Time
		v    valueByMetric
	}
	slice := make([]st, len(valuesByTime))
	i := 0
	for t, v := range valuesByTime {
		slice[i] = st{t, v}
		i++
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].time.Before(slice[j].time)
	})

	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// header
	fmt.Fprintf(w, "time%s%s\n", delimiter, strings.Join(metrics, delimiter))

	// print rows
	for _, s := range slice {
		values := make([]string, len(metrics))
		for i, m := range metrics {
			if v, ok := s.v[m]; ok {
				values[i] = fmt.Sprintf("%f", v)
			} else {
				values[i] = ""
			}
		}
		fmt.Fprintf(w, "%d%s%s\n", s.time.Unix(), delimiter, strings.Join(values, delimiter))
	}
	w.Flush()
	return buf.String(), nil
}

func stringMapToString(m map[string]string, delimiter string) string {
	s := make([][]string, len(m))
	i := 0
	for k, v := range m {
		s[i] = []string{k, v}
		i++
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})

	ss := make([]string, len(s))
	for i, v := range s {
		ss[i] = fmt.Sprintf("%s:%s", v[0], v[1])
	}

	return strings.Join(ss, delimiter)
}
