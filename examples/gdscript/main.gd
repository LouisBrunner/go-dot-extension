extends Node2D

var time: float = 0

@export var factor: float = 100

@onready var my: MyNode2D = $MyNode2D

func _ready() -> void:
  my.monitor($ColorRect)
  my.secret_printed.connect(func(x): print(x))
  my.report({ # non-string keys will be converted to strings
    "String Key": 5,
    4: [1, 2, 3],
    7: "Hello",
    "sub_dict": {"sub_key": "Nested value"},
  })
  my.report({
    "String Key": 5,
    "sub_dict": {"sub_key": "Nested value"},
  })

func _process(delta: float) -> void:
  time += delta
  my.move(Vector2(sin(time) * factor, cos(time) * factor) + Vector2(500, 300))
