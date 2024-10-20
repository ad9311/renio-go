package seed

func Run() error {
	if err := seedEntryClasses(); err != nil {
		return err
	}

	return nil
}
