package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
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
func Random(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max-1))
	return n.Int64() + 1
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

func escribir_file(path string, numb int) {
	d1 := []byte("")
	c := 0

	aux := numb / 100

	now := time.Now()
	for n := 0; n < aux; n++ {
		folder := getFolder64(int64(n * 100))
		newpath := filepath.Join(path, folder)
		err := os.MkdirAll(newpath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			fmt.Println("FOLDER ERROR: ", err)
		}
		for i := 0; i < 100; i++ {
			err := os.WriteFile(path+"/"+folder+"/"+strconv.Itoa(i), d1, 0644)
			if err != nil {
				fmt.Println(err)
			}
			c++
		}
	}
	elapsed := time.Since(now)
	fmt.Printf("WRITES FILES %v [%s] c/u total %v\n", c, Time_cu(elapsed, c), elapsed)
}
func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}
func getFolder64(num int64) string {
	c1, n1 := divmod(num, 1000000)
	c2, n2 := divmod(n1, 10000)
	c3, _ := divmod(n2, 100)
	return strconv.FormatInt(c1, 10) + "/" + strconv.FormatInt(c2, 10) + "/" + strconv.FormatInt(c3, 10)
}
func getFolderFile64(num int64) string {
	c1, n1 := divmod(num, 1000000)
	c2, n2 := divmod(n1, 10000)
	c3, c4 := divmod(n2, 100)
	return strconv.FormatInt(c1, 10) + "/" + strconv.FormatInt(c2, 10) + "/" + strconv.FormatInt(c3, 10) + "/" + strconv.FormatInt(c4, 10)
}
