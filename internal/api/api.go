package api

//go:generate $GOPATH/bin/mockgen -source=api.go -destination ../mocks/api.go -package mocks -self_package github.com/khulnasoft-lab/go-application-framework/pkg/api/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/khulnasoft-lab/go-application-framework/internal/api/contract"
	"github.com/khulnasoft-lab/go-application-framework/internal/constants"
)

type ApiClient interface {
	GetDefaultOrgId() (orgID string, err error)
	GetOrgIdFromSlug(slugName string) (string, error)
	Init(url string, client *http.Client)
}

type vulnmapApiClient struct {
	url    string
	client *http.Client
}

func (a *vulnmapApiClient) GetOrgIdFromSlug(slugName string) (string, error) {
	url := a.url + "/v1/orgs"
	res, err := a.client.Get(url)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var userOrgInfo contract.OrganizationsResponse
	err = json.Unmarshal(body, &userOrgInfo)
	if err != nil {
		return "", err
	}

	for _, org := range userOrgInfo.Organizations {
		if org.Slug == slugName {
			return org.ID, nil
		}
	}

	return "", fmt.Errorf("org ID not found for slug %v", slugName)
}

func (a *vulnmapApiClient) GetDefaultOrgId() (string, error) {
	url := a.url + "/rest/self?version=" + constants.VULNMAP_API_VERSION
	res, err := a.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve org ID: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("unable to retrieve org ID: %w", err)
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("unable to retrieve org ID (status: %d)", res.StatusCode)
	}

	var userInfo contract.SelfResponse
	if err = json.Unmarshal(body, &userInfo); err != nil {
		return "", fmt.Errorf("unable to retrieve org ID (status: %d): %w", res.StatusCode, err)
	}

	return userInfo.Data.Attributes.DefaultOrgContext, nil
}

func (a *vulnmapApiClient) Init(url string, client *http.Client) {
	a.url = url
	a.client = client
}

func NewApi(url string, httpClient *http.Client) ApiClient {
	client := NewApiInstance()
	client.Init(url, httpClient)
	return client
}

func NewApiInstance() ApiClient {
	return &vulnmapApiClient{}
}
