// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  ENetPacketPeerPACKET_LOSS_SCALE = 65536
  ENetPacketPeerPACKET_THROTTLE_SCALE = 32
  ENetPacketPeerFLAG_RELIABLE = 1
  ENetPacketPeerFLAG_UNSEQUENCED = 2
  ENetPacketPeerFLAG_UNRELIABLE_FRAGMENT = 8
)

type ENetPacketPeerPeerState int
const (
  ENetPacketPeerStateDisconnected ENetPacketPeerPeerState = 0
  ENetPacketPeerStateConnecting ENetPacketPeerPeerState = 1
  ENetPacketPeerStateAcknowledgingConnect ENetPacketPeerPeerState = 2
  ENetPacketPeerStateConnectionPending ENetPacketPeerPeerState = 3
  ENetPacketPeerStateConnectionSucceeded ENetPacketPeerPeerState = 4
  ENetPacketPeerStateConnected ENetPacketPeerPeerState = 5
  ENetPacketPeerStateDisconnectLater ENetPacketPeerPeerState = 6
  ENetPacketPeerStateDisconnecting ENetPacketPeerPeerState = 7
  ENetPacketPeerStateAcknowledgingDisconnect ENetPacketPeerPeerState = 8
  ENetPacketPeerStateZombie ENetPacketPeerPeerState = 9
)

type ENetPacketPeerPeerStatistic int
const (
  ENetPacketPeerPeerPacketLoss ENetPacketPeerPeerStatistic = 0
  ENetPacketPeerPeerPacketLossVariance ENetPacketPeerPeerStatistic = 1
  ENetPacketPeerPeerPacketLossEpoch ENetPacketPeerPeerStatistic = 2
  ENetPacketPeerPeerRoundTripTime ENetPacketPeerPeerStatistic = 3
  ENetPacketPeerPeerRoundTripTimeVariance ENetPacketPeerPeerStatistic = 4
  ENetPacketPeerPeerLastRoundTripTime ENetPacketPeerPeerStatistic = 5
  ENetPacketPeerPeerLastRoundTripTimeVariance ENetPacketPeerPeerStatistic = 6
  ENetPacketPeerPeerPacketThrottle ENetPacketPeerPeerStatistic = 7
  ENetPacketPeerPeerPacketThrottleLimit ENetPacketPeerPeerStatistic = 8
  ENetPacketPeerPeerPacketThrottleCounter ENetPacketPeerPeerStatistic = 9
  ENetPacketPeerPeerPacketThrottleEpoch ENetPacketPeerPeerStatistic = 10
  ENetPacketPeerPeerPacketThrottleAcceleration ENetPacketPeerPeerStatistic = 11
  ENetPacketPeerPeerPacketThrottleDeceleration ENetPacketPeerPeerStatistic = 12
  ENetPacketPeerPeerPacketThrottleInterval ENetPacketPeerPeerStatistic = 13
)
