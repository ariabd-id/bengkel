package customer

import "bengkel/domain"

type Service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &Service{
		customerRepository: customerRepository,
	}
}
