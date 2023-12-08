// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ENetPacketPeer struct {
  obj gdc.ObjectPtr
}

func (me *ENetPacketPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ENetPacketPeer) BaseClass() string {
  return "ENetPacketPeer"
}

// TODO: needed?
// const (
// // )

var (
  ENetPacketPeerPacketLossScale = "65536" // TODO: construct correctly
  ENetPacketPeerPacketThrottleScale = "32" // TODO: construct correctly
  ENetPacketPeerFlagReliable = "1" // TODO: construct correctly
  ENetPacketPeerFlagUnsequenced = "2" // TODO: construct correctly
  ENetPacketPeerFlagUnreliableFragment = "8" // TODO: construct correctly
)

type ENetPacketPeerPeerState int
const (
  ENetPacketPeerPeerStateStateDisconnected ENetPacketPeerPeerState = 0
  ENetPacketPeerPeerStateStateConnecting ENetPacketPeerPeerState = 1
  ENetPacketPeerPeerStateStateAcknowledgingConnect ENetPacketPeerPeerState = 2
  ENetPacketPeerPeerStateStateConnectionPending ENetPacketPeerPeerState = 3
  ENetPacketPeerPeerStateStateConnectionSucceeded ENetPacketPeerPeerState = 4
  ENetPacketPeerPeerStateStateConnected ENetPacketPeerPeerState = 5
  ENetPacketPeerPeerStateStateDisconnectLater ENetPacketPeerPeerState = 6
  ENetPacketPeerPeerStateStateDisconnecting ENetPacketPeerPeerState = 7
  ENetPacketPeerPeerStateStateAcknowledgingDisconnect ENetPacketPeerPeerState = 8
  ENetPacketPeerPeerStateStateZombie ENetPacketPeerPeerState = 9
)

type ENetPacketPeerPeerStatistic int
const (
  ENetPacketPeerPeerStatisticPeerPacketLoss ENetPacketPeerPeerStatistic = 0
  ENetPacketPeerPeerStatisticPeerPacketLossVariance ENetPacketPeerPeerStatistic = 1
  ENetPacketPeerPeerStatisticPeerPacketLossEpoch ENetPacketPeerPeerStatistic = 2
  ENetPacketPeerPeerStatisticPeerRoundTripTime ENetPacketPeerPeerStatistic = 3
  ENetPacketPeerPeerStatisticPeerRoundTripTimeVariance ENetPacketPeerPeerStatistic = 4
  ENetPacketPeerPeerStatisticPeerLastRoundTripTime ENetPacketPeerPeerStatistic = 5
  ENetPacketPeerPeerStatisticPeerLastRoundTripTimeVariance ENetPacketPeerPeerStatistic = 6
  ENetPacketPeerPeerStatisticPeerPacketThrottle ENetPacketPeerPeerStatistic = 7
  ENetPacketPeerPeerStatisticPeerPacketThrottleLimit ENetPacketPeerPeerStatistic = 8
  ENetPacketPeerPeerStatisticPeerPacketThrottleCounter ENetPacketPeerPeerStatistic = 9
  ENetPacketPeerPeerStatisticPeerPacketThrottleEpoch ENetPacketPeerPeerStatistic = 10
  ENetPacketPeerPeerStatisticPeerPacketThrottleAcceleration ENetPacketPeerPeerStatistic = 11
  ENetPacketPeerPeerStatisticPeerPacketThrottleDeceleration ENetPacketPeerPeerStatistic = 12
  ENetPacketPeerPeerStatisticPeerPacketThrottleInterval ENetPacketPeerPeerStatistic = 13
)

func  (me *ENetPacketPeer) PeerDisconnect(data int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) PeerDisconnectLater(data int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) PeerDisconnectNow(data int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) Ping() { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) PingInterval(ping_interval int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) Reset() { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) Send(channel int, packet PackedByteArray, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) ThrottleConfigure(interval int, acceleration int, deceleration int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) SetTimeout(timeout int, timeout_min int, timeout_max int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) GetRemoteAddress() { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) GetRemotePort() { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) GetStatistic(statistic ENetPacketPeerPeerStatistic, ) { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) GetState() { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) GetChannels() { // TODO: return value
  // TODO: implement
}

func  (me *ENetPacketPeer) IsActive() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
