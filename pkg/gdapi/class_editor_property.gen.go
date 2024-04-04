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

type EditorProperty struct {
  Container
}

func (me *EditorProperty) BaseClass() string {
  return "EditorProperty"
}

func NewEditorProperty() *EditorProperty {
  str := StringNameFromStr("EditorProperty") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorProperty{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) GetLabel() String {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorProperty) SetReadOnly(read_only bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_read_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&read_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsReadOnly() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_read_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetCheckable(checkable bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&checkable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsCheckable() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetChecked(checked bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&checked) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsChecked() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetDrawWarning(draw_warning bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_warning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_warning) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsDrawWarning() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_warning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetKeying(keying bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keying) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsKeying() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keying")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetDeletable(deletable bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deletable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deletable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsDeletable() bool {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deletable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) GetEditedProperty() StringName {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorProperty) GetEditedObject() Object {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edited_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2050059866) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorProperty) UpdateProperty()  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) AddFocusable(control Control, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_focusable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) SetBottomEditor(editor Control, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bottom_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{editor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) EmitChanged(property StringName, value Variant, field StringName, changing bool, )  {
  classNameV := StringNameFromStr("EditorProperty")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("emit_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3069422438) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr(), field.AsCTypePtr(), gdc.ConstTypePtr(&changing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorPropertyPropertyChangedSignalFn func(property StringName, value Variant, field StringName, changing bool, )

func (me *EditorProperty) ConnectPropertyChanged(subs SignalSubscribers, fn EditorPropertyPropertyChangedSignalFn) {
  sig := StringNameFromStr("property_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyChanged(subs SignalSubscribers, fn EditorPropertyPropertyChangedSignalFn) {
  sig := StringNameFromStr("property_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyMultiplePropertiesChangedSignalFn func(properties PackedStringArray, value Array, )

func (me *EditorProperty) ConnectMultiplePropertiesChanged(subs SignalSubscribers, fn EditorPropertyMultiplePropertiesChangedSignalFn) {
  sig := StringNameFromStr("multiple_properties_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectMultiplePropertiesChanged(subs SignalSubscribers, fn EditorPropertyMultiplePropertiesChangedSignalFn) {
  sig := StringNameFromStr("multiple_properties_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyPropertyKeyedSignalFn func(property StringName, )

func (me *EditorProperty) ConnectPropertyKeyed(subs SignalSubscribers, fn EditorPropertyPropertyKeyedSignalFn) {
  sig := StringNameFromStr("property_keyed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyKeyed(subs SignalSubscribers, fn EditorPropertyPropertyKeyedSignalFn) {
  sig := StringNameFromStr("property_keyed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyPropertyDeletedSignalFn func(property StringName, )

func (me *EditorProperty) ConnectPropertyDeleted(subs SignalSubscribers, fn EditorPropertyPropertyDeletedSignalFn) {
  sig := StringNameFromStr("property_deleted")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyDeleted(subs SignalSubscribers, fn EditorPropertyPropertyDeletedSignalFn) {
  sig := StringNameFromStr("property_deleted")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyPropertyKeyedWithValueSignalFn func(property StringName, value Variant, )

func (me *EditorProperty) ConnectPropertyKeyedWithValue(subs SignalSubscribers, fn EditorPropertyPropertyKeyedWithValueSignalFn) {
  sig := StringNameFromStr("property_keyed_with_value")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyKeyedWithValue(subs SignalSubscribers, fn EditorPropertyPropertyKeyedWithValueSignalFn) {
  sig := StringNameFromStr("property_keyed_with_value")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyPropertyCheckedSignalFn func(property StringName, checked bool, )

func (me *EditorProperty) ConnectPropertyChecked(subs SignalSubscribers, fn EditorPropertyPropertyCheckedSignalFn) {
  sig := StringNameFromStr("property_checked")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyChecked(subs SignalSubscribers, fn EditorPropertyPropertyCheckedSignalFn) {
  sig := StringNameFromStr("property_checked")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyPropertyPinnedSignalFn func(property StringName, pinned bool, )

func (me *EditorProperty) ConnectPropertyPinned(subs SignalSubscribers, fn EditorPropertyPropertyPinnedSignalFn) {
  sig := StringNameFromStr("property_pinned")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyPinned(subs SignalSubscribers, fn EditorPropertyPropertyPinnedSignalFn) {
  sig := StringNameFromStr("property_pinned")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyPropertyCanRevertChangedSignalFn func(property StringName, can_revert bool, )

func (me *EditorProperty) ConnectPropertyCanRevertChanged(subs SignalSubscribers, fn EditorPropertyPropertyCanRevertChangedSignalFn) {
  sig := StringNameFromStr("property_can_revert_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectPropertyCanRevertChanged(subs SignalSubscribers, fn EditorPropertyPropertyCanRevertChangedSignalFn) {
  sig := StringNameFromStr("property_can_revert_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyResourceSelectedSignalFn func(path String, resource Resource, )

func (me *EditorProperty) ConnectResourceSelected(subs SignalSubscribers, fn EditorPropertyResourceSelectedSignalFn) {
  sig := StringNameFromStr("resource_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectResourceSelected(subs SignalSubscribers, fn EditorPropertyResourceSelectedSignalFn) {
  sig := StringNameFromStr("resource_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertyObjectIdSelectedSignalFn func(property StringName, id int, )

func (me *EditorProperty) ConnectObjectIdSelected(subs SignalSubscribers, fn EditorPropertyObjectIdSelectedSignalFn) {
  sig := StringNameFromStr("object_id_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectObjectIdSelected(subs SignalSubscribers, fn EditorPropertyObjectIdSelectedSignalFn) {
  sig := StringNameFromStr("object_id_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorPropertySelectedSignalFn func(path String, focusable_idx int, )

func (me *EditorProperty) ConnectSelected(subs SignalSubscribers, fn EditorPropertySelectedSignalFn) {
  sig := StringNameFromStr("selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorProperty) DisconnectSelected(subs SignalSubscribers, fn EditorPropertySelectedSignalFn) {
  sig := StringNameFromStr("selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
