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

type MissingResource struct {
  Resource
}

func (me *MissingResource) BaseClass() string {
  return "MissingResource"
}

func NewMissingResource() *MissingResource {
  str := StringNameFromStr("MissingResource") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MissingResource{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("MissingResource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_original_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MissingResource) GetOriginalClass() String {
  classNameV := StringNameFromStr("MissingResource")
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

func  (me *MissingResource) SetRecordingProperties(enable bool, )  {
  classNameV := StringNameFromStr("MissingResource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_recording_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MissingResource) IsRecordingProperties() bool {
  classNameV := StringNameFromStr("MissingResource")
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
