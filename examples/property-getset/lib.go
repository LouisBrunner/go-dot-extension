package main

import (
	"fmt"

	gde "github.com/LouisBrunner/go-dot-extension"
	"github.com/LouisBrunner/go-dot-extension/pkg/entry"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
)

type MyNode2D struct {
	gdapi.Node2D

	Speed     int           `gde:"name=my_speed"`                              // Can rename the property
	Direction gdapi.Vector2 `gde:"get=GetCurrentDirection,set=SetDirectionTo"` // Can also change the getter and setter names
	secret    string
}

func (n *MyNode2D) Move(vec gdapi.Vector2) {
	n.Node2D.SetPosition(vec.MultiplyInt(n.Speed))
}

// These will be used instead of implicit getters and setters
func (n *MyNode2D) SetSpeed(speed int) {
	n.Speed = speed
}

func (n *MyNode2D) GetSpeed() int {
	return n.Speed
}

func (n *MyNode2D) SetDirectionTo(vec gdapi.Vector2) {
	n.Direction = vec
}

func (n *MyNode2D) GetCurrentDirection() gdapi.Vector2 {
	return n.Direction
}

func (n *MyNode2D) X_Ready() {
	n.printSecret()
	n.Speed = 100
}

func (n *MyNode2D) printSecret() {
	fmt.Println(n.secret)
}

func newMyNode2D() gde.Class {
	return &MyNode2D{
		secret: "123",
	}
}

func init() {
	entry.DebugMode = true
	gde.Register(newMyNode2D)
}

func main() {}
