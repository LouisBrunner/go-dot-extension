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

type ResourceLoader struct {
  Object
}

func (me *ResourceLoader) BaseClass() string {
  return "ResourceLoader"
}

func NewResourceLoader() *ResourceLoader {
  str := StringNameFromStr("ResourceLoader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceLoader{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ResourceLoaderThreadLoadStatus int
const (
  ResourceLoaderThreadLoadStatusThreadLoadInvalidResource ResourceLoaderThreadLoadStatus = 0
  ResourceLoaderThreadLoadStatusThreadLoadInProgress ResourceLoaderThreadLoadStatus = 1
  ResourceLoaderThreadLoadStatusThreadLoadFailed ResourceLoaderThreadLoadStatus = 2
  ResourceLoaderThreadLoadStatusThreadLoadLoaded ResourceLoaderThreadLoadStatus = 3
)

type ResourceLoaderCacheMode int
const (
  ResourceLoaderCacheModeCacheModeIgnore ResourceLoaderCacheMode = 0
  ResourceLoaderCacheModeCacheModeReuse ResourceLoaderCacheMode = 1
  ResourceLoaderCacheModeCacheModeReplace ResourceLoaderCacheMode = 2
)

func (me *ResourceLoader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceLoader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceLoader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ResourceLoader) LoadThreadedRequest(path String, type_hint String, use_sub_threads bool, cache_mode ResourceLoaderCacheMode, ) Error {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_threaded_request")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614384323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), type_hint.AsCTypePtr(), gdc.ConstTypePtr(&use_sub_threads) , gdc.ConstTypePtr(&cache_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&use_sub_threads)
  pinner.Pin(&cache_mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ResourceLoader) LoadThreadedGetStatus(path String, progress Array, ) ResourceLoaderThreadLoadStatus {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_threaded_get_status")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4137685479) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), progress.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ResourceLoaderThreadLoadStatus

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ResourceLoader) LoadThreadedGet(path String, ) Resource {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_threaded_get")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1748875256) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ResourceLoader) Load(path String, type_hint String, cache_mode ResourceLoaderCacheMode, ) Resource {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3358495409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), type_hint.AsCTypePtr(), gdc.ConstTypePtr(&cache_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()
  pinner.Pin(&cache_mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ResourceLoader) GetRecognizedExtensionsForType(type_ String, ) PackedStringArray {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_recognized_extensions_for_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3538744774) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{type_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ResourceLoader) AddResourceFormatLoader(format_loader ResourceFormatLoader, at_front bool, )  {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_resource_format_loader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2896595483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{format_loader.AsCTypePtr(), gdc.ConstTypePtr(&at_front) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourceLoader) RemoveResourceFormatLoader(format_loader ResourceFormatLoader, )  {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_resource_format_loader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 405397102) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{format_loader.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourceLoader) SetAbortOnMissingResources(abort bool, )  {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_abort_on_missing_resources")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&abort) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ResourceLoader) GetDependencies(path String, ) PackedStringArray {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dependencies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3538744774) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ResourceLoader) HasCached(path String, ) bool {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_cached")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ResourceLoader) Exists(path String, type_hint String, ) bool {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("exists")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4185558881) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), type_hint.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ResourceLoader) GetResourceUid(path String, ) int64 {
  classNameV := StringNameFromStr("ResourceLoader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resource_uid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1597066294) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
