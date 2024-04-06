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

type ptrsForGraphEditList struct {
	fnXIsInInputHotzone              gdc.MethodBindPtr
	fnXIsInOutputHotzone             gdc.MethodBindPtr
	fnXGetConnectionLine             gdc.MethodBindPtr
	fnXIsNodeHoverValid              gdc.MethodBindPtr
	fnConnectNode                    gdc.MethodBindPtr
	fnIsNodeConnected                gdc.MethodBindPtr
	fnDisconnectNode                 gdc.MethodBindPtr
	fnSetConnectionActivity          gdc.MethodBindPtr
	fnGetConnectionList              gdc.MethodBindPtr
	fnClearConnections               gdc.MethodBindPtr
	fnForceConnectionDragEnd         gdc.MethodBindPtr
	fnGetScrollOffset                gdc.MethodBindPtr
	fnSetScrollOffset                gdc.MethodBindPtr
	fnAddValidRightDisconnectType    gdc.MethodBindPtr
	fnRemoveValidRightDisconnectType gdc.MethodBindPtr
	fnAddValidLeftDisconnectType     gdc.MethodBindPtr
	fnRemoveValidLeftDisconnectType  gdc.MethodBindPtr
	fnAddValidConnectionType         gdc.MethodBindPtr
	fnRemoveValidConnectionType      gdc.MethodBindPtr
	fnIsValidConnectionType          gdc.MethodBindPtr
	fnGetConnectionLine              gdc.MethodBindPtr
	fnSetPanningScheme               gdc.MethodBindPtr
	fnGetPanningScheme               gdc.MethodBindPtr
	fnSetZoom                        gdc.MethodBindPtr
	fnGetZoom                        gdc.MethodBindPtr
	fnSetZoomMin                     gdc.MethodBindPtr
	fnGetZoomMin                     gdc.MethodBindPtr
	fnSetZoomMax                     gdc.MethodBindPtr
	fnGetZoomMax                     gdc.MethodBindPtr
	fnSetZoomStep                    gdc.MethodBindPtr
	fnGetZoomStep                    gdc.MethodBindPtr
	fnSetShowGrid                    gdc.MethodBindPtr
	fnIsShowingGrid                  gdc.MethodBindPtr
	fnSetSnappingEnabled             gdc.MethodBindPtr
	fnIsSnappingEnabled              gdc.MethodBindPtr
	fnSetSnappingDistance            gdc.MethodBindPtr
	fnGetSnappingDistance            gdc.MethodBindPtr
	fnSetConnectionLinesCurvature    gdc.MethodBindPtr
	fnGetConnectionLinesCurvature    gdc.MethodBindPtr
	fnSetConnectionLinesThickness    gdc.MethodBindPtr
	fnGetConnectionLinesThickness    gdc.MethodBindPtr
	fnSetConnectionLinesAntialiased  gdc.MethodBindPtr
	fnIsConnectionLinesAntialiased   gdc.MethodBindPtr
	fnSetMinimapSize                 gdc.MethodBindPtr
	fnGetMinimapSize                 gdc.MethodBindPtr
	fnSetMinimapOpacity              gdc.MethodBindPtr
	fnGetMinimapOpacity              gdc.MethodBindPtr
	fnSetMinimapEnabled              gdc.MethodBindPtr
	fnIsMinimapEnabled               gdc.MethodBindPtr
	fnSetShowMenu                    gdc.MethodBindPtr
	fnIsShowingMenu                  gdc.MethodBindPtr
	fnSetShowZoomLabel               gdc.MethodBindPtr
	fnIsShowingZoomLabel             gdc.MethodBindPtr
	fnSetShowGridButtons             gdc.MethodBindPtr
	fnIsShowingGridButtons           gdc.MethodBindPtr
	fnSetShowZoomButtons             gdc.MethodBindPtr
	fnIsShowingZoomButtons           gdc.MethodBindPtr
	fnSetShowMinimapButton           gdc.MethodBindPtr
	fnIsShowingMinimapButton         gdc.MethodBindPtr
	fnSetShowArrangeButton           gdc.MethodBindPtr
	fnIsShowingArrangeButton         gdc.MethodBindPtr
	fnSetRightDisconnects            gdc.MethodBindPtr
	fnIsRightDisconnectsEnabled      gdc.MethodBindPtr
	fnGetMenuHbox                    gdc.MethodBindPtr
	fnArrangeNodes                   gdc.MethodBindPtr
	fnSetSelected                    gdc.MethodBindPtr
}

