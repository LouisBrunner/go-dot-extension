// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *DTLSServer) Setup(server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("DTLSServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262296096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(server_options.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *DTLSServer) TakeConnection(udp_peer PacketPeerUDP, ) PacketPeerDTLS {
  classNameV := StringNameFromStr("DTLSServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("take_connection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3946580474) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(udp_peer.AsCTypePtr()), }
  ret := NewPacketPeerDTLS()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
