package gohubspot

type Associations struct {
	AssociatedVids       string `json:"associatedVids"`
	AssociatedCompanyIds string `json:"associatedCompanyIds"`
	AssociatedDealIds    string `json:"associatedDealIds"`
}

func (a *Associations) SetAssociation(value Associations) {
	a.AssociatedVids = value.AssociatedVids
	a.AssociatedCompanyIds = value.AssociatedCompanyIds
	a.AssociatedDealIds = value.AssociatedDealIds
}
