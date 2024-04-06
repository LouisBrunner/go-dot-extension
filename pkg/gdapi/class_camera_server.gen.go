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

type ptrsForCameraServerList struct {
	fnGetFeed      gdc.MethodBindPtr
	fnGetFeedCount gdc.MethodBindPtr
	fnFeeds        gdc.MethodBindPtr
	fnAddFeed      gdc.MethodBindPtr
	fnRemoveFeed   gdc.MethodBindPtr
}

var ptrsForCameraServer ptrsForCameraServerList

func initCameraServerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CameraServer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_feed")
		defer methodName.Destroy()
		ptrsForCameraServer.fnGetFeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 361927068))
	}
	{
		methodName := StringNameFromStr("get_feed_count")
		defer methodName.Destroy()
		ptrsForCameraServer.fnGetFeedCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("feeds")
		defer methodName.Destroy()
		ptrsForCameraServer.fnFeeds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("add_feed")
		defer methodName.Destroy()
		ptrsForCameraServer.fnAddFeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3204782488))
	}
	{
		methodName := StringNameFromStr("remove_feed")
		defer methodName.Destroy()
		ptrsForCameraServer.fnRemoveFeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3204782488))
	}
}

type CameraServer struct {
	Object
}

func (me *CameraServer) BaseClass() string {
	return "CameraServer"
}

func NewCameraServer() *CameraServer {
	str := StringNameFromStr("CameraServer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CameraServer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CameraServerFeedImage int

const (
	CameraServerFeedImageFeedRgbaImage  CameraServerFeedImage = 0
	CameraServerFeedImageFeedYcbcrImage CameraServerFeedImage = 0
	CameraServerFeedImageFeedYImage     CameraServerFeedImage = 0
	CameraServerFeedImageFeedCbcrImage  CameraServerFeedImage = 1
)

func (me *CameraServer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CameraServer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CameraServer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CameraServer) GetFeed(index int64) CameraFeed {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCameraFeed()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraServer.fnGetFeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CameraServer) GetFeedCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraServer.fnGetFeedCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CameraServer) Feeds() []CameraFeed {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraServer.fnFeeds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[CameraFeed](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *CameraServer) AddFeed(feed CameraFeed) {
	cargs := []gdc.ConstTypePtr{feed.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraServer.fnAddFeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraServer) RemoveFeed(feed CameraFeed) {
	cargs := []gdc.ConstTypePtr{feed.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraServer.fnRemoveFeed), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type CameraServerCameraFeedAddedSignalFn func(id int)

func (me *CameraServer) ConnectCameraFeedAdded(subs SignalSubscribers, fn CameraServerCameraFeedAddedSignalFn) {
	sig := StringNameFromStr("camera_feed_added")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CameraServer) DisconnectCameraFeedAdded(subs SignalSubscribers, fn CameraServerCameraFeedAddedSignalFn) {
	sig := StringNameFromStr("camera_feed_added")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type CameraServerCameraFeedRemovedSignalFn func(id int)

func (me *CameraServer) ConnectCameraFeedRemoved(subs SignalSubscribers, fn CameraServerCameraFeedRemovedSignalFn) {
	sig := StringNameFromStr("camera_feed_removed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *CameraServer) DisconnectCameraFeedRemoved(subs SignalSubscribers, fn CameraServerCameraFeedRemovedSignalFn) {
	sig := StringNameFromStr("camera_feed_removed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
