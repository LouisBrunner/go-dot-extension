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

type ptrsForResourceSaverList struct {
	fnSave                      gdc.MethodBindPtr
	fnGetRecognizedExtensions   gdc.MethodBindPtr
	fnAddResourceFormatSaver    gdc.MethodBindPtr
	fnRemoveResourceFormatSaver gdc.MethodBindPtr
}

var ptrsForResourceSaver ptrsForResourceSaverList

func initResourceSaverPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceSaver")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("save")
		defer methodName.Destroy()
		ptrsForResourceSaver.fnSave = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2983274697))
	}
	{
		methodName := StringNameFromStr("get_recognized_extensions")
		defer methodName.Destroy()
		ptrsForResourceSaver.fnGetRecognizedExtensions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4223597960))
	}
	{
		methodName := StringNameFromStr("add_resource_format_saver")
		defer methodName.Destroy()
		ptrsForResourceSaver.fnAddResourceFormatSaver = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 362894272))
	}
	{
		methodName := StringNameFromStr("remove_resource_format_saver")
		defer methodName.Destroy()
		ptrsForResourceSaver.fnRemoveResourceFormatSaver = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3373026878))
	}

}

type ResourceSaver struct {
	Object
}

func (me *ResourceSaver) BaseClass() string {
	return "ResourceSaver"
}

func NewResourceSaver() *ResourceSaver {
	str := StringNameFromStr("ResourceSaver") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceSaver{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ResourceSaverSaverFlags int

const (
	ResourceSaverSaverFlagsFlagNone                    ResourceSaverSaverFlags = 0
	ResourceSaverSaverFlagsFlagRelativePaths           ResourceSaverSaverFlags = 1
	ResourceSaverSaverFlagsFlagBundleResources         ResourceSaverSaverFlags = 2
	ResourceSaverSaverFlagsFlagChangePath              ResourceSaverSaverFlags = 4
	ResourceSaverSaverFlagsFlagOmitEditorProperties    ResourceSaverSaverFlags = 8
	ResourceSaverSaverFlagsFlagSaveBigEndian           ResourceSaverSaverFlags = 16
	ResourceSaverSaverFlagsFlagCompress                ResourceSaverSaverFlags = 32
	ResourceSaverSaverFlagsFlagReplaceSubresourcePaths ResourceSaverSaverFlags = 64
)

func (me *ResourceSaver) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceSaver) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceSaver) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ResourceSaver) Save(resource Resource, path String, flags ResourceSaverSaverFlags) Error {
	cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), path.AsCTypePtr(), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceSaver.fnSave), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ResourceSaver) GetRecognizedExtensions(type_ Resource) PackedStringArray {
	cargs := []gdc.ConstTypePtr{type_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceSaver.fnGetRecognizedExtensions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ResourceSaver) AddResourceFormatSaver(format_saver ResourceFormatSaver, at_front bool) {
	cargs := []gdc.ConstTypePtr{format_saver.AsCTypePtr(), gdc.ConstTypePtr(&at_front)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceSaver.fnAddResourceFormatSaver), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ResourceSaver) RemoveResourceFormatSaver(format_saver ResourceFormatSaver) {
	cargs := []gdc.ConstTypePtr{format_saver.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForResourceSaver.fnRemoveResourceFormatSaver), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
