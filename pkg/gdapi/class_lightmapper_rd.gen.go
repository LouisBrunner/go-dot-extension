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

type ptrsForLightmapperRDList struct {
}

var ptrsForLightmapperRD ptrsForLightmapperRDList

func initLightmapperRDPtrs(iface gdc.Interface) {

	className := StringNameFromStr("LightmapperRD")
	defer className.Destroy()

}

type LightmapperRD struct {
	Lightmapper
}

func (me *LightmapperRD) BaseClass() string {
	return "LightmapperRD"
}

func NewLightmapperRD() *LightmapperRD {
	str := StringNameFromStr("LightmapperRD") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &LightmapperRD{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *LightmapperRD) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *LightmapperRD) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *LightmapperRD) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
