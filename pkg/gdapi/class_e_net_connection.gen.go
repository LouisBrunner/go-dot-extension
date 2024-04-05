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

type ptrsForENetConnectionList struct {
  fnCreateHostBound gdc.MethodBindPtr
  fnCreateHost gdc.MethodBindPtr
  fnDestroy gdc.MethodBindPtr
  fnConnectToHost gdc.MethodBindPtr
  fnService gdc.MethodBindPtr
  fnFlush gdc.MethodBindPtr
  fnBandwidthLimit gdc.MethodBindPtr
  fnChannelLimit gdc.MethodBindPtr
  fnBroadcast gdc.MethodBindPtr
  fnCompress gdc.MethodBindPtr
  fnDtlsServerSetup gdc.MethodBindPtr
  fnDtlsClientSetup gdc.MethodBindPtr
  fnRefuseNewConnections gdc.MethodBindPtr
  fnPopStatistic gdc.MethodBindPtr
  fnGetMaxChannels gdc.MethodBindPtr
  fnGetLocalPort gdc.MethodBindPtr
  fnGetPeers gdc.MethodBindPtr
  fnSocketSend gdc.MethodBindPtr
}

var ptrsForENetConnection ptrsForENetConnectionList

func initENetConnectionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ENetConnection")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_host_bound")
    defer methodName.Destroy()
    ptrsForENetConnection.fnCreateHostBound = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1515002313))
  }
  {
    methodName := StringNameFromStr("create_host")
    defer methodName.Destroy()
    ptrsForENetConnection.fnCreateHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 117198950))
  }
  {
    methodName := StringNameFromStr("destroy")
    defer methodName.Destroy()
    ptrsForENetConnection.fnDestroy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("connect_to_host")
    defer methodName.Destroy()
    ptrsForENetConnection.fnConnectToHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2171300490))
  }
  {
    methodName := StringNameFromStr("service")
    defer methodName.Destroy()
    ptrsForENetConnection.fnService = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2402345344))
  }
  {
    methodName := StringNameFromStr("flush")
    defer methodName.Destroy()
    ptrsForENetConnection.fnFlush = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bandwidth_limit")
    defer methodName.Destroy()
    ptrsForENetConnection.fnBandwidthLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2302169788))
  }
  {
    methodName := StringNameFromStr("channel_limit")
    defer methodName.Destroy()
    ptrsForENetConnection.fnChannelLimit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("broadcast")
    defer methodName.Destroy()
    ptrsForENetConnection.fnBroadcast = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2772371345))
  }
  {
    methodName := StringNameFromStr("compress")
    defer methodName.Destroy()
    ptrsForENetConnection.fnCompress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2660215187))
  }
  {
    methodName := StringNameFromStr("dtls_server_setup")
    defer methodName.Destroy()
    ptrsForENetConnection.fnDtlsServerSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1262296096))
  }
  {
    methodName := StringNameFromStr("dtls_client_setup")
    defer methodName.Destroy()
    ptrsForENetConnection.fnDtlsClientSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1966198364))
  }
  {
    methodName := StringNameFromStr("refuse_new_connections")
    defer methodName.Destroy()
    ptrsForENetConnection.fnRefuseNewConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("pop_statistic")
    defer methodName.Destroy()
    ptrsForENetConnection.fnPopStatistic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2166904170))
  }
  {
    methodName := StringNameFromStr("get_max_channels")
    defer methodName.Destroy()
    ptrsForENetConnection.fnGetMaxChannels = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_local_port")
    defer methodName.Destroy()
    ptrsForENetConnection.fnGetLocalPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_peers")
    defer methodName.Destroy()
    ptrsForENetConnection.fnGetPeers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("socket_send")
    defer methodName.Destroy()
    ptrsForENetConnection.fnSocketSend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1100646812))
  }
}

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
  cargs := []gdc.ConstTypePtr{bind_address.AsCTypePtr(), gdc.ConstTypePtr(&bind_port) , gdc.ConstTypePtr(&max_peers) , gdc.ConstTypePtr(&max_channels) , gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&bind_port)
  pinner.Pin(&max_peers)
  pinner.Pin(&max_channels)
  pinner.Pin(&in_bandwidth)
  pinner.Pin(&out_bandwidth)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnCreateHostBound), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) CreateHost(max_peers int64, max_channels int64, in_bandwidth int64, out_bandwidth int64, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_peers) , gdc.ConstTypePtr(&max_channels) , gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&max_peers)
  pinner.Pin(&max_channels)
  pinner.Pin(&in_bandwidth)
  pinner.Pin(&out_bandwidth)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnCreateHost), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) Destroy()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnDestroy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) ConnectToHost(address String, port int64, channels int64, data int64, ) ENetPacketPeer {
  cargs := []gdc.ConstTypePtr{address.AsCTypePtr(), gdc.ConstTypePtr(&port) , gdc.ConstTypePtr(&channels) , gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewENetPacketPeer()
  pinner.Pin(&port)
  pinner.Pin(&channels)
  pinner.Pin(&data)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnConnectToHost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetConnection) Service(timeout int64, ) Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&timeout)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnService), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetConnection) Flush()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnFlush), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) BandwidthLimit(in_bandwidth int64, out_bandwidth int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&in_bandwidth) , gdc.ConstTypePtr(&out_bandwidth) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnBandwidthLimit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) ChannelLimit(limit int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnChannelLimit), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) Broadcast(channel int64, packet PackedByteArray, flags int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , packet.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnBroadcast), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) Compress(mode ENetConnectionCompressionMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnCompress), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) DtlsServerSetup(server_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{server_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnDtlsServerSetup), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) DtlsClientSetup(hostname String, client_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{hostname.AsCTypePtr(), client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnDtlsClientSetup), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetConnection) RefuseNewConnections(refuse bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&refuse) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnRefuseNewConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetConnection) PopStatistic(statistic ENetConnectionHostStatistic, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&statistic) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&statistic)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnPopStatistic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetConnection) GetMaxChannels() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnGetMaxChannels), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetConnection) GetLocalPort() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnGetLocalPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetConnection) GetPeers() []ENetPacketPeer {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnGetPeers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[ENetPacketPeer](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *ENetConnection) SocketSend(destination_address String, destination_port int64, packet PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{destination_address.AsCTypePtr(), gdc.ConstTypePtr(&destination_port) , packet.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetConnection.fnSocketSend), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
