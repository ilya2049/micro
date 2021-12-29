package cleanup

type Func func()

type Funcs []Func

func (fns Funcs) Invoke() {
	if len(fns) == 0 {
		return
	}

	for i := len(fns) - 1; i >= 0; i-- {
		fns[i]()
	}
}
