package api

import "cowin/pkg/contracts"

type Notify interface {
	Notify(centers []contracts.Center, district int)
	Add(filter contracts.Filter) error
	Remove(filterID string, district int) error
}
