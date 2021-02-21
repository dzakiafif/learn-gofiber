package config

func PanicHandler(err interface{}) {
	if err != nil {
		panic(err)
	}
}