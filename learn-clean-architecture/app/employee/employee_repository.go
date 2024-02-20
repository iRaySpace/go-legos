package employee

import (
	"errors"

	"github.com/google/uuid"
)

type (
	EmployeeRepository interface {
		GetAll() []Employee
		Save(e Employee) Employee
		DeleteByID(id string) error
	}
	LocalEmployeeRepository struct {
		data []Employee
	}
)

func (r *LocalEmployeeRepository) GetAll() []Employee {
	return r.data
}

func (r *LocalEmployeeRepository) Save(e Employee) Employee {
	e.ID = uuid.NewString()
	r.data = append(r.data, e)
	return e
}

func (r *LocalEmployeeRepository) DeleteByID(id string) error {
	idx := -1
	for i, e := range r.data {
		if e.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("not found")
	}
	r.data = append(r.data[:idx], r.data[idx+1:]...)
	return nil
}
