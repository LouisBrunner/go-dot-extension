// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MissingNode) GetOriginalClass() String {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_original_class")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MissingNode) SetRecordingProperties(enable bool, )  {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_recording_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MissingNode) IsRecordingProperties() bool {
  classNameV := StringNameFromStr("MissingNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_recording_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *MissingNode) GetPropOriginalClass() String {
  panic("TODO: implement")
}

func (me *MissingNode) SetPropOriginalClass(value String) {
  panic("TODO: implement")
}

func (me *MissingNode) GetPropRecordingProperties() bool {
  panic("TODO: implement")
}

func (me *MissingNode) SetPropRecordingProperties(value bool) {
  panic("TODO: implement")
}