// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceLoader struct {
  obj gdc.ObjectPtr
}

func (me *ResourceLoader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceLoader) BaseClass() string {
  return "ResourceLoader"
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

func (me *ResourceLoader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceLoader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ResourceLoader) LoadThreadedRequest(path String, type_hint String, use_sub_threads bool, cache_mode ResourceLoaderCacheMode, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) LoadThreadedGetStatus(path String, progress Array, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) LoadThreadedGet(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) Load(path String, type_hint String, cache_mode ResourceLoaderCacheMode, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) GetRecognizedExtensionsForType(type_ String, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) AddResourceFormatLoader(format_loader ResourceFormatLoader, at_front bool, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) RemoveResourceFormatLoader(format_loader ResourceFormatLoader, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) SetAbortOnMissingResources(abort bool, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) GetDependencies(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) HasCached(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) Exists(path String, type_hint String, )  {
  panic("TODO: implement")
}

func  (me *ResourceLoader) GetResourceUid(path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
