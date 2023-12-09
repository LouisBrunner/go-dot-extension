// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CallbackTweener struct {
  obj gdc.ObjectPtr
}

func (me *CallbackTweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CallbackTweener) BaseClass() string {
  return "CallbackTweener"
}



// Enums

func (me *CallbackTweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CallbackTweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CallbackTweener) SetDelay(delay float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
