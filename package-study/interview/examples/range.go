package main

type student struct {
	Name string
	Age  int
}

func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// error
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}
	println()

	// correct
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

func main() {
	paseStudent()
}
