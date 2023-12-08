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

type ResourceFormatLoaderCacheMode int
const (
  ResourceFormatLoaderCacheModeCacheModeIgnore ResourceFormatLoaderCacheMode = 0
  ResourceFormatLoaderCacheModeCacheModeReuse ResourceFormatLoaderCacheMode = 1
  ResourceFormatLoaderCacheModeCacheModeReplace ResourceFormatLoaderCacheMode = 2
)

func  (me *ResourceFormatLoader) XGetRecognizedExtensions() { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XRecognizePath(path String, type_ StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XHandlesType(type_ StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XGetResourceType(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XGetResourceScriptClass(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XGetResourceUid(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XGetDependencies(path String, add_types bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XRenameDependencies(path String, renames Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XExists(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XGetClassesUsed(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ResourceFormatLoader) XLoad(path String, original_path String, use_sub_threads bool, cache_mode int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
