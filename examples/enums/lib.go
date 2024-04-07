package main

import (
	"fmt"

	gde "github.com/LouisBrunner/go-dot-extension"
	"github.com/LouisBrunner/go-dot-extension/pkg/entry"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
)

type MyNode2D struct {
	gdapi.Node2D

	Speed     int
	Direction Direction
	secret    string
}

type Direction int

func (n *MyNode2D) Constants() gde.ConstantsDefinition {
	return gde.ConstantsDefinition{
		Constants: map[string]int{
			"MaxSpeed": 100, // all value names are automatically transformed to SNAKE_CASE automatically
		},
		Enums: map[string]map[string]int{
			// all enum names are transformed to CamelCase automatically,
			// if you want GDScript to have the right type hint, you will need this name to match your type name
			"Direction": {
				"UP":   0,
				"DOWN": 1,
			},
		},
	}
}

func (n *MyNode2D) Heading() Direction {
	return Direction(0)
}

func (n *MyNode2D) SetHeading(dir Direction) {
	fmt.Println("Heading set to", dir)
}

func (n *MyNode2D) Move(vec gdapi.Vector2) {
	n.Node2D.SetPosition(vec.MultiplyInt(int64(n.Speed)))
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