var ptrsForGraphEdit ptrsForGraphEditList

func initGraphEditPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GraphEdit")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("connect_node")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnConnectNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 195065850))
	}
	{
		methodName := StringNameFromStr("is_node_connected")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsNodeConnected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4216241294))
	}
	{
		methodName := StringNameFromStr("disconnect_node")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnDisconnectNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1933654315))
	}
	{
		methodName := StringNameFromStr("set_connection_activity")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetConnectionActivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1141899943))
	}
	{
		methodName := StringNameFromStr("get_connection_list")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetConnectionList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("clear_connections")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnClearConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("force_connection_drag_end")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnForceConnectionDragEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_scroll_offset")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_scroll_offset")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("add_valid_right_disconnect_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnAddValidRightDisconnectType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_valid_right_disconnect_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnRemoveValidRightDisconnectType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("add_valid_left_disconnect_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnAddValidLeftDisconnectType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_valid_left_disconnect_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnRemoveValidLeftDisconnectType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("add_valid_connection_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnAddValidConnectionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_valid_connection_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnRemoveValidConnectionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("is_valid_connection_type")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsValidConnectionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("get_connection_line")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetConnectionLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1562168077))
	}
	{
		methodName := StringNameFromStr("set_panning_scheme")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetPanningScheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 18893313))
	}
	{
		methodName := StringNameFromStr("get_panning_scheme")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetPanningScheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 549924446))
	}
	{
		methodName := StringNameFromStr("set_zoom")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetZoom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_zoom")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetZoom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_zoom_min")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetZoomMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_zoom_min")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetZoomMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_zoom_max")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetZoomMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_zoom_max")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetZoomMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_zoom_step")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetZoomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_zoom_step")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetZoomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_show_grid")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowGrid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_grid")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingGrid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_snapping_enabled")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetSnappingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_snapping_enabled")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsSnappingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_snapping_distance")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetSnappingDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_snapping_distance")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetSnappingDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_connection_lines_curvature")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetConnectionLinesCurvature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_connection_lines_curvature")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetConnectionLinesCurvature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_connection_lines_thickness")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetConnectionLinesThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_connection_lines_thickness")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetConnectionLinesThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_connection_lines_antialiased")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetConnectionLinesAntialiased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_connection_lines_antialiased")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsConnectionLinesAntialiased = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_minimap_size")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetMinimapSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_minimap_size")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetMinimapSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_minimap_opacity")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetMinimapOpacity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_minimap_opacity")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetMinimapOpacity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_minimap_enabled")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetMinimapEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_minimap_enabled")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsMinimapEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_show_menu")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_menu")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_show_zoom_label")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowZoomLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_zoom_label")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingZoomLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_show_grid_buttons")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowGridButtons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_grid_buttons")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingGridButtons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_show_zoom_buttons")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowZoomButtons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_zoom_buttons")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingZoomButtons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_show_minimap_button")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowMinimapButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_minimap_button")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingMinimapButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_show_arrange_button")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetShowArrangeButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_arrange_button")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsShowingArrangeButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_right_disconnects")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetRightDisconnects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_right_disconnects_enabled")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnIsRightDisconnectsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_menu_hbox")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnGetMenuHbox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3590609951))
	}
	{
		methodName := StringNameFromStr("arrange_nodes")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnArrangeNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_selected")
		defer methodName.Destroy()
		ptrsForGraphEdit.fnSetSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}

}

