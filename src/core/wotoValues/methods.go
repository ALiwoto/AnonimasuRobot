package wotoValues

//---------------------------------------------------------

func (w *WotoCore) LockDb() {
	w.DbMutex.Lock()
}

func (w *WotoCore) UnlockDb() {
	w.DbMutex.Unlock()
}

func (w *WotoCore) AutoMigrateDB(models ...interface{}) error {
	if len(models) == 0 {
		return nil
	}

	return w.DB.AutoMigrate(models...)
}

//---------------------------------------------------------

//---------------------------------------------------------
