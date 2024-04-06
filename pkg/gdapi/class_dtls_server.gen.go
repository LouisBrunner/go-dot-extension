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

type ptrsForDTLSServerList struct {
	fnSetup          gdc.MethodBindPtr
	fnTakeConnection gdc.MethodBindPtr
}

var ptrsForDTLSServer ptrsForDTLSServerList

func initDTLSServerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("DTLSServer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("setup")
		defer methodName.Destroy()
		ptrsForDTLSServer.fnSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262296096))
	}
	{
		methodName := StringNameFromStr("take_connection")
		defer methodName.Destroy()
		ptrsForDTLSServer.fnTakeConnection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3946580474))
	}
}

type DTLSServer struct {
	RefCounted
}

func (me *DTLSServer) BaseClass() string {
	return "DTLSServer"
}

func NewDTLSServer() *DTLSServer {
	str := StringNameFromStr("DTLSServer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &DTLSServer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *DTLSServer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *DTLSServer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *DTLSServer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *DTLSServer) Setup(server_options TLSOptions) Error {
	cargs := []gdc.ConstTypePtr{server_options.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDTLSServer.fnSetup), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DTLSServer) TakeConnection(udp_peer PacketPeerUDP) PacketPeerDTLS {
	cargs := []gdc.ConstTypePtr{udp_peer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPacketPeerDTLS()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDTLSServer.fnTakeConnection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