type GraphEdit struct {
	Control
}

func (me *GraphEdit) BaseClass() string {
	return "GraphEdit"
}

func NewGraphEdit() *GraphEdit {
	str := StringNameFromStr("GraphEdit") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GraphEdit{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GraphEditPanningScheme int

const (
	GraphEditPanningSchemeScrollZooms GraphEditPanningScheme = 0
	GraphEditPanningSchemeScrollPans  GraphEditPanningScheme = 1
)

func (me *GraphEdit) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GraphEdit) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GraphEdit) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GraphEdit) ConnectNode(from_node StringName, from_port int64, to_node StringName, to_port int64) Error {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port), to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&from_port)
	pinner.Pin(&to_port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnConnectNode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GraphEdit) IsNodeConnected(from_node StringName, from_port int64, to_node StringName, to_port int64) bool {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port), to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&from_port)
	pinner.Pin(&to_port)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsNodeConnected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) DisconnectNode(from_node StringName, from_port int64, to_node StringName, to_port int64) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port), to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnDisconnectNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) SetConnectionActivity(from_node StringName, from_port int64, to_node StringName, to_port int64, amount float64) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&from_port), to_node.AsCTypePtr(), gdc.ConstTypePtr(&to_port), gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetConnectionActivity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetConnectionList() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetConnectionList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *GraphEdit) ClearConnections() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnClearConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) ForceConnectionDragEnd() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnForceConnectionDragEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetScrollOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetScrollOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphEdit) SetScrollOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetScrollOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) AddValidRightDisconnectType(type_ int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnAddValidRightDisconnectType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) RemoveValidRightDisconnectType(type_ int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnRemoveValidRightDisconnectType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) AddValidLeftDisconnectType(type_ int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnAddValidLeftDisconnectType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) RemoveValidLeftDisconnectType(type_ int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnRemoveValidLeftDisconnectType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) AddValidConnectionType(from_type int64, to_type int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type), gdc.ConstTypePtr(&to_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnAddValidConnectionType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) RemoveValidConnectionType(from_type int64, to_type int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type), gdc.ConstTypePtr(&to_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnRemoveValidConnectionType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsValidConnectionType(from_type int64, to_type int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_type), gdc.ConstTypePtr(&to_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&from_type)
	pinner.Pin(&to_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsValidConnectionType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) GetConnectionLine(from_node Vector2, to_node Vector2) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), to_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetConnectionLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphEdit) SetPanningScheme(scheme GraphEditPanningScheme) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scheme)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetPanningScheme), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetPanningScheme() GraphEditPanningScheme {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GraphEditPanningScheme

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetPanningScheme), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GraphEdit) SetZoom(zoom float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetZoom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetZoom() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetZoom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetZoomMin(zoom_min float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_min)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetZoomMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetZoomMin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetZoomMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetZoomMax(zoom_max float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_max)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetZoomMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetZoomMax() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetZoomMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetZoomStep(zoom_step float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&zoom_step)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetZoomStep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetZoomStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetZoomStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowGrid(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowGrid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingGrid() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingGrid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetSnappingEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetSnappingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsSnappingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsSnappingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetSnappingDistance(pixels int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetSnappingDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetSnappingDistance() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetSnappingDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetConnectionLinesCurvature(curvature float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curvature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetConnectionLinesCurvature), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetConnectionLinesCurvature() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetConnectionLinesCurvature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetConnectionLinesThickness(pixels float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetConnectionLinesThickness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetConnectionLinesThickness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetConnectionLinesThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetConnectionLinesAntialiased(pixels bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetConnectionLinesAntialiased), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsConnectionLinesAntialiased() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsConnectionLinesAntialiased), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetMinimapSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetMinimapSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetMinimapSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetMinimapSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphEdit) SetMinimapOpacity(opacity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&opacity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetMinimapOpacity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) GetMinimapOpacity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetMinimapOpacity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetMinimapEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetMinimapEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsMinimapEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsMinimapEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowMenu(hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowMenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingMenu() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowZoomLabel(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowZoomLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingZoomLabel() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingZoomLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowGridButtons(hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowGridButtons), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingGridButtons() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingGridButtons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowZoomButtons(hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowZoomButtons), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingZoomButtons() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingZoomButtons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowMinimapButton(hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowMinimapButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingMinimapButton() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingMinimapButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetShowArrangeButton(hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetShowArrangeButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsShowingArrangeButton() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsShowingArrangeButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) SetRightDisconnects(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetRightDisconnects), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) IsRightDisconnectsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnIsRightDisconnectsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphEdit) GetMenuHbox() HBoxContainer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewHBoxContainer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnGetMenuHbox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphEdit) ArrangeNodes() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnArrangeNodes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphEdit) SetSelected(node Node) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphEdit.fnSetSelected), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphEditConnectionRequestSignalFn func(from_node StringName, from_port int, to_node StringName, to_port int)

