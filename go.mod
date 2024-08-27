module monigo-example

replace github.com/iyashjayesh/monigo => ../monigo

go 1.22

toolchain go1.22.6

require golang.org/x/exp v0.0.0-20240823005443-9b4947da3948

require github.com/iyashjayesh/monigo v0.0.0-20240827131359-59d07b15d1e6

require (
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.etcd.io/bbolt v1.3.11 // indirect
	golang.org/x/sys v0.24.0 // indirect
)
