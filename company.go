package gohubspot

type CompaniesService service

type Company struct {
	PortalId   int               `json:"portalId"`
	CompanyId  int               `json:"companyId"`
	IsDeleted  bool              `json:"isDeleted"`
	Properties CompanyProperties `json:"properties"`
}

type CompanyProperties struct {
}
