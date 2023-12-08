// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

type Side int
const (
  SideLeft Side = 0
  SideTop Side = 1
  SideRight Side = 2
  SideBottom Side = 3
)

type Corner int
const (
  CornerTopLeft Corner = 0
  CornerTopRight Corner = 1
  CornerBottomRight Corner = 2
  CornerBottomLeft Corner = 3
)

type Orientation int
const (
  Vertical Orientation = 1
  Horizontal Orientation = 0
)

type ClockDirection int
const (
  Clockwise ClockDirection = 0
  Counterclockwise ClockDirection = 1
)

type HorizontalAlignment int
const (
  HorizontalAlignmentLeft HorizontalAlignment = 0
  HorizontalAlignmentCenter HorizontalAlignment = 1
  HorizontalAlignmentRight HorizontalAlignment = 2
  HorizontalAlignmentFill HorizontalAlignment = 3
)

type VerticalAlignment int
const (
  VerticalAlignmentTop VerticalAlignment = 0
  VerticalAlignmentCenter VerticalAlignment = 1
  VerticalAlignmentBottom VerticalAlignment = 2
  VerticalAlignmentFill VerticalAlignment = 3
)

type InlineAlignment int
const (
  InlineAlignmentTopTo InlineAlignment = 0
  InlineAlignmentCenterTo InlineAlignment = 1
  InlineAlignmentBaselineTo InlineAlignment = 3
  InlineAlignmentBottomTo InlineAlignment = 2
  InlineAlignmentToTop InlineAlignment = 0
  InlineAlignmentToCenter InlineAlignment = 4
  InlineAlignmentToBaseline InlineAlignment = 8
  InlineAlignmentToBottom InlineAlignment = 12
  InlineAlignmentTop InlineAlignment = 0
  InlineAlignmentCenter InlineAlignment = 5
  InlineAlignmentBottom InlineAlignment = 14
  InlineAlignmentImageMask InlineAlignment = 3
  InlineAlignmentTextMask InlineAlignment = 12
)

type EulerOrder int
const (
  EulerOrderXyz EulerOrder = 0
  EulerOrderXzy EulerOrder = 1
  EulerOrderYxz EulerOrder = 2
  EulerOrderYzx EulerOrder = 3
  EulerOrderZxy EulerOrder = 4
  EulerOrderZyx EulerOrder = 5
)

