// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

// Properties