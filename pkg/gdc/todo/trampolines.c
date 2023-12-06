#include <stdint.h>

void* go_dot_gdextension_get_proc_address(void *(*p_get_proc_address)(const char *name), const char *name) {
	return p_get_proc_address(name);
}

void go_dot_gdextension_classdb_register_extension_class(void (*fn)(void *lib, const void *name, const void *parent_name, const void *def), void *lib, const void *name, const void *parent_name, const void *def) {
	fn(lib, name, parent_name, def);
}

void go_dot_gdextension_classdb_unregister_extension_class(void (*fn)(void *lib, const void *name), void *lib, const void *name) {
	fn(lib, name);
}

void go_dot_gdextension_print_error(void (*fn)(const char *p_description, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify), const char *p_description, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify) {
  fn(p_description, p_function, p_file, p_line, p_editor_notify);
}

void go_dot_gdextension_print_error_with_message(void (*fn)(const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify), const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify) {
  fn(p_description, p_message, p_function, p_file, p_line, p_editor_notify);
}

void go_dot_gdextension_print_warning(void (*fn)(const char *p_description, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify), const char *p_description, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify) {
  fn(p_description, p_function, p_file, p_line, p_editor_notify);
}

void go_dot_gdextension_print_warning_with_message(void (*fn)(const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify), const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, char p_editor_notify) {
  fn(p_description, p_message, p_function, p_file, p_line, p_editor_notify);
}

void* go_dot_gdextension_variant_get_ptr_destructor(void* (*fn)(int p_type), int p_type) {
  return fn(p_type);
}

void* go_dot_gdextension_variant_get_ptr_constructor(void* (*fn)(int p_type, int32_t p_constructor), int p_type, int32_t p_constructor) {
  return fn(p_type, p_constructor);
}

void* go_dot_gdextension_get_variant_from_type_constructor(void* (*fn)(int p_type), int p_type) {
  return fn(p_type);
}

void go_dot_gdextension_string_new_with_utf8_chars(void (*fn)(void *r_dest, const char *p_contents), void *r_dest, const char *p_contents) {
  fn(r_dest, p_contents);
}

void* go_dot_gdextension_variant_get_ptr_utility_function(void* (*fn)(const void *p_function, int p_hash), const void *p_function, int p_hash) {
  return fn(p_function, p_hash);
}

void go_dot_gdextension_variant_destroy(void (*fn)(void *p_self), void *p_self) {
  fn(p_self);
}

void go_dot_gdextension_call_variant_constructor(void (*fn)(void *p_self, const void *p_args, int p_argcount), void *p_self, const void *p_args, int p_argcount) {
  fn(p_self, p_args, p_argcount);
}

void go_dot_gdextension_call_variant_type_constructor(void (*fn)(void *p_self, const void *inside), void *p_self, const void *inside) {
  fn(p_self, inside);
}

void go_dot_gdextension_call_variant_destructor(void (*fn)(void *p_self), void *p_self) {
  fn(p_self);
}

void go_dot_gdextension_call_utility_function(void (*fn)(void** ret, const void *p_args, int p_argcount), void** ret, const void *p_args, int p_argcount) {
  fn(ret, p_args, p_argcount);
}

void* go_dot_gdextension_classdb_construct_object(void* (*fn)(const void *p_classname), const void *p_classname) {
  return fn(p_classname);
}

void go_dot_gdextension_object_set_instance(void (*fn)(void *p_o, const void *p_classname, void *p_instance), void *p_o, const void *p_classname, void *p_instance) {
  fn(p_o, p_classname, p_instance);
}

uint64_t go_dot_gdextension_object_get_instance_id(uint64_t (*fn)(void *p_o), void *p_o) {
  return fn(p_o);
}
