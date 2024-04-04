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

type AspectRatioContainer struct {
  Container
}

func (me *AspectRatioContainer) BaseClass() string {
  return "AspectRatioContainer"
}

func NewAspectRatioContainer() *AspectRatioContainer {
  str := StringNameFromStr("AspectRatioContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AspectRatioContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AspectRatioContainerStretchMode int
const (
  AspectRatioContainerStretchModeStretchWidthControlsHeight AspectRatioContainerStretchMode = 0
  AspectRatioContainerStretchModeStretchHeightControlsWidth AspectRatioContainerStretchMode = 1
  AspectRatioContainerStretchModeStretchFit AspectRatioContainerStretchMode = 2
  AspectRatioContainerStretchModeStretchCover AspectRatioContainerStretchMode = 3
)

type AspectRatioContainerAlignmentMode int
const (
  AspectRatioContainerAlignmentModeAlignmentBegin AspectRatioContainerAlignmentMode = 0
  AspectRatioContainerAlignmentModeAlignmentCenter AspectRatioContainerAlignmentMode = 1
  AspectRatioContainerAlignmentModeAlignmentEnd AspectRatioContainerAlignmentMode = 2
)

func (me *AspectRatioContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AspectRatioContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AspectRatioContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AspectRatioContainer) SetRatio(ratio float64, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AspectRatioContainer) GetRatio() float64 {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AspectRatioContainer) SetStretchMode(stretch_mode AspectRatioContainerStretchMode, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1876743467) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AspectRatioContainer) GetStretchMode() AspectRatioContainerStretchMode {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3416449033) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AspectRatioContainerStretchMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AspectRatioContainer) SetAlignmentHorizontal(alignment_horizontal AspectRatioContainerAlignmentMode, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment_horizontal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2147829016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment_horizontal) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AspectRatioContainer) GetAlignmentHorizontal() AspectRatioContainerAlignmentMode {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment_horizontal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3838875429) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AspectRatioContainerAlignmentMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AspectRatioContainer) SetAlignmentVertical(alignment_vertical AspectRatioContainerAlignmentMode, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2147829016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment_vertical) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AspectRatioContainer) GetAlignmentVertical() AspectRatioContainerAlignmentMode {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3838875429) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AspectRatioContainerAlignmentMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
