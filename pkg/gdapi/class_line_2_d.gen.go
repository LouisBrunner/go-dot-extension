// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Line2D struct {
  obj gdc.ObjectPtr
}

func (me *Line2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Line2D) BaseClass() string {
  return "Line2D"
}



// Enums

type Line2DLineJointMode int
const (
  Line2DLineJointModeLineJointSharp Line2DLineJointMode = 0
  Line2DLineJointModeLineJointBevel Line2DLineJointMode = 1
  Line2DLineJointModeLineJointRound Line2DLineJointMode = 2
)

type Line2DLineCapMode int
const (
  Line2DLineCapModeLineCapNone Line2DLineCapMode = 0
  Line2DLineCapModeLineCapBox Line2DLineCapMode = 1
  Line2DLineCapModeLineCapRound Line2DLineCapMode = 2
)

type Line2DLineTextureMode int
const (
  Line2DLineTextureModeLineTextureNone Line2DLineTextureMode = 0
  Line2DLineTextureModeLineTextureTile Line2DLineTextureMode = 1
  Line2DLineTextureModeLineTextureStretch Line2DLineTextureMode = 2
)

func (me *Line2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Line2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Line2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Line2D) SetPoints(points PackedVector2Array, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetPoints() PackedVector2Array {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetPointPosition(index int, position Vector2, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetPointPosition(index int, ) Vector2 {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) GetPointCount() int {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) AddPoint(position Vector2, index int, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 468506575) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) RemovePoint(index int, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) ClearPoints()  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) SetWidth(width float32, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetWidth() float32 {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetCurve(curve Curve, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetCurve() Curve {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetDefaultColor(color Color, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetDefaultColor() Color {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetGradient(color Gradient, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetGradient() Gradient {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetTexture(texture Texture2D, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetTexture() Texture2D {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetTextureMode(mode Line2DLineTextureMode, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1952559516) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetTextureMode() Line2DLineTextureMode {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2341040722) // FIXME: should cache?
  var ret Line2DLineTextureMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetJointMode(mode Line2DLineJointMode, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 604292979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetJointMode() Line2DLineJointMode {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546544037) // FIXME: should cache?
  var ret Line2DLineJointMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetBeginCapMode(mode Line2DLineCapMode, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_begin_cap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1669024546) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetBeginCapMode() Line2DLineCapMode {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_begin_cap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1107511441) // FIXME: should cache?
  var ret Line2DLineCapMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetEndCapMode(mode Line2DLineCapMode, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_end_cap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1669024546) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetEndCapMode() Line2DLineCapMode {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_end_cap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1107511441) // FIXME: should cache?
  var ret Line2DLineCapMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetSharpLimit(limit float32, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sharp_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetSharpLimit() float32 {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sharp_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetRoundPrecision(precision int, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_round_precision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&precision), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetRoundPrecision() int {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_round_precision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Line2D) SetAntialiased(antialiased bool, )  {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&antialiased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Line2D) GetAntialiased() bool {
  classNameV := StringNameFromStr("Line2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_antialiased")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Line2D) GetPropPoints() PackedVector2Array {
  panic("TODO: implement")
}

func (me *Line2D) SetPropPoints(value PackedVector2Array) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropWidth() float32 {
  panic("TODO: implement")
}

func (me *Line2D) SetPropWidth(value float32) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropWidthCurve() Curve {
  panic("TODO: implement")
}

func (me *Line2D) SetPropWidthCurve(value Curve) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropDefaultColor() Color {
  panic("TODO: implement")
}

func (me *Line2D) SetPropDefaultColor(value Color) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropGradient() Gradient {
  panic("TODO: implement")
}

func (me *Line2D) SetPropGradient(value Gradient) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *Line2D) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropTextureMode() int {
  panic("TODO: implement")
}

func (me *Line2D) SetPropTextureMode(value int) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropJointMode() int {
  panic("TODO: implement")
}

func (me *Line2D) SetPropJointMode(value int) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropBeginCapMode() int {
  panic("TODO: implement")
}

func (me *Line2D) SetPropBeginCapMode(value int) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropEndCapMode() int {
  panic("TODO: implement")
}

func (me *Line2D) SetPropEndCapMode(value int) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropSharpLimit() float32 {
  panic("TODO: implement")
}

func (me *Line2D) SetPropSharpLimit(value float32) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropRoundPrecision() int {
  panic("TODO: implement")
}

func (me *Line2D) SetPropRoundPrecision(value int) {
  panic("TODO: implement")
}

func (me *Line2D) GetPropAntialiased() bool {
  panic("TODO: implement")
}

func (me *Line2D) SetPropAntialiased(value bool) {
  panic("TODO: implement")
}