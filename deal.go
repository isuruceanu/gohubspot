package gohubspot

import "fmt"

type DealsService service

type Deal struct {
	Properties   Properties   `json:"properties"`
	Associations Associations `json:"associations"`
}

type DealResponse struct {
	PortalId     int          `json:"portalId"`
	DealId       int          `json:"dealId"`
	IsDeleted    bool         `json:"isDeleted"`
	Properties   Properties   `json:"properties"`
	Associations Associations `json:"associations"`
}

func (s *DealsService) Create(deal Deal) (*DealResponse, error) {
	url := "/deals/v1/deal"
	res := new(DealResponse)
	err := s.client.RunPost(url, deal, res)
	return res, err
}

func (s *DealsService) Update(dealId int, deal Deal) error {
	url := fmt.Sprintf("/deals/v1/deal/%d", dealId)
	res := new(DealResponse)
	return s.client.RunPost(url, deal, res)
}

func (s *DealsService) associateDealToContact(dealId int, contactId int, deal Deal) error {
	url := fmt.Sprintf("/deals/v1/deal/%d/associations/CONTACT?id=%d", dealId, contactId)
	return s.client.RunPost(url, deal, nil)
}
