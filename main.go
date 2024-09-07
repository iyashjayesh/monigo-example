package main

import (
	"net/http"
	"time"

	"github.com/iyashjayesh/monigo"
	"github.com/iyashjayesh/monigo/models"
	"golang.org/x/exp/rand"
)

func main() {

	monigoInstance := &monigo.Monigo{
		ServiceName:        "data-api", // Compulsory
		PurgeMonigoStorage: true,       // Default is false
		DashboardPort:      8080,       // Default is 8080
		DbSyncFrequency:    "1m",       // Default is 5 Minutes
		RetentionPeriod:    "4d",       // Default is 14 days (2 weeks)
	}

	// 	### Weight Configuration

	// In the health scoring system, weights determine the importance of each metric:

	// - **Weight of `1.0`**: Indicates maximum importance. Metrics with this weight have the highest impact on the overall health score.
	// - **Weights Less Than `1.0`**: Reflect decreasing levels of importance. Metrics with lower weights contribute less to the overall score.

	// **Example**:
	// - Set `MaxLoad.Weight` to `1.0` if CPU load is critical.
	// - Set `MaxMemory.Weight` to `0.5` if memory usage is moderately important.
	// - Set `MaxGoroutines.Weight` to `0.2` for less critical metrics.

	// 1.0 is the maximum weight and 0.0 is the minimum weight.
	// critical - 1.0
	// moderate - 0.5
	// less critical - 0.2

	// to check overall health of the service
	monigoInstance.SetServiceThresholds(&models.ServiceHealthThresholds{
		MaxGoroutines: models.Thresholds{
			Value:  100,
			Weight: 1.0,
		},
		MaxCPULoad: models.Thresholds{
			Value:  2,
			Weight: 0.5,
		},
		MaxMemory: models.Thresholds{
			Value:  80,
			Weight: 0.2,
		},
	})

	monigoInstance.Start()

	// monigoInstance.DeleteMonigoStorage() // Delete monigo storage
	// monigoInstance.SetDbSyncFrequency("1m")
	// routinesStats := monigoInstance.PrintGoRoutinesStats() // Print go routines stats
	// log.Println(routinesStats)

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api2", apiHandler2)
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
