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

type ptrsForEditorExportPlatformWebList struct {
}

var ptrsForEditorExportPlatformWeb ptrsForEditorExportPlatformWebList

func initEditorExportPlatformWebPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorExportPlatformWeb")
	defer className.Destroy()
}

type EditorExportPlatformWeb struct {
	EditorExportPlatform
}

func (me *EditorExportPlatformWeb) BaseClass() string {
	return "EditorExportPlatformWeb"
}

func NewEditorExportPlatformWeb() *EditorExportPlatformWeb {
	str := StringNameFromStr("EditorExportPlatformWeb") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorExportPlatformWeb{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorExportPlatformWeb) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorExportPlatformWeb) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformWeb) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
