package main

import (
	"io"
	"os"
	"strings"
)

var ENV map[string]string

func parseEnv() {
	ENV = make(map[string]string)
	env, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer env.Close()
	data, err := io.ReadAll(env)
	if err != nil {
		panic(err)
	}
	line_data := strings.Split(string(data), "\n")
	for _, line := range line_data {
		if strings.Contains(line, "=") {
			key_value := strings.SplitAfterN(line, "=", 2)
			key_value[0] = strings.Trim(key_value[0], "=")
			key_value[1] = strings.Trim(key_value[1], "\r")
			ENV[key_value[0]] = key_value[1]
		}
	}
}
