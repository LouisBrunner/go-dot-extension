package main

import (
	"fmt"

	gde "github.com/LouisBrunner/go-dot-extension"
	"github.com/LouisBrunner/go-dot-extension/pkg/entry"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
)

type MyNode2D struct {
	gdapi.Node2D

	Speed     int           `godot:"name=my_speed"`                   // Can rename the property
	Direction gdapi.Vector2 `godot:"get=GetCurrentDirection,set=nil"` // Can also change the getter and setter functions
	secret    string
}

func (n *MyNode2D) Move(vec gdapi.Vector2) {
	n.Node2D.SetPosition(vec.MultiplyInt(int64(n.Speed)))
}

// These will be used instead of implicit getters and setters
func (n *MyNode2D) SetSpeed(speed int) {
	n.Speed = speed
}

func (n *MyNode2D) GetSpeed() int {
	return n.Speed
}

func (n *MyNode2D) GetCurrentDirection() gdapi.Vector2 {
	return n.Direction
}

func (n *MyNode2D) X_Ready() {
	n.printSecret()
	n.Speed *= 10
}

func (n *MyNode2D) printSecret() {
	fmt.Println(n.secret)
}

func newMyNode2D() gde.Class {
	return &MyNode2D{
		secret: "123",
		Speed:  10,
	}
}

func init() {
	entry.DebugMode = true
	gde.Register(newMyNode2D)
}

func main() {}
