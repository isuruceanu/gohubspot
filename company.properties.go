package gohubspot

import "fmt"

// CompanyPropertiesService Company Properties Service definition
type CompanyPropertiesService service

// GetAll gets all company properties
func (s *CompanyPropertiesService) GetAll() (*ItemProperties, error) {
	url := "/properties/v1/companies/properties/"
	res := new(ItemProperties)
	err := s.client.RunGet(url, res)

	return res, err
}

// GetByName gets company property by name
func (s *CompanyPropertiesService) GetByName(name string) (*ItemProperty, error) {
	url := fmt.Sprintf("/properties/v1/companies/properties/named/%s", name)
	res := new(ItemProperty)
	err := s.client.RunGet(url, res)
	return res, err
}

// Create create a contact property
func (s *CompanyPropertiesService) Create(property ItemProperty) (*ItemProperty, error) {
	url := "/properties/v1/companies/properties/"
	res := new(ItemProperty)
	if property.Options == nil {
		property.Options = []ItemPropertyOption{}
	}
	err := s.client.RunPost(url, property, res)
	return res, err
}

// Update updates an existing property by it name
func (s *CompanyPropertiesService) Update(name string, update ItemProperty) (*ItemProperty, error) {
	url := fmt.Sprintf("/properties/v1/companies/properties/named/%s", name)
	res := new(ItemProperty)
	err := s.client.RunPut(url, update, res)
	return res, err
}

// Delete deletes a contact property by name
func (s *CompanyPropertiesService) Delete(name string) error {
	url := fmt.Sprintf("/properties/v1/companies/properties/named/%s", name)
	err := s.client.RunDelete(url, nil)
	return err
}

// GetGroups gets all contacts property groups
func (s *CompanyPropertiesService) GetGroups() (*ItemPropertyGroups, error) {
	url := "/properties/v1/companies/groups"
	res := new(ItemPropertyGroups)
	err := s.client.RunGet(url, res)
	return res, err
}

// CreateGroup creates a contact property group
func (s *CompanyPropertiesService) CreateGroup(group ItemPropertyGroup) (*ItemPropertyGroup, error) {
	url := "/properties/v1/companies/groups"
	res := new(ItemPropertyGroup)
	err := s.client.RunPost(url, group, res)
	return res, err
}

// UpdateGroup updates contact property group
func (s *CompanyPropertiesService) UpdateGroup(name string, update ItemPropertyGroup) (*ItemPropertyGroup, error) {
	url := fmt.Sprintf("/properties/v1/companies/groups/named/%s", name)
	res := new(ItemPropertyGroup)
	err := s.client.RunPut(url, update, res)
	return res, err
}

// DeleteGroup deletes a contact property group
func (s *CompanyPropertiesService) DeleteGroup(name string) error {
	url := fmt.Sprintf("/properties/v1/companies/groups/named/%s", name)
	err := s.client.RunDelete(url, nil)
	return err
}
