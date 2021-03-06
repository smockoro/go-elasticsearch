// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// CountOption is a non-required Count option that gets applied to an HTTP request.
type CountOption func(r *transport.Request)

// WithCountIndex - a comma-separated list of index names to limit the returned information.
func WithCountIndex(index []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountFormat - a short version of the Accept header, e.g. json, yaml.
func WithCountFormat(format string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountH - comma-separated list of column names to display.
func WithCountH(h []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountHelp - return help information.
func WithCountHelp(help bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountLocal - return local information, do not retrieve the state from master node (default: false).
func WithCountLocal(local bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountMasterTimeout - explicit operation timeout for connection to master node.
func WithCountMasterTimeout(masterTimeout time.Duration) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountS - comma-separated list of column names or column aliases to sort by.
func WithCountS(s []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountV - verbose mode. Display column headers.
func WithCountV(v bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountErrorTrace - include the stack trace of returned errors.
func WithCountErrorTrace(errorTrace bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountFilterPath - a comma-separated list of filters used to reduce the respone.
func WithCountFilterPath(filterPath []string) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountHuman - return human readable values for statistics.
func WithCountHuman(human bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountIgnore - ignores the specified HTTP status codes.
func WithCountIgnore(ignore []int) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountPretty - pretty format the returned JSON response.
func WithCountPretty(pretty bool) CountOption {
	return func(r *transport.Request) {
	}
}

// WithCountSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithCountSourceParam(sourceParam string) CountOption {
	return func(r *transport.Request) {
	}
}

// Count - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-count.html for more info.
//
// options: optional parameters.
func (c *Cat) Count(options ...CountOption) (*CountResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &CountResponse{resp}, err
}

// CountResponse is the response for Count.
type CountResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *CountResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
