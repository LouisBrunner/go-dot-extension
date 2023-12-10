// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RootMotionView struct {
  obj gdc.ObjectPtr
}

func (me *RootMotionView) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RootMotionView) BaseClass() string {
  return "RootMotionView"
}



// Enums

func (me *RootMotionView) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RootMotionView) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RootMotionView) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RootMotionView) SetAnimationPath(path NodePath, )  {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RootMotionView) GetAnimationPath() NodePath {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RootMotionView) SetColor(color Color, )  {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RootMotionView) GetColor() Color {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RootMotionView) SetCellSize(size float32, )  {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RootMotionView) GetCellSize() float32 {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RootMotionView) SetRadius(size float32, )  {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RootMotionView) GetRadius() float32 {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RootMotionView) SetZeroY(enable bool, )  {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_zero_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RootMotionView) GetZeroY() bool {
  classNameV := StringNameFromStr("RootMotionView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_zero_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RootMotionView) GetPropAnimationPath() NodePath {
  panic("TODO: implement")
}

func (me *RootMotionView) SetPropAnimationPath(value NodePath) {
  panic("TODO: implement")
}

func (me *RootMotionView) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *RootMotionView) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *RootMotionView) GetPropCellSize() float32 {
  panic("TODO: implement")
}

func (me *RootMotionView) SetPropCellSize(value float32) {
  panic("TODO: implement")
}

func (me *RootMotionView) GetPropRadius() float32 {
  panic("TODO: implement")
}

func (me *RootMotionView) SetPropRadius(value float32) {
  panic("TODO: implement")
}

func (me *RootMotionView) GetPropZeroY() bool {
  panic("TODO: implement")
}

func (me *RootMotionView) SetPropZeroY(value bool) {
  panic("TODO: implement")
}