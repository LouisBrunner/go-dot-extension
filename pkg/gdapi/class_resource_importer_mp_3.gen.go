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

type ptrsForResourceImporterMP3List struct {
}

var ptrsForResourceImporterMP3 ptrsForResourceImporterMP3List

func initResourceImporterMP3Ptrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceImporterMP3")
	defer className.Destroy()

}

type ResourceImporterMP3 struct {
	ResourceImporter
}

func (me *ResourceImporterMP3) BaseClass() string {
	return "ResourceImporterMP3"
}

func NewResourceImporterMP3() *ResourceImporterMP3 {
	str := StringNameFromStr("ResourceImporterMP3") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceImporterMP3{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ResourceImporterMP3) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceImporterMP3) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterMP3) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
