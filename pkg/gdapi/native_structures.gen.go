// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

type AudioFrame struct {
  Left float32
  Right float32
}

func DefaultAudioFrame() AudioFrame {
  return AudioFrame{
  }
}

type CaretInfo struct {
  LeadingCaret Rect2
  TrailingCaret Rect2
  LeadingDirection TextServerDirection
  TrailingDirection TextServerDirection
}

func DefaultCaretInfo() CaretInfo {
  return CaretInfo{
  }
}

type Glyph struct {
  Start int
  End int
  Count uint8
  Repeat uint8
  Flags uint16
  XOff float32
  YOff float32
  Advance float32
  FontRid RID
  FontSize int
  Index int32
}

func DefaultGlyph() Glyph {
  return Glyph{
      Start: -1,
      End: -1,
      Count: 0,
      Repeat: 1,
      Flags: 0,
      XOff: 0,
      YOff: 0,
      Advance: 0,
      FontSize: 0,
      Index: 0,
  }
}

type ObjectID struct {
  Id uint64
}

func DefaultObjectID() ObjectID {
  return ObjectID{
      Id: 0,
  }
}

type PhysicsServer2DExtensionMotionResult struct {
  Travel Vector2
  Remainder Vector2
  CollisionPoint Vector2
  CollisionNormal Vector2
  ColliderVelocity Vector2
  CollisionDepth RealType
  CollisionSafeFraction RealType
  CollisionUnsafeFraction RealType
  CollisionLocalShape int
  ColliderId ObjectID
  Collider RID
  ColliderShape int
}

func DefaultPhysicsServer2DExtensionMotionResult() PhysicsServer2DExtensionMotionResult {
  return PhysicsServer2DExtensionMotionResult{
  }
}

type PhysicsServer2DExtensionRayResult struct {
  Position Vector2
  Normal Vector2
  Rid RID
  ColliderId ObjectID
  Collider *Object
  Shape int
}

func DefaultPhysicsServer2DExtensionRayResult() PhysicsServer2DExtensionRayResult {
  return PhysicsServer2DExtensionRayResult{
  }
}

type PhysicsServer2DExtensionShapeRestInfo struct {
  Point Vector2
  Normal Vector2
  Rid RID
  ColliderId ObjectID
  Shape int
  LinearVelocity Vector2
}

func DefaultPhysicsServer2DExtensionShapeRestInfo() PhysicsServer2DExtensionShapeRestInfo {
  return PhysicsServer2DExtensionShapeRestInfo{
  }
}

type PhysicsServer2DExtensionShapeResult struct {
  Rid RID
  ColliderId ObjectID
  Collider *Object
  Shape int
}

func DefaultPhysicsServer2DExtensionShapeResult() PhysicsServer2DExtensionShapeResult {
  return PhysicsServer2DExtensionShapeResult{
  }
}

type PhysicsServer3DExtensionMotionCollision struct {
  Position Vector3
  Normal Vector3
  ColliderVelocity Vector3
  ColliderAngularVelocity Vector3
  Depth RealType
  LocalShape int
  ColliderId ObjectID
  Collider RID
  ColliderShape int
}

func DefaultPhysicsServer3DExtensionMotionCollision() PhysicsServer3DExtensionMotionCollision {
  return PhysicsServer3DExtensionMotionCollision{
  }
}

type PhysicsServer3DExtensionMotionResult struct {
  Travel Vector3
  Remainder Vector3
  CollisionDepth RealType
  CollisionSafeFraction RealType
  CollisionUnsafeFraction RealType
  Collisions [32]PhysicsServer3DExtensionMotionCollision
  CollisionCount int
}

func DefaultPhysicsServer3DExtensionMotionResult() PhysicsServer3DExtensionMotionResult {
  return PhysicsServer3DExtensionMotionResult{
  }
}

type PhysicsServer3DExtensionRayResult struct {
  Position Vector3
  Normal Vector3
  Rid RID
  ColliderId ObjectID
  Collider *Object
  Shape int
}

func DefaultPhysicsServer3DExtensionRayResult() PhysicsServer3DExtensionRayResult {
  return PhysicsServer3DExtensionRayResult{
  }
}

type PhysicsServer3DExtensionShapeRestInfo struct {
  Point Vector3
  Normal Vector3
  Rid RID
  ColliderId ObjectID
  Shape int
  LinearVelocity Vector3
}

func DefaultPhysicsServer3DExtensionShapeRestInfo() PhysicsServer3DExtensionShapeRestInfo {
  return PhysicsServer3DExtensionShapeRestInfo{
  }
}

type PhysicsServer3DExtensionShapeResult struct {
  Rid RID
  ColliderId ObjectID
  Collider *Object
  Shape int
}

func DefaultPhysicsServer3DExtensionShapeResult() PhysicsServer3DExtensionShapeResult {
  return PhysicsServer3DExtensionShapeResult{
  }
}

type ScriptLanguageExtensionProfilingInfo struct {
  Signature StringName
  CallCount uint64
  TotalTime uint64
  SelfTime uint64
}

func DefaultScriptLanguageExtensionProfilingInfo() ScriptLanguageExtensionProfilingInfo {
  return ScriptLanguageExtensionProfilingInfo{
  }
}
