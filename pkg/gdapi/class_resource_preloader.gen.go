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

type ptrsForResourcePreloaderList struct {
  fnAddResource gdc.MethodBindPtr
  fnRemoveResource gdc.MethodBindPtr
  fnRenameResource gdc.MethodBindPtr
  fnHasResource gdc.MethodBindPtr
  fnGetResource gdc.MethodBindPtr
  fnGetResourceList gdc.MethodBindPtr
}

var ptrsForResourcePreloader ptrsForResourcePreloaderList

func initResourcePreloaderPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ResourcePreloader")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_resource")
    defer methodName.Destroy()
    ptrsForResourcePreloader.fnAddResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1168801743))
  }
  {
    methodName := StringNameFromStr("remove_resource")
    defer methodName.Destroy()
    ptrsForResourcePreloader.fnRemoveResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("rename_resource")
    defer methodName.Destroy()
    ptrsForResourcePreloader.fnRenameResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
  }
  {
    methodName := StringNameFromStr("has_resource")
    defer methodName.Destroy()
    ptrsForResourcePreloader.fnHasResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
  }
  {
    methodName := StringNameFromStr("get_resource")
    defer methodName.Destroy()
    ptrsForResourcePreloader.fnGetResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3742749261))
  }
  {
    methodName := StringNameFromStr("get_resource_list")
    defer methodName.Destroy()
    ptrsForResourcePreloader.fnGetResourceList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
}

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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), resource.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourcePreloader.fnAddResource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourcePreloader) RemoveResource(name StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourcePreloader.fnRemoveResource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourcePreloader) RenameResource(name StringName, newname StringName, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), newname.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourcePreloader.fnRenameResource), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourcePreloader) HasResource(name StringName, ) bool {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourcePreloader.fnHasResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ResourcePreloader) GetResource(name StringName, ) Resource {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourcePreloader.fnGetResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ResourcePreloader) GetResourceList() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourcePreloader.fnGetResourceList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
