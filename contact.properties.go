package gohubspot

import (
	"fmt"
)

// ContactPropertiesService - Contact properties are used to store specific information for each of your contact records
type ContactPropertiesService service

// GetAll gets all contacts properties
func (s *ContactPropertiesService) GetAll() (*ItemProperties, error) {
	url := "/properties/v1/contacts/properties"
	res := new(ItemProperties)
	err := s.client.RunGet(url, res)

	return res, err
}

// GetByName gets property by name
func (s *ContactPropertiesService) GetByName(name string) (*ItemProperty, error) {
	url := fmt.Sprintf("/properties/v1/contacts/properties/named/%s", name)
	res := new(ItemProperty)
	err := s.client.RunGet(url, res)
	return res, err
}

// Create create a contact property
func (s *ContactPropertiesService) Create(property ItemProperty) (*ItemProperty, error) {
	url := "/properties/v1/contacts/properties"
	res := new(ItemProperty)
	if property.Options == nil {
		property.Options = []ItemPropertyOption{}
	}
	err := s.client.RunPost(url, property, res)
	return res, err
}

// Update updates an existing property by it name
func (s *ContactPropertiesService) Update(name string, update ItemProperty) (*ItemProperty, error) {
	url := fmt.Sprintf("/properties/v1/contacts/properties/named/%s", name)
	res := new(ItemProperty)
	err := s.client.RunPut(url, update, res)
	return res, err
}

// Delete deletes a contact property by name
func (s *ContactPropertiesService) Delete(name string) error {
	url := fmt.Sprintf("/properties/v1/contacts/properties/named/%s", name)
	err := s.client.RunDelete(url, nil)
	return err
}

// GetGroups gets all contacts property groups
func (s *ContactPropertiesService) GetGroups() (*ItemPropertyGroups, error) {
	url := "/properties/v1/contacts/groups"
	res := new(ItemPropertyGroups)
	err := s.client.RunGet(url, res)
	return res, err
}

// CreateGroup creates a contact property group
func (s *ContactPropertiesService) CreateGroup(group ItemPropertyGroup) (*ItemPropertyGroup, error) {
	url := "/properties/v1/contacts/groups"
	res := new(ItemPropertyGroup)
	err := s.client.RunPost(url, group, res)
	return res, err
}

// UpdateGroup updates contact property group
func (s *ContactPropertiesService) UpdateGroup(name string, update ItemPropertyGroup) (*ItemPropertyGroup, error) {
	url := fmt.Sprintf("/properties/v1/contacts/groups/named/%s", name)
	res := new(ItemPropertyGroup)
	err := s.client.RunPut(url, update, res)
	return res, err
}

// DeleteGroup deletes a contact property group
func (s *ContactPropertiesService) DeleteGroup(name string) error {
	url := fmt.Sprintf("/properties/v1/contacts/groups/named/%s", name)
	err := s.client.RunDelete(url, nil)
	return err
}
