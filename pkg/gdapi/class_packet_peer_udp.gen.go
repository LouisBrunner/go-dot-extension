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

type ptrsForPacketPeerUDPList struct {
	fnBind                gdc.MethodBindPtr
	fnClose               gdc.MethodBindPtr
	fnWait                gdc.MethodBindPtr
	fnIsBound             gdc.MethodBindPtr
	fnConnectToHost       gdc.MethodBindPtr
	fnIsSocketConnected   gdc.MethodBindPtr
	fnGetPacketIp         gdc.MethodBindPtr
	fnGetPacketPort       gdc.MethodBindPtr
	fnGetLocalPort        gdc.MethodBindPtr
	fnSetDestAddress      gdc.MethodBindPtr
	fnSetBroadcastEnabled gdc.MethodBindPtr
	fnJoinMulticastGroup  gdc.MethodBindPtr
	fnLeaveMulticastGroup gdc.MethodBindPtr
}

var ptrsForPacketPeerUDP ptrsForPacketPeerUDPList

func initPacketPeerUDPPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PacketPeerUDP")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("bind")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnBind = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051239242))
	}
	{
		methodName := StringNameFromStr("close")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("wait")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnWait = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("is_bound")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnIsBound = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("connect_to_host")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnConnectToHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 993915709))
	}
	{
		methodName := StringNameFromStr("is_socket_connected")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnIsSocketConnected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_packet_ip")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnGetPacketIp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_packet_port")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnGetPacketPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_local_port")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnGetLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_dest_address")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnSetDestAddress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 993915709))
	}
	{
		methodName := StringNameFromStr("set_broadcast_enabled")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnSetBroadcastEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("join_multicast_group")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnJoinMulticastGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("leave_multicast_group")
		defer methodName.Destroy()
		ptrsForPacketPeerUDP.fnLeaveMulticastGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
}

type PacketPeerUDP struct {
	PacketPeer
}

func (me *PacketPeerUDP) BaseClass() string {
	return "PacketPeerUDP"
}

func NewPacketPeerUDP() *PacketPeerUDP {
	str := StringNameFromStr("PacketPeerUDP") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PacketPeerUDP{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PacketPeerUDP) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PacketPeerUDP) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PacketPeerUDP) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PacketPeerUDP) Bind(port int64, bind_address String, recv_buf_size int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), bind_address.AsCTypePtr(), gdc.ConstTypePtr(&recv_buf_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&port)
	pinner.Pin(&recv_buf_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnBind), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PacketPeerUDP) Close() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PacketPeerUDP) Wait() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnWait), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PacketPeerUDP) IsBound() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnIsBound), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PacketPeerUDP) ConnectToHost(host String, port int64) Error {
	cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnConnectToHost), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PacketPeerUDP) IsSocketConnected() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnIsSocketConnected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PacketPeerUDP) GetPacketIp() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnGetPacketIp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PacketPeerUDP) GetPacketPort() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnGetPacketPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PacketPeerUDP) GetLocalPort() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnGetLocalPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PacketPeerUDP) SetDestAddress(host String, port int64) Error {
	cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnSetDestAddress), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PacketPeerUDP) SetBroadcastEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnSetBroadcastEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PacketPeerUDP) JoinMulticastGroup(multicast_address String, interface_name String) Error {
	cargs := []gdc.ConstTypePtr{multicast_address.AsCTypePtr(), interface_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnJoinMulticastGroup), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PacketPeerUDP) LeaveMulticastGroup(multicast_address String, interface_name String) Error {
	cargs := []gdc.ConstTypePtr{multicast_address.AsCTypePtr(), interface_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerUDP.fnLeaveMulticastGroup), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
