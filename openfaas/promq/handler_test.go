package function

import (
	_ "regexp"
	"testing"
	"time"

	"github.com/ymotongpoo/datemaki"
)

func TestRequestDefaults(t *testing.T) {
	var json = []byte(`{"query": "up"}`)
	req, err := NewRequest(json)
	if err != nil {
		t.Fatalf("\nUnexpected error\nGot: \n%v", err)
	}

	expectedUrl := "http://prometheus.openfaas:9090"
	if req.Server != expectedUrl {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expectedUrl, req.Server)
	}

	expectedStep := "1m"
	if req.Step != expectedStep {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expectedStep, req.Step)
	}

	_, _, step, err := req.GetQueryRange()
	if err != nil {
		t.Fatalf("\nUnexpected error\nGot: \n%v", err)
	}

	expectedStepDuration, _ := time.ParseDuration(expectedStep)
	if step != expectedStepDuration {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expectedStepDuration, step)
	}
}

func TestRequestQueryRange(t *testing.T) {
	var json = []byte(`{"query": "up", "start": "12 hours ago", "end": "1 hour ago", "step": "15s"}`)
	req, err := NewRequest(json)
	if err != nil {
		t.Fatalf("\nUnexpected error\nGot: \n%v", err)
	}

	start, stop, step, err := req.GetQueryRange()
	if err != nil {
		t.Fatalf("\nUnexpected error\nGot: \n%v", err)
	}

	expectedStart, _ := datemaki.Parse("12 hours ago")
	if start.Hour() != expectedStart.Hour() {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expectedStart.Hour(), start.Hour())
	}

	expectedStop, _ := datemaki.Parse("1 hour ago")
	if stop.Hour() != expectedStop.Hour() {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expectedStop.Hour(), stop.Hour())
	}

	expectedStep, _ := time.ParseDuration("15s")
	if step != expectedStep {
		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expectedStep, step)
	}
}

//func TestHandlerJsonWithLocalProm(t *testing.T) {
//	var json = []byte(`{"server": "http://localhost:9090", "query": "sum(up) by (job)", "format": "json"}`)
//	expected := "prometheus"
//	resp := Handle(json)
//	r := regexp.MustCompile("(?m:" + expected + ")")
//	if !r.MatchString(resp) {
//		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
//	}
//}
//
//func TestHandlerTableWithRemoteProm(t *testing.T) {
//	var json = []byte(`{"server": "http://localhost:9090", "query": "sum(up) by (job)", "format": "table"}`)
//	expected := "job:prometheus"
//	resp := Handle(json)
//	r := regexp.MustCompile("(?m:" + expected + ")")
//	if !r.MatchString(resp) {
//		t.Fatalf("\nExpected: \n%v\nGot: \n%v", expected, resp)
//	}
//}
