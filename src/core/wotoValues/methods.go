package wotoValues

//---------------------------------------------------------

func (w *WotoCore) LockDb() {
	w.DbMutex.Lock()
}

func (w *WotoCore) UnlockDb() {
	w.DbMutex.Unlock()
}

//---------------------------------------------------------
