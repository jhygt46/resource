package consul

import (
	"fmt"
	"net"
	"time"
    "github.com/hashicorp/consul/api"
)
type ConsulRegister struct {
	Address                        string
	Name                           string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}
func ConsulRegisters(name string, consuladress string) bool {

	s := NewConsulRegister(name, consuladress)
	config := api.DefaultConfig()
	config.Address = s.Address
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}
	agent := client.Agent()

	IP := LocalIP()
	reg := &api.AgentServiceRegistration{
		 ID: fmt.Sprintf("%v-%v-%v", s.Name, IP, s.Port), // Name of the service node
		 Name: s.Name, // service name
		 Tags: s.Tag, // tag, can be empty
		 Port: s.Port, // service port
		 Address: IP, // Service IP
		 Check: &api.AgentServiceCheck{ // Health Check
			 Interval: s.Interval.String(), // Health check interval
			 HTTP: fmt.Sprintf("http://%s:%d/%s", IP, s.Port, s.Name), // grpc support, address to perform health check, service will be passed to Health.Check function
			 DeregisterCriticalServiceAfter: s.DeregisterCriticalServiceAfter.String(), // Deregistration time, equivalent to expiration time
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return false
	}else{
		return true
	}

}
func NewConsulRegister(name string, consuladress string) *ConsulRegister {
	return &ConsulRegister{
		Address:                        consuladress,
		Name:                           name,
		Tag:                            []string{},
		Port:                           80,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
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