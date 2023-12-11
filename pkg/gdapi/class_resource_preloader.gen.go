// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourcePreloader struct {
  obj gdc.ObjectPtr
}

func (me *ResourcePreloader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourcePreloader) BaseClass() string {
  return "ResourcePreloader"
}



// Enums

func (me *ResourcePreloader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourcePreloader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourcePreloader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ResourcePreloader) AddResource(name StringName, resource Resource, )  {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1168801743) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(resource.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ResourcePreloader) RemoveResource(name StringName, )  {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ResourcePreloader) RenameResource(name StringName, newname StringName, )  {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(newname.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ResourcePreloader) HasResource(name StringName, ) bool {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourcePreloader) GetResource(name StringName, ) Resource {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3742749261) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ResourcePreloader) GetResourceList() PackedStringArray {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
