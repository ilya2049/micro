package cleanup

type Func func()

type Funcs []Func

func (fns Funcs) Invoke() {
	for i := len(fns); i >= 0; i-- {
		fns[i]()
	}
}
