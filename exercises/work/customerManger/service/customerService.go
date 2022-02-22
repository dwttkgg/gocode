package service

import (
	"gocode/exercises/work/customerManger/model"
)

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

// 返回一个 *CustomerService
func NewCuntomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customers := model.NewCustomer(1, "zhangsan", "nan", 22, "1111111", "qq@c.com")
	customerService.customers = append(customerService.customers, customers)
	return customerService
}

// 返回客户切片
func (this *CustomerService) ListCustomer() []model.Customer {
	return this.customers
}

// 增加新的客户
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}
