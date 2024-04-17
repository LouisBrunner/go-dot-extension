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

type ptrsForFBXDocumentList struct {
}

var ptrsForFBXDocument ptrsForFBXDocumentList

func initFBXDocumentPtrs(iface gdc.Interface) {

	className := StringNameFromStr("FBXDocument")
	defer className.Destroy()

}

type FBXDocument struct {
	GLTFDocument
}

func (me *FBXDocument) BaseClass() string {
	return "FBXDocument"
}

func NewFBXDocument() *FBXDocument {
	str := StringNameFromStr("FBXDocument") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FBXDocument{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *FBXDocument) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FBXDocument) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FBXDocument) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
