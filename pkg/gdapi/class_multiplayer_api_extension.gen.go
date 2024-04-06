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

type ptrsForMultiplayerAPIExtensionList struct {
	fnXPoll                      gdc.MethodBindPtr
	fnXSetMultiplayerPeer        gdc.MethodBindPtr
	fnXGetMultiplayerPeer        gdc.MethodBindPtr
	fnXGetUniqueId               gdc.MethodBindPtr
	fnXGetPeerIds                gdc.MethodBindPtr
	fnXRpc                       gdc.MethodBindPtr
	fnXGetRemoteSenderId         gdc.MethodBindPtr
	fnXObjectConfigurationAdd    gdc.MethodBindPtr
	fnXObjectConfigurationRemove gdc.MethodBindPtr
}

var ptrsForMultiplayerAPIExtension ptrsForMultiplayerAPIExtensionList

func initMultiplayerAPIExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MultiplayerAPIExtension")
	defer className.Destroy()

}

type MultiplayerAPIExtension struct {
	MultiplayerAPI
}

func (me *MultiplayerAPIExtension) BaseClass() string {
	return "MultiplayerAPIExtension"
}

func NewMultiplayerAPIExtension() *MultiplayerAPIExtension {
	str := StringNameFromStr("MultiplayerAPIExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MultiplayerAPIExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MultiplayerAPIExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MultiplayerAPIExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MultiplayerAPIExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
