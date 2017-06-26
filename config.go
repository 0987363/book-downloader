package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)


func ReadRules(name string) []*Rule {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panic("Open file failed, ", err)
	}

	rules := []*Rule{}
	err = json.Unmarshal(data, &rules)
	if err != nil {
		log.Panic("Parse config failed:", err)
	}

	return rules
}



