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

type ptrsForResourceLoaderList struct {
	fnLoadThreadedRequest            gdc.MethodBindPtr
	fnLoadThreadedGetStatus          gdc.MethodBindPtr
	fnLoadThreadedGet                gdc.MethodBindPtr
	fnLoad                           gdc.MethodBindPtr
	fnGetRecognizedExtensionsForType gdc.MethodBindPtr
	fnAddResourceFormatLoader        gdc.MethodBindPtr
	fnRemoveResourceFormatLoader     gdc.MethodBindPtr
	fnSetAbortOnMissingResources     gdc.MethodBindPtr
	fnGetDependencies                gdc.MethodBindPtr
	fnHasCached                      gdc.MethodBindPtr
	fnExists                         gdc.MethodBindPtr
	fnGetResourceUid                 gdc.MethodBindPtr
}

var ptrsForResourceLoader ptrsForResourceLoaderList

func initResourceLoaderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceLoader")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("load_threaded_request")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnLoadThreadedRequest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614384323))
	}
	{
		methodName := StringNameFromStr("load_threaded_get_status")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnLoadThreadedGetStatus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4137685479))
	}
	{
		methodName := StringNameFromStr("load_threaded_get")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnLoadThreadedGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1748875256))
	}
	{
		methodName := StringNameFromStr("load")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnLoad = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3358495409))
	}
	{
		methodName := StringNameFromStr("get_recognized_extensions_for_type")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnGetRecognizedExtensionsForType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3538744774))
	}
	{
		methodName := StringNameFromStr("add_resource_format_loader")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnAddResourceFormatLoader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2896595483))
	}
	{
		methodName := StringNameFromStr("remove_resource_format_loader")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnRemoveResourceFormatLoader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 405397102))
	}
	{
		methodName := StringNameFromStr("set_abort_on_missing_resources")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnSetAbortOnMissingResources = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_dependencies")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnGetDependencies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3538744774))
	}
	{
		methodName := StringNameFromStr("has_cached")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnHasCached = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("exists")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnExists = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4185558881))
	}
	{
		methodName := StringNameFromStr("get_resource_uid")
		defer methodName.Destroy()
		ptrsForResourceLoader.fnGetResourceUid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1597066294))
	}

}

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
	ResourceLoaderThreadLoadStatusThreadLoadInProgress      ResourceLoaderThreadLoadStatus = 1
	ResourceLoaderThreadLoadStatusThreadLoadFailed          ResourceLoaderThreadLoadStatus = 2
	ResourceLoaderThreadLoadStatusThreadLoadLoaded          ResourceLoaderThreadLoadStatus = 3
)

type ResourceLoaderCacheMode int

const (
	ResourceLoaderCacheModeCacheModeIgnore      ResourceLoaderCacheMode = 0
	ResourceLoaderCacheModeCacheModeReuse       ResourceLoaderCacheMode = 1
	ResourceLoaderCacheModeCacheModeReplace     ResourceLoaderCacheMode = 2
	ResourceLoaderCacheModeCacheModeIgnoreDeep  ResourceLoaderCacheMode = 3
	ResourceLoaderCacheModeCacheModeReplaceDeep ResourceLoaderCacheMode = 4
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

func (me *ResourceLoader) LoadThreadedRequest(path String, type_hint String, use_sub_threads bool, cache_mode ResourceLoaderCacheMode) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), type_hint.AsCTypePtr(), gdc.ConstTypePtr(&use_sub_threads), gdc.ConstTypePtr(&cache_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&use_sub_threads)
	pinner.Pin(&cache_mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnLoadThreadedRequest), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ResourceLoader) LoadThreadedGetStatus(path String, progress Array) ResourceLoaderThreadLoadStatus {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), progress.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ResourceLoaderThreadLoadStatus

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnLoadThreadedGetStatus), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ResourceLoader) LoadThreadedGet(path String) Resource {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnLoadThreadedGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceLoader) Load(path String, type_hint String, cache_mode ResourceLoaderCacheMode) Resource {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), type_hint.AsCTypePtr(), gdc.ConstTypePtr(&cache_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()
	pinner.Pin(&cache_mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnLoad), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceLoader) GetRecognizedExtensionsForType(type_ String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnGetRecognizedExtensionsForType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceLoader) AddResourceFormatLoader(format_loader ResourceFormatLoader, at_front bool) {
	cargs := []gdc.ConstTypePtr{format_loader.AsCTypePtr(), gdc.ConstTypePtr(&at_front)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnAddResourceFormatLoader), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ResourceLoader) RemoveResourceFormatLoader(format_loader ResourceFormatLoader) {
	cargs := []gdc.ConstTypePtr{format_loader.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnRemoveResourceFormatLoader), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ResourceLoader) SetAbortOnMissingResources(abort bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&abort)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnSetAbortOnMissingResources), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ResourceLoader) GetDependencies(path String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnGetDependencies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceLoader) HasCached(path String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnHasCached), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ResourceLoader) Exists(path String, type_hint String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), type_hint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnExists), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ResourceLoader) GetResourceUid(path String) int64 {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceLoader.fnGetResourceUid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
