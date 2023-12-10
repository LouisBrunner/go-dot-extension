// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CanvasLayer struct {
  obj gdc.ObjectPtr
}

func (me *CanvasLayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasLayer) BaseClass() string {
  return "CanvasLayer"
}



// Enums

func (me *CanvasLayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasLayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasLayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CanvasLayer) SetLayer(layer int, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetLayer() int {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetVisible(visible bool, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) IsVisible() bool {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) Show()  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) Hide()  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) SetTransform(transform Transform2D, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) GetFinalTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_final_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetOffset() Vector2 {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetRotation(radians float32, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetRotation() float32 {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetScale(scale Vector2, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(scale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetScale() Vector2 {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetFollowViewport(enable bool, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_follow_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) IsFollowingViewport() bool {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_following_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetFollowViewportScale(scale float32, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_follow_viewport_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetFollowViewportScale() float32 {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_follow_viewport_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) SetCustomViewport(viewport Node, )  {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(viewport.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasLayer) GetCustomViewport() Node {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_viewport")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160264692) // FIXME: should cache?
  var ret Node
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasLayer) GetCanvas() RID {
  classNameV := StringNameFromStr("CanvasLayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CanvasLayer) GetPropLayer() int {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropLayer(value int) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropVisible() bool {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropVisible(value bool) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropRotation() float32 {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropRotation(value float32) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropScale() Vector2 {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropScale(value Vector2) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropTransform() Transform2D {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropTransform(value Transform2D) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropCustomViewport() Viewport {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropCustomViewport(value Viewport) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropFollowViewportEnabled() bool {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropFollowViewportEnabled(value bool) {
  panic("TODO: implement")
}

func (me *CanvasLayer) GetPropFollowViewportScale() float32 {
  panic("TODO: implement")
}

func (me *CanvasLayer) SetPropFollowViewportScale(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API