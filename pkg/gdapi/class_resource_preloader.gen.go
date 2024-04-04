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

type ResourcePreloader struct {
  Node
}

func (me *ResourcePreloader) BaseClass() string {
  return "ResourcePreloader"
}

func NewResourcePreloader() *ResourcePreloader {
  str := StringNameFromStr("ResourcePreloader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourcePreloader{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), resource.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourcePreloader) RemoveResource(name StringName, )  {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourcePreloader) RenameResource(name StringName, newname StringName, )  {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), newname.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourcePreloader) HasResource(name StringName, ) bool {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ResourcePreloader) GetResource(name StringName, ) Resource {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3742749261) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ResourcePreloader) GetResourceList() PackedStringArray {
  classNameV := StringNameFromStr("ResourcePreloader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
