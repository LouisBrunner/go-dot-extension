// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TouchScreenButton struct {
  obj gdc.ObjectPtr
}

func (me *TouchScreenButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TouchScreenButton) BaseClass() string {
  return "TouchScreenButton"
}



// Enums

type TouchScreenButtonVisibilityMode int
const (
  TouchScreenButtonVisibilityModeVisibilityAlways TouchScreenButtonVisibilityMode = 0
  TouchScreenButtonVisibilityModeVisibilityTouchscreenOnly TouchScreenButtonVisibilityMode = 1
)

func (me *TouchScreenButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TouchScreenButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TouchScreenButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TouchScreenButton) SetTextureNormal(texture Texture2D, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) GetTextureNormal() Texture2D {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetTexturePressed(texture Texture2D, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) GetTexturePressed() Texture2D {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetBitmask(bitmask BitMap, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bitmask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 698588216) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bitmask.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) GetBitmask() BitMap {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bitmask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2459671998) // FIXME: should cache?
  var ret BitMap
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetShape(shape Shape2D, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 771364740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) GetShape() Shape2D {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 522005891) // FIXME: should cache?
  var ret Shape2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetShapeCentered(bool bool, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bool), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) IsShapeCentered() bool {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_centered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetShapeVisible(bool bool, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bool), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) IsShapeVisible() bool {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shape_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetAction(action String, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) GetAction() String {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetVisibilityMode(mode TouchScreenButtonVisibilityMode, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3031128463) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) GetVisibilityMode() TouchScreenButtonVisibilityMode {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2558996468) // FIXME: should cache?
  var ret TouchScreenButtonVisibilityMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) SetPassbyPress(enabled bool, )  {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_passby_press")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TouchScreenButton) IsPassbyPressEnabled() bool {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_passby_press_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TouchScreenButton) IsPressed() bool {
  classNameV := StringNameFromStr("TouchScreenButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TouchScreenButtonPressedSignalFn func()

func (me *TouchScreenButton) ConnectPressed(subs SignalSubscribers, fn TouchScreenButtonPressedSignalFn) {
  sig := StringNameFromStr("pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TouchScreenButton) DisconnectPressed(subs SignalSubscribers, fn TouchScreenButtonPressedSignalFn) {
  sig := StringNameFromStr("pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TouchScreenButtonReleasedSignalFn func()

func (me *TouchScreenButton) ConnectReleased(subs SignalSubscribers, fn TouchScreenButtonReleasedSignalFn) {
  sig := StringNameFromStr("released")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TouchScreenButton) DisconnectReleased(subs SignalSubscribers, fn TouchScreenButtonReleasedSignalFn) {
  sig := StringNameFromStr("released")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
