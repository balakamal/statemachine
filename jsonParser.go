package main

import (
	"encoding/json"
	"io/ioutil"
)

type Statemachinejson struct {
	Name 	string	`json:"name"`
	Version	string	`json:"version"`
	Status	string	`json:"status"`
	InitialState string		`json:"initial_state"`
	States		[]Statejson 	`json:"states"`
	Transitions	[]Transitionjson	`json:"transactions"`

}

// State structure to store States
type Statejson struct {
	Name   string `json:"name"`
	DoAction   []string `json:"doAction"`
	EntryGuard    string    `json:"entryGuard"`
	GuardFailState string	`json:"guard_fail_state"`
	ExitAction   []string `json:"doAction"`
}

//Structure for transactions
type Transitionjson struct {
	Type 	string	`json:"type"`
	Event	string	`json:"event"`
	FromState	string	`json:"from_state"`
	TargetState string	`json:"target_state"`
	Implementation string	`json:"implementation"`

}


func parseStateMachineJson() Statemachinejson {
	file, _ := ioutil.ReadFile("example.json")

	data := Statemachinejson{}

	_ = json.Unmarshal([]byte(file), &data)
	return data
}


