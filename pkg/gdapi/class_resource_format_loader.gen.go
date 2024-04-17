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

type ptrsForResourceFormatLoaderList struct {
	fnXGetRecognizedExtensions gdc.MethodBindPtr
	fnXRecognizePath           gdc.MethodBindPtr
	fnXHandlesType             gdc.MethodBindPtr
	fnXGetResourceType         gdc.MethodBindPtr
	fnXGetResourceScriptClass  gdc.MethodBindPtr
	fnXGetResourceUid          gdc.MethodBindPtr
	fnXGetDependencies         gdc.MethodBindPtr
	fnXRenameDependencies      gdc.MethodBindPtr
	fnXExists                  gdc.MethodBindPtr
	fnXGetClassesUsed          gdc.MethodBindPtr
	fnXLoad                    gdc.MethodBindPtr
}

var ptrsForResourceFormatLoader ptrsForResourceFormatLoaderList

func initResourceFormatLoaderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceFormatLoader")
	defer className.Destroy()

}

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
	ResourceFormatLoaderCacheModeCacheModeIgnore      ResourceFormatLoaderCacheMode = 0
	ResourceFormatLoaderCacheModeCacheModeReuse       ResourceFormatLoaderCacheMode = 1
	ResourceFormatLoaderCacheModeCacheModeReplace     ResourceFormatLoaderCacheMode = 2
	ResourceFormatLoaderCacheModeCacheModeIgnoreDeep  ResourceFormatLoaderCacheMode = 3
	ResourceFormatLoaderCacheModeCacheModeReplaceDeep ResourceFormatLoaderCacheMode = 4
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
