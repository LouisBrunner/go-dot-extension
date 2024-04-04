package main

import (
	"fmt"

	gde "github.com/LouisBrunner/go-dot-extension"
	"github.com/LouisBrunner/go-dot-extension/pkg/entry"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
)

type MyNode2D struct {
	gdapi.Node2D

	Speed  int
	secret string

	SecretPrinted gdapi.Signal
	subs          gdapi.SignalSubscribers
}

func (n *MyNode2D) Move(vec gdapi.Vector2) {
	n.Node2D.SetPosition(vec.MultiplyInt(int64(n.Speed)))
	n.printSecret()
}

func (n *MyNode2D) X_Ready() {
	n.printSecret()
	n.Speed *= 10
}

func (n *MyNode2D) Monitor(other gdapi.CanvasItem) {
	var sub gdapi.CanvasItemDrawSignalFn
	sub = func() {
		fmt.Println("Other has been drawn")
		other.DisconnectDraw(n.subs, sub)
	}
	other.ConnectDraw(n.subs, sub)
}

func (n *MyNode2D) printSecret() {
	fmt.Println(n.secret)
	n.SecretPrinted.Emit(*gdapi.NewBoolFromBool(true).AsVariant())
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
