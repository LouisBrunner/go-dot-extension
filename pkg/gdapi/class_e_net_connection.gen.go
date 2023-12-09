// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetConnection struct {
  obj gdc.ObjectPtr
}

func (me *ENetConnection) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *ENetConnection) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ENetConnection) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ENetConnection) CreateHostBound(bind_address String, bind_port int, max_peers int, max_channels int, in_bandwidth int, out_bandwidth int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) CreateHost(max_peers int, max_channels int, in_bandwidth int, out_bandwidth int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) Destroy()  {
  panic("TODO: implement")
}

func  (me *ENetConnection) ConnectToHost(address String, port int, channels int, data int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) Service(timeout int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) Flush()  {
  panic("TODO: implement")
}

func  (me *ENetConnection) BandwidthLimit(in_bandwidth int, out_bandwidth int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) ChannelLimit(limit int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) Broadcast(channel int, packet PackedByteArray, flags int, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) Compress(mode ENetConnectionCompressionMode, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) DtlsServerSetup(server_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) DtlsClientSetup(hostname String, client_options TLSOptions, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) RefuseNewConnections(refuse bool, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) PopStatistic(statistic ENetConnectionHostStatistic, )  {
  panic("TODO: implement")
}

func  (me *ENetConnection) GetMaxChannels()  {
  panic("TODO: implement")
}

func  (me *ENetConnection) GetLocalPort()  {
  panic("TODO: implement")
}

func  (me *ENetConnection) GetPeers()  {
  panic("TODO: implement")
}

func  (me *ENetConnection) SocketSend(destination_address String, destination_port int, packet PackedByteArray, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
