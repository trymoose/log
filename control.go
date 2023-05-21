package log

type (
	_Controller func(Arg) (arg Arg, ok bool)
)

// _ArgHandler creates a handler that checks if the given [Arg]'s type == T.
// If the types match the value is passed to fn and true is returned.
func _ArgHandler[T Arg](fn func(T) Arg) _Controller {
	return func(ctrl Arg) (Arg, bool) {
		if v, ok := any(ctrl).(T); ok && fn != nil {
			return fn(v), true
		}
		return Skip(), false
	}
}

// _ArgHandlerSkip is the same as [_ArgHandler] except all values it catches returns [Skip].
func _ArgHandlerSkip[T Arg](fn func(T)) _Controller {
	return _ArgHandler[T](func(t T) Arg {
		fn(t)
		return Skip()
	})
}

// _HandleArgs loops through in calling each handler on the value until it is handled.
// If it is handled it will not be present in out.
// Out will be a new slice from in.
func _HandleArgs(in []Arg, handlers ..._Controller) (out []Arg) {
	out = make([]Arg, 0, len(in))
	for _, e := range in {
		if e = handleArg(e, handlers); e != Skip() {
			out = append(out, e)
		}
	}
	return
}

func handleArg(arg Arg, handlers []_Controller) Arg {
	for _, h := range handlers {
		if r, ok := h(arg); ok {
			return r
		}
	}
	return arg
}
