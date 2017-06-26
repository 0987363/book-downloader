package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
)

// Server is main loop service
func Server(rule *Rule, dir string) {
	client := http.Client {
		Timeout: time.Second * 10,
	}
	resp, err := client.Get(dir)
	if err != nil {
		log.Panic("Get directory failed.", err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024 * 1024)
	n, err := resp.Body.Read(buf)
	if err != nil {
		log.Panic("Read directory failed.", err)
	}

	buf = buf[:n]

	fmt.Println("Read directory success.", string(buf))







}
