// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ButtonGroup struct {
  obj gdc.ObjectPtr
}

func (me *ButtonGroup) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ButtonGroup) BaseClass() string {
  return "ButtonGroup"
}



// Enums

func (me *ButtonGroup) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ButtonGroup) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ButtonGroup) GetPressedButton()  {
  panic("TODO: implement")
}

func  (me *ButtonGroup) GetButtons()  {
  panic("TODO: implement")
}

func  (me *ButtonGroup) SetAllowUnpress(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ButtonGroup) IsAllowUnpress()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
