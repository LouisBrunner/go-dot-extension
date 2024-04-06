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

type ptrsForEditorExportPlatformWindowsList struct {
}

var ptrsForEditorExportPlatformWindows ptrsForEditorExportPlatformWindowsList

func initEditorExportPlatformWindowsPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorExportPlatformWindows")
	defer className.Destroy()
}

type EditorExportPlatformWindows struct {
	EditorExportPlatformPC
}

func (me *EditorExportPlatformWindows) BaseClass() string {
	return "EditorExportPlatformWindows"
}

func NewEditorExportPlatformWindows() *EditorExportPlatformWindows {
	str := StringNameFromStr("EditorExportPlatformWindows") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorExportPlatformWindows{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorExportPlatformWindows) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorExportPlatformWindows) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformWindows) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
