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

type ptrsForSpotLight3DList struct {
}

var ptrsForSpotLight3D ptrsForSpotLight3DList

func initSpotLight3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SpotLight3D")
	defer className.Destroy()
}

type SpotLight3D struct {
	Light3D
}

func (me *SpotLight3D) BaseClass() string {
	return "SpotLight3D"
}

func NewSpotLight3D() *SpotLight3D {
	str := StringNameFromStr("SpotLight3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SpotLight3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SpotLight3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SpotLight3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SpotLight3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
