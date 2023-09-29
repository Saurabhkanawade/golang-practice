package model

type Employee struct {
	EmpId      string      `json="id"`
	Lastname   string      `json="firstname"`
	Firstname  string      `json="lastname"`
	Department *Department `json:"department"`
}

type Department struct {
	DeptId   string `json="dept_id"`
	Deptname string `json="dept_name"`
}
