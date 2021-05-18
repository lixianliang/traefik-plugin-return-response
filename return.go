package traefik_plugin_return_response

import (
	"context"
	"log"
	"net/http"
	"os"
)

// Config the plugin configuration
type Config struct {
	Code    int               `json:code,omitempty`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}

// CreateConfig creates the deafult plugin configuration
func CreateConfig() *Config {
	return &Config{
		Code:    200,
		Headers: map[string]string{},
		Body:    "",
	}
}

// Demo a Demo plugin
type returnResponse struct {
	next    http.Handler
	name    string
	Code    int
	Headers map[string]string
	Body    string
	logger  log.Logger
}

// New created a new plugin
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	/*if config.Body == "" {
		return nil, fmt.Errorf("config Body is nil")
	}*/
	return &returnResponse{
		next:    next,
		name:    name,
		Code:    config.Code,
		Headers: config.Headers,
		Body:    config.Body,
		logger:  *log.New(os.Stdout, "plugin:returnResponse ", log.Ldate|log.Ltime),
	}, nil
}

func (r *returnResponse) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Printf("return response code:%d headers:%v body: %s", r.Code, r.Headers, r.Body)
	rw.WriteHeader(r.Code)
	rw.Write([]byte(r.Body))
}
