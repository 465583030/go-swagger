// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package petstore

import (
	"io"
	gotest "testing"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/security"
	"github.com/go-swagger/go-swagger/httpkit/untyped"
	testingutil "github.com/go-swagger/go-swagger/internal/testing"
	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/toolkit"
	"github.com/stretchr/testify/assert"
)

// NewAPI registers a stub api for the pet store
func NewAPI(t *gotest.T) (*spec.Document, *untyped.API) {
	spec, err := spec.New(testingutil.PetStoreJSONMessage, "")
	assert.NoError(t, err)
	api := untyped.NewAPI(spec)

	api.RegisterConsumer("application/json", toolkit.JSONConsumer())
	api.RegisterProducer("application/json", toolkit.JSONProducer())
	api.RegisterConsumer("application/xml", new(stubConsumer))
	api.RegisterProducer("application/xml", new(stubProducer))
	api.RegisterProducer("text/plain", new(stubProducer))
	api.RegisterProducer("text/html", new(stubProducer))
	api.RegisterConsumer("application/x-yaml", toolkit.YAMLConsumer())
	api.RegisterProducer("application/x-yaml", toolkit.YAMLProducer())

	api.RegisterAuth("basic", security.BasicAuth(func(username, password string) (interface{}, error) {
		if username == "admin" && password == "admin" {
			return "admin", nil
		}
		return nil, errors.Unauthenticated("basic")
	}))
	api.RegisterAuth("apiKey", security.APIKeyAuth("X-API-KEY", "header", func(token string) (interface{}, error) {
		if token == "token123" {
			return "admin", nil
		}
		return nil, errors.Unauthenticated("token")
	}))
	api.RegisterOperation("get", "/pets", new(stubOperationHandler))
	api.RegisterOperation("post", "/pets", new(stubOperationHandler))
	api.RegisterOperation("delete", "/pets/{id}", new(stubOperationHandler))
	api.RegisterOperation("get", "/pets/{id}", new(stubOperationHandler))

	api.Models["pet"] = func() interface{} { return new(Pet) }
	api.Models["newPet"] = func() interface{} { return new(Pet) }
	api.Models["tag"] = func() interface{} { return new(Tag) }

	return spec, api
}

// NewRootAPI registers a stub api for the pet store
func NewRootAPI(t *gotest.T) (*spec.Document, *untyped.API) {
	spec, err := spec.New(testingutil.RootPetStoreJSONMessage, "")
	assert.NoError(t, err)
	api := untyped.NewAPI(spec)

	api.RegisterConsumer("application/json", toolkit.JSONConsumer())
	api.RegisterProducer("application/json", toolkit.JSONProducer())
	api.RegisterConsumer("application/xml", new(stubConsumer))
	api.RegisterProducer("application/xml", new(stubProducer))
	api.RegisterProducer("text/plain", new(stubProducer))
	api.RegisterProducer("text/html", new(stubProducer))
	api.RegisterConsumer("application/x-yaml", toolkit.YAMLConsumer())
	api.RegisterProducer("application/x-yaml", toolkit.YAMLProducer())

	api.RegisterAuth("basic", security.BasicAuth(func(username, password string) (interface{}, error) {
		if username == "admin" && password == "admin" {
			return "admin", nil
		}
		return nil, errors.Unauthenticated("basic")
	}))
	api.RegisterAuth("apiKey", security.APIKeyAuth("X-API-KEY", "header", func(token string) (interface{}, error) {
		if token == "token123" {
			return "admin", nil
		}
		return nil, errors.Unauthenticated("token")
	}))
	api.RegisterOperation("get", "/pets", new(stubOperationHandler))
	api.RegisterOperation("post", "/pets", new(stubOperationHandler))
	api.RegisterOperation("delete", "/pets/{id}", new(stubOperationHandler))
	api.RegisterOperation("get", "/pets/{id}", new(stubOperationHandler))

	api.Models["pet"] = func() interface{} { return new(Pet) }
	api.Models["newPet"] = func() interface{} { return new(Pet) }
	api.Models["tag"] = func() interface{} { return new(Tag) }

	return spec, api
}

// Tag the tag model
type Tag struct {
	ID   int64
	Name string
}

// Pet the pet model
type Pet struct {
	ID        int64
	Name      string
	PhotoURLs []string
	Status    string
	Tags      []Tag
}

type stubConsumer struct {
}

func (s *stubConsumer) Consume(_ io.Reader, _ interface{}) error {
	return nil
}

type stubProducer struct {
}

func (s *stubProducer) Produce(_ io.Writer, _ interface{}) error {
	return nil
}

type stubOperationHandler struct {
}

func (s *stubOperationHandler) ParameterModel() interface{} {
	return nil
}

func (s *stubOperationHandler) Handle(params interface{}) (interface{}, error) {
	return nil, nil
}
