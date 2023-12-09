// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HMACContext struct {
  obj gdc.ObjectPtr
}

func (me *HMACContext) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HMACContext) BaseClass() string {
  return "HMACContext"
}



// Enums

func (me *HMACContext) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HMACContext) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *HMACContext) Start(hash_type HashingContextHashType, key PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *HMACContext) Update(data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *HMACContext) Finish()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
