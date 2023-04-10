package main

import "fmt"

var eventName = make(map[string] []func(interface{}))

func main() {
	a := new(Actor)
	RegisterEvent("OnSkill", a.OnEvent)
	RegisterEvent("OnSkill", a.GlobalEvent)
	CallEvent("OnSkill", 100)
}

func CallEvent(name string, params interface{}) {
	list :=  eventName[name]
	for _, callback :=  range list {
		callback(params)
	}
}

func RegisterEvent(name string, callback func(interface{})) {
	list := eventName[name]
	list = append(list, callback)
	eventName[name] = list
}


type Actor struct {
}

func (a *Actor) OnEvent(param interface{})  {
	fmt.Println("actor event:", param)
}

func (a *Actor) GlobalEvent(param interface{}){
	fmt.Println("global parameter:", param)
}