type Key int
const (
  KeyNone Key = 0
  KeySpecial Key = 4194304
  KeyEscape Key = 4194305
  KeyTab Key = 4194306
  KeyBacktab Key = 4194307
  KeyBackspace Key = 4194308
  KeyEnter Key = 4194309
  KeyKpEnter Key = 4194310
  KeyInsert Key = 4194311
  KeyDelete Key = 4194312
  KeyPause Key = 4194313
  KeyPrint Key = 4194314
  KeySysreq Key = 4194315
  KeyClear Key = 4194316
  KeyHome Key = 4194317
  KeyEnd Key = 4194318
  KeyLeft Key = 4194319
  KeyUp Key = 4194320
  KeyRight Key = 4194321
  KeyDown Key = 4194322
  KeyPageup Key = 4194323
  KeyPagedown Key = 4194324
  KeyShift Key = 4194325
  KeyCtrl Key = 4194326
  KeyMeta Key = 4194327
  KeyAlt Key = 4194328
  KeyCapslock Key = 4194329
  KeyNumlock Key = 4194330
  KeyScrolllock Key = 4194331
  KeyF1 Key = 4194332
  KeyF2 Key = 4194333
  KeyF3 Key = 4194334
  KeyF4 Key = 4194335
  KeyF5 Key = 4194336
  KeyF6 Key = 4194337
  KeyF7 Key = 4194338
  KeyF8 Key = 4194339
  KeyF9 Key = 4194340
  KeyF10 Key = 4194341
  KeyF11 Key = 4194342
  KeyF12 Key = 4194343
  KeyF13 Key = 4194344
  KeyF14 Key = 4194345
  KeyF15 Key = 4194346
  KeyF16 Key = 4194347
  KeyF17 Key = 4194348
  KeyF18 Key = 4194349
  KeyF19 Key = 4194350
  KeyF20 Key = 4194351
  KeyF21 Key = 4194352
  KeyF22 Key = 4194353
  KeyF23 Key = 4194354
  KeyF24 Key = 4194355
  KeyF25 Key = 4194356
  KeyF26 Key = 4194357
  KeyF27 Key = 4194358
  KeyF28 Key = 4194359
  KeyF29 Key = 4194360
  KeyF30 Key = 4194361
  KeyF31 Key = 4194362
  KeyF32 Key = 4194363
  KeyF33 Key = 4194364
  KeyF34 Key = 4194365
  KeyF35 Key = 4194366
  KeyKpMultiply Key = 4194433
  KeyKpDivide Key = 4194434
  KeyKpSubtract Key = 4194435
  KeyKpPeriod Key = 4194436
  KeyKpAdd Key = 4194437
  KeyKp0 Key = 4194438
  KeyKp1 Key = 4194439
  KeyKp2 Key = 4194440
  KeyKp3 Key = 4194441
  KeyKp4 Key = 4194442
  KeyKp5 Key = 4194443
  KeyKp6 Key = 4194444
  KeyKp7 Key = 4194445
  KeyKp8 Key = 4194446
  KeyKp9 Key = 4194447
  KeyMenu Key = 4194370
  KeyHyper Key = 4194371
  KeyHelp Key = 4194373
  KeyBack Key = 4194376
  KeyForward Key = 4194377
  KeyStop Key = 4194378
  KeyRefresh Key = 4194379
  KeyVolumedown Key = 4194380
  KeyVolumemute Key = 4194381
  KeyVolumeup Key = 4194382
  KeyMediaplay Key = 4194388
  KeyMediastop Key = 4194389
  KeyMediaprevious Key = 4194390
  KeyMedianext Key = 4194391
  KeyMediarecord Key = 4194392
  KeyHomepage Key = 4194393
  KeyFavorites Key = 4194394
  KeySearch Key = 4194395
  KeyStandby Key = 4194396
  KeyOpenurl Key = 4194397
  KeyLaunchmail Key = 4194398
  KeyLaunchmedia Key = 4194399
  KeyLaunch0 Key = 4194400
  KeyLaunch1 Key = 4194401
  KeyLaunch2 Key = 4194402
  KeyLaunch3 Key = 4194403
  KeyLaunch4 Key = 4194404
  KeyLaunch5 Key = 4194405
  KeyLaunch6 Key = 4194406
  KeyLaunch7 Key = 4194407
  KeyLaunch8 Key = 4194408
  KeyLaunch9 Key = 4194409
  KeyLauncha Key = 4194410
  KeyLaunchb Key = 4194411
  KeyLaunchc Key = 4194412
  KeyLaunchd Key = 4194413
  KeyLaunche Key = 4194414
  KeyLaunchf Key = 4194415
  KeyGlobe Key = 4194416
  KeyKeyboard Key = 4194417
  KeyJisEisu Key = 4194418
  KeyJisKana Key = 4194419
  KeyUnknown Key = 8388607
  KeySpace Key = 32
  KeyExclam Key = 33
  KeyQuotedbl Key = 34
  KeyNumbersign Key = 35
  KeyDollar Key = 36
  KeyPercent Key = 37
  KeyAmpersand Key = 38
  KeyApostrophe Key = 39
  KeyParenleft Key = 40
  KeyParenright Key = 41
  KeyAsterisk Key = 42
  KeyPlus Key = 43
  KeyComma Key = 44
  KeyMinus Key = 45
  KeyPeriod Key = 46
  KeySlash Key = 47
  Key0 Key = 48
  Key1 Key = 49
  Key2 Key = 50
  Key3 Key = 51
  Key4 Key = 52
  Key5 Key = 53
  Key6 Key = 54
  Key7 Key = 55
  Key8 Key = 56
  Key9 Key = 57
  KeyColon Key = 58
  KeySemicolon Key = 59
  KeyLess Key = 60
  KeyEqual Key = 61
  KeyGreater Key = 62
  KeyQuestion Key = 63
  KeyAt Key = 64
  KeyA Key = 65
  KeyB Key = 66
  KeyC Key = 67
  KeyD Key = 68
  KeyE Key = 69
  KeyF Key = 70
  KeyG Key = 71
  KeyH Key = 72
  KeyI Key = 73
  KeyJ Key = 74
  KeyK Key = 75
  KeyL Key = 76
  KeyM Key = 77
  KeyN Key = 78
  KeyO Key = 79
  KeyP Key = 80
  KeyQ Key = 81
  KeyR Key = 82
  KeyS Key = 83
  KeyT Key = 84
  KeyU Key = 85
  KeyV Key = 86
  KeyW Key = 87
  KeyX Key = 88
  KeyY Key = 89
  KeyZ Key = 90
  KeyBracketleft Key = 91
  KeyBackslash Key = 92
  KeyBracketright Key = 93
  KeyAsciicircum Key = 94
  KeyUnderscore Key = 95
  KeyQuoteleft Key = 96
  KeyBraceleft Key = 123
  KeyBar Key = 124
  KeyBraceright Key = 125
  KeyAsciitilde Key = 126
  KeyYen Key = 165
  KeySection Key = 167
)

