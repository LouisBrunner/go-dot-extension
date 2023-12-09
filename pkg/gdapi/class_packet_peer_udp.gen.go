// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PacketPeerUDP struct {
  obj gdc.ObjectPtr
}

func (me *PacketPeerUDP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func  (me *PacketPeerUDP) Bind(port int, bind_address String, recv_buf_size int, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) Close()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) Wait()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) IsBound()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) ConnectToHost(host String, port int, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) IsSocketConnected()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) GetPacketIp()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) GetPacketPort()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) GetLocalPort()  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) SetDestAddress(host String, port int, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) SetBroadcastEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) JoinMulticastGroup(multicast_address String, interface_name String, )  {
  panic("TODO: implement")
}

func  (me *PacketPeerUDP) LeaveMulticastGroup(multicast_address String, interface_name String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
