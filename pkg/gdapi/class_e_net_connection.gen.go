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

func  (me *ENetConnection) CreateHostBound(bind_address String, bind_port int, max_peers int, max_channels int, in_bandwidth int, out_bandwidth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) CreateHost(max_peers int, max_channels int, in_bandwidth int, out_bandwidth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) Destroy() { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) ConnectToHost(address String, port int, channels int, data int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) Service(timeout int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) Flush() { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) BandwidthLimit(in_bandwidth int, out_bandwidth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) ChannelLimit(limit int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) Broadcast(channel int, packet PackedByteArray, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) Compress(mode ENetConnectionCompressionMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) DtlsServerSetup(server_options TLSOptions, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) DtlsClientSetup(hostname String, client_options TLSOptions, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) RefuseNewConnections(refuse bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) PopStatistic(statistic ENetConnectionHostStatistic, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) GetMaxChannels() { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) GetLocalPort() { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) GetPeers() { // TODO: return value
  // TODO: implement
}

func  (me *ENetConnection) SocketSend(destination_address String, destination_port int, packet PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
