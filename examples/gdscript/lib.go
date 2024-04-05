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

	SecretPrinted gde.Signal
	subs          gdapi.SignalSubscribers
}

func (n *MyNode2D) Move(vec gdapi.Vector2) {
	n.Node2D.SetPosition(vec.MultiplyInt(int64(n.Speed)))
}

func (n *MyNode2D) X_Ready() {
	n.printSecret()
	n.Speed *= 10
}

func (n *MyNode2D) Monitor(other gdapi.CanvasItem) {
	fmt.Printf("Other properties: %+v\n", other.GetPropertyList())

	var sub gdapi.CanvasItemDrawSignalFn
	sub = func() {
		fmt.Println("Other has been drawn")
		other.DisconnectDraw(n.subs, sub)
		n.printSecret()
	}
	other.ConnectDraw(n.subs, sub)
}

type receivedDataSub struct {
	SubKey string `json:"sub_key"`
}

type receivedData struct {
	StringKey int64           `json:"String Key"`
	SubDict   receivedDataSub `json:"sub_dict"`
}

func (n *MyNode2D) Report(data gdapi.Dictionary) {
	var receivedMap map[string]interface{}
	err := gde.UnmarshalDict(data, &receivedMap)
	if err != nil {
		fmt.Printf("Error unmarshalling data (dict): %v\n", err)
		return
	}
	fmt.Printf("Received data (map): %v\n", receivedMap)

	var received receivedData
	err = gde.UnmarshalDict(data, &received)
	if err != nil {
		fmt.Printf("Error unmarshalling data (struct): %v\n", err)
		return
	}
	fmt.Printf("Received data (struct): %v\n", received)

	n.printSecret()
}

type secretData struct {
	SomeInfo      float32
	SomeOtherData []string `json:"some_other_data2"`
}

func (n *MyNode2D) printSecret() {
	fmt.Println(n.secret)
	dict, err := gde.MarshalDict(secretData{
		SomeInfo:      3.14,
		SomeOtherData: []string{"Hello", "World"},
	})
	if err != nil {
		fmt.Println("Error marshalling secret data:", err)
		return
	}
	n.SecretPrinted.Emit(*dict.AsVariant())
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
