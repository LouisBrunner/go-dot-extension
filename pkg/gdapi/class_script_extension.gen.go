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

type ptrsForScriptExtensionList struct {
	fnXEditorCanReloadFromFile      gdc.MethodBindPtr
	fnXPlaceholderErased            gdc.MethodBindPtr
	fnXCanInstantiate               gdc.MethodBindPtr
	fnXGetBaseScript                gdc.MethodBindPtr
	fnXGetGlobalName                gdc.MethodBindPtr
	fnXInheritsScript               gdc.MethodBindPtr
	fnXGetInstanceBaseType          gdc.MethodBindPtr
	fnXInstanceCreate               gdc.MethodBindPtr
	fnXPlaceholderInstanceCreate    gdc.MethodBindPtr
	fnXInstanceHas                  gdc.MethodBindPtr
	fnXHasSourceCode                gdc.MethodBindPtr
	fnXGetSourceCode                gdc.MethodBindPtr
	fnXSetSourceCode                gdc.MethodBindPtr
	fnXReload                       gdc.MethodBindPtr
	fnXGetDocumentation             gdc.MethodBindPtr
	fnXGetClassIconPath             gdc.MethodBindPtr
	fnXHasMethod                    gdc.MethodBindPtr
	fnXHasStaticMethod              gdc.MethodBindPtr
	fnXGetMethodInfo                gdc.MethodBindPtr
	fnXIsTool                       gdc.MethodBindPtr
	fnXIsValid                      gdc.MethodBindPtr
	fnXIsAbstract                   gdc.MethodBindPtr
	fnXGetLanguage                  gdc.MethodBindPtr
	fnXHasScriptSignal              gdc.MethodBindPtr
	fnXGetScriptSignalList          gdc.MethodBindPtr
	fnXHasPropertyDefaultValue      gdc.MethodBindPtr
	fnXGetPropertyDefaultValue      gdc.MethodBindPtr
	fnXUpdateExports                gdc.MethodBindPtr
	fnXGetScriptMethodList          gdc.MethodBindPtr
	fnXGetScriptPropertyList        gdc.MethodBindPtr
	fnXGetMemberLine                gdc.MethodBindPtr
	fnXGetConstants                 gdc.MethodBindPtr
	fnXGetMembers                   gdc.MethodBindPtr
	fnXIsPlaceholderFallbackEnabled gdc.MethodBindPtr
	fnXGetRpcConfig                 gdc.MethodBindPtr
}

var ptrsForScriptExtension ptrsForScriptExtensionList

func initScriptExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ScriptExtension")
	defer className.Destroy()

}

type ScriptExtension struct {
	Script
}

func (me *ScriptExtension) BaseClass() string {
	return "ScriptExtension"
}

func NewScriptExtension() *ScriptExtension {
	str := StringNameFromStr("ScriptExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ScriptExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ScriptExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ScriptExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ScriptExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
