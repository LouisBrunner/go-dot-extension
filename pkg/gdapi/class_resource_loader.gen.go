// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type ResourceLoaderThreadLoadStatus int
const (
  ResourceLoaderThreadLoadInvalidResource ResourceLoaderThreadLoadStatus = 0
  ResourceLoaderThreadLoadInProgress ResourceLoaderThreadLoadStatus = 1
  ResourceLoaderThreadLoadFailed ResourceLoaderThreadLoadStatus = 2
  ResourceLoaderThreadLoadLoaded ResourceLoaderThreadLoadStatus = 3
)

type ResourceLoaderCacheMode int
const (
  ResourceLoaderCacheModeIgnore ResourceLoaderCacheMode = 0
  ResourceLoaderCacheModeReuse ResourceLoaderCacheMode = 1
  ResourceLoaderCacheModeReplace ResourceLoaderCacheMode = 2
)
