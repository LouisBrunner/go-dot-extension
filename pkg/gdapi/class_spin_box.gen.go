// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SpinBox struct {
  obj gdc.ObjectPtr
}

func (me *SpinBox) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpinBox) BaseClass() string {
  return "SpinBox"
}



// Enums

func (me *SpinBox) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpinBox) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpinBox) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SpinBox) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) GetHorizontalAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) SetSuffix(suffix String, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(suffix.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) GetSuffix() String {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) SetPrefix(prefix String, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_prefix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(prefix.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) GetPrefix() String {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_prefix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) SetEditable(enabled bool, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) SetCustomArrowStep(arrow_step float32, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_arrow_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&arrow_step), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) GetCustomArrowStep() float32 {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_arrow_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) IsEditable() bool {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) SetUpdateOnTextChanged(enabled bool, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_update_on_text_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) GetUpdateOnTextChanged() bool {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_update_on_text_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) SetSelectAllOnFocus(enabled bool, )  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_select_all_on_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) IsSelectAllOnFocus() bool {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_select_all_on_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SpinBox) Apply()  {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SpinBox) GetLineEdit() LineEdit {
  classNameV := StringNameFromStr("SpinBox")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_edit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4071694264) // FIXME: should cache?
  var ret LineEdit
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SpinBox) GetPropAlignment() int {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropAlignment(value int) {
  panic("TODO: implement")
}

func (me *SpinBox) GetPropEditable() bool {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropEditable(value bool) {
  panic("TODO: implement")
}

func (me *SpinBox) GetPropUpdateOnTextChanged() bool {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropUpdateOnTextChanged(value bool) {
  panic("TODO: implement")
}

func (me *SpinBox) GetPropPrefix() String {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropPrefix(value String) {
  panic("TODO: implement")
}

func (me *SpinBox) GetPropSuffix() String {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropSuffix(value String) {
  panic("TODO: implement")
}

func (me *SpinBox) GetPropCustomArrowStep() float32 {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropCustomArrowStep(value float32) {
  panic("TODO: implement")
}

func (me *SpinBox) GetPropSelectAllOnFocus() bool {
  panic("TODO: implement")
}

func (me *SpinBox) SetPropSelectAllOnFocus(value bool) {
  panic("TODO: implement")
}