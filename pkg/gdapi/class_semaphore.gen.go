// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Semaphore struct {
  obj gdc.ObjectPtr
}

func (me *Semaphore) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Semaphore) BaseClass() string {
  return "Semaphore"
}



// Enums

func (me *Semaphore) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Semaphore) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Semaphore) Wait()  {
  panic("TODO: implement")
}

func  (me *Semaphore) TryWait()  {
  panic("TODO: implement")
}

func  (me *Semaphore) Post()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
