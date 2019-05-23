package main

func main() {
	// demo()
	beegoOrm()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
