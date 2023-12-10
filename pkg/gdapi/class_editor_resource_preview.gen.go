// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorResourcePreview struct {
  obj gdc.ObjectPtr
}

func (me *EditorResourcePreview) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorResourcePreview) BaseClass() string {
  return "EditorResourcePreview"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(receiver.AsCTypePtr()), gdc.ConstTypePtr(receiver_func.AsCTypePtr()), gdc.ConstTypePtr(userdata.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePreview) QueueEditedResourcePreview(resource Resource, receiver Object, receiver_func StringName, userdata Variant, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_edited_resource_preview")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1608376650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(resource.AsCTypePtr()), gdc.ConstTypePtr(receiver.AsCTypePtr()), gdc.ConstTypePtr(receiver_func.AsCTypePtr()), gdc.ConstTypePtr(userdata.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePreview) AddPreviewGenerator(generator EditorResourcePreviewGenerator, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_preview_generator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 332288124) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(generator.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePreview) RemovePreviewGenerator(generator EditorResourcePreviewGenerator, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_preview_generator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 332288124) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(generator.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorResourcePreview) CheckForInvalidation(path String, )  {
  classNameV := StringNameFromStr("EditorResourcePreview")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("check_for_invalidation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API