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

type ptrsForCompressedCubemapArrayList struct {
}

var ptrsForCompressedCubemapArray ptrsForCompressedCubemapArrayList

func initCompressedCubemapArrayPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CompressedCubemapArray")
	defer className.Destroy()

}

type CompressedCubemapArray struct {
	CompressedTextureLayered
}

func (me *CompressedCubemapArray) BaseClass() string {
	return "CompressedCubemapArray"
}

func NewCompressedCubemapArray() *CompressedCubemapArray {
	str := StringNameFromStr("CompressedCubemapArray") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CompressedCubemapArray{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CompressedCubemapArray) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CompressedCubemapArray) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CompressedCubemapArray) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