type KeyModifierMask int
const (
  KeyCodeMask KeyModifierMask = 8388607
  KeyModifierMask_ KeyModifierMask = 532676608
  KeyMaskCmdOrCtrl KeyModifierMask = 16777216
  KeyMaskShift KeyModifierMask = 33554432
  KeyMaskAlt KeyModifierMask = 67108864
  KeyMaskMeta KeyModifierMask = 134217728
  KeyMaskCtrl KeyModifierMask = 268435456
  KeyMaskKpad KeyModifierMask = 536870912
  KeyMaskGroupSwitch KeyModifierMask = 1073741824
)

type MouseButton int
const (
  MouseButtonNone MouseButton = 0
  MouseButtonLeft MouseButton = 1
  MouseButtonRight MouseButton = 2
  MouseButtonMiddle MouseButton = 3
  MouseButtonWheelUp MouseButton = 4
  MouseButtonWheelDown MouseButton = 5
  MouseButtonWheelLeft MouseButton = 6
  MouseButtonWheelRight MouseButton = 7
  MouseButtonXbutton1 MouseButton = 8
  MouseButtonXbutton2 MouseButton = 9
)

type MouseButtonMask int
const (
  MouseButtonMaskLeft MouseButtonMask = 1
  MouseButtonMaskRight MouseButtonMask = 2
  MouseButtonMaskMiddle MouseButtonMask = 4
  MouseButtonMaskMbXbutton1 MouseButtonMask = 128
  MouseButtonMaskMbXbutton2 MouseButtonMask = 256
)

type JoyButton int
const (
  JoyButtonInvalid JoyButton = -1
  JoyButtonA JoyButton = 0
  JoyButtonB JoyButton = 1
  JoyButtonX JoyButton = 2
  JoyButtonY JoyButton = 3
  JoyButtonBack JoyButton = 4
  JoyButtonGuide JoyButton = 5
  JoyButtonStart JoyButton = 6
  JoyButtonLeftStick JoyButton = 7
  JoyButtonRightStick JoyButton = 8
  JoyButtonLeftShoulder JoyButton = 9
  JoyButtonRightShoulder JoyButton = 10
  JoyButtonDpadUp JoyButton = 11
  JoyButtonDpadDown JoyButton = 12
  JoyButtonDpadLeft JoyButton = 13
  JoyButtonDpadRight JoyButton = 14
  JoyButtonMisc1 JoyButton = 15
  JoyButtonPaddle1 JoyButton = 16
  JoyButtonPaddle2 JoyButton = 17
  JoyButtonPaddle3 JoyButton = 18
  JoyButtonPaddle4 JoyButton = 19
  JoyButtonTouchpad JoyButton = 20
  JoyButtonSdlMax JoyButton = 21
  JoyButtonMax JoyButton = 128
)

