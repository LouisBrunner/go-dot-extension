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

type ptrsForPanelContainerList struct {
}

var ptrsForPanelContainer ptrsForPanelContainerList

func initPanelContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PanelContainer")
	defer className.Destroy()
}

type PanelContainer struct {
	Container
}

func (me *PanelContainer) BaseClass() string {
	return "PanelContainer"
}

func NewPanelContainer() *PanelContainer {
	str := StringNameFromStr("PanelContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PanelContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PanelContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PanelContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PanelContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
