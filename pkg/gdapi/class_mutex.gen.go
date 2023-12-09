// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Mutex struct {
  obj gdc.ObjectPtr
}

func (me *Mutex) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Mutex) BaseClass() string {
  return "Mutex"
}



// Enums

func (me *Mutex) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Mutex) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Mutex) Lock()  {
  panic("TODO: implement")
}

func  (me *Mutex) TryLock()  {
  panic("TODO: implement")
}

func  (me *Mutex) Unlock()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
