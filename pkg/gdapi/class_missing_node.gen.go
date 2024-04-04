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

type MissingNode struct {
  Node
}

func (me *MissingNode) BaseClass() string {
  return "MissingNode"
}

func NewMissingNode() *MissingNode {
  str := StringNameFromStr("MissingNode") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MissingNode{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MissingNode) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MissingNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MissingNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MissingNode) SetOriginalClass(name String, )  {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_original_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MissingNode) GetOriginalClass() String {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_original_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MissingNode) SetRecordingProperties(enable bool, )  {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_recording_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MissingNode) IsRecordingProperties() bool {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_recording_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
