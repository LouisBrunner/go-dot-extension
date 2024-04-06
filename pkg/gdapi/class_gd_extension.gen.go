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

type ptrsForGDExtensionList struct {
	fnOpenLibrary                          gdc.MethodBindPtr
	fnCloseLibrary                         gdc.MethodBindPtr
	fnIsLibraryOpen                        gdc.MethodBindPtr
	fnGetMinimumLibraryInitializationLevel gdc.MethodBindPtr
	fnInitializeLibrary                    gdc.MethodBindPtr
}

var ptrsForGDExtension ptrsForGDExtensionList

func initGDExtensionPtrs(iface gdc.Interface) {
	return

	className := StringNameFromStr("GDExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("open_library")
		defer methodName.Destroy()
		ptrsForGDExtension.fnOpenLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("close_library")
		defer methodName.Destroy()
		ptrsForGDExtension.fnCloseLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_library_open")
		defer methodName.Destroy()
		ptrsForGDExtension.fnIsLibraryOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_minimum_library_initialization_level")
		defer methodName.Destroy()
		ptrsForGDExtension.fnGetMinimumLibraryInitializationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 964858755))
	}
	{
		methodName := StringNameFromStr("initialize_library")
		defer methodName.Destroy()
		ptrsForGDExtension.fnInitializeLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3409922941))
	}
}

type GDExtension struct {
	Resource
}

func (me *GDExtension) BaseClass() string {
	return "GDExtension"
}

func NewGDExtension() *GDExtension {
	str := StringNameFromStr("GDExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GDExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GDExtensionInitializationLevel int

const (
	GDExtensionInitializationLevelInitializationLevelCore    GDExtensionInitializationLevel = 0
	GDExtensionInitializationLevelInitializationLevelServers GDExtensionInitializationLevel = 1
	GDExtensionInitializationLevelInitializationLevelScene   GDExtensionInitializationLevel = 2
	GDExtensionInitializationLevelInitializationLevelEditor  GDExtensionInitializationLevel = 3
)

func (me *GDExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GDExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GDExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GDExtension) OpenLibrary(path String, entry_symbol String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), entry_symbol.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtension.fnOpenLibrary), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GDExtension) CloseLibrary() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtension.fnCloseLibrary), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GDExtension) IsLibraryOpen() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtension.fnIsLibraryOpen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GDExtension) GetMinimumLibraryInitializationLevel() GDExtensionInitializationLevel {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GDExtensionInitializationLevel

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtension.fnGetMinimumLibraryInitializationLevel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GDExtension) InitializeLibrary(level GDExtensionInitializationLevel) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&level)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGDExtension.fnInitializeLibrary), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
