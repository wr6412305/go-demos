package main

import (
	"encoding/json"
	"fmt"
)

func marshal1() {
	m := map[string][]string{
		"level":   {"debug"},
		"message": {"File not found", "Stack overflow"},
	}

	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n", data)
	}
}
