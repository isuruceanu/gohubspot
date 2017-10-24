package gohubspot

import (
	"fmt"
	"gohubspot/contactproperty"
)

type ContactPropertiesService service

type ContactPropertyOption struct {
	Description  string      `json:"description"`
	Label        string      `json:"label"`
	DisplayOrder int         `json:"displayOrder"`
	Hidden       bool        `json:"hidden"`
	Value        interface{} `json:"value"`
}

type ContactProperty struct {
	Name                          string                    `json:"name"`
	Label                         string                    `json:"label,omitempty"`
	Description                   string                    `json:"description,omitempty"`
	DataType                      contactproperty.DataType  `json:"type,omitempty"`
	FieldType                     contactproperty.FieldType `json:"fieldType,omitempty"`
	GroupName                     string                    `json:"groupName,omitempty"`
	Options                       []ContactPropertyOption   `json:"options"`
	Deleted                       bool                      `json:"deleted"`
	FormField                     bool                      `json:"formField"`
	DisplayOrder                  int                       `json:"displayOrder"`
	ReadOnlyValue                 bool                      `json:"readOnlyValue"`
	ReadOnlyDefinition            bool                      `json:"readOnlyDefinition"`
	Hidden                        bool                      `json:"hidden"`
	MutableDefinitionNotDeletable bool                      `json:"mutableDefinitionNotDeletable"`
	Calculated                    bool                      `json:"calculated"`
	ExternalOptions               bool                      `json:"externalOptions"`
}

type ContactProperties []ContactProperty

type ContactPropertyGroup struct {
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	DisplayOrder   int    `json:"displayOrder"`
	HubspotDefined bool   `json:"hubspotDefined"`
}

type ContactPropertyGroups []ContactPropertyGroup

func (s *ContactPropertiesService) GetAll() (*ContactProperties, error) {
	url := "/properties/v1/contacts/properties"
	res := new(ContactProperties)
	err := s.client.RunGet(url, res)

	return res, err
}

func (s *ContactPropertiesService) GetByName(name string) (*ContactProperty, error) {
	url := fmt.Sprintf("/properties/v1/contacts/properties/named/%s", name)
	res := new(ContactProperty)
	err := s.client.RunGet(url, res)
	return res, err
}

func (s *ContactPropertiesService) Create(property ContactProperty) (*ContactProperty, error) {
	url := "/properties/v1/contacts/properties"
	res := new(ContactProperty)
	if property.Options == nil {
		property.Options = []ContactPropertyOption{}
	}
	err := s.client.RunPost(url, property, res)
	return res, err
}

func (s *ContactPropertiesService) Update(name string, update ContactProperty) (*ContactProperty, error) {
	url := fmt.Sprintf("/properties/v1/contacts/properties/named/%s", name)
	res := new(ContactProperty)
	err := s.client.RunPut(url, update, res)
	return res, err
}

func (s *ContactPropertiesService) Delete(name string) error {
	url := fmt.Sprintf("/properties/v1/contacts/properties/named/%s", name)
	err := s.client.RunDelete(url, nil)
	return err
}

func (s *ContactPropertiesService) GetGroups() (*ContactPropertyGroups, error) {
	url := "/properties/v1/contacts/groups"
	res := new(ContactPropertyGroups)
	err := s.client.RunGet(url, res)
	return res, err
}

func (s *ContactPropertiesService) CreateGroup(group ContactPropertyGroup) (*ContactPropertyGroup, error) {
	url := "/properties/v1/contacts/groups"
	res := new(ContactPropertyGroup)
	err := s.client.RunPost(url, group, res)
	return res, err
}

func (s *ContactPropertiesService) UpdateGroup(name string, update ContactPropertyGroup) (*ContactPropertyGroup, error) {
	url := fmt.Sprintf("/properties/v1/contacts/groups/named/%s", name)
	res := new(ContactPropertyGroup)
	err := s.client.RunPut(url, update, res)
	return res, err
}

func (s *ContactPropertiesService) DeleteGroup(name string) error {
	url := fmt.Sprintf("/properties/v1/contacts/groups/named/%s", name)
	err := s.client.RunDelete(url, nil)
	return err
}
