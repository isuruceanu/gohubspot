package gohubspot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// FormService The Forms API's principle method is the submit form method,
// which allows you to pass any information captured on your website or application to HubSpot,
// including any custom data.
// This endpoint doesn't require authentication, so you can make the call from a form to our API
// from the client without needing to worry about insecurity.
type FormService struct {
	service
	hubspotutk string
	options    *url.Values
	hsContext  *HsContext
}

type HsContext struct {
	Hutk        string `json:"hutk"`
	IPAddress   string `json:"ipAddress"`
	PageURL     string `json:"pageUrl"`
	PageName    string `json:"pageName"`
	RedirectURL string `json:"redirectUrl"`
}

func (s *FormService) AddOption(key, value string) *FormService {
	if s.options == nil {
		s.options = &url.Values{}
	}
	s.options.Add(key, value)
	return s
}

func (s *FormService) AddOptions(data url.Values) *FormService {
	s.options = &data
	return s
}

func (s *FormService) SetHubspotCookie(cookie string) *FormService {
	if cookie == "" {
		return s
	}
	s.checkHsContext()
	s.hsContext.Hutk = cookie
	return s
}

func (s *FormService) SetRemoteIpAddress(url string) *FormService {
	if url == "" {
		return s
	}

	s.checkHsContext()
	s.hsContext.IPAddress = url
	return s
}

func (s *FormService) SetPageUrl(url string) *FormService {

	if url == "" {
		return s
	}

	s.checkHsContext()
	s.hsContext.PageURL = url
	return s
}

func (s *FormService) SetPageName(name string) *FormService {
	if name == "" {
		return s
	}

	s.checkHsContext()
	s.hsContext.PageName = name
	return s
}

func (s *FormService) SetReturnUrl(url string) *FormService {
	if url == "" {
		return s
	}

	s.checkHsContext()
	s.hsContext.RedirectURL = url
	return s
}

func (s *FormService) SetHsContext(context HsContext) *FormService {
	s.hsContext = &context
	return s
}

func (s *FormService) SubmitForm(portalID int, formID string) error {
	url := fmt.Sprintf("https://forms.hubspot.com/uploads/form/v2/%d/%s", portalID, formID)

	req, err := s.newFormRequest(url, s.getBody())
	if err != nil {
		return err
	}

	DumpRequest(req, true)

	return s.client.Do(req, nil)
}

func (s *FormService) checkHsContext() {
	if s.hsContext == nil {
		s.hsContext = &HsContext{}
	}
}

func (s *FormService) newFormRequest(url string, body string) (*http.Request, error) {
	u, err := s.client.BaseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	var buf io.Reader
	if len(body) > 0 {
		buf = strings.NewReader(body)
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if s.client.UserAgent != "" {
		req.Header.Set("User-Agent", s.client.UserAgent)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Content-Length", strconv.Itoa(len(body)))

	return req, nil
}

func (s *FormService) getBody() string {

	if s.hsContext != nil {
		hsBody, _ := json.Marshal(s.hsContext)
		s.AddOption("hs_context", string(hsBody))
	}

	if s.options != nil {
		body := s.options.Encode()
		//fmt.Println(body)
		return body
	}

	return ""
}
