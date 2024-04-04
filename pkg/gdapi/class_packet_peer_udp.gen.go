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

func  (me *PacketPeerUDP) Bind(port int64, bind_address String, recv_buf_size int64, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051239242) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , bind_address.AsCTypePtr(), gdc.ConstTypePtr(&recv_buf_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)
  pinner.Pin(&recv_buf_size)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerUDP) Close()  {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerUDP) Wait() Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerUDP) IsBound() bool {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bound")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeerUDP) ConnectToHost(host String, port int64, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 993915709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerUDP) IsSocketConnected() bool {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_socket_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeerUDP) GetPacketIp() String {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_ip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PacketPeerUDP) GetPacketPort() int64 {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeerUDP) GetLocalPort() int64 {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PacketPeerUDP) SetDestAddress(host String, port int64, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dest_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 993915709) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{host.AsCTypePtr(), gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerUDP) SetBroadcastEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_broadcast_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PacketPeerUDP) JoinMulticastGroup(multicast_address String, interface_name String, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("join_multicast_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{multicast_address.AsCTypePtr(), interface_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PacketPeerUDP) LeaveMulticastGroup(multicast_address String, interface_name String, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("leave_multicast_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{multicast_address.AsCTypePtr(), interface_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
