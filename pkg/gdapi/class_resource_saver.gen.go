// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceSaver struct {
  obj gdc.ObjectPtr
}

func (me *ResourceSaver) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceSaver) BaseClass() string {
  return "ResourceSaver"
}



// Enums

type ResourceSaverSaverFlags int
const (
  ResourceSaverSaverFlagsFlagNone ResourceSaverSaverFlags = 0
  ResourceSaverSaverFlagsFlagRelativePaths ResourceSaverSaverFlags = 1
  ResourceSaverSaverFlagsFlagBundleResources ResourceSaverSaverFlags = 2
  ResourceSaverSaverFlagsFlagChangePath ResourceSaverSaverFlags = 4
  ResourceSaverSaverFlagsFlagOmitEditorProperties ResourceSaverSaverFlags = 8
  ResourceSaverSaverFlagsFlagSaveBigEndian ResourceSaverSaverFlags = 16
  ResourceSaverSaverFlagsFlagCompress ResourceSaverSaverFlags = 32
  ResourceSaverSaverFlagsFlagReplaceSubresourcePaths ResourceSaverSaverFlags = 64
)

func (me *ResourceSaver) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceSaver) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceSaver) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ResourceSaver) Save(resource Resource, path String, flags ResourceSaverSaverFlags, ) Error {
  classNameV := StringNameFromStr("ResourceSaver")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2303056517) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(resource.AsCTypePtr()), gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceSaver) GetRecognizedExtensions(type_ Resource, ) PackedStringArray {
  classNameV := StringNameFromStr("ResourceSaver")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_recognized_extensions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4223597960) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(type_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourceSaver) AddResourceFormatSaver(format_saver ResourceFormatSaver, at_front bool, )  {
  classNameV := StringNameFromStr("ResourceSaver")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_resource_format_saver")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 362894272) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(format_saver.AsCTypePtr()), gdc.ConstTypePtr(&at_front), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ResourceSaver) RemoveResourceFormatSaver(format_saver ResourceFormatSaver, )  {
  classNameV := StringNameFromStr("ResourceSaver")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource_format_saver")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3373026878) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(format_saver.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
