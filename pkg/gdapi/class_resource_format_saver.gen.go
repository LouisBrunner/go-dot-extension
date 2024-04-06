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

type ptrsForResourceFormatSaverList struct {
	fnXSave                    gdc.MethodBindPtr
	fnXSetUid                  gdc.MethodBindPtr
	fnXRecognize               gdc.MethodBindPtr
	fnXGetRecognizedExtensions gdc.MethodBindPtr
	fnXRecognizePath           gdc.MethodBindPtr
}

var ptrsForResourceFormatSaver ptrsForResourceFormatSaverList

func initResourceFormatSaverPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ResourceFormatSaver")
	defer className.Destroy()
}

type ResourceFormatSaver struct {
	RefCounted
}

func (me *ResourceFormatSaver) BaseClass() string {
	return "ResourceFormatSaver"
}

func NewResourceFormatSaver() *ResourceFormatSaver {
	str := StringNameFromStr("ResourceFormatSaver") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ResourceFormatSaver{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ResourceFormatSaver) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ResourceFormatSaver) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ResourceFormatSaver) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
