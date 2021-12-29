package utils

import (
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

func FileExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}
func Time_cu(t time.Duration, c int) string {
	ms := float64(t / time.Nanosecond)
	res := ms / float64(c)
	var s string
	if res < 1000 {
		s = fmt.Sprintf("%.2f NanoSec", res)
	} else if res >= 1000 && res < 1000000 {
		s = fmt.Sprintf("%.2f MicroSec", res/1000)
	} else {
		s = fmt.Sprintf("%.2f MilliSec", res/1000000)
	}
	return s
}
func Read_uint32(data []byte) uint32 {
	var x uint32
	for _, c := range data {
		x = x*10 + uint32(c-'0')
	}
	return x
}
func GetNum(post []int32, letras []int32, lens float64) (int32, error) {
	var x int32 = 0
	for k, _ := range post {
		value := post[len(post)-1-k]
		index := IndexOf(letras, value)
		if index == -1 {
			return value, errors.New("LETRA NO ENCONTRADA: ")
		}
		pos := int32(math.Pow(lens, float64(k)))
		x = x + index*pos
	}
	x++
	return x, nil
}
func IndexOf(arr []int32, candidate int32) int32 {
	for index, c := range arr {
		if c == candidate {
			return int32(index)
		}
	}
	return -1
}

// AWS METADATA //
func GetInstanceMeta(meta string) string {

	url := "http://169.254.169.254/latest/meta-data/" + meta
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)
	bodyBytes := resp.Body()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	return string(bodyBytes)

}

// AWS METADATA //
