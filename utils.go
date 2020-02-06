package main

func ifError(err error) error {
	if err != nil {
		return err
	} else {
		return nil
	}
}
