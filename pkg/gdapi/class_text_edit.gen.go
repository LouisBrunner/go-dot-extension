// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextEdit struct {
  obj gdc.ObjectPtr
}

func (me *TextEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextEdit) BaseClass() string {
  return "TextEdit"
}

type TextEditMenuItems int
const (
  TextEditMenuCut TextEditMenuItems = 0
  TextEditMenuCopy TextEditMenuItems = 1
  TextEditMenuPaste TextEditMenuItems = 2
  TextEditMenuClear TextEditMenuItems = 3
  TextEditMenuSelectAll TextEditMenuItems = 4
  TextEditMenuUndo TextEditMenuItems = 5
  TextEditMenuRedo TextEditMenuItems = 6
  TextEditMenuSubmenuTextDir TextEditMenuItems = 7
  TextEditMenuDirInherited TextEditMenuItems = 8
  TextEditMenuDirAuto TextEditMenuItems = 9
  TextEditMenuDirLtr TextEditMenuItems = 10
  TextEditMenuDirRtl TextEditMenuItems = 11
  TextEditMenuDisplayUcc TextEditMenuItems = 12
  TextEditMenuSubmenuInsertUcc TextEditMenuItems = 13
  TextEditMenuInsertLrm TextEditMenuItems = 14
  TextEditMenuInsertRlm TextEditMenuItems = 15
  TextEditMenuInsertLre TextEditMenuItems = 16
  TextEditMenuInsertRle TextEditMenuItems = 17
  TextEditMenuInsertLro TextEditMenuItems = 18
  TextEditMenuInsertRlo TextEditMenuItems = 19
  TextEditMenuInsertPdf TextEditMenuItems = 20
  TextEditMenuInsertAlm TextEditMenuItems = 21
  TextEditMenuInsertLri TextEditMenuItems = 22
  TextEditMenuInsertRli TextEditMenuItems = 23
  TextEditMenuInsertFsi TextEditMenuItems = 24
  TextEditMenuInsertPdi TextEditMenuItems = 25
  TextEditMenuInsertZwj TextEditMenuItems = 26
  TextEditMenuInsertZwnj TextEditMenuItems = 27
  TextEditMenuInsertWj TextEditMenuItems = 28
  TextEditMenuInsertShy TextEditMenuItems = 29
  TextEditMenuMax TextEditMenuItems = 30
)

type TextEditEditAction int
const (
  TextEditActionNone TextEditEditAction = 0
  TextEditActionTyping TextEditEditAction = 1
  TextEditActionBackspace TextEditEditAction = 2
  TextEditActionDelete TextEditEditAction = 3
)

type TextEditSearchFlags int
const (
  TextEditSearchMatchCase TextEditSearchFlags = 1
  TextEditSearchWholeWords TextEditSearchFlags = 2
  TextEditSearchBackwards TextEditSearchFlags = 4
)

type TextEditCaretType int
const (
  TextEditCaretTypeLine TextEditCaretType = 0
  TextEditCaretTypeBlock TextEditCaretType = 1
)

type TextEditSelectionMode int
const (
  TextEditSelectionModeNone TextEditSelectionMode = 0
  TextEditSelectionModeShift TextEditSelectionMode = 1
  TextEditSelectionModePointer TextEditSelectionMode = 2
  TextEditSelectionModeWord TextEditSelectionMode = 3
  TextEditSelectionModeLine TextEditSelectionMode = 4
)

type TextEditLineWrappingMode int
const (
  TextEditLineWrappingNone TextEditLineWrappingMode = 0
  TextEditLineWrappingBoundary TextEditLineWrappingMode = 1
)

type TextEditGutterType int
const (
  TextEditGutterTypeString TextEditGutterType = 0
  TextEditGutterTypeIcon TextEditGutterType = 1
  TextEditGutterTypeCustom TextEditGutterType = 2
)
