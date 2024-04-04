// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type EncodedObjectAsID struct {
  RefCounted
}

func (me *EncodedObjectAsID) BaseClass() string {
  return "EncodedObjectAsID"
}

func NewEncodedObjectAsID() *EncodedObjectAsID {
  str := StringNameFromStr("EncodedObjectAsID") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EncodedObjectAsID{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *EncodedObjectAsID) SetObjectId(id int64, )  {
  classNameV := StringNameFromStr("EncodedObjectAsID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_object_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EncodedObjectAsID) GetObjectId() int64 {
  classNameV := StringNameFromStr("EncodedObjectAsID")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_object_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
