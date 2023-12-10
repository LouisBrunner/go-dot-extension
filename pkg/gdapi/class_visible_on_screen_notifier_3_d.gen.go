// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisibleOnScreenNotifier3D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenNotifier3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenNotifier3D) BaseClass() string {
  return "VisibleOnScreenNotifier3D"
}



// Enums

func (me *VisibleOnScreenNotifier3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisibleOnScreenNotifier3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisibleOnScreenNotifier3D) SetAabb(rect AABB, )  {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisibleOnScreenNotifier3D) IsOnScreen() bool {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisibleOnScreenNotifier3D) GetPropAabb() AABB {
  panic("TODO: implement")
}

func (me *VisibleOnScreenNotifier3D) SetPropAabb(value AABB) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API