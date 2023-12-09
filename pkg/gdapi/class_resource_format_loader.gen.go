// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceFormatLoader struct {
  obj gdc.ObjectPtr
}

func (me *ResourceFormatLoader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceFormatLoader) BaseClass() string {
  return "ResourceFormatLoader"
}



// Enums

type ResourceFormatLoaderCacheMode int
const (
  ResourceFormatLoaderCacheModeCacheModeIgnore ResourceFormatLoaderCacheMode = 0
  ResourceFormatLoaderCacheModeCacheModeReuse ResourceFormatLoaderCacheMode = 1
  ResourceFormatLoaderCacheModeCacheModeReplace ResourceFormatLoaderCacheMode = 2
)

func (me *ResourceFormatLoader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceFormatLoader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceFormatLoader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ResourceFormatLoader) XGetRecognizedExtensions()  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XRecognizePath(path String, type_ StringName, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XHandlesType(type_ StringName, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XGetResourceType(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XGetResourceScriptClass(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XGetResourceUid(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XGetDependencies(path String, add_types bool, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XRenameDependencies(path String, renames Dictionary, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XExists(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XGetClassesUsed(path String, )  {
  panic("TODO: implement")
}

func  (me *ResourceFormatLoader) XLoad(path String, original_path String, use_sub_threads bool, cache_mode int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
