// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForSegmentShape2DList struct {
  fnSetA gdc.MethodBindPtr
  fnGetA gdc.MethodBindPtr
  fnSetB gdc.MethodBindPtr
  fnGetB gdc.MethodBindPtr
}

var ptrsForSegmentShape2D ptrsForSegmentShape2DList

func initSegmentShape2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SegmentShape2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_a")
    defer methodName.Destroy()
    ptrsForSegmentShape2D.fnSetA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_a")
    defer methodName.Destroy()
    ptrsForSegmentShape2D.fnGetA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_b")
    defer methodName.Destroy()
    ptrsForSegmentShape2D.fnSetB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_b")
    defer methodName.Destroy()
    ptrsForSegmentShape2D.fnGetB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
}

type SegmentShape2D struct {
  Shape2D
}

func (me *SegmentShape2D) BaseClass() string {
  return "SegmentShape2D"
}

func NewSegmentShape2D() *SegmentShape2D {
  str := StringNameFromStr("SegmentShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SegmentShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{a.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSegmentShape2D.fnSetA), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SegmentShape2D) GetA() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSegmentShape2D.fnGetA), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SegmentShape2D) SetB(b Vector2, )  {
  cargs := []gdc.ConstTypePtr{b.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSegmentShape2D.fnSetB), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SegmentShape2D) GetB() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSegmentShape2D.fnGetB), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