type JoyAxis int
const (
  JoyAxisInvalid JoyAxis = -1
  JoyAxisLeftX JoyAxis = 0
  JoyAxisLeftY JoyAxis = 1
  JoyAxisRightX JoyAxis = 2
  JoyAxisRightY JoyAxis = 3
  JoyAxisTriggerLeft JoyAxis = 4
  JoyAxisTriggerRight JoyAxis = 5
  JoyAxisSdlMax JoyAxis = 6
  JoyAxisMax JoyAxis = 10
)

type MIDIMessage int
const (
  MidiMessageNone MIDIMessage = 0
  MidiMessageNoteOff MIDIMessage = 8
  MidiMessageNoteOn MIDIMessage = 9
  MidiMessageAftertouch MIDIMessage = 10
  MidiMessageControlChange MIDIMessage = 11
  MidiMessageProgramChange MIDIMessage = 12
  MidiMessageChannelPressure MIDIMessage = 13
  MidiMessagePitchBend MIDIMessage = 14
  MidiMessageSystemExclusive MIDIMessage = 240
  MidiMessageQuarterFrame MIDIMessage = 241
  MidiMessageSongPositionPointer MIDIMessage = 242
  MidiMessageSongSelect MIDIMessage = 243
  MidiMessageTuneRequest MIDIMessage = 246
  MidiMessageTimingClock MIDIMessage = 248
  MidiMessageStart MIDIMessage = 250
  MidiMessageContinue MIDIMessage = 251
  MidiMessageStop MIDIMessage = 252
  MidiMessageActiveSensing MIDIMessage = 254
  MidiMessageSystemReset MIDIMessage = 255
)

type Error int
const (
  Ok Error = 0
  Failed Error = 1
  ErrUnavailable Error = 2
  ErrUnconfigured Error = 3
  ErrUnauthorized Error = 4
  ErrParameterRangeError Error = 5
  ErrOutOfMemory Error = 6
  ErrFileNotFound Error = 7
  ErrFileBadDrive Error = 8
  ErrFileBadPath Error = 9
  ErrFileNoPermission Error = 10
  ErrFileAlreadyInUse Error = 11
  ErrFileCantOpen Error = 12
  ErrFileCantWrite Error = 13
  ErrFileCantRead Error = 14
  ErrFileUnrecognized Error = 15
  ErrFileCorrupt Error = 16
  ErrFileMissingDependencies Error = 17
  ErrFileEof Error = 18
  ErrCantOpen Error = 19
  ErrCantCreate Error = 20
  ErrQueryFailed Error = 21
  ErrAlreadyInUse Error = 22
  ErrLocked Error = 23
  ErrTimeout Error = 24
  ErrCantConnect Error = 25
  ErrCantResolve Error = 26
  ErrConnectionError Error = 27
  ErrCantAcquireResource Error = 28
  ErrCantFork Error = 29
  ErrInvalidData Error = 30
  ErrInvalidParameter Error = 31
  ErrAlreadyExists Error = 32
  ErrDoesNotExist Error = 33
  ErrDatabaseCantRead Error = 34
  ErrDatabaseCantWrite Error = 35
  ErrCompilationFailed Error = 36
  ErrMethodNotFound Error = 37
  ErrLinkFailed Error = 38
  ErrScriptFailed Error = 39
  ErrCyclicLink Error = 40
  ErrInvalidDeclaration Error = 41
  ErrDuplicateSymbol Error = 42
  ErrParseError Error = 43
  ErrBusy Error = 44
  ErrSkip Error = 45
  ErrHelp Error = 46
  ErrBug Error = 47
  ErrPrinterOnFire Error = 48
)

