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

type ptrsForResourceImporterCSVTranslationList struct {
}

var ptrsForResourceImporterCSVTranslation ptrsForResourceImporterCSVTranslationList

func initResourceImporterCSVTranslationPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceImporterCSVTranslation")
	defer className.Destroy()
}

type ResourceImporterCSVTranslation struct {
	ResourceImporter
}

func (me *ResourceImporterCSVTranslation) BaseClass() string {
	return "ResourceImporterCSVTranslation"
}

func NewResourceImporterCSVTranslation() *ResourceImporterCSVTranslation {
	str := StringNameFromStr("ResourceImporterCSVTranslation") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceImporterCSVTranslation{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ResourceImporterCSVTranslation) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceImporterCSVTranslation) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterCSVTranslation) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
