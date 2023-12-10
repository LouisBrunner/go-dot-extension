// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Button struct {
  obj gdc.ObjectPtr
}

func (me *Button) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Button) BaseClass() string {
  return "Button"
}



// Enums

func (me *Button) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Button) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Button) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Button) SetText(text String, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetText() String {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1008890932) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetTextOverrunBehavior() TextServerOverrunBehavior {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3779142101) // FIXME: should cache?
  var ret TextServerOverrunBehavior
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetLanguage() String {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetButtonIcon(texture Texture2D, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_button_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetButtonIcon() Texture2D {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_button_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetFlat(enabled bool, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) IsFlat() bool {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetClipText(enabled bool, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetClipText() bool {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetTextAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetTextAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetIconAlignment(icon_alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&icon_alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetIconAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetVerticalIconAlignment(vertical_icon_alignment VerticalAlignment, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical_icon_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1796458609) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical_icon_alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) GetVerticalIconAlignment() VerticalAlignment {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertical_icon_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3274884059) // FIXME: should cache?
  var ret VerticalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Button) SetExpandIcon(enabled bool, )  {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Button) IsExpandIcon() bool {
  classNameV := StringNameFromStr("Button")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_expand_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Button) GetPropText() String {
  panic("TODO: implement")
}

func (me *Button) SetPropText(value String) {
  panic("TODO: implement")
}

func (me *Button) GetPropIcon() Texture2D {
  panic("TODO: implement")
}

func (me *Button) SetPropIcon(value Texture2D) {
  panic("TODO: implement")
}

func (me *Button) GetPropFlat() bool {
  panic("TODO: implement")
}

func (me *Button) SetPropFlat(value bool) {
  panic("TODO: implement")
}

func (me *Button) GetPropAlignment() int {
  panic("TODO: implement")
}

func (me *Button) SetPropAlignment(value int) {
  panic("TODO: implement")
}

func (me *Button) GetPropTextOverrunBehavior() int {
  panic("TODO: implement")
}

func (me *Button) SetPropTextOverrunBehavior(value int) {
  panic("TODO: implement")
}

func (me *Button) GetPropClipText() bool {
  panic("TODO: implement")
}

func (me *Button) SetPropClipText(value bool) {
  panic("TODO: implement")
}

func (me *Button) GetPropIconAlignment() int {
  panic("TODO: implement")
}

func (me *Button) SetPropIconAlignment(value int) {
  panic("TODO: implement")
}

func (me *Button) GetPropVerticalIconAlignment() int {
  panic("TODO: implement")
}

func (me *Button) SetPropVerticalIconAlignment(value int) {
  panic("TODO: implement")
}

func (me *Button) GetPropExpandIcon() bool {
  panic("TODO: implement")
}

func (me *Button) SetPropExpandIcon(value bool) {
  panic("TODO: implement")
}

func (me *Button) GetPropTextDirection() int {
  panic("TODO: implement")
}

func (me *Button) SetPropTextDirection(value int) {
  panic("TODO: implement")
}

func (me *Button) GetPropLanguage() String {
  panic("TODO: implement")
}

func (me *Button) SetPropLanguage(value String) {
  panic("TODO: implement")
}