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

type ptrsForPlaceholderTexture2DList struct {
	fnSetSize gdc.MethodBindPtr
}

var ptrsForPlaceholderTexture2D ptrsForPlaceholderTexture2DList

func initPlaceholderTexture2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PlaceholderTexture2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForPlaceholderTexture2D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}

}

type PlaceholderTexture2D struct {
	Texture2D
}

func (me *PlaceholderTexture2D) BaseClass() string {
	return "PlaceholderTexture2D"
}

func NewPlaceholderTexture2D() *PlaceholderTexture2D {
	str := StringNameFromStr("PlaceholderTexture2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PlaceholderTexture2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PlaceholderTexture2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PlaceholderTexture2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PlaceholderTexture2D) SetSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderTexture2D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
