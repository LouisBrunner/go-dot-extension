// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisibleOnScreenNotifier2D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenNotifier2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenNotifier2D) BaseClass() string {
  return "VisibleOnScreenNotifier2D"
}



// Enums

func (me *VisibleOnScreenNotifier2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisibleOnScreenNotifier2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisibleOnScreenNotifier2D) SetRect(rect Rect2, )  {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisibleOnScreenNotifier2D) GetRect() Rect2 {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisibleOnScreenNotifier2D) IsOnScreen() bool {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier2D")
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

func (me *VisibleOnScreenNotifier2D) GetPropRect() Rect2 {
  panic("TODO: implement")
}

func (me *VisibleOnScreenNotifier2D) SetPropRect(value Rect2) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API