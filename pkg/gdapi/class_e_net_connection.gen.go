// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ENetConnection struct {
  RefCounted
}

func (me *ENetConnection) BaseClass() string {
  return "ENetConnection"
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

func  (me *ENetConnection) CreateHostBound(bind_address String, bind_port int, max_peers int, max_channels int, in_bandwidth int, out_bandwidth int, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_host_bound")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1515002313) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bind_address.AsCTypePtr()), gdc.ConstTypePtr(&bind_port), gdc.ConstTypePtr(&max_peers), gdc.ConstTypePtr(&max_channels), gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) CreateHost(max_peers int, max_channels int, in_bandwidth int, out_bandwidth int, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 117198950) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_peers), gdc.ConstTypePtr(&max_channels), gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) Destroy()  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("destroy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) ConnectToHost(address String, port int, channels int, data int, ) ENetPacketPeer {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2171300490) // FIXME: should cache?
  var ret ENetPacketPeer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(address.AsCTypePtr()), gdc.ConstTypePtr(&port), gdc.ConstTypePtr(&channels), gdc.ConstTypePtr(&data), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) Service(timeout int, ) Array {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("service")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2402345344) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) Flush()  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("flush")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) BandwidthLimit(in_bandwidth int, out_bandwidth int, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bandwidth_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2302169788) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&in_bandwidth), gdc.ConstTypePtr(&out_bandwidth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) ChannelLimit(limit int, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("channel_limit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) Broadcast(channel int, packet PackedByteArray, flags int, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("broadcast")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2772371345) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel), gdc.ConstTypePtr(packet.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) Compress(mode ENetConnectionCompressionMode, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2660215187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) DtlsServerSetup(server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dtls_server_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262296096) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(server_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) DtlsClientSetup(hostname String, client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dtls_client_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1966198364) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(hostname.AsCTypePtr()), gdc.ConstTypePtr(client_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) RefuseNewConnections(refuse bool, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("refuse_new_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refuse), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetConnection) PopStatistic(statistic ENetConnectionHostStatistic, ) float32 {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pop_statistic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2166904170) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&statistic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) GetMaxChannels() int {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_channels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) GetLocalPort() int {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) GetPeers() ENetPacketPeer {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret ENetPacketPeer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetConnection) SocketSend(destination_address String, destination_port int, packet PackedByteArray, )  {
  classNameV := StringNameFromStr("ENetConnection")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("socket_send")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1100646812) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(destination_address.AsCTypePtr()), gdc.ConstTypePtr(&destination_port), gdc.ConstTypePtr(packet.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
