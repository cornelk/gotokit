package httpclient

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProxyTransportFromConfig_Success(t *testing.T) {
	tests := []struct {
		name        string
		proxyConfig string
		wantScheme  string
	}{
		{
			name:        "HTTP proxy",
			proxyConfig: "http://proxy.example.com:8080",
			wantScheme:  "http",
		},
		{
			name:        "HTTPS proxy",
			proxyConfig: "https://proxy.example.com:8080",
			wantScheme:  "https",
		},
		{
			name:        "SOCKS5 proxy",
			proxyConfig: "socks5://proxy.example.com:1080",
			wantScheme:  "socks5",
		},
		{
			name:        "TCP proxy (SOCKS5 alias)",
			proxyConfig: "tcp://proxy.example.com:1080",
			wantScheme:  "tcp",
		},
		{
			name:        "HTTP proxy with authentication",
			proxyConfig: "http://user:pass@proxy.example.com:8080",
			wantScheme:  "http",
		},
		{
			name:        "SOCKS5 proxy with authentication",
			proxyConfig: "socks5://user:pass@proxy.example.com:1080",
			wantScheme:  "socks5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport, err := ProxyTransportFromConfig(tt.proxyConfig)
			require.NoError(t, err)
			require.NotNil(t, transport)

			// Verify transport is properly configured
			assert.IsType(t, &http.Transport{}, transport)

			// For HTTP proxies, verify proxy function is set
			if tt.wantScheme == "http" || tt.wantScheme == "https" {
				assert.NotNil(t, transport.Proxy)

				// Test proxy function with a dummy request
				req := &http.Request{URL: &url.URL{Host: "example.com"}}
				proxyURL, err := transport.Proxy(req)
				require.NoError(t, err)
				assert.Equal(t, tt.wantScheme, proxyURL.Scheme)
			}

			// For SOCKS5 proxies, verify dial context is set
			if tt.wantScheme == "socks5" || tt.wantScheme == "tcp" {
				assert.NotNil(t, transport.DialContext)
				assert.True(t, transport.DisableKeepAlives)
			}
		})
	}
}

func TestProxyTransportFromConfig_EmptyConfig(t *testing.T) {
	// Empty config should return default transport
	transport, err := ProxyTransportFromConfig("")
	require.NoError(t, err)
	require.NotNil(t, transport)

	// Should be a clone of default transport
	defaultTransport := http.DefaultTransport.(*http.Transport)
	assert.Equal(t, defaultTransport.MaxIdleConns, transport.MaxIdleConns)
	assert.Equal(t, defaultTransport.IdleConnTimeout, transport.IdleConnTimeout)

	// Should be different instances (clone, not same object)
	assert.NotSame(t, defaultTransport, transport)
}

func TestProxyTransportFromConfig_Errors(t *testing.T) {
	tests := []struct {
		name        string
		proxyConfig string
		wantErr     error
	}{
		{
			name:        "invalid URL",
			proxyConfig: "://invalid-url",
			wantErr:     ErrInvalidProxyURL,
		},
		{
			name:        "missing host",
			proxyConfig: "http://",
			wantErr:     ErrInvalidProxyURL,
		},
		{
			name:        "unsupported scheme",
			proxyConfig: "ftp://proxy.example.com:21",
			wantErr:     ErrUnsupportedScheme,
		},
		{
			name:        "unsupported scheme with auth",
			proxyConfig: "ldap://user:pass@proxy.example.com:389",
			wantErr:     ErrUnsupportedScheme,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport, err := ProxyTransportFromConfig(tt.proxyConfig)
			assert.Nil(t, transport)
			require.Error(t, err)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestProxyTransportFromConfig_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		proxyConfig string
		expectError bool
	}{
		{
			name:        "port only",
			proxyConfig: "http://:8080",
			expectError: false, // ":8080" is a valid host in Go
		},
		{
			name:        "IPv6 host",
			proxyConfig: "http://[::1]:8080",
			expectError: false,
		},
		{
			name:        "uppercase scheme",
			proxyConfig: "HTTP://proxy.example.com:8080",
			expectError: false,
		},
		{
			name:        "mixed case scheme",
			proxyConfig: "SoCkS5://proxy.example.com:1080",
			expectError: false,
		},
		{
			name:        "user without password",
			proxyConfig: "http://user@proxy.example.com:8080",
			expectError: false,
		},
		{
			name:        "password with special characters",
			proxyConfig: "http://user:p%40ssw0rd@proxy.example.com:8080",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transport, err := ProxyTransportFromConfig(tt.proxyConfig)

			if tt.expectError {
				require.Error(t, err)
				assert.Nil(t, transport)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, transport)
			}
		})
	}
}

func TestCreateSOCKS5Transport_AuthHandling(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		hasAuth  bool
		username string
		password string
	}{
		{
			name:    "no authentication",
			url:     "socks5://proxy.example.com:1080",
			hasAuth: false,
		},
		{
			name:     "with username and password",
			url:      "socks5://user:pass@proxy.example.com:1080",
			hasAuth:  true,
			username: "user",
			password: "pass",
		},
		{
			name:     "with username only",
			url:      "socks5://user@proxy.example.com:1080",
			hasAuth:  true,
			username: "user",
			password: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proxyURL, err := url.Parse(tt.url)
			require.NoError(t, err)

			transport, err := createSOCKS5Transport(proxyURL)
			require.NoError(t, err)
			require.NotNil(t, transport)

			// Verify transport configuration
			assert.NotNil(t, transport.DialContext)
			assert.True(t, transport.DisableKeepAlives)
		})
	}
}

func TestCreateHTTPProxyTransport(t *testing.T) {
	proxyURL, err := url.Parse("http://proxy.example.com:8080")
	require.NoError(t, err)

	transport, err := createHTTPProxyTransport(proxyURL)
	require.NoError(t, err)
	require.NotNil(t, transport)

	// Verify proxy function is set
	assert.NotNil(t, transport.Proxy)

	// Test proxy function
	req := &http.Request{URL: &url.URL{Host: "example.com"}}
	resultURL, err := transport.Proxy(req)
	require.NoError(t, err)
	assert.Equal(t, "http", resultURL.Scheme)
	assert.Equal(t, "proxy.example.com:8080", resultURL.Host)
}
