package kubernet

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

type Kubernet struct {
	Configuracion Configuracion      `json:"Configuracion"`
	IpBlocks      []IpBlocks         `json:"IpBlocks"`
	Servicios     []Servicio         `json:"Servicios"`
	Servers       map[string]*Server `json:"Servers"`
	Archivos      []Archivo          `json:"Archivos"`
}
type Archivo struct {
	Tipo   int8   `json:"Tipo"`
	File   string `json:"File"`
	Ip     string `json:"Ip"`
	Rango1 int64  `json:"Rango1"`
	Rango2 int64  `json:"Rango2"`
}
type Configuracion struct {
	Ip           string    `json:"Ip"`
	Port         string    `json:"Port"`
	UltimoCambio time.Time `json:"UltimoCambio"`
}
type IpBlocks struct {
	Tiempo time.Time `json:"Tiempo"`
	Ip     string    `json:"IP"`
}
type Servicio struct {
	Tipo            int8             `json:"Tipo"` // 0 FILTROS - 1 AUTOCOMPLETE - 2 BUSQUEDAS
	Nombre          string           `json:"Nombre"`
	Valor           string           `json:"Valor"`
	ListadeBackends []ListadeBackend `json:"ListadeBackends"`
}
type ListadeBackend struct {
	Activo   bool      `json:"Activo"`
	Backends []Backend `json:"Backends"`
}
type Backend struct {
	Acls    []Acl      `json:"Acls"`
	Servers []ServerId `json:"Servers"`
}
type Acl struct {
	Param  string `json:"Param"`
	Tipo   int8   `json:"Tipo"`
	Valor1 int64  `json:"Valor1"`
	Valor2 int64  `json:"Valor2"`
}
type ServerId struct {
	Id string `json:"Id"`
}
type Server struct {
	Ip                   string   `json:"Ip"`
	Cpu                  int8     `json:"Cpu"`
	Memory               int8     `json:"Memory"`
	DiskMb               float64  `json:"Disk"`
	PosicionServicio     int      `json:"PosicionServicio"`
	PosicionListaBackend int      `json:"PosicionListaBackend"`
	PosicionBackend      int      `json:"PosicionBackend"`
	Iniciado             Iniciado `json:"Iniciado"`
}
type Iniciado struct {
	Consul bool `json:"Consul"`
	Scp    bool `json:"Scp"`
	Init   bool `json:"Init"`
}
type StatusCpu struct {
	Fecha      time.Time `json:"Fecha"`
	CpuUsage   float64   `json:"CpuUsage"`
	IdleTicks  float64   `json:"IdleTicks"`
	TotalTicks float64   `json:"TotalTicks"`
}
type StatusMemory struct {
	Fecha      time.Time `json:"Fecha"`
	Alloc      uint64    `json:"Alloc"`
	TotalAlloc uint64    `json:"TotalAlloc"`
	Sys        uint64    `json:"Sys"`
	NumGC      uint32    `json:"NumGC"`
}

func Create_config_file(Kubernet Kubernet) string {

	var b strings.Builder
	fmt.Fprintf(&b, "global\n\tlog /dev/log    local0\n\tlog /dev/log    local1 notice\n\tchroot /var/lib/haproxy\n\tstats socket ipv4@127.0.0.1:9999 level admin\n\tstats socket /var/run/haproxy.sock mode 666 level admin\n\tstats timeout 2m\n\tuser haproxy\n\tgroup haproxy\n\tdaemon\n\n\tca-base /etc/ssl/certs\n\tcrt-base /etc/ssl/private\n\n\tssl-default-bind-ciphers ECDH+AESGCM:DH+AESGCM:ECDH+AES256:DH+AES256:ECDH+AES128:DH+AES:RSA+AESGCM:RSA+AES:!aNULL:!MD5:!DSS\n\tssl-default-bind-options no-sslv3\n\n")
	fmt.Fprintf(&b, "defaults\n\tlog     global\n\tmode    http\n\toption  httplog\n\toption  dontlognull\n\ttimeout connect 5000\n\ttimeout client  50000\n\ttimeout server  50000\n\terrorfile 400 /etc/haproxy/errors/400.http\n\terrorfile 403 /etc/haproxy/errors/403.http\n\terrorfile 408 /etc/haproxy/errors/408.http\n\terrorfile 500 /etc/haproxy/errors/500.http\n\terrorfile 502 /etc/haproxy/errors/502.http\n\terrorfile 503 /etc/haproxy/errors/503.http\n\terrorfile 504 /etc/haproxy/errors/504.http\n\n")
	fmt.Fprintf(&b, "frontend apache_front\n\n\tbind *:80\n\n")

	for _, servicio := range Kubernet.Servicios {
		for i, lista := range servicio.ListadeBackends {
			if lista.Activo {
				switch servicio.Tipo {
				case 1:
					fmt.Fprintf(&b, "\tacl is_%s path_beg %s\n", servicio.Nombre, servicio.Valor)
				default:
					fmt.Fprintf(&b, "")
				}
				for j, backend := range lista.Backends {
					for _, acl := range backend.Acls {
						fmt.Fprintf(&b, "\tacl %s%d%d ", servicio.Nombre, i, j)
						switch acl.Tipo {
						case 1:
							fmt.Fprintf(&b, "urlp_val(%s) %d:%d\n", acl.Param, acl.Valor1, acl.Valor2)
						case 2:
							fmt.Fprintf(&b, "urlp_reg(%s) ^[%d-%d]\n", acl.Param, acl.Valor1, acl.Valor2)
						default:
							fmt.Fprintf(&b, "")
						}
					}
				}
				fmt.Fprintf(&b, "\n")
			}
		}
	}

	for _, servicio := range Kubernet.Servicios {
		for i, lista := range servicio.ListadeBackends {
			if lista.Activo {
				for j, backend := range lista.Backends {
					fmt.Fprintf(&b, "\tuse_backend bn%s%d if is_%s", servicio.Nombre, j, servicio.Nombre)
					for x := 0; x < len(backend.Acls); x++ {
						fmt.Fprintf(&b, " %s%d%d", servicio.Nombre, i, j)
					}
					fmt.Fprintf(&b, "\n")
				}
			}
		}
	}

	fmt.Fprintf(&b, "\n")
	for _, servicio := range Kubernet.Servicios {
		for i, lista := range servicio.ListadeBackends {
			if lista.Activo {
				for j, _ := range lista.Backends {
					fmt.Fprintf(&b, "backend bn%s%d\n\tbalance roundrobin\n\tserver-template mywebapp 10 _cn%s%d%d._tcp.service.consul resolvers consul resolve-opts allow-dup-ip resolve-prefer ipv4 check\n\ttimeout connect 1m\n\ttimeout server 1m\n\n", servicio.Nombre, j, servicio.Nombre, i, j)
				}
			}
		}
	}

	fmt.Fprintf(&b, "resolvers consul\n\tnameserver consul %s:%s\n\taccepted_payload_size 8192\n\thold valid 5s\n\n", Kubernet.Configuracion.Ip, Kubernet.Configuracion.Port)
	return b.String()

}

type ReqInitServer struct {
	Id           string    `json:"Id"`
	Ip           string    `json:"Ip"`
	UltimoCambio time.Time `json:"UltimoCambio"`
}
type ResInitServer struct {
}

func NewConfig(url string, post ReqInitServer) (ResInitServer, error) {

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
func RemoveServerId(s []ServerId, index int) []ServerId {
	return append(s[:index], s[index+1:]...)
}
