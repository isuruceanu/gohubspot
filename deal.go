package gohubspot

import "fmt"

type DealsService service

type Deal struct {
	Properties   Properties `json:"properties"`
	Associations Properties `json:"associations"`
}

func (s *DealsService) Create(properties Properties) (*IdentityProfile, error) {
	url := "/deals/v1/deal"
	res := new(IdentityProfile)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *DealsService) Update(dealId int, properties Properties) error {
	url := fmt.Sprintf("/deals/v1/deal/%d", dealId)
	return s.client.RunPost(url, properties, nil)
}

func (s *DealsService) associateDealToContact(dealId int, contactId int, properties Properties) error {
	url := fmt.Sprintf("/deals/v1/deal/%d/associations/CONTACT?id=%d", dealId, contactId)
	return s.client.RunPost(url, properties, nil)
}