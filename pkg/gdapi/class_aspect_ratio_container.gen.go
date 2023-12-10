// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AspectRatioContainer struct {
  obj gdc.ObjectPtr
}

func (me *AspectRatioContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AspectRatioContainer) BaseClass() string {
  return "AspectRatioContainer"
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

func  (me *AspectRatioContainer) SetRatio(ratio float32, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AspectRatioContainer) GetRatio() float32 {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AspectRatioContainer) SetStretchMode(stretch_mode AspectRatioContainerStretchMode, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1876743467) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AspectRatioContainer) GetStretchMode() AspectRatioContainerStretchMode {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3416449033) // FIXME: should cache?
  var ret AspectRatioContainerStretchMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AspectRatioContainer) SetAlignmentHorizontal(alignment_horizontal AspectRatioContainerAlignmentMode, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment_horizontal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2147829016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment_horizontal), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AspectRatioContainer) GetAlignmentHorizontal() AspectRatioContainerAlignmentMode {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment_horizontal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3838875429) // FIXME: should cache?
  var ret AspectRatioContainerAlignmentMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AspectRatioContainer) SetAlignmentVertical(alignment_vertical AspectRatioContainerAlignmentMode, )  {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alignment_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2147829016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment_vertical), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AspectRatioContainer) GetAlignmentVertical() AspectRatioContainerAlignmentMode {
  classNameV := StringNameFromStr("AspectRatioContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alignment_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3838875429) // FIXME: should cache?
  var ret AspectRatioContainerAlignmentMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AspectRatioContainer) GetPropRatio() float32 {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) SetPropRatio(value float32) {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) GetPropStretchMode() int {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) SetPropStretchMode(value int) {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) GetPropAlignmentHorizontal() int {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) SetPropAlignmentHorizontal(value int) {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) GetPropAlignmentVertical() int {
  panic("TODO: implement")
}

func (me *AspectRatioContainer) SetPropAlignmentVertical(value int) {
  panic("TODO: implement")
}