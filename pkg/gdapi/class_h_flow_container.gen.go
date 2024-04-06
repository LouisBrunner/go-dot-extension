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

type ptrsForHFlowContainerList struct {
}

var ptrsForHFlowContainer ptrsForHFlowContainerList

func initHFlowContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("HFlowContainer")
	defer className.Destroy()
}

type HFlowContainer struct {
	FlowContainer
}

func (me *HFlowContainer) BaseClass() string {
	return "HFlowContainer"
}

func NewHFlowContainer() *HFlowContainer {
	str := StringNameFromStr("HFlowContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &HFlowContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *HFlowContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *HFlowContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *HFlowContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
