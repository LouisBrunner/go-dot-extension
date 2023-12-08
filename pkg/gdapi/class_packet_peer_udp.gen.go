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

func  (me *PacketPeerUDP) Bind(port int, bind_address String, recv_buf_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) Close() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) Wait() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) IsBound() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) ConnectToHost(host String, port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) IsSocketConnected() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) GetPacketIp() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) GetPacketPort() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) GetLocalPort() { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) SetDestAddress(host String, port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) SetBroadcastEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) JoinMulticastGroup(multicast_address String, interface_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *PacketPeerUDP) LeaveMulticastGroup(multicast_address String, interface_name String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
