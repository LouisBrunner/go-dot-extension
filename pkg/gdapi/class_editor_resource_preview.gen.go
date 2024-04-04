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

type EditorResourcePreview struct {
  Node
}

func (me *EditorResourcePreview) BaseClass() string {
  return "EditorResourcePreview"
}

func NewEditorResourcePreview() *EditorResourcePreview {
  str := StringNameFromStr("EditorResourcePreview") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorResourcePreview{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorResourcePreview) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorResourcePreview) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePreview) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorResourcePreview) QueueResourcePreview(path String, receiver Object, receiver_func StringName, userdata Variant, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_resource_preview")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 233177534) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), receiver.AsCTypePtr(), receiver_func.AsCTypePtr(), userdata.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorResourcePreview) QueueEditedResourcePreview(resource Resource, receiver Object, receiver_func StringName, userdata Variant, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_edited_resource_preview")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1608376650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), receiver.AsCTypePtr(), receiver_func.AsCTypePtr(), userdata.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorResourcePreview) AddPreviewGenerator(generator EditorResourcePreviewGenerator, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_preview_generator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 332288124) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{generator.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorResourcePreview) RemovePreviewGenerator(generator EditorResourcePreviewGenerator, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_preview_generator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 332288124) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{generator.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorResourcePreview) CheckForInvalidation(path String, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("check_for_invalidation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type EditorResourcePreviewPreviewInvalidatedSignalFn func(path String, )

func (me *EditorResourcePreview) ConnectPreviewInvalidated(subs SignalSubscribers, fn EditorResourcePreviewPreviewInvalidatedSignalFn) {
  sig := StringNameFromStr("preview_invalidated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorResourcePreview) DisconnectPreviewInvalidated(subs SignalSubscribers, fn EditorResourcePreviewPreviewInvalidatedSignalFn) {
  sig := StringNameFromStr("preview_invalidated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
