package main

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

const token = "yT2-kf6ncn4CAvrpPWkMzIoywvLCwiknLEf-W3LBsTo-tRkpPwokkaRskvo2M7noGSyhhJdxzTNU9WDP-kiPXQ=="

func main() {
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	org := "org"
	bucket := "tutorial"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 10; value++ {
		tags := map[string]string{
			"environment": "local",
		}
		fields := map[string]interface{}{
			"counter": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
		log.Printf("done with %d", value)
	}

	// query
	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "tutorial")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	// aggregate query
	query = `from(bucket: "tutorial")
              |> range(start: -10m)
              |> filter(fn: (r) => r._measurement == "measurement1")
              |> mean()`
	results, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}
