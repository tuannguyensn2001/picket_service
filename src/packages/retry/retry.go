package retrypkg

type Options struct {
	Attempt int
}

func Do(fun func() error, opt Options) error {

	var err error
	for i := 1; i <= opt.Attempt; i++ {
		err = fun()
		if err == nil {
			return nil
		}
	}
	return err
}
