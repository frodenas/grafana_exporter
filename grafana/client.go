package grafana

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/common/version"
)

type Client struct {
	url        *url.URL
	username   string
	password   string
	httpClient *http.Client
}

func NewClient(uri string, username string, password string, skipSSLVerify bool) (*Client, error) {
	grafanaURL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipSSLVerify,
		},
	}
	httpClient := &http.Client{
		Timeout:   time.Duration(10 * time.Second),
		Transport: transport,
	}
	grafanaClient := &Client{
		url:        grafanaURL,
		username:   username,
		password:   password,
		httpClient: httpClient,
	}

	return grafanaClient, nil
}

func (c *Client) GetAdminStats() (AdminStats, error) {
	var adminStats AdminStats

	uri := c.url
	uri.Path = "/api/admin/stats"
	request, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return adminStats, err
	}
	request.Header.Set("User-Agent", "grafana_exporter "+version.Version)
	if c.username != "" && c.password != "" {
		request.SetBasicAuth(c.username, c.password)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return adminStats, errors.New(fmt.Sprintf("Error getting admin stats: %s", err))
	}

	if response.StatusCode != http.StatusOK {
		return adminStats, errors.New(fmt.Sprintf("Error getting admin stats, http status code: %d", response.StatusCode))
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return adminStats, errors.New(fmt.Sprintf("Error reading admin stats response: %s", err))
	}
	defer response.Body.Close()

	if err := json.Unmarshal(responseBody, &adminStats); err != nil {
		return adminStats, errors.New(fmt.Sprintf("Error unmarshalling admin stats response: %s", err))
	}

	return adminStats, nil
}

func (c *Client) GetMetrics() (Metrics, error) {
	var metrics Metrics

	uri := c.url
	uri.Path = "/api/metrics"
	request, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return metrics, err
	}
	request.Header.Set("User-Agent", "grafana_exporter "+version.Version)
	if c.username != "" && c.password != "" {
		request.SetBasicAuth(c.username, c.password)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return metrics, errors.New(fmt.Sprintf("Error getting metrics: %s", err))
	}

	if response.StatusCode != http.StatusOK {
		return metrics, errors.New(fmt.Sprintf("Error getting metrics, http status code: %d", response.StatusCode))
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return metrics, errors.New(fmt.Sprintf("Error reading metrics response: %s", err))
	}
	defer response.Body.Close()

	if err := json.Unmarshal(responseBody, &metrics); err != nil {
		return metrics, errors.New(fmt.Sprintf("Error unmarshalling metrics response: %s", err))
	}

	return metrics, nil
}
