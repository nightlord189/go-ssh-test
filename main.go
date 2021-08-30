package main

import (
	"fmt"
	"log"

	"github.com/melbahja/goph"
)

func main() {
	fmt.Println("start")

	auth, err := goph.Key("/home/<user>/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", "<ip>", auth)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Printf(runCommand(client, `free -m | awk 'NR==2{printf "Memory Usage: %s/%sMB (%.2f%%)\n", $3,$2,$3*100/$2 }'`))
	fmt.Printf(runCommand(client, `df -h | awk '$NF=="/"{printf "Disk Usage: %d/%dGB (%s)\n", $3,$2,$5}'`))
	fmt.Printf(runCommand(client, `top -bn1 | grep load | awk '{printf "CPU Load: %.2f\n", $(NF-2)}'`))
}

func runCommand(client *goph.Client, cmd string) (string, error) {
	outBytes, err := client.Run(cmd)

	if err != nil {
		return "", err
	}
	return string(outBytes), nil
}
