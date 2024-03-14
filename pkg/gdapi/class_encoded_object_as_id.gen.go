// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EncodedObjectAsID struct {
  RefCounted
}

func (me *EncodedObjectAsID) BaseClass() string {
  return "EncodedObjectAsID"
}



// Enums

func (me *EncodedObjectAsID) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EncodedObjectAsID) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EncodedObjectAsID) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EncodedObjectAsID) SetObjectId(id int, )  {
  classNameV := StringNameFromStr("EncodedObjectAsID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_object_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EncodedObjectAsID) GetObjectId() int {
  classNameV := StringNameFromStr("EncodedObjectAsID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_object_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
