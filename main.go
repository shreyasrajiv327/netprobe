package main

import(
  "fmt"
  "net"
  "os"
  "time"
)


func main(){
	if len(os.Args) != 2{
		fmt.Println("Usage: go run . <hostport>")
		os.Exit(1)
	}

	target := os.Args[1]
	start := time.Now()

	conn, err := net.DialTimeout("tcp", target, 3*time.Second)
	latency := time.Since(start)

	if err!=nil{
		fmt.Println("Target:", target)
		fmt.Println("Status: UNREACHABLE")
		fmt.Println("Error:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Target:", target)
	fmt.Println("Status: REACHABLE")
	fmt.Println("Latency:", latency)
}