// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	SpecTableName = "specs"
)

// Spec represents a spec
type Spec struct {
	ID        string    `json:"id,omitempty" gorm:"primaryKey"`
	Doc       SpecDoc   `json:"doc" gorm:"column:doc"`
	DocType   string    `json:"doc_type" gorm:"column:doc_type"`
	Revision  string    `json:"revision" gorm:"column:revision;index"`
	Score     *int      `json:"score" gorm:"column:score"`
	ServiceID string    `json:"service_id" gorm:"column:service_id;index"`
	State     string    `json:"state" gorm:"column:state;index"` // Archive, Release, Development, Latest
	Valid     string    `json:"valid" gorm:"column:valid"`
	Version   string    `json:"version" gorm:"column:version;index"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`

	DocOAS *openapi3.T `json:"-" gorm:"-"`
}

// TableName implements gorm Tabler interface
func (m *Spec) TableName() string {
	return SpecTableName
}

func (m *Spec) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = shared.TimeUUID()
	return
}

// GetID returns the ID of spec object
func (m *Spec) GetID() string {
	return fmt.Sprintf("%v", m.ID)
}

// GetTags returns all the tags
func (m *Spec) GetTags() []string {
	tags := make([]string, 0, 10)
	tags = append(tags, m.ServiceID)
	tags = append(tags, m.Version)
	tags = append(tags, m.Revision)
	tags = append(tags, m.State)
	return tags
}

// String returns the text representation of spec object
func (m *Spec) String() string {
	return fmt.Sprintf("%v", *m)
}

// GetIndex returns an index for specific field
func (m *Spec) GetIndex(field string) string {
	return m.GetIndexes()[field]
}

// GetIndexes returns all the field indexes
func (m *Spec) GetIndexes() map[string]string {
	return map[string]string{
		"service_id": "idx_service_id",
		"version":    "idx_version",
		"revision":   "idx_revision",
		"state":      "idx_state",
	}
}

// GetIndexValue return index value for specified field
func (m *Spec) GetIndexValue(field string) string {
	return m.GetIndexValues()[field]
}

// GetIndexValues return all field index values
func (m *Spec) GetIndexValues() map[string]string {
	return map[string]string{
		"service_id": m.ServiceID,
		"version":    m.Version,
		"revision":   m.Revision,
		"state":      m.State,
	}
}

// Sortable checks if field is sortable.
func (m *Spec) Sortable(field string) bool {
	_, found := m.SortableFields()[field]
	return found
}

// SortableFields returns all sortable fields
func (m *Spec) SortableFields() map[string]struct{} {
	return map[string]struct{}{
		"created_at": {},
		"updated_at": {},
	}
}

// LoadDocAsOAS loads Spec.Doc as an OpenAPI spec & stores it as Spec.DocOAS.
// Set validate to validate Spec.DocOAS.
// Set setDocType to derive Spec.DocType from Spec.DocOAS.Version.
// Set setVersion to derive Spec.Version from Spec.DocOAS.Info.Version.
func (m *Spec) LoadDocAsOAS(ctx context.Context, validate, setDocType, setVersion bool) (*openapi3.T, error) {
	if m.DocOAS != nil {
		return m.DocOAS, nil
	}
	if m.Doc == nil || len(*m.Doc) == 0 {
		return nil, fmt.Errorf("spec: missing Spec.Doc")
	}
	var err error
	m.DocOAS, err = openapi3.NewLoader().LoadFromData([]byte(*m.Doc))
	if err != nil {
		return nil, fmt.Errorf("spec: failed to load Spec.Doc as OAS")
	}
	if validate {
		err := m.DocOAS.Validate(ctx)
		if err != nil {
			return nil, fmt.Errorf("spec: invalid Spec.Doc: %v", err)
		}
	}
	if setDocType && m.DocType == "" {
		var version string
		if m.DocOAS.OpenAPI != "" {
			version = m.DocOAS.OpenAPI
		} else {
			if swagger, ok := m.DocOAS.Extensions["swagger"]; ok {
				if versionBytes, ok := swagger.(json.RawMessage); ok && len(versionBytes) > 0 {
					version = strings.ReplaceAll(string(versionBytes), "\"", "")
				}
			}
		}
		m.DocType = fmt.Sprintf("oas-%v", version)
	}
	if setVersion && m.Version == "" {
		if m.DocOAS.Info != nil && m.DocOAS.Info.Version != "" {
			m.Version = m.DocOAS.Info.Version
		}
	}
	return m.DocOAS, nil
}

// GetDocAsMap unmarshals Spec.Doc into a map by first parsing as a JSON & if that fails, as a YAML.
func (m *Spec) GetDocAsMap() (docMap map[string]interface{}, isJSON bool, err error) {
	if m.Doc == nil {
		return nil, false, fmt.Errorf("spec: missing Spec.Doc")
	}

	err = json.Unmarshal([]byte(*m.Doc), &docMap)

	if err == nil {
		isJSON = true
		return
	}

	err = yaml.Unmarshal([]byte(*m.Doc), &docMap)

	if err != nil {
		return nil, false, fmt.Errorf("spec: invalid Spec.Doc: %v", err)
	}

	return
}

// SpecResponse wrappers spec response
type SpecResponse struct {
	Pagination
	Data []Spec `json:"data"`
}

type SpecDoc *string

func NewSpecDocFromBytes(data []byte) SpecDoc {
	s := string(data)
	return &s
}
