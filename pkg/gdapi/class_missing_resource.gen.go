// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MissingResource struct {
  obj gdc.ObjectPtr
}

func (me *MissingResource) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MissingResource) BaseClass() string {
  return "MissingResource"
}



// Enums

func (me *MissingResource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MissingResource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MissingResource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MissingResource) SetOriginalClass(name String, )  {
  panic("TODO: implement")
}

func  (me *MissingResource) GetOriginalClass()  {
  panic("TODO: implement")
}

func  (me *MissingResource) SetRecordingProperties(enable bool, )  {
  panic("TODO: implement")
}

func  (me *MissingResource) IsRecordingProperties()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
