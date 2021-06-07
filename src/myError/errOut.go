package myError

func ErrorOut(err error) {
	if err != nil {
		panic(err)
	}
}
