## 读文件

#方式一：
func f1() {
	fileObj, err := os.Open("main.go")
	if err != nil {
		fmt.Println("file open fail")
		return
	}
	// var by [1024]byte
	for {
		by := make([]byte, 1024)
		bylen, err := fileObj.Read(by)
		if err != nil && err != io.EOF {
			fmt.Println("read file fail")
			return
		}
		if bylen < 1024 {
			fmt.Print(string(by))
			break
		}
		fmt.Print(string(by))

		defer fileObj.Close()
	}
	fileObj, err := os.Open("main.go")
	if err != nil {
		fmt.Println("read failed")
	}
}

#方式二：
func f2(){
	buf := bufio.NewReader(fileObj)
	for {
		by := make([]byte, 1024)
		bylen, err := buf.Read(by)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		if bylen < 1024 {
			fmt.Println(string(by))
			break
		}
	}
	defer fileObj.Close()
}

#方式三
func f3(){
	bys, err := ioutil.ReadFile("main.go")
	if err != nil {
		fmt.Println("read failed")
	}
	fmt.Println(string(bys))
}
