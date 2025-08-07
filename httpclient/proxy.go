// Package httpclient provides HTTP client utilities and proxy configuration support.
package httpclient

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

var (
	// ErrInvalidProxyURL is returned when the proxy URL cannot be parsed.
	ErrInvalidProxyURL = errors.New("invalid proxy URL")
	// ErrUnsupportedScheme is returned when the proxy scheme is not supported.
	ErrUnsupportedScheme = errors.New("unsupported proxy scheme")
)

// ProxyTransportFromConfig creates an HTTP transport configured with the specified proxy.
// The proxyConfig parameter should be a URL in the format: scheme://[user:password@]host:port
// If proxyConfig is empty, returns a copy of the default HTTP transport for direct connections.
// Supported schemes: http, https, socks5, tcp (alias for socks5).
func ProxyTransportFromConfig(proxyConfig string) (*http.Transport, error) {
	if proxyConfig == "" {
		// Return a copy of the default transport for direct connections
		return http.DefaultTransport.(*http.Transport).Clone(), nil
	}

	proxyURL, err := url.Parse(proxyConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidProxyURL, err)
	}

	if proxyURL.Host == "" {
		return nil, fmt.Errorf("%w: missing host", ErrInvalidProxyURL)
	}

	scheme := strings.ToLower(proxyURL.Scheme)
	switch scheme {
	case "tcp", "socks5":
		return createSOCKS5Transport(proxyURL)
	case "http", "https":
		return createHTTPProxyTransport(proxyURL)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedScheme, scheme)
	}
}

// createSOCKS5Transport creates an HTTP transport using SOCKS5 proxy.
func createSOCKS5Transport(proxyURL *url.URL) (*http.Transport, error) {
	var auth *proxy.Auth
	if proxyURL.User != nil {
		password, _ := proxyURL.User.Password()
		auth = &proxy.Auth{
			User:     proxyURL.User.Username(),
			Password: password,
		}
	}

	dialer, err := proxy.SOCKS5("tcp", proxyURL.Host, auth, proxy.Direct)
	if err != nil {
		return nil, fmt.Errorf("creating socks proxy client: %w", err)
	}

	dialContext := func(_ context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}

	return &http.Transport{
		DialContext:       dialContext,
		DisableKeepAlives: true,
	}, nil
}

// createHTTPProxyTransport creates an HTTP transport using HTTP proxy.
func createHTTPProxyTransport(proxyURL *url.URL) (*http.Transport, error) {
	return &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}, nil
}
