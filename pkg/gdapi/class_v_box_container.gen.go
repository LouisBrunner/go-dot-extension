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

type ptrsForVBoxContainerList struct {
}

var ptrsForVBoxContainer ptrsForVBoxContainerList

func initVBoxContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VBoxContainer")
	defer className.Destroy()

}

type VBoxContainer struct {
	BoxContainer
}

func (me *VBoxContainer) BaseClass() string {
	return "VBoxContainer"
}

func NewVBoxContainer() *VBoxContainer {
	str := StringNameFromStr("VBoxContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VBoxContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VBoxContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VBoxContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VBoxContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
