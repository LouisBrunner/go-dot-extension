// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRController3D struct {
  obj gdc.ObjectPtr
}

func (me *XRController3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRController3D) BaseClass() string {
  return "XRController3D"
}



// Enums

func (me *XRController3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRController3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRController3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XRController3D) IsButtonPressed(name StringName, ) bool {
  classNameV := StringNameFromStr("XRController3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_button_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRController3D) GetInput(name StringName, ) Variant {
  classNameV := StringNameFromStr("XRController3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRController3D) GetFloat(name StringName, ) float32 {
  classNameV := StringNameFromStr("XRController3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_float")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2349060816) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRController3D) GetVector2(name StringName, ) Vector2 {
  classNameV := StringNameFromStr("XRController3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vector2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3100822709) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XRController3D) GetTrackerHand() XRPositionalTrackerTrackerHand {
  classNameV := StringNameFromStr("XRController3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tracker_hand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4181770860) // FIXME: should cache?
  var ret XRPositionalTrackerTrackerHand
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API