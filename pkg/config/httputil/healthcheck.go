package http

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type endpoint 		string
type datacenter 	string
type environment 	string
type urltype 		string


type endpoints 		[]endpoint
type datacenters	[]datacenter
type environments	[]environment
type urltypes		[]urltype

type DCEndpoints 			map[datacenter](endpoints)
type EnvEndpoints 			map[environment](endpoints)
type TypeEndpoints 			map[urltype](endpoints)

var ProdCheck = map[string][]string{}

type HealthCmd struct {
	Endpoint 		endpoint
	Environment		EnvEndpoints
	Datacenter  	string
	Type 			string
}




type HealCheckCmd interface {
	CheckSingle() bool
	CheckEnvironment() bool
	CheckDatacenter() bool
	CheckType() bool
	CheckAll() bool
}

func (h *HealthCmd) CheckSingle() bool {
	return len(h.Single) != 0
}

func (h *HealthCmd) HasEnvironment() bool {
	return len(h.Environment) != 0
}

func (h *HealthCmd) HasDatacenter() bool {
	return len(h.Datacenter) != 0
}

func (h *HealthCmd) HasType() bool {
	return len(h.Type) != 0
}

func (h *HealthCmd) HasAll() bool {
	return h.All != false
}

func (h *HealthCmd) FQDNHealth() {
	_, err := flags.Parse(&HealthCmd{})
	if err != nil {
		os.Exit(1)
	}
	if h.FQDN != "" {
		CheckAll()
	}
}

//Returns the URL, SUCCESS/FAIL, and Status Code
func CheckOne() {
	url := health.FQDN
	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	resp, err := c.Get(url)
	if err != nil {
		fmt.Println(url, err, "😱")
	}
	if err == nil {
		if resp.StatusCode != 200 {
			fmt.Println(url, "😱", resp.Status)
		} else {
			fmt.Println(url, "👍", resp.Status)
		}
	}
}

//Returns the URL, SUCCESS/FAIL, and Status Code
func CheckAll() {
	for _, url := range CertdEndpoints {
		c := &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 10 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
		resp, err := c.Get(url)
		if err != nil {
			fmt.Println(url, err, "😱")
		}
		if err == nil {
			if resp.StatusCode != 200 {
				fmt.Println(url, "😱", resp.Status)
			} else {
				fmt.Println(url, "👍", resp.Status)
			}
		}
	}
}

func init() {

	fmt.Println("initializing endpoints tests....")
	_, err := flags.Parse(&HealthCmd{})
	if err != nil {
		os.Exit(1)
	}
	if health.HasAll() {
		return urls = CertdEndpoints
	}
	CheckAll(urls)
}
