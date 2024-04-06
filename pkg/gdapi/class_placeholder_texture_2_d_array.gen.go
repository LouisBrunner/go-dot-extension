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

type ptrsForPlaceholderTexture2DArrayList struct {
}

var ptrsForPlaceholderTexture2DArray ptrsForPlaceholderTexture2DArrayList

func initPlaceholderTexture2DArrayPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PlaceholderTexture2DArray")
	defer className.Destroy()
}

type PlaceholderTexture2DArray struct {
	PlaceholderTextureLayered
}

func (me *PlaceholderTexture2DArray) BaseClass() string {
	return "PlaceholderTexture2DArray"
}

func NewPlaceholderTexture2DArray() *PlaceholderTexture2DArray {
	str := StringNameFromStr("PlaceholderTexture2DArray") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PlaceholderTexture2DArray{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PlaceholderTexture2DArray) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PlaceholderTexture2DArray) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture2DArray) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
