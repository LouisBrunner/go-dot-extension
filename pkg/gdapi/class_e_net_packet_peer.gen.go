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

type ptrsForENetPacketPeerList struct {
  fnPeerDisconnect gdc.MethodBindPtr
  fnPeerDisconnectLater gdc.MethodBindPtr
  fnPeerDisconnectNow gdc.MethodBindPtr
  fnPing gdc.MethodBindPtr
  fnPingInterval gdc.MethodBindPtr
  fnReset gdc.MethodBindPtr
  fnSend gdc.MethodBindPtr
  fnThrottleConfigure gdc.MethodBindPtr
  fnSetTimeout gdc.MethodBindPtr
  fnGetRemoteAddress gdc.MethodBindPtr
  fnGetRemotePort gdc.MethodBindPtr
  fnGetStatistic gdc.MethodBindPtr
  fnGetState gdc.MethodBindPtr
  fnGetChannels gdc.MethodBindPtr
  fnIsActive gdc.MethodBindPtr
}

var ptrsForENetPacketPeer ptrsForENetPacketPeerList

func initENetPacketPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ENetPacketPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("peer_disconnect")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnPeerDisconnect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
  }
  {
    methodName := StringNameFromStr("peer_disconnect_later")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnPeerDisconnectLater = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
  }
  {
    methodName := StringNameFromStr("peer_disconnect_now")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnPeerDisconnectNow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
  }
  {
    methodName := StringNameFromStr("ping")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnPing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("ping_interval")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnPingInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("reset")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("send")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnSend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 120522849))
  }
  {
    methodName := StringNameFromStr("throttle_configure")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnThrottleConfigure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1649997291))
  }
  {
    methodName := StringNameFromStr("set_timeout")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnSetTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1649997291))
  }
  {
    methodName := StringNameFromStr("get_remote_address")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnGetRemoteAddress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_remote_port")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnGetRemotePort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_statistic")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnGetStatistic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1642578323))
  }
  {
    methodName := StringNameFromStr("get_state")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnGetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711068532))
  }
  {
    methodName := StringNameFromStr("get_channels")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnGetChannels = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("is_active")
    defer methodName.Destroy()
    ptrsForENetPacketPeer.fnIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type ENetPacketPeer struct {
  PacketPeer
}

func (me *ENetPacketPeer) BaseClass() string {
  return "ENetPacketPeer"
}

func NewENetPacketPeer() *ENetPacketPeer {
  str := StringNameFromStr("ENetPacketPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ENetPacketPeer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *ENetPacketPeer) PeerDisconnect(data int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnPeerDisconnect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) PeerDisconnectLater(data int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnPeerDisconnectLater), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) PeerDisconnectNow(data int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnPeerDisconnectNow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) Ping()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnPing), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) PingInterval(ping_interval int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ping_interval) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnPingInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) Reset()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnReset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) Send(channel int64, packet PackedByteArray, flags int64, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , packet.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&channel)
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnSend), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetPacketPeer) ThrottleConfigure(interval int64, acceleration int64, deceleration int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval) , gdc.ConstTypePtr(&acceleration) , gdc.ConstTypePtr(&deceleration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnThrottleConfigure), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) SetTimeout(timeout int64, timeout_min int64, timeout_max int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , gdc.ConstTypePtr(&timeout_min) , gdc.ConstTypePtr(&timeout_max) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnSetTimeout), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ENetPacketPeer) GetRemoteAddress() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnGetRemoteAddress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ENetPacketPeer) GetRemotePort() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnGetRemotePort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetPacketPeer) GetStatistic(statistic ENetPacketPeerPeerStatistic, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&statistic) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&statistic)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnGetStatistic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetPacketPeer) GetState() ENetPacketPeerPeerState {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ENetPacketPeerPeerState

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnGetState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ENetPacketPeer) GetChannels() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnGetChannels), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ENetPacketPeer) IsActive() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForENetPacketPeer.fnIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
