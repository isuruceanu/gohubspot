package gohubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.hubapi.com"
)

type HubspotClient struct {
	authenticator Authenticator
	common        service
	client        *http.Client
	BaseURL       *url.URL
	UserAgent     string
	// Services used for talking to different parts of the API
	ContactLists      *ContactListsService
	Contacts          *ContactsService
	ContactProperties *ContactPropertiesService
	CompanyProperties *CompanyPropertiesService
	Companies         *CompaniesService
	Forms             *FormService
	Deals             *DealsService
}

type service struct {
	client *HubspotClient
}

func NewHubspotOAuthClient(token string) *HubspotClient {
	return NewHubspotClient(NewOAuth2(token))
}

func NewHubspotApiClient(apikey string) *HubspotClient {
	return NewHubspotClient(NewAPIKeyAuth(apikey))
}

func NewHubspotClient(auth Authenticator) *HubspotClient {
	r := &HubspotClient{authenticator: auth}
	url, e := url.Parse(defaultBaseURL)
	if e != nil {
		panic(e)
	}

	r.BaseURL = url
	r.client = http.DefaultClient
	r.common.client = r
	r.ContactLists = (*ContactListsService)(&r.common)
	r.Contacts = (*ContactsService)(&r.common)
	r.Deals = (*DealsService)(&r.common)
	r.ContactProperties = (*ContactPropertiesService)(&r.common)
	r.CompanyProperties = (*CompanyPropertiesService)(&r.common)
	r.Companies = (*CompaniesService)(&r.common)
	//r.Forms = (*FormService)(&r.common)
	r.Forms = &FormService{service: r.common}

	return r
}

// Post creates a new POST request
func (c *HubspotClient) Post(url string, body interface{}) (*http.Request, error) {
	return c.NewRequest(http.MethodPost, url, body)
}

// Get creates a new GET API request
func (c *HubspotClient) Get(url string) (*http.Request, error) {
	return c.NewRequest(http.MethodGet, url, nil)
}

func (c *HubspotClient) RunGet(url string, res interface{}) error {
	if req, err := c.Get(url); err != nil {
		return err
	} else {
		return c.Do(req, res)
	}
}

func (c *HubspotClient) RunPost(url string, body, res interface{}) error {
	if req, err := c.Post(url, body); err != nil {
		return err
	} else {
		return c.Do(req, res)
	}
}

func (c *HubspotClient) RunPut(url string, body, res interface{}) error {
	if req, err := c.NewRequest(http.MethodPut, url, body); err != nil {
		return err
	} else {
		return c.Do(req, res)
	}
}

func (c *HubspotClient) RunDelete(url string, res interface{}) error {
	if req, err := c.NewRequest(http.MethodDelete, url, nil); err != nil {
		return err
	} else {
		return c.Do(req, res)
	}
}

// NewRequest creates a new API request.
func (c *HubspotClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must NOT have a trailing slash, but %q does not", c.BaseURL)
	}

	if !strings.HasPrefix(urlStr, "/") {
		return nil, fmt.Errorf("urlStr must have begin with a slash, but %q does not", urlStr)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if e := json.NewEncoder(buf).Encode(body); e != nil {
			return nil, e
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// body, err = ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	log.Fatalf("ERROR: %s", err)
	// }

	// fmt.Printf("%s", body)

	if body != nil {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	if e := c.authenticator.Authenticate(req); e != nil {
		return nil, e
	}
	return req, nil
}

// Do sends an API request and returns the API response
func (c *HubspotClient) Do(req *http.Request, v interface{}) error {

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil //ignore EOF erros caused by empty response body
			}
		}
	}

	return nil
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("Server error %d \n can not read body: %v", r.StatusCode, err)
	}

	return fmt.Errorf("Server error %d \n Body: %v", r.StatusCode, string(data[:]))
}
