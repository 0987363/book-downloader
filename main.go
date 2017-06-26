package main

import (
	"runtime"
	"fmt"
	"time"
	"flag"
	"log"
)

func init() {
	time.Local = time.UTC
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main() {
	var config = flag.String("c", "./config.json", "Rule config file")
	var directory = flag.String("d", "", "xiaoshuo directory page url")
	flag.Parse()

	if *directory == "" || *config == "" {
		log.Panic("Need valid derectory, ", *directory, *config)
	}

	rules := ReadRules(*config)

	rule := FindRule(rules, *directory)
	if rule == nil {
		log.Panic("Unknown site.")
	}
	fmt.Println("Found rule:", *rule)

	Server(rule, *directory)

}
