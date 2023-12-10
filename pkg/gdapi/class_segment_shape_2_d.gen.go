// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SegmentShape2D struct {
  obj gdc.ObjectPtr
}

func (me *SegmentShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SegmentShape2D) BaseClass() string {
  return "SegmentShape2D"
}



// Enums

func (me *SegmentShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SegmentShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SegmentShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SegmentShape2D) SetA(a Vector2, )  {
  classNameV := StringNameFromStr("SegmentShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(a.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SegmentShape2D) GetA() Vector2 {
  classNameV := StringNameFromStr("SegmentShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SegmentShape2D) SetB(b Vector2, )  {
  classNameV := StringNameFromStr("SegmentShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(b.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SegmentShape2D) GetB() Vector2 {
  classNameV := StringNameFromStr("SegmentShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SegmentShape2D) GetPropA() Vector2 {
  panic("TODO: implement")
}

func (me *SegmentShape2D) SetPropA(value Vector2) {
  panic("TODO: implement")
}

func (me *SegmentShape2D) GetPropB() Vector2 {
  panic("TODO: implement")
}

func (me *SegmentShape2D) SetPropB(value Vector2) {
  panic("TODO: implement")
}