// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MissingNode struct {
  obj gdc.ObjectPtr
}

func (me *MissingNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MissingNode) BaseClass() string {
  return "MissingNode"
}



// Enums

func (me *MissingNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MissingNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MissingNode) SetOriginalClass(name String, )  {
  panic("TODO: implement")
}

func  (me *MissingNode) GetOriginalClass()  {
  panic("TODO: implement")
}

func  (me *MissingNode) SetRecordingProperties(enable bool, )  {
  panic("TODO: implement")
}

func  (me *MissingNode) IsRecordingProperties()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
