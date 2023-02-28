package ch7

type TypeAssertRunner int64
type TA struct {
}

func (TA) Write(b []byte) (n int, err error) {
	println("Write")
	return 0, nil
}

func (TA) Read(p []byte) (n int, err error) {
	println("Read")
	return 0, nil
}

func (tar TypeAssertRunner) Run() {
	//var w io.Writer
	//w = os.Stdout
	//f := w.(*os.File)
	//fmt.Printf("---> %t > %v\n", f == w, reflect.TypeOf(w))
	//c := w.(*bytes.Buffer)
	//println(c)

	//f, err := os.Open("/no/such/file")
	//fmt.Println(os.IsNotExist(err))
}
