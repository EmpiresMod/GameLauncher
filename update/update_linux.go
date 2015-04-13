package update

func (u *Update) UpdateExec() (err error) {

	// We only need to do magic stuff to the executable on windows to replace it
	return u.Update()
}