type PropertyHint int
const (
  PropertyHintNone PropertyHint = 0
  PropertyHintRange PropertyHint = 1
  PropertyHintEnum PropertyHint = 2
  PropertyHintEnumSuggestion PropertyHint = 3
  PropertyHintExpEasing PropertyHint = 4
  PropertyHintLink PropertyHint = 5
  PropertyHintFlags PropertyHint = 6
  PropertyHintLayers2DRender PropertyHint = 7
  PropertyHintLayers2DPhysics PropertyHint = 8
  PropertyHintLayers2DNavigation PropertyHint = 9
  PropertyHintLayers3DRender PropertyHint = 10
  PropertyHintLayers3DPhysics PropertyHint = 11
  PropertyHintLayers3DNavigation PropertyHint = 12
  PropertyHintLayersAvoidance PropertyHint = 37
  PropertyHintFile PropertyHint = 13
  PropertyHintDir PropertyHint = 14
  PropertyHintGlobalFile PropertyHint = 15
  PropertyHintGlobalDir PropertyHint = 16
  PropertyHintResourceType PropertyHint = 17
  PropertyHintMultilineText PropertyHint = 18
  PropertyHintExpression PropertyHint = 19
  PropertyHintPlaceholderText PropertyHint = 20
  PropertyHintColorNoAlpha PropertyHint = 21
  PropertyHintObjectId PropertyHint = 22
  PropertyHintTypeString PropertyHint = 23
  PropertyHintNodePathToEditedNode PropertyHint = 24
  PropertyHintObjectTooBig PropertyHint = 25
  PropertyHintNodePathValidTypes PropertyHint = 26
  PropertyHintSaveFile PropertyHint = 27
  PropertyHintGlobalSaveFile PropertyHint = 28
  PropertyHintIntIsObjectid PropertyHint = 29
  PropertyHintIntIsPointer PropertyHint = 30
  PropertyHintArrayType PropertyHint = 31
  PropertyHintLocaleId PropertyHint = 32
  PropertyHintLocalizableString PropertyHint = 33
  PropertyHintNodeType PropertyHint = 34
  PropertyHintHideQuaternionEdit PropertyHint = 35
  PropertyHintPassword PropertyHint = 36
  PropertyHintMax PropertyHint = 38
)

type PropertyUsageFlags int
const (
  PropertyUsageNone PropertyUsageFlags = 0
  PropertyUsageStorage PropertyUsageFlags = 2
  PropertyUsageEditor PropertyUsageFlags = 4
  PropertyUsageInternal PropertyUsageFlags = 8
  PropertyUsageCheckable PropertyUsageFlags = 16
  PropertyUsageChecked PropertyUsageFlags = 32
  PropertyUsageGroup PropertyUsageFlags = 64
  PropertyUsageCategory PropertyUsageFlags = 128
  PropertyUsageSubgroup PropertyUsageFlags = 256
  PropertyUsageClassIsBitfield PropertyUsageFlags = 512
  PropertyUsageNoInstanceState PropertyUsageFlags = 1024
  PropertyUsageRestartIfChanged PropertyUsageFlags = 2048
  PropertyUsageScriptVariable PropertyUsageFlags = 4096
  PropertyUsageStoreIfNull PropertyUsageFlags = 8192
  PropertyUsageUpdateAllIfModified PropertyUsageFlags = 16384
  PropertyUsageScriptDefaultValue PropertyUsageFlags = 32768
  PropertyUsageClassIsEnum PropertyUsageFlags = 65536
  PropertyUsageNilIsVariant PropertyUsageFlags = 131072
  PropertyUsageArray PropertyUsageFlags = 262144
  PropertyUsageAlwaysDuplicate PropertyUsageFlags = 524288
  PropertyUsageNeverDuplicate PropertyUsageFlags = 1048576
  PropertyUsageHighEndGfx PropertyUsageFlags = 2097152
  PropertyUsageNodePathFromSceneRoot PropertyUsageFlags = 4194304
  PropertyUsageResourceNotPersistent PropertyUsageFlags = 8388608
  PropertyUsageKeyingIncrements PropertyUsageFlags = 16777216
  PropertyUsageDeferredSetResource PropertyUsageFlags = 33554432
  PropertyUsageEditorInstantiateObject PropertyUsageFlags = 67108864
  PropertyUsageEditorBasicSetting PropertyUsageFlags = 134217728
  PropertyUsageReadOnly PropertyUsageFlags = 268435456
  PropertyUsageSecret PropertyUsageFlags = 536870912
  PropertyUsageDefault PropertyUsageFlags = 6
  PropertyUsageNoEditor PropertyUsageFlags = 2
)

