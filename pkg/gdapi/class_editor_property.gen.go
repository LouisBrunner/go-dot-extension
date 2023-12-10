// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorProperty struct {
  obj gdc.ObjectPtr
}

func (me *EditorProperty) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorProperty) BaseClass() string {
  return "EditorProperty"
}



// Enums

func (me *EditorProperty) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorProperty) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorProperty) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorProperty) SetLabel(text String, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) GetLabel() String {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) SetReadOnly(read_only bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_read_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&read_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) IsReadOnly() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_read_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) SetCheckable(checkable bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&checkable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) IsCheckable() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) SetChecked(checked bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&checked), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) IsChecked() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) SetDrawWarning(draw_warning bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_warning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_warning), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) IsDrawWarning() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_warning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) SetKeying(keying bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keying), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) IsKeying() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) SetDeletable(deletable bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deletable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deletable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) IsDeletable() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deletable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) GetEditedProperty() StringName {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) GetEditedObject() Object {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2050059866) // FIXME: should cache?
  var ret Object
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorProperty) UpdateProperty()  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) AddFocusable(control Control, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_focusable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) SetBottomEditor(editor Control, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bottom_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(editor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorProperty) EmitChanged(property StringName, value Variant, field StringName, changing bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("emit_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3069422438) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), gdc.ConstTypePtr(field.AsCTypePtr()), gdc.ConstTypePtr(&changing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *EditorProperty) GetPropLabel() String {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropLabel(value String) {
  panic("TODO: implement")
}

func (me *EditorProperty) GetPropReadOnly() bool {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropReadOnly(value bool) {
  panic("TODO: implement")
}

func (me *EditorProperty) GetPropCheckable() bool {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropCheckable(value bool) {
  panic("TODO: implement")
}

func (me *EditorProperty) GetPropChecked() bool {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropChecked(value bool) {
  panic("TODO: implement")
}

func (me *EditorProperty) GetPropDrawWarning() bool {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropDrawWarning(value bool) {
  panic("TODO: implement")
}

func (me *EditorProperty) GetPropKeying() bool {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropKeying(value bool) {
  panic("TODO: implement")
}

func (me *EditorProperty) GetPropDeletable() bool {
  panic("TODO: implement")
}

func (me *EditorProperty) SetPropDeletable(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API