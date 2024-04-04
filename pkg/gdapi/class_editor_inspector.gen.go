// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type EditorInspector struct {
  ScrollContainer
}

func (me *EditorInspector) BaseClass() string {
  return "EditorInspector"
}

func NewEditorInspector() *EditorInspector {
  str := StringNameFromStr("EditorInspector") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorInspector{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorInspector) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorInspector) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorInspector) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorInspector) GetSelectedPath() String {
  classNameV := StringNameFromStr("EditorInspector")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorInspector) GetEditedObject() Object {
  classNameV := StringNameFromStr("EditorInspector")
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

// Signals

type EditorInspectorPropertySelectedSignalFn func(property String, )

func (me *EditorInspector) ConnectPropertySelected(subs SignalSubscribers, fn EditorInspectorPropertySelectedSignalFn) {
  sig := StringNameFromStr("property_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectPropertySelected(subs SignalSubscribers, fn EditorInspectorPropertySelectedSignalFn) {
  sig := StringNameFromStr("property_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorPropertyKeyedSignalFn func(property String, value Variant, advance bool, )

func (me *EditorInspector) ConnectPropertyKeyed(subs SignalSubscribers, fn EditorInspectorPropertyKeyedSignalFn) {
  sig := StringNameFromStr("property_keyed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectPropertyKeyed(subs SignalSubscribers, fn EditorInspectorPropertyKeyedSignalFn) {
  sig := StringNameFromStr("property_keyed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorPropertyDeletedSignalFn func(property String, )

func (me *EditorInspector) ConnectPropertyDeleted(subs SignalSubscribers, fn EditorInspectorPropertyDeletedSignalFn) {
  sig := StringNameFromStr("property_deleted")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectPropertyDeleted(subs SignalSubscribers, fn EditorInspectorPropertyDeletedSignalFn) {
  sig := StringNameFromStr("property_deleted")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorResourceSelectedSignalFn func(resource Resource, path String, )

func (me *EditorInspector) ConnectResourceSelected(subs SignalSubscribers, fn EditorInspectorResourceSelectedSignalFn) {
  sig := StringNameFromStr("resource_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectResourceSelected(subs SignalSubscribers, fn EditorInspectorResourceSelectedSignalFn) {
  sig := StringNameFromStr("resource_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorObjectIdSelectedSignalFn func(id int, )

func (me *EditorInspector) ConnectObjectIdSelected(subs SignalSubscribers, fn EditorInspectorObjectIdSelectedSignalFn) {
  sig := StringNameFromStr("object_id_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectObjectIdSelected(subs SignalSubscribers, fn EditorInspectorObjectIdSelectedSignalFn) {
  sig := StringNameFromStr("object_id_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorPropertyEditedSignalFn func(property String, )

func (me *EditorInspector) ConnectPropertyEdited(subs SignalSubscribers, fn EditorInspectorPropertyEditedSignalFn) {
  sig := StringNameFromStr("property_edited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectPropertyEdited(subs SignalSubscribers, fn EditorInspectorPropertyEditedSignalFn) {
  sig := StringNameFromStr("property_edited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorPropertyToggledSignalFn func(property String, checked bool, )

func (me *EditorInspector) ConnectPropertyToggled(subs SignalSubscribers, fn EditorInspectorPropertyToggledSignalFn) {
  sig := StringNameFromStr("property_toggled")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectPropertyToggled(subs SignalSubscribers, fn EditorInspectorPropertyToggledSignalFn) {
  sig := StringNameFromStr("property_toggled")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorEditedObjectChangedSignalFn func()

func (me *EditorInspector) ConnectEditedObjectChanged(subs SignalSubscribers, fn EditorInspectorEditedObjectChangedSignalFn) {
  sig := StringNameFromStr("edited_object_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectEditedObjectChanged(subs SignalSubscribers, fn EditorInspectorEditedObjectChangedSignalFn) {
  sig := StringNameFromStr("edited_object_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorInspectorRestartRequestedSignalFn func()

func (me *EditorInspector) ConnectRestartRequested(subs SignalSubscribers, fn EditorInspectorRestartRequestedSignalFn) {
  sig := StringNameFromStr("restart_requested")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorInspector) DisconnectRestartRequested(subs SignalSubscribers, fn EditorInspectorRestartRequestedSignalFn) {
  sig := StringNameFromStr("restart_requested")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
