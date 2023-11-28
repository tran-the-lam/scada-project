package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type EValue struct {
	Value     int `json:"value"`
	Threshold int `json:"threshold"`
}

var sensorIDs = []string{"sensor1", "sensor2", "sensor3", "sensor4", "sensor5", "sensor6", "sensor7", "sensor8", "sensor9", "sensor10"}
var parameters = []string{"temperature", "humidity"}
var evalues = []EValue{{Value: 40, Threshold: 25}, {Value: 60, Threshold: 35}, {Value: 80, Threshold: 30}, {Value: 100, Threshold: 34}, {Value: 120, Threshold: 30}, {Value: 140, Threshold: 30}, {Value: 160, Threshold: 30}, {Value: 180, Threshold: 30}, {Value: 200, Threshold: 30}}
var roles = []string{"employee", "manager"}

func sendReq(url, payload, token string) {
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers if needed
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-scada-api-key", "scada-api-key")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}

func gen50event() {
	for i := 1; i <= 50; i++ {

		rand.Seed(time.Now().UnixNano())
		randSIndex := rand.Intn(len(sensorIDs))
		randPIndex := rand.Intn(len(parameters))
		randEIndex := rand.Intn(len(evalues))

		sensorID := sensorIDs[randSIndex]
		parameter := parameters[randPIndex]
		value := evalues[randEIndex].Value
		threshold := evalues[randEIndex].Threshold
		now := time.Now().Unix()

		payload := fmt.Sprintf(`{"event": "limit_exceeded", "sensor_id": "%s", "parameter": "%s", "value": %d, "threshold": %d, "timestamp": %d}`, sensorID, parameter, value, threshold, now)
		fmt.Println(payload)
		sendReq("http://localhost:3000/events", payload, "")

		time.Sleep(3 * time.Second)
	}
}

func gen20user() {
	for i := 1; i <= 20; i++ {
		rand.Seed(time.Now().UnixNano())
		randRIndex := rand.Intn(len(roles))
		role := roles[randRIndex]
		payload := fmt.Sprintf(`{"user_id": "u_%d", "role": "%s"}`, i, role)
		fmt.Println(payload)
		sendReq("http://localhost:3000/users", payload, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE0NDE1NjIsInVzZXJfaWQiOiJhZG1pbiIsInVzZXJfcm9sZSI6ImFkbWluIn0.1SEmxcWzb3pd3i4h8gxshWawnovMtSeCuUPoArBJnf0")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// gen50event()
	gen20user()
}
