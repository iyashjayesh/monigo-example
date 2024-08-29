package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iyashjayesh/monigo"
	"golang.org/x/exp/rand"
)

func main() {

	monigoInstance := &monigo.Monigo{
		ServiceName:   "Yash-MicroService",
		DashboardPort: 8080,
	}

	monigoInstance.PurgeMonigoStorage()
	monigoInstance.SetDbSyncFrequency("1m")
	monigoInstance.StartDashboard()

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api2", apiHandler2)
	log.Println("Server started at :8000")
	http.ListenAndServe(":8000", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	monigo.MeasureExecutionTime("MemExpensiveFunc", memexpensiveFunc)
	monigo.MeasureExecutionTime("CpuExpensiveFunc", cpuexpensiveFunc)
	monigo.RecordRequestDuration(time.Since(start))
	w.Write([]byte("API response"))
}

func apiHandler2(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	monigo.MeasureExecutionTime("MemExpensiveFunc", memexpensiveFunc)
	monigo.MeasureExecutionTime("CpuExpensiveFunc", cpuexpensiveFunc)
	monigo.RecordRequestDuration(time.Since(start))
	w.Write([]byte("API response"))
}

//go:noinline
func memexpensiveFunc() {
	m := make([]int, 10_000_000)
	for i := range m {
		m[i] = rand.Intn(127)
	}

	memanotherExpensiveFunc()
}

//go:noinline
func memanotherExpensiveFunc() {
	m := make(map[int]float64, 1_000_000)

	for key := range m {
		m[key] = rand.Float64()
	}
}

//go:noinline
func cpuexpensiveFunc() {
	var sum float64
	for i := 0; i < 10_000_000; i++ {
		sum += rand.Float64()
	}

	anotherExpensiveFunc()
}

//go:noinline
func anotherExpensiveFunc() {
	var sum int
	for i := 0; i < 1_000_000; i++ {
		sum += rand.Intn(10)
	}
}
