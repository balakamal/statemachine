package main

import (
	_ "Go/parser"
	_ "fmt"
)

type Param struct {
	validationResult	string
}

type Transition struct {
	Type 	string
	Event	string
	fromState	State
	targetState State
	Implementation string
	Params 	Param
}

type State struct {
	Name string
	DoAction []string
	EntryGaurd	string
	GuardFailState string
	Transitions []Transition

}

type StateMachineDefinition struct {
	InitialState string
	Name string
	statemap map[string]State
}

func main() {
	data := parseStateMachineJson()
	buildStateMachineDefinition(data)

}

func buildStateMachineDefinition(data Statemachinejson) {
	definition := StateMachineDefinition{
		InitialState: data.InitialState,
		Name: data.Name,
	}
	definition.statemap = make(map[string]State)
	states := data.States
	for i:=0;i<len(states);i++{
		state := states[i]
		targetState := State{}
		targetState.Name = state.Name
		targetState.DoAction=state.DoAction
		targetState.EntryGaurd = state.EntryGuard
		targetState.GuardFailState = state.GuardFailState
		definition.statemap[state.Name] = targetState
	}
	transitionsStateMachineDefinition(data,definition)
}
func transitionsStateMachineDefinition(data Statemachinejson,definition StateMachineDefinition){
	for i:=0;i<len(data.Transitions);i++{
		transition:=data.Transitions[i]
		defTransition := Transition{
			Type: transition.Type,
			Event: transition.Event,
			Implementation: transition.Implementation,

		}
		fromState := definition.statemap[transition.FromState]
		defTransition.fromState = fromState
		fromState.Transitions =  append(fromState.Transitions,defTransition)
		toState := definition.statemap[transition.TargetState]
		defTransition.targetState = toState
		toState.Transitions = append(toState.Transitions,defTransition)
	}
}
