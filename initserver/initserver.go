package initserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

type ResStatus struct {
	Cpu    []StatusCpu    `json:"Cpu"`
	SizeMb float64        `json:"DiskMb"`
	Memory []StatusMemory `json:"Memory"`
	Consul bool           `json:"Consul"`
	Scp    bool           `json:"Scp"`
	Init   bool           `json:"Init"`
}
type StatusCpu struct {
	Fecha              time.Time `json:"Fecha"`
	CpuUsage           float64   `json:"CpuUsage"`
	IdleTicks          float64   `json:"IdleTicks"`
	TotalTicks         float64   `json:"TotalTicks"`
	CountCacheperMilli int32     `json:"CountCacheperMilli"`
	CountDbperMilli    int32     `json:"CountDbperMilli"`
}
type StatusMemory struct {
	Fecha      time.Time `json:"Fecha"`
	Alloc      uint64    `json:"Alloc"`
	TotalAlloc uint64    `json:"TotalAlloc"`
	Sys        uint64    `json:"Sys"`
	NumGC      uint32    `json:"NumGC"`
}
type ResInitServer struct {
	Encontrado bool    `json:"Encontrado"`
	Consulname string  `json:"Consulname"`
	Consulhost string  `json:"Consulhost"`
	AutoCache  bool    `json:"AutoCache"` // 0 AUTOMATICO - 1 LISTA CACHE
	ListaCache []int64 `json:"ListaCache"`
	TotalCache int32   `json:"TotalCache"`
	Files      []File  `json:"Files"`
}
type File struct {
	File string `json:"Files"`
	Ip   string `json:"Ip"`
}
type ReqInitServer struct {
	Id    string `json:"Id"`
	Ip    string `json:"Ip"`
	Token string `json:"Token"`
}

func Init(url string, post ReqInitServer) (ResInitServer, error) {

	var resp ResInitServer
	u, err := json.Marshal(post)
	if err != nil {
		return resp, errors.New("Marshal Error")
	}
	req := fasthttp.AcquireRequest()
	req.SetBody(u)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err == nil {
		body := res.Body()
		json.Unmarshal(body, &resp)
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)
		return resp, nil
	} else {
		return resp, errors.New("Request Error")
	}
}

type ReqStatus struct {
	Id         string `json:"Id"`
	Ip         string `json:"Ip"`
	Token      string `json:"Token"`
	PrimeraVez bool   `json:"PrimeraVez"`
}

func Status(url string, post ReqStatus) (ResStatus, error) {

	var resp ResStatus
	u, err := json.Marshal(post)
	if err != nil {
		return resp, errors.New("Marshal Error")
	}
	req := fasthttp.AcquireRequest()
	req.SetBody(u)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()

	if err := fasthttp.Do(req, res); err == nil {
		body := res.Body()
		json.Unmarshal(body, &resp)
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)
		return resp, nil
	} else {
		return resp, errors.New("Request Error")
	}
}
func GetMonitoringsCpu(countcache, countdb int32) StatusCpu {

	StatusCpu := StatusCpu{}

	idle0, total0 := GetCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := GetCPUSample()

	StatusCpu.Fecha = time.Now()
	StatusCpu.CountDbperMilli = countdb
	StatusCpu.CountCacheperMilli = countcache
	StatusCpu.IdleTicks = float64(idle1 - idle0)
	StatusCpu.TotalTicks = float64(total1 - total0)
	StatusCpu.CpuUsage = 100 * (StatusCpu.TotalTicks - StatusCpu.IdleTicks) / StatusCpu.TotalTicks

	return StatusCpu
}
func GetCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}
func GetMemoryPonderacion(mem []StatusMemory) int8 {

	/*
		for _, v := range mem {
			fmt.Printf("MemoryFecha %v / Alloc: %v / TotalAlloc: %v / Sys: %v / NumGC: %v\n", v.Fecha, v.Alloc, v.TotalAlloc, v.Sys, v.NumGC)
		}
	*/
	return 32
}
func GetCpuPonderacion(cpu []StatusCpu) int8 {
	/*
		for _, v := range cpu {
			fmt.Printf("Fecha: %v / CountCache: %v / CountDb: %v\n", v.Fecha, v.CountCacheperMilli, v.CountDbperMilli)
			fmt.Printf("CpuUsage: %v / IdleTicks: %v / TotalTicks: %v\n", v.CpuUsage, v.IdleTicks, v.TotalTicks)
		}
	*/
	return 32
}
func DirSize(path string) (float64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	sizeMB := float64(size) / 1024.0 / 1024.0
	return sizeMB, err
}
func PrintMemUsage() StatusMemory {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	StatusMemory := StatusMemory{}
	StatusMemory.Fecha = time.Now()
	StatusMemory.Alloc = BToMb(m.Alloc)
	StatusMemory.TotalAlloc = BToMb(m.TotalAlloc)
	StatusMemory.Sys = BToMb(m.Sys)
	StatusMemory.NumGC = m.NumGC
	return StatusMemory
}
func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func RemoveIndexCpu(s []StatusCpu, index int) []StatusCpu {
	return append(s[:index], s[index+1:]...)
}
func RemoveIndexMem(s []StatusMemory, index int) []StatusMemory {
	return append(s[:index], s[index+1:]...)
}
func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
