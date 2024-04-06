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

type ptrsForOfflineMultiplayerPeerList struct {
}

var ptrsForOfflineMultiplayerPeer ptrsForOfflineMultiplayerPeerList

func initOfflineMultiplayerPeerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OfflineMultiplayerPeer")
	defer className.Destroy()

}

type OfflineMultiplayerPeer struct {
	MultiplayerPeer
}

func (me *OfflineMultiplayerPeer) BaseClass() string {
	return "OfflineMultiplayerPeer"
}

func NewOfflineMultiplayerPeer() *OfflineMultiplayerPeer {
	str := StringNameFromStr("OfflineMultiplayerPeer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OfflineMultiplayerPeer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OfflineMultiplayerPeer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OfflineMultiplayerPeer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OfflineMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
