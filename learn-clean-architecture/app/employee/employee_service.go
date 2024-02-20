package employee

type EmployeeService struct {
	employeeRepository EmployeeRepository
}

func (s *EmployeeService) Create(e Employee) (Employee, error) {
	return s.employeeRepository.Save(e), nil
}

func (s *EmployeeService) GetAll() []Employee {
	return s.employeeRepository.GetAll()
}

func (s *EmployeeService) DeleteByID(id string) error {
	return s.employeeRepository.DeleteByID(id)
}
