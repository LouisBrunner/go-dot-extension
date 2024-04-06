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

type ptrsForTextureList struct {
}

var ptrsForTexture ptrsForTextureList

func initTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Texture")
	defer className.Destroy()
}

type Texture struct {
	Resource
}

func (me *Texture) BaseClass() string {
	return "Texture"
}

func NewTexture() *Texture {
	str := StringNameFromStr("Texture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Texture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Texture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Texture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Texture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
