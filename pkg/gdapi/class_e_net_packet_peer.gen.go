// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ENetPacketPeer struct {
  PacketPeer
}

func (me *ENetPacketPeer) BaseClass() string {
  return "ENetPacketPeer"
}



// Constants

var (
  ENetPacketPeerPacketLossScale = "65536" // TODO: construct correctly
  ENetPacketPeerPacketThrottleScale = "32" // TODO: construct correctly
  ENetPacketPeerFlagReliable = "1" // TODO: construct correctly
  ENetPacketPeerFlagUnsequenced = "2" // TODO: construct correctly
  ENetPacketPeerFlagUnreliableFragment = "8" // TODO: construct correctly
)

// Enums

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

func (me *ENetPacketPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ENetPacketPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ENetPacketPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ENetPacketPeer) PeerDisconnect(data int, )  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("peer_disconnect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) PeerDisconnectLater(data int, )  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("peer_disconnect_later")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) PeerDisconnectNow(data int, )  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("peer_disconnect_now")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) Ping()  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) PingInterval(ping_interval int, )  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ping_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ping_interval), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) Reset()  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) Send(channel int, packet PackedByteArray, flags int, ) Error {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 120522849) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel), gdc.ConstTypePtr(packet.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetPacketPeer) ThrottleConfigure(interval int, acceleration int, deceleration int, )  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("throttle_configure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1649997291) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval), gdc.ConstTypePtr(&acceleration), gdc.ConstTypePtr(&deceleration), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) SetTimeout(timeout int, timeout_min int, timeout_max int, )  {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1649997291) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout), gdc.ConstTypePtr(&timeout_min), gdc.ConstTypePtr(&timeout_max), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ENetPacketPeer) GetRemoteAddress() String {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_remote_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetPacketPeer) GetRemotePort() int {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_remote_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetPacketPeer) GetStatistic(statistic ENetPacketPeerPeerStatistic, ) float32 {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_statistic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1642578323) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&statistic), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetPacketPeer) GetState() ENetPacketPeerPeerState {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711068532) // FIXME: should cache?
  var ret ENetPacketPeerPeerState
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetPacketPeer) GetChannels() int {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_channels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ENetPacketPeer) IsActive() bool {
  classNameV := StringNameFromStr("ENetPacketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
