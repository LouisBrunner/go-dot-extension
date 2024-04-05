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

type ptrsForEditorPropertyList struct {
  fnXUpdateProperty gdc.MethodBindPtr
  fnXSetReadOnly gdc.MethodBindPtr
  fnSetLabel gdc.MethodBindPtr
  fnGetLabel gdc.MethodBindPtr
  fnSetReadOnly gdc.MethodBindPtr
  fnIsReadOnly gdc.MethodBindPtr
  fnSetCheckable gdc.MethodBindPtr
  fnIsCheckable gdc.MethodBindPtr
  fnSetChecked gdc.MethodBindPtr
  fnIsChecked gdc.MethodBindPtr
  fnSetDrawWarning gdc.MethodBindPtr
  fnIsDrawWarning gdc.MethodBindPtr
  fnSetKeying gdc.MethodBindPtr
  fnIsKeying gdc.MethodBindPtr
  fnSetDeletable gdc.MethodBindPtr
  fnIsDeletable gdc.MethodBindPtr
  fnGetEditedProperty gdc.MethodBindPtr
  fnGetEditedObject gdc.MethodBindPtr
  fnUpdateProperty gdc.MethodBindPtr
  fnAddFocusable gdc.MethodBindPtr
  fnSetBottomEditor gdc.MethodBindPtr
  fnEmitChanged gdc.MethodBindPtr
}

var ptrsForEditorProperty ptrsForEditorPropertyList

func initEditorPropertyPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorProperty")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_label")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_label")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnGetLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_read_only")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetReadOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_read_only")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnIsReadOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_checkable")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_checkable")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnIsCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_checked")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_checked")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnIsChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_draw_warning")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetDrawWarning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_draw_warning")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnIsDrawWarning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_keying")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetKeying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_keying")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnIsKeying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_deletable")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetDeletable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_deletable")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnIsDeletable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_edited_property")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnGetEditedProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("get_edited_object")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnGetEditedObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2050059866))
  }
  {
    methodName := StringNameFromStr("update_property")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnUpdateProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("add_focusable")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnAddFocusable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
  }
  {
    methodName := StringNameFromStr("set_bottom_editor")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnSetBottomEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
  }
  {
    methodName := StringNameFromStr("emit_changed")
    defer methodName.Destroy()
    ptrsForEditorProperty.fnEmitChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3069422438))
  }
}

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
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) GetLabel() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnGetLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorProperty) SetReadOnly(read_only bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&read_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetReadOnly), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsReadOnly() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnIsReadOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetCheckable(checkable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&checkable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsCheckable() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnIsCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetChecked(checked bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&checked) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetChecked), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsChecked() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnIsChecked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetDrawWarning(draw_warning bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_warning) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetDrawWarning), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsDrawWarning() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnIsDrawWarning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetKeying(keying bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keying) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetKeying), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsKeying() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnIsKeying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) SetDeletable(deletable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deletable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetDeletable), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) IsDeletable() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnIsDeletable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorProperty) GetEditedProperty() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnGetEditedProperty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorProperty) GetEditedObject() Object {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnGetEditedObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorProperty) UpdateProperty()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnUpdateProperty), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) AddFocusable(control Control, )  {
  cargs := []gdc.ConstTypePtr{control.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnAddFocusable), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) SetBottomEditor(editor Control, )  {
  cargs := []gdc.ConstTypePtr{editor.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnSetBottomEditor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorProperty) EmitChanged(property StringName, value Variant, field StringName, changing bool, )  {
  cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr(), field.AsCTypePtr(), gdc.ConstTypePtr(&changing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorProperty.fnEmitChanged), me.obj, unsafe.SliceData(cargs), nil)

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
