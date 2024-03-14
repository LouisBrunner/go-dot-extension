// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PacketPeerUDP struct {
  PacketPeer
}

func (me *PacketPeerUDP) BaseClass() string {
  return "PacketPeerUDP"
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

func  (me *PacketPeerUDP) Bind(port int, bind_address String, recv_buf_size int, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051239242) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(bind_address.AsCTypePtr()), gdc.ConstTypePtr(&recv_buf_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) Close()  {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PacketPeerUDP) Wait() Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) IsBound() bool {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bound")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) ConnectToHost(host String, port int, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 993915709) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(host.AsCTypePtr()), gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) IsSocketConnected() bool {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_socket_connected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) GetPacketIp() String {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_ip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) GetPacketPort() int {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) GetLocalPort() int {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) SetDestAddress(host String, port int, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dest_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 993915709) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(host.AsCTypePtr()), gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) SetBroadcastEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_broadcast_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PacketPeerUDP) JoinMulticastGroup(multicast_address String, interface_name String, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("join_multicast_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(multicast_address.AsCTypePtr()), gdc.ConstTypePtr(interface_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PacketPeerUDP) LeaveMulticastGroup(multicast_address String, interface_name String, ) Error {
  classNameV := StringNameFromStr("PacketPeerUDP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("leave_multicast_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 852856452) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(multicast_address.AsCTypePtr()), gdc.ConstTypePtr(interface_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
