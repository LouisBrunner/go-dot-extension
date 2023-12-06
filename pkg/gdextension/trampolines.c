#include <stdint.h>

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
