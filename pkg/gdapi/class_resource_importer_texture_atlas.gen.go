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

type ptrsForResourceImporterTextureAtlasList struct {
}

var ptrsForResourceImporterTextureAtlas ptrsForResourceImporterTextureAtlasList

func initResourceImporterTextureAtlasPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceImporterTextureAtlas")
	defer className.Destroy()

}

type ResourceImporterTextureAtlas struct {
	ResourceImporter
}

func (me *ResourceImporterTextureAtlas) BaseClass() string {
	return "ResourceImporterTextureAtlas"
}

func NewResourceImporterTextureAtlas() *ResourceImporterTextureAtlas {
	str := StringNameFromStr("ResourceImporterTextureAtlas") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceImporterTextureAtlas{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ResourceImporterTextureAtlas) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceImporterTextureAtlas) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterTextureAtlas) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
