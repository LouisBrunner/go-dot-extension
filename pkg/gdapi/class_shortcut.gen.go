// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shortcut struct {
  obj gdc.ObjectPtr
}

func (me *Shortcut) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shortcut) BaseClass() string {
  return "Shortcut"
}



// Enums

func (me *Shortcut) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shortcut) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Shortcut) SetEvents(events Array, )  {
  panic("TODO: implement")
}

func  (me *Shortcut) GetEvents()  {
  panic("TODO: implement")
}

func  (me *Shortcut) HasValidEvent()  {
  panic("TODO: implement")
}

func  (me *Shortcut) MatchesEvent(event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *Shortcut) GetAsText()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
