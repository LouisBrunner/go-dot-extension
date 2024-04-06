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

type ptrsForUDPServerList struct {
	fnListen                   gdc.MethodBindPtr
	fnPoll                     gdc.MethodBindPtr
	fnIsConnectionAvailable    gdc.MethodBindPtr
	fnGetLocalPort             gdc.MethodBindPtr
	fnIsListening              gdc.MethodBindPtr
	fnTakeConnection           gdc.MethodBindPtr
	fnStop                     gdc.MethodBindPtr
	fnSetMaxPendingConnections gdc.MethodBindPtr
	fnGetMaxPendingConnections gdc.MethodBindPtr
}

var ptrsForUDPServer ptrsForUDPServerList

func initUDPServerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("UDPServer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("listen")
		defer methodName.Destroy()
		ptrsForUDPServer.fnListen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3167955072))
	}
	{
		methodName := StringNameFromStr("poll")
		defer methodName.Destroy()
		ptrsForUDPServer.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("is_connection_available")
		defer methodName.Destroy()
		ptrsForUDPServer.fnIsConnectionAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_local_port")
		defer methodName.Destroy()
		ptrsForUDPServer.fnGetLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("is_listening")
		defer methodName.Destroy()
		ptrsForUDPServer.fnIsListening = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("take_connection")
		defer methodName.Destroy()
		ptrsForUDPServer.fnTakeConnection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 808734560))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForUDPServer.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_max_pending_connections")
		defer methodName.Destroy()
		ptrsForUDPServer.fnSetMaxPendingConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_pending_connections")
		defer methodName.Destroy()
		ptrsForUDPServer.fnGetMaxPendingConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type UDPServer struct {
	RefCounted
}

func (me *UDPServer) BaseClass() string {
	return "UDPServer"
}

func NewUDPServer() *UDPServer {
	str := StringNameFromStr("UDPServer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &UDPServer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *UDPServer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *UDPServer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *UDPServer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *UDPServer) Listen(port int64, bind_address String) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), bind_address.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnListen), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *UDPServer) Poll() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnPoll), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *UDPServer) IsConnectionAvailable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnIsConnectionAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UDPServer) GetLocalPort() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnGetLocalPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UDPServer) IsListening() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnIsListening), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *UDPServer) TakeConnection() PacketPeerUDP {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPacketPeerUDP()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnTakeConnection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *UDPServer) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UDPServer) SetMaxPendingConnections(max_pending_connections int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_pending_connections)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnSetMaxPendingConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *UDPServer) GetMaxPendingConnections() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUDPServer.fnGetMaxPendingConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
