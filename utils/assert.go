package utils

// Assert panic if err!=nil
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}
