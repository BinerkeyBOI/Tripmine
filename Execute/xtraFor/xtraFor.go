package xtrafor

type loop struct {
	i      int
	paused bool
	call   func(i int)
}

func (l *loop) Step() {
	if !(l.paused) {
		l.call(l.i)
		l.i += 1
	}
}

func (l *loop) Jump(o int) {
	l.i = o
	l.call(l.i)
}

func Loop(i int, paused bool, call func(i int)) loop {
	l := loop{
		i:      i,
		paused: paused,
		call:   call,
	}
	return l
}
