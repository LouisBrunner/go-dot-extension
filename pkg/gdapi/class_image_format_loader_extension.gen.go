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

type ptrsForImageFormatLoaderExtensionList struct {
	fnXGetRecognizedExtensions gdc.MethodBindPtr
	fnXLoadImage               gdc.MethodBindPtr
	fnAddFormatLoader          gdc.MethodBindPtr
	fnRemoveFormatLoader       gdc.MethodBindPtr
}

var ptrsForImageFormatLoaderExtension ptrsForImageFormatLoaderExtensionList

func initImageFormatLoaderExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ImageFormatLoaderExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_format_loader")
		defer methodName.Destroy()
		ptrsForImageFormatLoaderExtension.fnAddFormatLoader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("remove_format_loader")
		defer methodName.Destroy()
		ptrsForImageFormatLoaderExtension.fnRemoveFormatLoader = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type ImageFormatLoaderExtension struct {
	ImageFormatLoader
}

func (me *ImageFormatLoaderExtension) BaseClass() string {
	return "ImageFormatLoaderExtension"
}

func NewImageFormatLoaderExtension() *ImageFormatLoaderExtension {
	str := StringNameFromStr("ImageFormatLoaderExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ImageFormatLoaderExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ImageFormatLoaderExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ImageFormatLoaderExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ImageFormatLoaderExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ImageFormatLoaderExtension) AddFormatLoader() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageFormatLoaderExtension.fnAddFormatLoader), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImageFormatLoaderExtension) RemoveFormatLoader() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageFormatLoaderExtension.fnRemoveFormatLoader), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
