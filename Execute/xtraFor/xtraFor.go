package xtrafor

type Loop struct {
	i      int
	paused bool
	call   func(i int)
}

func (l *Loop) Step() {
	if !(l.paused) {
		l.call(l.i)
		l.i += 1
	}
}

func (l *Loop) Jump(o int) {
	l.i = o
	l.call(l.i)
}

func (l *Loop) ChangeAttributes(i int, p bool, call func(i int)) {
	l.i = i
	l.paused = p
	l.call = call
}
