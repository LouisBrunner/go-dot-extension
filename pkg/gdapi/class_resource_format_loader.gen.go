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

type ResourceFormatLoader struct {
  RefCounted
}

func (me *ResourceFormatLoader) BaseClass() string {
  return "ResourceFormatLoader"
}

func NewResourceFormatLoader() *ResourceFormatLoader {
  str := StringNameFromStr("ResourceFormatLoader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceFormatLoader{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
