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
	fnIsLibraryOpen                        gdc.MethodBindPtr
	fnGetMinimumLibraryInitializationLevel gdc.MethodBindPtr
}

var ptrsForGDExtension ptrsForGDExtensionList

func initGDExtensionPtrs(iface gdc.Interface) {

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

// Signals