type MethodFlags int
const (
  MethodFlagNormal MethodFlags = 1
  MethodFlagEditor MethodFlags = 2
  MethodFlagConst MethodFlags = 4
  MethodFlagVirtual MethodFlags = 8
  MethodFlagVararg MethodFlags = 16
  MethodFlagStatic MethodFlags = 32
  MethodFlagObjectCore MethodFlags = 64
  MethodFlagsDefault MethodFlags = 1
)

type VariantType int
const (
  TypeNil VariantType = 0
  TypeBool VariantType = 1
  TypeInt VariantType = 2
  TypeFloat VariantType = 3
  TypeString VariantType = 4
  TypeVector2 VariantType = 5
  TypeVector2I VariantType = 6
  TypeRect2 VariantType = 7
  TypeRect2I VariantType = 8
  TypeVector3 VariantType = 9
  TypeVector3I VariantType = 10
  TypeTransform2D VariantType = 11
  TypeVector4 VariantType = 12
  TypeVector4I VariantType = 13
  TypePlane VariantType = 14
  TypeQuaternion VariantType = 15
  TypeAabb VariantType = 16
  TypeBasis VariantType = 17
  TypeTransform3D VariantType = 18
  TypeProjection VariantType = 19
  TypeColor VariantType = 20
  TypeStringName VariantType = 21
  TypeNodePath VariantType = 22
  TypeRid VariantType = 23
  TypeObject VariantType = 24
  TypeCallable VariantType = 25
  TypeSignal VariantType = 26
  TypeDictionary VariantType = 27
  TypeArray VariantType = 28
  TypePackedByteArray VariantType = 29
  TypePackedInt32Array VariantType = 30
  TypePackedInt64Array VariantType = 31
  TypePackedFloat32Array VariantType = 32
  TypePackedFloat64Array VariantType = 33
  TypePackedStringArray VariantType = 34
  TypePackedVector2Array VariantType = 35
  TypePackedVector3Array VariantType = 36
  TypePackedColorArray VariantType = 37
  TypeMax VariantType = 38
)

type VariantOperator int
const (
  OpEqual VariantOperator = 0
  OpNotEqual VariantOperator = 1
  OpLess VariantOperator = 2
  OpLessEqual VariantOperator = 3
  OpGreater VariantOperator = 4
  OpGreaterEqual VariantOperator = 5
  OpAdd VariantOperator = 6
  OpSubtract VariantOperator = 7
  OpMultiply VariantOperator = 8
  OpDivide VariantOperator = 9
  OpNegate VariantOperator = 10
  OpPositive VariantOperator = 11
  OpModule VariantOperator = 12
  OpPower VariantOperator = 13
  OpShiftLeft VariantOperator = 14
  OpShiftRight VariantOperator = 15
  OpBitAnd VariantOperator = 16
  OpBitOr VariantOperator = 17
  OpBitXor VariantOperator = 18
  OpBitNegate VariantOperator = 19
  OpAnd VariantOperator = 20
  OpOr VariantOperator = 21
  OpXor VariantOperator = 22
  OpNot VariantOperator = 23
  OpIn VariantOperator = 24
  OpMax VariantOperator = 25
)
