package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Tanawin Siriwan", "00000000", "kmutnb.ac.th"})
	fmt.Println(string(data))
}