func (me *GraphEdit) ConnectConnectionRequest(subs SignalSubscribers, fn GraphEditConnectionRequestSignalFn) {
	sig := StringNameFromStr("connection_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionRequest(subs SignalSubscribers, fn GraphEditConnectionRequestSignalFn) {
	sig := StringNameFromStr("connection_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditDisconnectionRequestSignalFn func(from_node StringName, from_port int, to_node StringName, to_port int)

func (me *GraphEdit) ConnectDisconnectionRequest(subs SignalSubscribers, fn GraphEditDisconnectionRequestSignalFn) {
	sig := StringNameFromStr("disconnection_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectDisconnectionRequest(subs SignalSubscribers, fn GraphEditDisconnectionRequestSignalFn) {
	sig := StringNameFromStr("disconnection_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionToEmptySignalFn func(from_node StringName, from_port int, release_position Vector2)

func (me *GraphEdit) ConnectConnectionToEmpty(subs SignalSubscribers, fn GraphEditConnectionToEmptySignalFn) {
	sig := StringNameFromStr("connection_to_empty")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionToEmpty(subs SignalSubscribers, fn GraphEditConnectionToEmptySignalFn) {
	sig := StringNameFromStr("connection_to_empty")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionFromEmptySignalFn func(to_node StringName, to_port int, release_position Vector2)

func (me *GraphEdit) ConnectConnectionFromEmpty(subs SignalSubscribers, fn GraphEditConnectionFromEmptySignalFn) {
	sig := StringNameFromStr("connection_from_empty")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionFromEmpty(subs SignalSubscribers, fn GraphEditConnectionFromEmptySignalFn) {
	sig := StringNameFromStr("connection_from_empty")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionDragStartedSignalFn func(from_node StringName, from_port int, is_output bool)

func (me *GraphEdit) ConnectConnectionDragStarted(subs SignalSubscribers, fn GraphEditConnectionDragStartedSignalFn) {
	sig := StringNameFromStr("connection_drag_started")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionDragStarted(subs SignalSubscribers, fn GraphEditConnectionDragStartedSignalFn) {
	sig := StringNameFromStr("connection_drag_started")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditConnectionDragEndedSignalFn func()

func (me *GraphEdit) ConnectConnectionDragEnded(subs SignalSubscribers, fn GraphEditConnectionDragEndedSignalFn) {
	sig := StringNameFromStr("connection_drag_ended")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectConnectionDragEnded(subs SignalSubscribers, fn GraphEditConnectionDragEndedSignalFn) {
	sig := StringNameFromStr("connection_drag_ended")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditCopyNodesRequestSignalFn func()

func (me *GraphEdit) ConnectCopyNodesRequest(subs SignalSubscribers, fn GraphEditCopyNodesRequestSignalFn) {
	sig := StringNameFromStr("copy_nodes_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectCopyNodesRequest(subs SignalSubscribers, fn GraphEditCopyNodesRequestSignalFn) {
	sig := StringNameFromStr("copy_nodes_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditPasteNodesRequestSignalFn func()

func (me *GraphEdit) ConnectPasteNodesRequest(subs SignalSubscribers, fn GraphEditPasteNodesRequestSignalFn) {
	sig := StringNameFromStr("paste_nodes_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectPasteNodesRequest(subs SignalSubscribers, fn GraphEditPasteNodesRequestSignalFn) {
	sig := StringNameFromStr("paste_nodes_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditDuplicateNodesRequestSignalFn func()

func (me *GraphEdit) ConnectDuplicateNodesRequest(subs SignalSubscribers, fn GraphEditDuplicateNodesRequestSignalFn) {
	sig := StringNameFromStr("duplicate_nodes_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectDuplicateNodesRequest(subs SignalSubscribers, fn GraphEditDuplicateNodesRequestSignalFn) {
	sig := StringNameFromStr("duplicate_nodes_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditDeleteNodesRequestSignalFn func(nodes []StringName)

func (me *GraphEdit) ConnectDeleteNodesRequest(subs SignalSubscribers, fn GraphEditDeleteNodesRequestSignalFn) {
	sig := StringNameFromStr("delete_nodes_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectDeleteNodesRequest(subs SignalSubscribers, fn GraphEditDeleteNodesRequestSignalFn) {
	sig := StringNameFromStr("delete_nodes_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditNodeSelectedSignalFn func(node Node)

func (me *GraphEdit) ConnectNodeSelected(subs SignalSubscribers, fn GraphEditNodeSelectedSignalFn) {
	sig := StringNameFromStr("node_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectNodeSelected(subs SignalSubscribers, fn GraphEditNodeSelectedSignalFn) {
	sig := StringNameFromStr("node_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditNodeDeselectedSignalFn func(node Node)

func (me *GraphEdit) ConnectNodeDeselected(subs SignalSubscribers, fn GraphEditNodeDeselectedSignalFn) {
	sig := StringNameFromStr("node_deselected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectNodeDeselected(subs SignalSubscribers, fn GraphEditNodeDeselectedSignalFn) {
	sig := StringNameFromStr("node_deselected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditPopupRequestSignalFn func(position Vector2)

func (me *GraphEdit) ConnectPopupRequest(subs SignalSubscribers, fn GraphEditPopupRequestSignalFn) {
	sig := StringNameFromStr("popup_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectPopupRequest(subs SignalSubscribers, fn GraphEditPopupRequestSignalFn) {
	sig := StringNameFromStr("popup_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditBeginNodeMoveSignalFn func()

func (me *GraphEdit) ConnectBeginNodeMove(subs SignalSubscribers, fn GraphEditBeginNodeMoveSignalFn) {
	sig := StringNameFromStr("begin_node_move")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectBeginNodeMove(subs SignalSubscribers, fn GraphEditBeginNodeMoveSignalFn) {
	sig := StringNameFromStr("begin_node_move")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditEndNodeMoveSignalFn func()

func (me *GraphEdit) ConnectEndNodeMove(subs SignalSubscribers, fn GraphEditEndNodeMoveSignalFn) {
	sig := StringNameFromStr("end_node_move")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectEndNodeMove(subs SignalSubscribers, fn GraphEditEndNodeMoveSignalFn) {
	sig := StringNameFromStr("end_node_move")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphEditScrollOffsetChangedSignalFn func(offset Vector2)

func (me *GraphEdit) ConnectScrollOffsetChanged(subs SignalSubscribers, fn GraphEditScrollOffsetChangedSignalFn) {
	sig := StringNameFromStr("scroll_offset_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphEdit) DisconnectScrollOffsetChanged(subs SignalSubscribers, fn GraphEditScrollOffsetChangedSignalFn) {
	sig := StringNameFromStr("scroll_offset_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
