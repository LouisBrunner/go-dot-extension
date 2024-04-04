// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type ENetConnection struct {
  RefCounted
}

func (me *ENetConnection) BaseClass() string {
  return "ENetConnection"
}

func NewENetConnection() *ENetConnection {
  str := StringNameFromStr("ENetConnection") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ENetConnection{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ENetConnectionCompressionMode int
const (
  ENetConnectionCompressionModeCompressNone ENetConnectionCompressionMode = 0
  ENetConnectionCompressionModeCompressRangeCoder ENetConnectionCompressionMode = 1
  ENetConnectionCompressionModeCompressFastlz ENetConnectionCompressionMode = 2
  ENetConnectionCompressionModeCompressZlib ENetConnectionCompressionMode = 3
  ENetConnectionCompressionModeCompressZstd ENetConnectionCompressionMode = 4
)

type ENetConnectionEventType int
const (
  ENetConnectionEventTypeEventError ENetConnectionEventType = -1
  ENetConnectionEventTypeEventNone ENetConnectionEventType = 0
  ENetConnectionEventTypeEventConnect ENetConnectionEventType = 1
  ENetConnectionEventTypeEventDisconnect ENetConnectionEventType = 2
  ENetConnectionEventTypeEventReceive ENetConnectionEventType = 3
)

type ENetConnectionHostStatistic int
const (
  ENetConnectionHostStatisticHostTotalSentData ENetConnectionHostStatistic = 0
  ENetConnectionHostStatisticHostTotalSentPackets ENetConnectionHostStatistic = 1
  ENetConnectionHostStatisticHostTotalReceivedData ENetConnectionHostStatistic = 2
  ENetConnectionHostStatisticHostTotalReceivedPackets ENetConnectionHostStatistic = 3
)

func (me *ENetConnection) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ENetConnection) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ENetConnection) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ENetConnection) CreateHostBound(bind_address String, bind_port int64, max_peers int64, max_channels int64, in_bandwidth int64, out_bandwidth int64, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_host_bound")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1515002313) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bind_address.AsCTypePtr(), gdc.ConstTypePtr(&bind_port) , gdc.ConstTypePtr(&max_peers) , gdc.ConstTypePtr(&max_channels) , gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&bind_port)
  pinner.Pin(&max_peers)
  pinner.Pin(&max_channels)
  pinner.Pin(&in_bandwidth)
  pinner.Pin(&out_bandwidth)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) CreateHost(max_peers int64, max_channels int64, in_bandwidth int64, out_bandwidth int64, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 117198950) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_peers) , gdc.ConstTypePtr(&max_channels) , gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&max_peers)
  pinner.Pin(&max_channels)
  pinner.Pin(&in_bandwidth)
  pinner.Pin(&out_bandwidth)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) Destroy()  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("destroy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) ConnectToHost(address String, port int64, channels int64, data int64, ) ENetPacketPeer {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2171300490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{address.AsCTypePtr(), gdc.ConstTypePtr(&port) , gdc.ConstTypePtr(&channels) , gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewENetPacketPeer()
  pinner.Pin(&port)
  pinner.Pin(&channels)
  pinner.Pin(&data)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetConnection) Service(timeout int64, ) Array {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("service")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2402345344) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&timeout)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetConnection) Flush()  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flush")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) BandwidthLimit(in_bandwidth int64, out_bandwidth int64, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bandwidth_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2302169788) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) ChannelLimit(limit int64, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("channel_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) Broadcast(channel int64, packet PackedByteArray, flags int64, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("broadcast")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2772371345) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , packet.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) Compress(mode ENetConnectionCompressionMode, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2660215187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) DtlsServerSetup(server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dtls_server_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262296096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{server_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) DtlsClientSetup(hostname String, client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dtls_client_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1966198364) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{hostname.AsCTypePtr(), client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) RefuseNewConnections(refuse bool, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("refuse_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refuse) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) PopStatistic(statistic ENetConnectionHostStatistic, ) float64 {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pop_statistic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2166904170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&statistic) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&statistic)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetConnection) GetMaxChannels() int64 {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_channels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetConnection) GetLocalPort() int64 {
  classNameV := StringNameFromStr("ENetConnection")
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

func  (me *ENetConnection) GetPeers() []ENetPacketPeer {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[ENetPacketPeer](ret)
}

func  (me *ENetConnection) SocketSend(destination_address String, destination_port int64, packet PackedByteArray, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("socket_send")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1100646812) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{destination_address.AsCTypePtr(), gdc.ConstTypePtr(&destination_port) , packet.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
