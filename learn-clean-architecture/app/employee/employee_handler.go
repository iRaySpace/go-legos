package employee

type EmployeeHandler struct {
	service *EmployeeService
}

func InitHandler() *EmployeeHandler {
	return &EmployeeHandler{
		service: &EmployeeService{
			employeeRepository: &LocalEmployeeRepository{},
		},
	}
}

func (h *EmployeeHandler) Create(e Employee) (Employee, error) {
	return h.service.Create(e)
}

func (h *EmployeeHandler) GetAll() []Employee {
	return h.service.GetAll()
}

func (h *EmployeeHandler) DeleteByID(id string) error {
	return h.service.DeleteByID(id)
}
