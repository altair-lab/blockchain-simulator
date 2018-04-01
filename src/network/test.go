package network

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Configuration struct {
	Nodes int
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("nodes: ", configuration.Nodes)

	// generate node instances
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := make([]*Node, configuration.Nodes)
	for i := 0; i < configuration.Nodes; i++ {
		n[i] = newNode(r.Intn(100) + 1)
	}

	for i := 0; i < configuration.Nodes; i++ {
		fmt.Println(n[i].performance)
	}

	// loop by frame
}
