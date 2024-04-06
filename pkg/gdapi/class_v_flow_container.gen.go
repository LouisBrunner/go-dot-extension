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

type ptrsForVFlowContainerList struct {
}

var ptrsForVFlowContainer ptrsForVFlowContainerList

func initVFlowContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VFlowContainer")
	defer className.Destroy()
}

type VFlowContainer struct {
	FlowContainer
}

func (me *VFlowContainer) BaseClass() string {
	return "VFlowContainer"
}

func NewVFlowContainer() *VFlowContainer {
	str := StringNameFromStr("VFlowContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VFlowContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VFlowContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VFlowContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VFlowContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
