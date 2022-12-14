// Package dialogflow provides access to the Dialogflow API.
//
// See https://cloud.google.com/dialogflow-enterprise/
//
// Usage example:
//
//   import "google.golang.org/api/dialogflow/v2beta1"
//   ...
//   dialogflowService, err := dialogflow.New(oauthHttpClient)
package dialogflow // import "google.golang.org/api/dialogflow/v2beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "dialogflow:v2beta1"
const apiName = "dialogflow"
const apiVersion = "v2beta1"
const basePath = "https://dialogflow.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Agent = NewProjectsAgentService(s)
	rs.Operations = NewProjectsOperationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Agent *ProjectsAgentService

	Operations *ProjectsOperationsService
}

func NewProjectsAgentService(s *Service) *ProjectsAgentService {
	rs := &ProjectsAgentService{s: s}
	rs.EntityTypes = NewProjectsAgentEntityTypesService(s)
	rs.Intents = NewProjectsAgentIntentsService(s)
	rs.Sessions = NewProjectsAgentSessionsService(s)
	return rs
}

type ProjectsAgentService struct {
	s *Service

	EntityTypes *ProjectsAgentEntityTypesService

	Intents *ProjectsAgentIntentsService

	Sessions *ProjectsAgentSessionsService
}

func NewProjectsAgentEntityTypesService(s *Service) *ProjectsAgentEntityTypesService {
	rs := &ProjectsAgentEntityTypesService{s: s}
	rs.Entities = NewProjectsAgentEntityTypesEntitiesService(s)
	return rs
}

type ProjectsAgentEntityTypesService struct {
	s *Service

	Entities *ProjectsAgentEntityTypesEntitiesService
}

func NewProjectsAgentEntityTypesEntitiesService(s *Service) *ProjectsAgentEntityTypesEntitiesService {
	rs := &ProjectsAgentEntityTypesEntitiesService{s: s}
	return rs
}

type ProjectsAgentEntityTypesEntitiesService struct {
	s *Service
}

func NewProjectsAgentIntentsService(s *Service) *ProjectsAgentIntentsService {
	rs := &ProjectsAgentIntentsService{s: s}
	return rs
}

type ProjectsAgentIntentsService struct {
	s *Service
}

func NewProjectsAgentSessionsService(s *Service) *ProjectsAgentSessionsService {
	rs := &ProjectsAgentSessionsService{s: s}
	rs.Contexts = NewProjectsAgentSessionsContextsService(s)
	rs.EntityTypes = NewProjectsAgentSessionsEntityTypesService(s)
	return rs
}

type ProjectsAgentSessionsService struct {
	s *Service

	Contexts *ProjectsAgentSessionsContextsService

	EntityTypes *ProjectsAgentSessionsEntityTypesService
}

func NewProjectsAgentSessionsContextsService(s *Service) *ProjectsAgentSessionsContextsService {
	rs := &ProjectsAgentSessionsContextsService{s: s}
	return rs
}

type ProjectsAgentSessionsContextsService struct {
	s *Service
}

func NewProjectsAgentSessionsEntityTypesService(s *Service) *ProjectsAgentSessionsEntityTypesService {
	rs := &ProjectsAgentSessionsEntityTypesService{s: s}
	return rs
}

type ProjectsAgentSessionsEntityTypesService struct {
	s *Service
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

// Agent: Represents a conversational agent.
type Agent struct {
	// AvatarUri: Optional. The URI of the agent's avatar.
	// Avatars are used throughout API.AI console and in the
	// self-hosted
	// [Web Demo](https://dialogflow.com/docs/integrations/web-demo)
	// integration.
	AvatarUri string `json:"avatarUri,omitempty"`

	// ClassificationThreshold: Optional. To filter out false positive
	// results and still get variety in
	// matched natural language inputs for your agent, you can tune the
	// machine
	// learning classification threshold. If the returned score value is
	// less than
	// the threshold value, then a fallback intent is be triggered or, if
	// there
	// are no fallback intents defined, no intent will be triggered. The
	// score
	// values range from 0.0 (completely uncertain) to 1.0 (completely
	// certain).
	// If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold float64 `json:"classificationThreshold,omitempty"`

	// DefaultLanguageCode: Required. The default language of the agent as a
	// language tag. See
	// [Language Support](https://dialogflow.com/docs/reference/language)
	// for a
	// list of the currently supported language codes.
	// This field cannot be set by the `Update` method.
	DefaultLanguageCode string `json:"defaultLanguageCode,omitempty"`

	// Description: Optional. The description of this agent.
	// The maximum length is 500 characters. If exceeded, the request is
	// rejected.
	Description string `json:"description,omitempty"`

	// DisplayName: Required. The name of this agent.
	DisplayName string `json:"displayName,omitempty"`

	// EnableLogging: Optional. Determines whether this agent should log
	// conversation queries.
	EnableLogging bool `json:"enableLogging,omitempty"`

	// MatchMode: Optional. Determines how intents are detected from user
	// queries.
	//
	// Possible values:
	//   "MATCH_MODE_UNSPECIFIED" - Not specified.
	//   "MATCH_MODE_HYBRID" - Best for agents with a small number of
	// examples in intents and/or wide
	// use of templates syntax and composite entities.
	//   "MATCH_MODE_ML_ONLY" - Can be used for agents with a large number
	// of examples in intents,
	// especially the ones using @sys.any or very large developer entities.
	MatchMode string `json:"matchMode,omitempty"`

	// Parent: Required. The project of this agent.
	// Format: `projects/<Project ID>`.
	Parent string `json:"parent,omitempty"`

	// SupportedLanguageCodes: Optional. The list of all languages supported
	// by this agent (except for the
	// `default_language_code`).
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty"`

	// TimeZone: Required. The time zone of this agent from the
	// [time zone database](https://www.iana.org/time-zones),
	// e.g.,
	// America/New_York, Europe/Paris.
	TimeZone string `json:"timeZone,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AvatarUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvatarUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Agent) MarshalJSON() ([]byte, error) {
	type NoMethod Agent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Agent) UnmarshalJSON(data []byte) error {
	type NoMethod Agent
	var s1 struct {
		ClassificationThreshold gensupport.JSONFloat64 `json:"classificationThreshold"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ClassificationThreshold = float64(s1.ClassificationThreshold)
	return nil
}

// BatchCreateEntitiesRequest: The request message for
// EntityTypes.BatchCreateEntities.
type BatchCreateEntitiesRequest struct {
	// Entities: Required. The collection of entities to create.
	Entities []*EntityTypeEntity `json:"entities,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entities`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchCreateEntitiesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchCreateEntitiesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchDeleteEntitiesRequest: The request message for
// EntityTypes.BatchDeleteEntities.
type BatchDeleteEntitiesRequest struct {
	// EntityValues: Required. The canonical `values` of the entities to
	// delete. Note that
	// these are not fully-qualified names, i.e. they don't start
	// with
	// `projects/<Project ID>`.
	EntityValues []string `json:"entityValues,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entities`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityValues") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityValues") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchDeleteEntitiesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchDeleteEntitiesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchDeleteEntityTypesRequest: The request message for
// EntityTypes.BatchDeleteEntityTypes.
type BatchDeleteEntityTypesRequest struct {
	// EntityTypeNames: Required. The names entity types to delete. All
	// names must point to the
	// same agent as `parent`.
	EntityTypeNames []string `json:"entityTypeNames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypeNames") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypeNames") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BatchDeleteEntityTypesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchDeleteEntityTypesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchDeleteIntentsRequest: The request message for
// Intents.BatchDeleteIntents.
type BatchDeleteIntentsRequest struct {
	// Intents: Required. The collection of intents to delete. Only intent
	// `name` must be
	// filled in.
	Intents []*Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchDeleteIntentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchDeleteIntentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchUpdateEntitiesRequest: The response message for
// EntityTypes.BatchCreateEntities.
type BatchUpdateEntitiesRequest struct {
	// Entities: Required. The collection of new entities to replace the
	// existing entities.
	Entities []*EntityTypeEntity `json:"entities,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entities`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// UpdateMask: Optional. The mask to control which fields get updated.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchUpdateEntitiesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchUpdateEntitiesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchUpdateEntityTypesRequest: The request message for
// EntityTypes.BatchUpdateEntityTypes.
type BatchUpdateEntityTypesRequest struct {
	// EntityTypeBatchInline: The collection of entity type to update or
	// create.
	EntityTypeBatchInline *EntityTypeBatch `json:"entityTypeBatchInline,omitempty"`

	// EntityTypeBatchUri: Warning: Importing entity types from a URI is not
	// implemented yet.
	// This feature is coming soon.
	// The URI to a Google Cloud Storage file containing entity types to
	// update
	// or create. The file format can either be a serialized proto
	// (of
	// EntityBatch type) or a JSON object. Note: The URI must start
	// with
	// "gs://".
	EntityTypeBatchUri string `json:"entityTypeBatchUri,omitempty"`

	// LanguageCode: Optional. The language of entity synonyms defined in
	// `entity_types`. If not
	// specified, the agent's default language is used.
	// [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// UpdateMask: Optional. The mask to control which fields get updated.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "EntityTypeBatchInline") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypeBatchInline") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BatchUpdateEntityTypesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchUpdateEntityTypesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchUpdateEntityTypesResponse: The response message for
// EntityTypes.BatchUpdateEntityTypes.
type BatchUpdateEntityTypesResponse struct {
	// EntityTypes: The collection of updated or created entity types.
	EntityTypes []*EntityType `json:"entityTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchUpdateEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchUpdateEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchUpdateIntentsRequest: The request message for
// Intents.BatchUpdateIntents.
type BatchUpdateIntentsRequest struct {
	// IntentBatchInline: The collection of intents to update or create.
	IntentBatchInline *IntentBatch `json:"intentBatchInline,omitempty"`

	// IntentBatchUri: Warning: Importing intents from a URI is not
	// implemented yet.
	// This feature is coming soon.
	// The URI to a Google Cloud Storage file containing intents to update
	// or
	// create. The file format can either be a serialized proto (of
	// IntentBatch
	// type) or JSON object. Note: The URI must start with "gs://".
	IntentBatchUri string `json:"intentBatchUri,omitempty"`

	// IntentView: Optional. The resource view to apply to the returned
	// intent.
	//
	// Possible values:
	//   "INTENT_VIEW_UNSPECIFIED" - Training phrases field is not populated
	// in the response.
	//   "INTENT_VIEW_FULL" - All fields are populated.
	IntentView string `json:"intentView,omitempty"`

	// LanguageCode: Optional. The language of training phrases, parameters
	// and rich messages
	// defined in `intents`. If not specified, the agent's default language
	// is
	// used. [More than a
	// dozen
	// languages](https://dialogflow.com/docs/reference/language) are
	// supported.
	// Note: languages must be enabled in the agent, before they can be
	// used.
	LanguageCode string `json:"languageCode,omitempty"`

	// UpdateMask: Optional. The mask to control which fields get updated.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IntentBatchInline")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IntentBatchInline") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BatchUpdateIntentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchUpdateIntentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchUpdateIntentsResponse: The response message for
// Intents.BatchUpdateIntents.
type BatchUpdateIntentsResponse struct {
	// Intents: The collection of updated or created intents.
	Intents []*Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchUpdateIntentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchUpdateIntentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Context: Represents a context.
type Context struct {
	// LifespanCount: Optional. The number of conversational query requests
	// after which the
	// context expires. If set to `0` (the default) the context
	// expires
	// immediately. Contexts expire automatically after 10 minutes even if
	// there
	// are no matching queries.
	LifespanCount int64 `json:"lifespanCount,omitempty"`

	// Name: Required. The unique identifier of the context.
	// Format:
	// `projects/<Project ID>/agent/sessions/<Session ID>/contexts/<Context
	// ID>`.
	// Note: The Context ID is always converted to lowercase.
	Name string `json:"name,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// this context.
	// Refer to [this
	// doc](https://dialogflow.com/docs/actions-and-parameters) for
	// syntax.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "LifespanCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LifespanCount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Context) MarshalJSON() ([]byte, error) {
	type NoMethod Context
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DetectIntentRequest: The request to detect user's intent.
type DetectIntentRequest struct {
	// InputAudio: Optional. The natural language speech audio to be
	// processed. This field
	// should be populated iff `query_input` is set to an input audio
	// config.
	// A single request can contain up to 1 minute of speech audio data.
	InputAudio string `json:"inputAudio,omitempty"`

	// QueryInput: Required. The input specification. It can be set to:
	//
	// 1.  an audio config
	//     which instructs the speech recognizer how to process the speech
	// audio,
	//
	// 2.  a conversational query in the form of text, or
	//
	// 3.  an event that specifies which intent to trigger.
	QueryInput *QueryInput `json:"queryInput,omitempty"`

	// QueryParams: Optional. The parameters of this query.
	QueryParams *QueryParameters `json:"queryParams,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputAudio") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "InputAudio") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DetectIntentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod DetectIntentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DetectIntentResponse: The message returned from the DetectIntent
// method.
type DetectIntentResponse struct {
	// QueryResult: The results of the conversational query or event
	// processing.
	QueryResult *QueryResult `json:"queryResult,omitempty"`

	// ResponseId: The unique identifier of the response. It can be used
	// to
	// locate a response in the training example set or for reporting
	// issues.
	ResponseId string `json:"responseId,omitempty"`

	// WebhookStatus: Specifies the status of the webhook request.
	// `webhook_status`
	// is never populated in webhook requests.
	WebhookStatus *Status `json:"webhookStatus,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "QueryResult") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "QueryResult") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DetectIntentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod DetectIntentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// EntityType: Represents an entity type.
// Entity types serve as a tool for extracting parameter values from
// natural
// language queries.
type EntityType struct {
	// AutoExpansionMode: Optional. Indicates whether the entity type can be
	// automatically
	// expanded.
	//
	// Possible values:
	//   "AUTO_EXPANSION_MODE_UNSPECIFIED" - Auto expansion disabled for the
	// entity.
	//   "AUTO_EXPANSION_MODE_DEFAULT" - Allows an agent to recognize values
	// that have not been explicitly
	// listed in the entity.
	AutoExpansionMode string `json:"autoExpansionMode,omitempty"`

	// DisplayName: Required. The name of the entity.
	DisplayName string `json:"displayName,omitempty"`

	// Entities: Optional. The collection of entities associated with the
	// entity type.
	Entities []*EntityTypeEntity `json:"entities,omitempty"`

	// Kind: Required. Indicates the kind of entity type.
	//
	// Possible values:
	//   "KIND_UNSPECIFIED" - Not specified. This value should be never
	// used.
	//   "KIND_MAP" - Map entity types allow mapping of a group of synonyms
	// to a canonical
	// value.
	//   "KIND_LIST" - List entity types contain a set of entries that do
	// not map to canonical
	// values. However, list entity types can contain references to other
	// entity
	// types (with or without aliases).
	Kind string `json:"kind,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically.
	// The unique identifier of the entity type. Format:
	// `projects/<Project ID>/agent/entityTypes/<Entity Type ID>`.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AutoExpansionMode")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutoExpansionMode") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *EntityType) MarshalJSON() ([]byte, error) {
	type NoMethod EntityType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EntityTypeBatch: This message is a wrapper around a collection of
// entity types.
type EntityTypeBatch struct {
	// EntityTypes: A collection of entity types.
	EntityTypes []*EntityType `json:"entityTypes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EntityTypeBatch) MarshalJSON() ([]byte, error) {
	type NoMethod EntityTypeBatch
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EntityTypeEntity: Optional. Represents an entity.
type EntityTypeEntity struct {
	// Synonyms: Required. A collection of synonyms. For `KIND_LIST` entity
	// types this
	// must contain exactly one synonym equal to `value`.
	Synonyms []string `json:"synonyms,omitempty"`

	// Value: Required.
	// For `KIND_MAP` entity types:
	//   A canonical name to be used in place of synonyms.
	// For `KIND_LIST` entity types:
	//   A string that can contain references to other entity types (with
	// or
	//   without aliases).
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Synonyms") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Synonyms") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EntityTypeEntity) MarshalJSON() ([]byte, error) {
	type NoMethod EntityTypeEntity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// EventInput: Events allow for matching intents by event name instead
// of the natural
// language input. For instance, input `<event: { name:
// ???welcome_event???,
// parameters: { name: ???Sam??? } }>` can trigger a personalized
// welcome response.
// The parameter `name` may be used by the agent in the
// response:
// `???Hello #welcome_event.name! What can I do for you today????`.
type EventInput struct {
	// LanguageCode: Required. The language of this query. See
	// [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// Name: Required. The unique identifier of the event.
	Name string `json:"name,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// the event.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EventInput) MarshalJSON() ([]byte, error) {
	type NoMethod EventInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExportAgentRequest: The request message for Agents.ExportAgent.
type ExportAgentRequest struct {
	// AgentUri: Warning: Exporting agents to a URI is not implemented
	// yet.
	// This feature is coming soon.
	// Optional. The Google Cloud Storage URI to export the agent to.
	// Note: The URI must start with
	// "gs://". If left unspecified, the serialized agent is returned
	// inline.
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExportAgentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ExportAgentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExportAgentResponse: The response message for Agents.ExportAgent.
type ExportAgentResponse struct {
	// AgentContent: The exported agent.
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: The URI to a file containing the exported agent. This field
	// is populated
	// only if `agent_uri`
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExportAgentResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ExportAgentResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImportAgentRequest: The request message for Agents.ImportAgent.
type ImportAgentRequest struct {
	// AgentContent: The agent to import.
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: Warning: Importing agents from a URI is not implemented
	// yet.
	// This feature is coming soon.
	// The URI to a Google Cloud Storage file containing the agent to
	// import.
	// Note: The URI must start with "gs://".
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImportAgentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ImportAgentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InputAudioConfig: Instructs the speech recognizer how to process the
// audio content.
type InputAudioConfig struct {
	// AudioEncoding: Required. Audio encoding of the audio content to
	// process.
	//
	// Possible values:
	//   "AUDIO_ENCODING_UNSPECIFIED" - Not specified.
	//   "AUDIO_ENCODING_LINEAR_16" - Uncompressed 16-bit signed
	// little-endian samples (Linear PCM).
	//   "AUDIO_ENCODING_FLAC" -
	// [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless
	// Audio
	// Codec) is the recommended encoding because it is lossless
	// (therefore
	// recognition is not compromised) and requires only about half
	// the
	// bandwidth of `LINEAR16`. `FLAC` stream encoding supports 16-bit
	// and
	// 24-bit samples, however, not all fields in `STREAMINFO` are
	// supported.
	//   "AUDIO_ENCODING_MULAW" - 8-bit samples that compand 14-bit audio
	// samples using G.711 PCMU/mu-law.
	//   "AUDIO_ENCODING_AMR" - Adaptive Multi-Rate Narrowband codec.
	// `sample_rate_hertz` must be 8000.
	//   "AUDIO_ENCODING_AMR_WB" - Adaptive Multi-Rate Wideband codec.
	// `sample_rate_hertz` must be 16000.
	//   "AUDIO_ENCODING_OGG_OPUS" - Opus encoded audio frames in Ogg
	// container
	// ([OggOpus](https://wiki.xiph.org/OggOpus)).
	// `sample_rate_her
	// tz` must be 16000.
	//   "AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE" - Although the use of lossy
	// encodings is not recommended, if a very low
	// bitrate encoding is required, `OGG_OPUS` is highly preferred
	// over
	// Speex encoding. The [Speex](https://speex.org/) encoding supported
	// by
	// Dialogflow API has a header byte in each block, as in MIME
	// type
	// `audio/x-speex-with-header-byte`.
	// It is a variant of the RTP Speex encoding defined in
	// [RFC 5574](https://tools.ietf.org/html/rfc5574).
	// The stream is a sequence of blocks, one block per RTP packet. Each
	// block
	// starts with a byte containing the length of the block, in bytes,
	// followed
	// by one or more frames of Speex data, padded to an integral number
	// of
	// bytes (octets) as specified in RFC 5574. In other words, each RTP
	// header
	// is replaced with a single byte containing the block length. Only
	// Speex
	// wideband is supported. `sample_rate_hertz` must be 16000.
	AudioEncoding string `json:"audioEncoding,omitempty"`

	// LanguageCode: Required. The language of the supplied audio.
	// Dialogflow does not do
	// translations. See
	// [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// PhraseHints: Optional. The collection of phrase hints which are used
	// to boost accuracy
	// of speech recognition.
	// Refer to [Cloud Speech API
	// documentation](/speech/docs/basics#phrase-hints)
	// for more details.
	PhraseHints []string `json:"phraseHints,omitempty"`

	// SampleRateHertz: Required. Sample rate (in Hertz) of the audio
	// content sent in the query.
	// Refer to [Cloud Speech API documentation](/speech/docs/basics) for
	// more
	// details.
	SampleRateHertz int64 `json:"sampleRateHertz,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioEncoding") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioEncoding") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *InputAudioConfig) MarshalJSON() ([]byte, error) {
	type NoMethod InputAudioConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Intent: Represents an intent.
// Intents convert a number of user expressions or patterns into an
// action. An
// action is an extraction of a user command or sentence semantics.
type Intent struct {
	// Action: Optional. The name of the action associated with the intent.
	Action string `json:"action,omitempty"`

	// DefaultResponsePlatforms: Optional. The list of platforms for which
	// the first response will be
	// taken from among the messages assigned to the DEFAULT_PLATFORM.
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - Not specified.
	//   "FACEBOOK" - Facebook.
	//   "SLACK" - Slack.
	//   "TELEGRAM" - Telegram.
	//   "KIK" - Kik.
	//   "SKYPE" - Skype.
	//   "LINE" - Line.
	//   "VIBER" - Viber.
	//   "ACTIONS_ON_GOOGLE" - Actions on Google.
	DefaultResponsePlatforms []string `json:"defaultResponsePlatforms,omitempty"`

	// DisplayName: Required. The name of this intent.
	DisplayName string `json:"displayName,omitempty"`

	// Events: Optional. The collection of event names that trigger the
	// intent.
	// If the collection of input contexts is not empty, all of the contexts
	// must
	// be present in the active user session for an event to trigger this
	// intent.
	Events []string `json:"events,omitempty"`

	// FollowupIntentInfo: Optional. Collection of information about all
	// followup intents that have
	// name of this intent as a root_name.
	FollowupIntentInfo []*IntentFollowupIntentInfo `json:"followupIntentInfo,omitempty"`

	// InputContextNames: Optional. The list of context names required for
	// this intent to be
	// triggered.
	// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context
	// ID>`.
	InputContextNames []string `json:"inputContextNames,omitempty"`

	// IsFallback: Optional. Indicates whether this is a fallback intent.
	IsFallback bool `json:"isFallback,omitempty"`

	// Messages: Optional. The collection of rich messages corresponding to
	// the
	// `Response` field in API.AI console.
	Messages []*IntentMessage `json:"messages,omitempty"`

	// MlEnabled: Optional. Indicates whether Machine Learning is enabled
	// for the intent.
	// Note: If `ml_enabled` setting is set to false, then this intent is
	// not
	// taken into account during inference in `ML ONLY` match mode.
	// Also,
	// auto-markup in the UI is turned off.
	MlEnabled bool `json:"mlEnabled,omitempty"`

	// Name: Required for all methods except `create` (`create` populates
	// the name
	// automatically.
	// The unique identifier of this intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	Name string `json:"name,omitempty"`

	// OutputContexts: Optional. The collection of contexts that are
	// activated when the intent
	// is matched. Context messages in this collection should not set
	// the
	// parameters field. Setting the `lifespan_count` to 0 will reset the
	// context
	// when the intent is matched.
	// Format: `projects/<Project ID>/agent/sessions/-/contexts/<Context
	// ID>`.
	OutputContexts []*Context `json:"outputContexts,omitempty"`

	// Parameters: Optional. The collection of parameters associated with
	// the intent.
	Parameters []*IntentParameter `json:"parameters,omitempty"`

	// ParentFollowupIntentName: The unique identifier of the parent intent
	// in the chain of followup
	// intents.
	// It identifies the parent followup intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty"`

	// Priority: Optional. The priority of this intent. Higher numbers
	// represent higher
	// priorities. Zero or negative numbers mean that the intent is
	// disabled.
	Priority int64 `json:"priority,omitempty"`

	// ResetContexts: Optional. Indicates whether to delete all contexts in
	// the current
	// session when this intent is matched.
	ResetContexts bool `json:"resetContexts,omitempty"`

	// RootFollowupIntentName: The unique identifier of the root intent in
	// the chain of followup intents.
	// It identifies the correct followup intents chain for this
	// intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	RootFollowupIntentName string `json:"rootFollowupIntentName,omitempty"`

	// TrainingPhrases: Optional. The collection of examples/templates that
	// the agent is
	// trained on.
	TrainingPhrases []*IntentTrainingPhrase `json:"trainingPhrases,omitempty"`

	// WebhookState: Required. Indicates whether webhooks are enabled for
	// the intent.
	//
	// Possible values:
	//   "WEBHOOK_STATE_UNSPECIFIED" - Webhook is disabled in the agent and
	// in the intent.
	//   "WEBHOOK_STATE_ENABLED" - Webhook is enabled in the agent and in
	// the intent.
	//   "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING" - Webhook is enabled in
	// the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook.
	WebhookState string `json:"webhookState,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Intent) MarshalJSON() ([]byte, error) {
	type NoMethod Intent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentBatch: This message is a wrapper around a collection of
// intents.
type IntentBatch struct {
	// Intents: A collection of intents.
	Intents []*Intent `json:"intents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentBatch) MarshalJSON() ([]byte, error) {
	type NoMethod IntentBatch
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentFollowupIntentInfo: Represents a single followup intent in the
// chain.
type IntentFollowupIntentInfo struct {
	// FollowupIntentName: The unique identifier of the followup
	// intent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	FollowupIntentName string `json:"followupIntentName,omitempty"`

	// ParentFollowupIntentName: The unique identifier of the followup
	// intent parent.
	// Format: `projects/<Project ID>/agent/intents/<Intent ID>`.
	ParentFollowupIntentName string `json:"parentFollowupIntentName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FollowupIntentName")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FollowupIntentName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *IntentFollowupIntentInfo) MarshalJSON() ([]byte, error) {
	type NoMethod IntentFollowupIntentInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessage: Corresponds to the `Response` field in API.AI console.
type IntentMessage struct {
	// BasicCard: The basic card response for Actions on Google.
	BasicCard *IntentMessageBasicCard `json:"basicCard,omitempty"`

	// Card: The card response.
	Card *IntentMessageCard `json:"card,omitempty"`

	// CarouselSelect: The carousel card response for Actions on Google.
	CarouselSelect *IntentMessageCarouselSelect `json:"carouselSelect,omitempty"`

	// Image: The image response.
	Image *IntentMessageImage `json:"image,omitempty"`

	// LinkOutSuggestion: The link out suggestion chip for Actions on
	// Google.
	LinkOutSuggestion *IntentMessageLinkOutSuggestion `json:"linkOutSuggestion,omitempty"`

	// ListSelect: The list card response for Actions on Google.
	ListSelect *IntentMessageListSelect `json:"listSelect,omitempty"`

	// Payload: The response containing a custom payload.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Platform: Optional. The platform that this message is intended for.
	//
	// Possible values:
	//   "PLATFORM_UNSPECIFIED" - Not specified.
	//   "FACEBOOK" - Facebook.
	//   "SLACK" - Slack.
	//   "TELEGRAM" - Telegram.
	//   "KIK" - Kik.
	//   "SKYPE" - Skype.
	//   "LINE" - Line.
	//   "VIBER" - Viber.
	//   "ACTIONS_ON_GOOGLE" - Actions on Google.
	Platform string `json:"platform,omitempty"`

	// QuickReplies: The quick replies response.
	QuickReplies *IntentMessageQuickReplies `json:"quickReplies,omitempty"`

	// SimpleResponses: The voice and text-only responses for Actions on
	// Google.
	SimpleResponses *IntentMessageSimpleResponses `json:"simpleResponses,omitempty"`

	// Suggestions: The suggestion chips for Actions on Google.
	Suggestions *IntentMessageSuggestions `json:"suggestions,omitempty"`

	// Text: The text response.
	Text *IntentMessageText `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BasicCard") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BasicCard") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessage) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageBasicCard: The basic card message. Useful for displaying
// information.
type IntentMessageBasicCard struct {
	// Buttons: Optional. The collection of card buttons.
	Buttons []*IntentMessageBasicCardButton `json:"buttons,omitempty"`

	// FormattedText: Required, unless image is present. The body text of
	// the card.
	FormattedText string `json:"formattedText,omitempty"`

	// Image: Optional. The image for the card.
	Image *IntentMessageImage `json:"image,omitempty"`

	// Subtitle: Optional. The subtitle of the card.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Optional. The title of the card.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageBasicCard) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageBasicCard
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageBasicCardButton: The button object that appears at the
// bottom of a card.
type IntentMessageBasicCardButton struct {
	// OpenUriAction: Required. Action to take when a user taps on the
	// button.
	OpenUriAction *IntentMessageBasicCardButtonOpenUriAction `json:"openUriAction,omitempty"`

	// Title: Required. The title of the button.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OpenUriAction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OpenUriAction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageBasicCardButton) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageBasicCardButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageBasicCardButtonOpenUriAction: Opens the given URI.
type IntentMessageBasicCardButtonOpenUriAction struct {
	// Uri: Required. The HTTP or HTTPS scheme URI.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Uri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Uri") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageBasicCardButtonOpenUriAction) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageBasicCardButtonOpenUriAction
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageCard: The card response message.
type IntentMessageCard struct {
	// Buttons: Optional. The collection of card buttons.
	Buttons []*IntentMessageCardButton `json:"buttons,omitempty"`

	// ImageUri: Optional. The public URI to an image file for the card.
	ImageUri string `json:"imageUri,omitempty"`

	// Subtitle: Optional. The subtitle of the card.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Optional. The title of the card.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageCard) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageCard
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageCardButton: Optional. Contains information about a
// button.
type IntentMessageCardButton struct {
	// Postback: Optional. The text to send back to the Dialogflow API or a
	// URI to
	// open.
	Postback string `json:"postback,omitempty"`

	// Text: Optional. The text to show on the button.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Postback") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Postback") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageCardButton) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageCardButton
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageCarouselSelect: The card for presenting a carousel of
// options to select from.
type IntentMessageCarouselSelect struct {
	// Items: Required. Carousel items.
	Items []*IntentMessageCarouselSelectItem `json:"items,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageCarouselSelect) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageCarouselSelect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageCarouselSelectItem: An item in the carousel.
type IntentMessageCarouselSelectItem struct {
	// Description: Optional. The body text of the card.
	Description string `json:"description,omitempty"`

	// Image: Optional. The image to display.
	Image *IntentMessageImage `json:"image,omitempty"`

	// Info: Required. Additional info about the option item.
	Info *IntentMessageSelectItemInfo `json:"info,omitempty"`

	// Title: Required. Title of the carousel item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageCarouselSelectItem) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageCarouselSelectItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageImage: The image response message.
type IntentMessageImage struct {
	// ImageUri: Optional. The public URI to an image file.
	ImageUri string `json:"imageUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageUri") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageImage) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageImage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageLinkOutSuggestion: The suggestion chip message that
// allows the user to jump out to the app
// or website associated with this agent.
type IntentMessageLinkOutSuggestion struct {
	// DestinationName: Required. The name of the app or site this chip is
	// linking to.
	DestinationName string `json:"destinationName,omitempty"`

	// Uri: Required. The URI of the app or site to open when the user taps
	// the
	// suggestion chip.
	Uri string `json:"uri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DestinationName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DestinationName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageLinkOutSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageLinkOutSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageListSelect: The card for presenting a list of options to
// select from.
type IntentMessageListSelect struct {
	// Items: Required. List items.
	Items []*IntentMessageListSelectItem `json:"items,omitempty"`

	// Title: Optional. The overall title of the list.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageListSelect) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageListSelect
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageListSelectItem: An item in the list.
type IntentMessageListSelectItem struct {
	// Description: Optional. The main text describing the item.
	Description string `json:"description,omitempty"`

	// Image: Optional. The image to display.
	Image *IntentMessageImage `json:"image,omitempty"`

	// Info: Required. Additional information about this option.
	Info *IntentMessageSelectItemInfo `json:"info,omitempty"`

	// Title: Required. The title of the list item.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageListSelectItem) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageListSelectItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageQuickReplies: The quick replies response message.
type IntentMessageQuickReplies struct {
	// QuickReplies: Optional. The collection of quick replies.
	QuickReplies []string `json:"quickReplies,omitempty"`

	// Title: Optional. The title of the collection of quick replies.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "QuickReplies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "QuickReplies") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageQuickReplies) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageQuickReplies
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageSelectItemInfo: Additional info about the select item
// for when it is triggered in a
// dialog.
type IntentMessageSelectItemInfo struct {
	// Key: Required. A unique key that will be sent back to the agent if
	// this
	// response is given.
	Key string `json:"key,omitempty"`

	// Synonyms: Optional. A list of synonyms that can also be used to
	// trigger this
	// item in dialog.
	Synonyms []string `json:"synonyms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageSelectItemInfo) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageSelectItemInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageSimpleResponse: The simple response message containing
// speech or text.
type IntentMessageSimpleResponse struct {
	// DisplayText: Optional. The text to display.
	DisplayText string `json:"displayText,omitempty"`

	// Ssml: One of text_to_speech or ssml must be provided. Structured
	// spoken
	// response to the user in the SSML format. Mutually exclusive
	// with
	// text_to_speech.
	Ssml string `json:"ssml,omitempty"`

	// TextToSpeech: One of text_to_speech or ssml must be provided. The
	// plain text of the
	// speech output. Mutually exclusive with ssml.
	TextToSpeech string `json:"textToSpeech,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayText") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayText") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageSimpleResponse) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageSimpleResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageSimpleResponses: The collection of simple response
// candidates.
// This message in `QueryResult.fulfillment_messages`
// and
// `WebhookResponse.fulfillment_messages` should contain only
// one
// `SimpleResponse`.
type IntentMessageSimpleResponses struct {
	// SimpleResponses: Required. The list of simple responses.
	SimpleResponses []*IntentMessageSimpleResponse `json:"simpleResponses,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SimpleResponses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SimpleResponses") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageSimpleResponses) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageSimpleResponses
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageSuggestion: The suggestion chip message that the user
// can tap to quickly post a reply
// to the conversation.
type IntentMessageSuggestion struct {
	// Title: Required. The text shown the in the suggestion chip.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Title") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Title") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageSuggestion) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageSuggestion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageSuggestions: The collection of suggestions.
type IntentMessageSuggestions struct {
	// Suggestions: Required. The list of suggested replies.
	Suggestions []*IntentMessageSuggestion `json:"suggestions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Suggestions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Suggestions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageSuggestions) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageSuggestions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentMessageText: The text response message.
type IntentMessageText struct {
	// Text: Optional. The collection of the agent's responses.
	Text []string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentMessageText) MarshalJSON() ([]byte, error) {
	type NoMethod IntentMessageText
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentParameter: Represents intent parameters.
type IntentParameter struct {
	// DefaultValue: Optional. The default value to use when the `value`
	// yields an empty
	// result.
	// Default values can be extracted from contexts by using the
	// following
	// syntax: `#context_name.parameter_name`.
	DefaultValue string `json:"defaultValue,omitempty"`

	// DisplayName: Required. The name of the parameter.
	DisplayName string `json:"displayName,omitempty"`

	// EntityTypeDisplayName: Optional. The name of the entity type,
	// prefixed with `@`, that
	// describes values of the parameter. If the parameter is
	// required, this must be provided.
	EntityTypeDisplayName string `json:"entityTypeDisplayName,omitempty"`

	// IsList: Optional. Indicates whether the parameter represents a list
	// of values.
	IsList bool `json:"isList,omitempty"`

	// Mandatory: Optional. Indicates whether the parameter is required.
	// That is,
	// whether the intent cannot be completed without collecting the
	// parameter
	// value.
	Mandatory bool `json:"mandatory,omitempty"`

	// Name: The unique identifier of this parameter.
	Name string `json:"name,omitempty"`

	// Prompts: Optional. The collection of prompts that the agent can
	// present to the
	// user in order to collect value for the parameter.
	Prompts []string `json:"prompts,omitempty"`

	// Value: Optional. The definition of the parameter value. It can be:
	// - a constant string,
	// - a parameter value defined as `$parameter_name`,
	// - an original parameter value defined as
	// `$parameter_name.original`,
	// - a parameter value from some context defined as
	//   `#context_name.parameter_name`.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DefaultValue") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentParameter) MarshalJSON() ([]byte, error) {
	type NoMethod IntentParameter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentTrainingPhrase: Represents an example or template that the
// agent is trained on.
type IntentTrainingPhrase struct {
	// Name: Required. The unique identifier of this training phrase.
	Name string `json:"name,omitempty"`

	// Parts: Required. The collection of training phrase parts (can be
	// annotated).
	// Fields: `entity_type`, `alias` and `user_defined` should be
	// populated
	// only for the annotated parts of the training phrase.
	Parts []*IntentTrainingPhrasePart `json:"parts,omitempty"`

	// TimesAddedCount: Optional. Indicates how many times this example or
	// template was added to
	// the intent. Each time a developer adds an existing sample by editing
	// an
	// intent or training, this counter is increased.
	TimesAddedCount int64 `json:"timesAddedCount,omitempty"`

	// Type: Required. The type of the training phrase.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Not specified. This value should never be
	// used.
	//   "EXAMPLE" - Examples do not contain @-prefixed entity type names,
	// but example parts
	// can be annotated with entity types.
	//   "TEMPLATE" - Templates are not annotated with entity types, but
	// they can contain
	// @-prefixed entity type names as substrings.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentTrainingPhrase) MarshalJSON() ([]byte, error) {
	type NoMethod IntentTrainingPhrase
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// IntentTrainingPhrasePart: Represents a part of a training phrase.
type IntentTrainingPhrasePart struct {
	// Alias: Optional. The parameter name for the value extracted from
	// the
	// annotated part of the example.
	Alias string `json:"alias,omitempty"`

	// EntityType: Optional. The entity type name prefixed with `@`. This
	// field is
	// required for the annotated part of the text and applies only
	// to
	// examples.
	EntityType string `json:"entityType,omitempty"`

	// Text: Required. The text corresponding to the example or template,
	// if there are no annotations. For
	// annotated examples, it is the text for one of the example's parts.
	Text string `json:"text,omitempty"`

	// UserDefined: Optional. Indicates whether the text was manually
	// annotated by the
	// developer.
	UserDefined bool `json:"userDefined,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alias") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alias") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *IntentTrainingPhrasePart) MarshalJSON() ([]byte, error) {
	type NoMethod IntentTrainingPhrasePart
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LatLng: An object representing a latitude/longitude pair. This is
// expressed as a pair
// of doubles representing degrees latitude and degrees longitude.
// Unless
// specified otherwise, this must conform to the
// <a
// href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// st
// andard</a>. Values must be within normalized ranges.
type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Latitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Latitude") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LatLng) MarshalJSON() ([]byte, error) {
	type NoMethod LatLng
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LatLng) UnmarshalJSON(data []byte) error {
	type NoMethod LatLng
	var s1 struct {
		Latitude  gensupport.JSONFloat64 `json:"latitude"`
		Longitude gensupport.JSONFloat64 `json:"longitude"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Latitude = float64(s1.Latitude)
	s.Longitude = float64(s1.Longitude)
	return nil
}

// ListContextsResponse: The response message for Contexts.ListContexts.
type ListContextsResponse struct {
	// Contexts: The list of contexts. There will be a maximum number of
	// items
	// returned based on the page_size field in the request.
	Contexts []*Context `json:"contexts,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Contexts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contexts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListContextsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListContextsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListEntityTypesResponse: The response message for
// EntityTypes.ListEntityTypes.
type ListEntityTypesResponse struct {
	// EntityTypes: The list of agent entity types. There will be a maximum
	// number of items
	// returned based on the page_size field in the request.
	EntityTypes []*EntityType `json:"entityTypes,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "EntityTypes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EntityTypes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListIntentsResponse: The response message for Intents.ListIntents.
type ListIntentsResponse struct {
	// Intents: The list of agent intents. There will be a maximum number of
	// items
	// returned based on the page_size field in the request.
	Intents []*Intent `json:"intents,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Intents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Intents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListIntentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListIntentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListSessionEntityTypesResponse: The response message for
// SessionEntityTypes.ListSessionEntityTypes.
type ListSessionEntityTypesResponse struct {
	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SessionEntityTypes: The list of session entity types. There will be a
	// maximum number of items
	// returned based on the page_size field in the request.
	SessionEntityTypes []*SessionEntityType `json:"sessionEntityTypes,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListSessionEntityTypesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListSessionEntityTypesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type NoMethod Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OriginalDetectIntentRequest: Represents the contents of the original
// request that was passed to
// the `[Streaming]DetectIntent` call.
type OriginalDetectIntentRequest struct {
	// Payload: Optional. This field is set to the value of
	// `QueryParameters.payload` field
	// passed in the request.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Source: The source of this request, e.g., `google`, `facebook`,
	// `slack`. It is set
	// by Dialogflow-owned servers. Possible values of this field correspond
	// to
	// Intent.Message.Platform.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Payload") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Payload") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OriginalDetectIntentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod OriginalDetectIntentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryInput: Represents the query input. It can contain either:
//
// 1.  an audio config which
//     instructs the speech recognizer how to process the speech
// audio,
//
// 2.  a conversational query in the form of text, or
//
// 3.  an event that specifies which intent to trigger.
type QueryInput struct {
	// AudioConfig: Instructs the speech recognizer how to process the
	// speech audio.
	AudioConfig *InputAudioConfig `json:"audioConfig,omitempty"`

	// Event: The event to be processed.
	Event *EventInput `json:"event,omitempty"`

	// Text: The natural language text to be processed.
	Text *TextInput `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AudioConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AudioConfig") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryInput) MarshalJSON() ([]byte, error) {
	type NoMethod QueryInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryParameters: Represents the parameters of the conversational
// query.
type QueryParameters struct {
	// Contexts: Optional. The collection of contexts to be activated before
	// this query is
	// executed.
	Contexts []*Context `json:"contexts,omitempty"`

	// GeoLocation: Optional. The geo location of this conversational query.
	GeoLocation *LatLng `json:"geoLocation,omitempty"`

	// Payload: Optional. This field can be used to pass custom data into
	// the webhook
	// associated with the agent. Arbitrary JSON objects are supported.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// ResetContexts: Optional. Specifies whether to delete all contexts in
	// the current session
	// before the new ones are activated.
	ResetContexts bool `json:"resetContexts,omitempty"`

	// SessionEntityTypes: Optional. The collection of session entity types
	// to replace or extend
	// developer entities with for this query only. The entity synonyms
	// apply
	// to all languages.
	SessionEntityTypes []*SessionEntityType `json:"sessionEntityTypes,omitempty"`

	// TimeZone: Optional. The time zone of this conversational query from
	// the
	// [time zone database](https://www.iana.org/time-zones),
	// e.g.,
	// America/New_York, Europe/Paris. If not provided, the time zone
	// specified in
	// agent settings is used.
	TimeZone string `json:"timeZone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contexts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contexts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryParameters) MarshalJSON() ([]byte, error) {
	type NoMethod QueryParameters
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryResult: Represents the result of conversational query or event
// processing.
type QueryResult struct {
	// Action: The action name from the matched intent.
	Action string `json:"action,omitempty"`

	// AllRequiredParamsPresent: This field is set to:
	// - `false` if the matched intent has required parameters and not all
	// of
	//    the required parameter values have been collected.
	// - `true` if all required parameter values have been collected, or if
	// the
	//    matched intent doesn't contain any required parameters.
	AllRequiredParamsPresent bool `json:"allRequiredParamsPresent,omitempty"`

	// DiagnosticInfo: The free-form diagnostic info. For example, this
	// field
	// could contain webhook call latency.
	DiagnosticInfo googleapi.RawMessage `json:"diagnosticInfo,omitempty"`

	// FulfillmentMessages: The collection of rich messages to present to
	// the user.
	FulfillmentMessages []*IntentMessage `json:"fulfillmentMessages,omitempty"`

	// FulfillmentText: The text to be pronounced to the user or shown on
	// the screen.
	FulfillmentText string `json:"fulfillmentText,omitempty"`

	// Intent: The intent that matched the conversational query. Some,
	// not
	// all fields are filled in this message, including but not limited
	// to:
	// `name`, `display_name` and `webhook_state`.
	Intent *Intent `json:"intent,omitempty"`

	// IntentDetectionConfidence: The intent detection confidence. Values
	// range from 0.0
	// (completely uncertain) to 1.0 (completely certain).
	IntentDetectionConfidence float64 `json:"intentDetectionConfidence,omitempty"`

	// LanguageCode: The language that was triggered during intent
	// detection.
	// See [Language
	// Support](https://dialogflow.com/docs/reference/language)
	// for a list of the currently supported language codes.
	LanguageCode string `json:"languageCode,omitempty"`

	// OutputContexts: The collection of output contexts. If
	// applicable,
	// `output_contexts.parameters` contains entries with name
	// `<parameter name>.original` containing the original parameter
	// values
	// before the query.
	OutputContexts []*Context `json:"outputContexts,omitempty"`

	// Parameters: The collection of extracted parameters.
	Parameters googleapi.RawMessage `json:"parameters,omitempty"`

	// QueryText: The original conversational query text:
	// - If natural language text was provided as input, `query_text`
	// contains
	//   a copy of the input.
	// - If natural language speech audio was provided as input,
	// `query_text`
	//   contains the speech recognition result. If speech recognizer
	// produced
	//   multiple alternatives, a particular one is picked.
	// - If an event was provided as input, `query_text` is not set.
	QueryText string `json:"queryText,omitempty"`

	// SpeechRecognitionConfidence: The confidence estimate between 0.0 and
	// 1.0. A higher number
	// indicates an estimated greater likelihood that the recognized words
	// are
	// correct. The default of 0.0 is a sentinel value indicating that
	// confidence
	// was not set. This field is populated if natural speech audio was
	// provided
	// as input.
	SpeechRecognitionConfidence float64 `json:"speechRecognitionConfidence,omitempty"`

	// WebhookPayload: If the query was fulfilled by a webhook call, this
	// field is set to the
	// value of the `payload` field returned in the webhook response.
	WebhookPayload googleapi.RawMessage `json:"webhookPayload,omitempty"`

	// WebhookSource: If the query was fulfilled by a webhook call, this
	// field is set to the
	// value of the `source` field returned in the webhook response.
	WebhookSource string `json:"webhookSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryResult) MarshalJSON() ([]byte, error) {
	type NoMethod QueryResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *QueryResult) UnmarshalJSON(data []byte) error {
	type NoMethod QueryResult
	var s1 struct {
		IntentDetectionConfidence   gensupport.JSONFloat64 `json:"intentDetectionConfidence"`
		SpeechRecognitionConfidence gensupport.JSONFloat64 `json:"speechRecognitionConfidence"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.IntentDetectionConfidence = float64(s1.IntentDetectionConfidence)
	s.SpeechRecognitionConfidence = float64(s1.SpeechRecognitionConfidence)
	return nil
}

// RestoreAgentRequest: The request message for Agents.RestoreAgent.
type RestoreAgentRequest struct {
	// AgentContent: The agent to restore.
	AgentContent string `json:"agentContent,omitempty"`

	// AgentUri: Warning: Restoring agents from a URI is not implemented
	// yet.
	// This feature is coming soon.
	// The URI to a Google Cloud Storage file containing the agent to
	// restore.
	// Note: The URI must start with "gs://".
	AgentUri string `json:"agentUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentContent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentContent") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RestoreAgentRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RestoreAgentRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SearchAgentsResponse: The response message for Agents.SearchAgents.
type SearchAgentsResponse struct {
	// Agents: The list of agents. There will be a maximum number of items
	// returned based
	// on the page_size field in the request.
	Agents []*Agent `json:"agents,omitempty"`

	// NextPageToken: Token to retrieve the next page of results, or empty
	// if there are no
	// more results in the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Agents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Agents") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SearchAgentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod SearchAgentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SessionEntityType: Represents a session entity type.
//
// Extends or replaces a developer entity type at the user session level
// (we
// refer to the entity types defined at the agent level as "developer
// entity
// types").
//
// Note: session entity types apply to all queries, regardless of the
// language.
type SessionEntityType struct {
	// Entities: Required. The collection of entities associated with this
	// session entity
	// type.
	Entities []*EntityTypeEntity `json:"entities,omitempty"`

	// EntityOverrideMode: Required. Indicates whether the additional data
	// should override or
	// supplement the developer entity type definition.
	//
	// Possible values:
	//   "ENTITY_OVERRIDE_MODE_UNSPECIFIED" - Not specified. This value
	// should be never used.
	//   "ENTITY_OVERRIDE_MODE_OVERRIDE" - The collection of session
	// entities overrides the collection of entities
	// in the corresponding developer entity type.
	//   "ENTITY_OVERRIDE_MODE_SUPPLEMENT" - The collection of session
	// entities extends the collection of entities in
	// the corresponding developer entity type.
	// Calls to `ListSessionEntityTypes`,
	// `GetSessionEntityType`,
	// `CreateSessionEntityType` and `UpdateSessionEntityType` return the
	// full
	// collection of entities from the developer entity type in the
	// agent's
	// default language and the session entity type.
	EntityOverrideMode string `json:"entityOverrideMode,omitempty"`

	// Name: Required. The unique identifier of this session entity type.
	// Format:
	// `projects/<Project ID>/agent/sessions/<Session
	// ID>/entityTypes/<Entity Type
	// Display Name>`.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Entities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Entities") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SessionEntityType) MarshalJSON() ([]byte, error) {
	type NoMethod SessionEntityType
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextInput: Represents the natural language text to be processed.
type TextInput struct {
	// LanguageCode: Required. The language of this conversational query.
	// See [Language
	// Support](https://dialogflow.com/docs/languages) for a list of
	// the
	// currently supported language codes. Note that queries in the same
	// session
	// do not necessarily need to specify the same language.
	LanguageCode string `json:"languageCode,omitempty"`

	// Text: Required. The UTF-8 encoded natural language text to be
	// processed.
	// Text length must not exceed 256 bytes.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TextInput) MarshalJSON() ([]byte, error) {
	type NoMethod TextInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TrainAgentRequest: The request message for Agents.TrainAgent.
type TrainAgentRequest struct {
}

// WebhookRequest: The request message for a webhook call.
type WebhookRequest struct {
	// OriginalDetectIntentRequest: Optional. The contents of the original
	// request that was passed to
	// `[Streaming]DetectIntent` call.
	OriginalDetectIntentRequest *OriginalDetectIntentRequest `json:"originalDetectIntentRequest,omitempty"`

	// QueryResult: The result of the conversational query or event
	// processing. Contains the
	// same value as `[Streaming]DetectIntentResponse.query_result`.
	QueryResult *QueryResult `json:"queryResult,omitempty"`

	// ResponseId: The unique identifier of the response. Contains the same
	// value as
	// `[Streaming]DetectIntentResponse.response_id`.
	ResponseId string `json:"responseId,omitempty"`

	// Session: The unique identifier of detectIntent request session.
	// Can be used to identify end-user inside webhook
	// implementation.
	// Format: `projects/<Project ID>/agent/sessions/<Session ID>`.
	Session string `json:"session,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "OriginalDetectIntentRequest") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "OriginalDetectIntentRequest") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WebhookRequest) MarshalJSON() ([]byte, error) {
	type NoMethod WebhookRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WebhookResponse: The response message for a webhook call.
type WebhookResponse struct {
	// FollowupEventInput: Optional. Makes the platform immediately invoke
	// another `DetectIntent` call
	// internally with the specified event as input.
	FollowupEventInput *EventInput `json:"followupEventInput,omitempty"`

	// FulfillmentMessages: Optional. The collection of rich messages to
	// present to the user. This
	// value is passed directly to `QueryResult.fulfillment_messages`.
	FulfillmentMessages []*IntentMessage `json:"fulfillmentMessages,omitempty"`

	// FulfillmentText: Optional. The text to be shown on the screen. This
	// value is passed directly
	// to `QueryResult.fulfillment_text`.
	FulfillmentText string `json:"fulfillmentText,omitempty"`

	// OutputContexts: Optional. The collection of output contexts. This
	// value is passed directly
	// to `QueryResult.output_contexts`.
	OutputContexts []*Context `json:"outputContexts,omitempty"`

	// Payload: Optional. This value is passed directly to
	// `QueryResult.webhook_payload`.
	Payload googleapi.RawMessage `json:"payload,omitempty"`

	// Source: Optional. This value is passed directly to
	// `QueryResult.webhook_source`.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FollowupEventInput")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FollowupEventInput") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *WebhookResponse) MarshalJSON() ([]byte, error) {
	type NoMethod WebhookResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "dialogflow.projects.getAgent":

type ProjectsGetAgentCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetAgent: Retrieves the specified agent.
func (r *ProjectsService) GetAgent(parent string) *ProjectsGetAgentCall {
	c := &ProjectsGetAgentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetAgentCall) Fields(s ...googleapi.Field) *ProjectsGetAgentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetAgentCall) IfNoneMatch(entityTag string) *ProjectsGetAgentCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetAgentCall) Context(ctx context.Context) *ProjectsGetAgentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetAgentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetAgentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.getAgent" call.
// Exactly one of *Agent or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Agent.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGetAgentCall) Do(opts ...googleapi.CallOption) (*Agent, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Agent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.getAgent",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to fetch is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent",
	//   "response": {
	//     "$ref": "Agent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.export":

type ProjectsAgentExportCall struct {
	s                  *Service
	parent             string
	exportagentrequest *ExportAgentRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Export: Exports the specified agent to a ZIP file.
//
//
// Operation <response: ExportAgentResponse,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Export(parent string, exportagentrequest *ExportAgentRequest) *ProjectsAgentExportCall {
	c := &ProjectsAgentExportCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.exportagentrequest = exportagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentExportCall) Fields(s ...googleapi.Field) *ProjectsAgentExportCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentExportCall) Context(ctx context.Context) *ProjectsAgentExportCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentExportCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentExportCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:export")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.export" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentExportCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports the specified agent to a ZIP file.\n\n\nOperation \u003cresponse: ExportAgentResponse,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:export",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.export",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to export is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:export",
	//   "request": {
	//     "$ref": "ExportAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.import":

type ProjectsAgentImportCall struct {
	s                  *Service
	parent             string
	importagentrequest *ImportAgentRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Import: Imports the specified agent from a ZIP file.
//
// Uploads new intents and entity types without deleting the existing
// ones.
// Intents and entity types with the same name are replaced with the
// new
// versions from ImportAgentRequest.
//
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Import(parent string, importagentrequest *ImportAgentRequest) *ProjectsAgentImportCall {
	c := &ProjectsAgentImportCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.importagentrequest = importagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentImportCall) Fields(s ...googleapi.Field) *ProjectsAgentImportCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentImportCall) Context(ctx context.Context) *ProjectsAgentImportCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentImportCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentImportCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:import")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.import" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentImportCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Imports the specified agent from a ZIP file.\n\nUploads new intents and entity types without deleting the existing ones.\nIntents and entity types with the same name are replaced with the new\nversions from ImportAgentRequest.\n\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:import",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.import",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to import is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:import",
	//   "request": {
	//     "$ref": "ImportAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.restore":

type ProjectsAgentRestoreCall struct {
	s                   *Service
	parent              string
	restoreagentrequest *RestoreAgentRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Restore: Restores the specified agent from a ZIP file.
//
// Replaces the current agent version with a new one. All the intents
// and
// entity types in the older version are deleted.
//
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Restore(parent string, restoreagentrequest *RestoreAgentRequest) *ProjectsAgentRestoreCall {
	c := &ProjectsAgentRestoreCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.restoreagentrequest = restoreagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentRestoreCall) Fields(s ...googleapi.Field) *ProjectsAgentRestoreCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentRestoreCall) Context(ctx context.Context) *ProjectsAgentRestoreCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentRestoreCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentRestoreCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.restoreagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:restore")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.restore" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentRestoreCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restores the specified agent from a ZIP file.\n\nReplaces the current agent version with a new one. All the intents and\nentity types in the older version are deleted.\n\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:restore",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.restore",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to restore is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:restore",
	//   "request": {
	//     "$ref": "RestoreAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.search":

type ProjectsAgentSearchCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Search: Returns the list of agents.
//
// Since there is at most one conversational agent per project, this
// method is
// useful primarily for listing all agents across projects the caller
// has
// access to. One can achieve that with a wildcard project collection id
// "-".
// Refer to
// [List
// Sub-Collections](https://cloud.google.com/apis/design/design_pat
// terns#list_sub-collections).
func (r *ProjectsAgentService) Search(parent string) *ProjectsAgentSearchCall {
	c := &ProjectsAgentSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentSearchCall) PageSize(pageSize int64) *ProjectsAgentSearchCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentSearchCall) PageToken(pageToken string) *ProjectsAgentSearchCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSearchCall) Fields(s ...googleapi.Field) *ProjectsAgentSearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSearchCall) IfNoneMatch(entityTag string) *ProjectsAgentSearchCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSearchCall) Context(ctx context.Context) *ProjectsAgentSearchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSearchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.search" call.
// Exactly one of *SearchAgentsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SearchAgentsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSearchCall) Do(opts ...googleapi.CallOption) (*SearchAgentsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SearchAgentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of agents.\n\nSince there is at most one conversational agent per project, this method is\nuseful primarily for listing all agents across projects the caller has\naccess to. One can achieve that with a wildcard project collection id \"-\".\nRefer to [List\nSub-Collections](https://cloud.google.com/apis/design/design_patterns#list_sub-collections).",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:search",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.search",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The project to list agents from.\nFormat: `projects/\u003cProject ID or '-'\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:search",
	//   "response": {
	//     "$ref": "SearchAgentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentSearchCall) Pages(ctx context.Context, f func(*SearchAgentsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dialogflow.projects.agent.train":

type ProjectsAgentTrainCall struct {
	s                 *Service
	parent            string
	trainagentrequest *TrainAgentRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Train: Trains the specified agent.
//
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentService) Train(parent string, trainagentrequest *TrainAgentRequest) *ProjectsAgentTrainCall {
	c := &ProjectsAgentTrainCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.trainagentrequest = trainagentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentTrainCall) Fields(s ...googleapi.Field) *ProjectsAgentTrainCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentTrainCall) Context(ctx context.Context) *ProjectsAgentTrainCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentTrainCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentTrainCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.trainagentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/agent:train")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.train" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentTrainCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Trains the specified agent.\n\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent:train",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.train",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project that the agent to train is associated with.\nFormat: `projects/\u003cProject ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/agent:train",
	//   "request": {
	//     "$ref": "TrainAgentRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.batchDelete":

type ProjectsAgentEntityTypesBatchDeleteCall struct {
	s                             *Service
	parent                        string
	batchdeleteentitytypesrequest *BatchDeleteEntityTypesRequest
	urlParams_                    gensupport.URLParams
	ctx_                          context.Context
	header_                       http.Header
}

// BatchDelete: Deletes entity types in the specified agent.
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesService) BatchDelete(parent string, batchdeleteentitytypesrequest *BatchDeleteEntityTypesRequest) *ProjectsAgentEntityTypesBatchDeleteCall {
	c := &ProjectsAgentEntityTypesBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchdeleteentitytypesrequest = batchdeleteentitytypesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Context(ctx context.Context) *ProjectsAgentEntityTypesBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchdeleteentitytypesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.batchDelete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesBatchDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes entity types in the specified agent.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.batchDelete",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to delete all entities types for. Format:\n`projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes:batchDelete",
	//   "request": {
	//     "$ref": "BatchDeleteEntityTypesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.batchUpdate":

type ProjectsAgentEntityTypesBatchUpdateCall struct {
	s                             *Service
	parent                        string
	batchupdateentitytypesrequest *BatchUpdateEntityTypesRequest
	urlParams_                    gensupport.URLParams
	ctx_                          context.Context
	header_                       http.Header
}

// BatchUpdate: Updates/Creates multiple entity types in the specified
// agent.
//
// Operation <response: BatchUpdateEntityTypesResponse,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesService) BatchUpdate(parent string, batchupdateentitytypesrequest *BatchUpdateEntityTypesRequest) *ProjectsAgentEntityTypesBatchUpdateCall {
	c := &ProjectsAgentEntityTypesBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchupdateentitytypesrequest = batchupdateentitytypesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchupdateentitytypesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.batchUpdate" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesBatchUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates/Creates multiple entity types in the specified agent.\n\nOperation \u003cresponse: BatchUpdateEntityTypesResponse,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes:batchUpdate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to update or create entity types in.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes:batchUpdate",
	//   "request": {
	//     "$ref": "BatchUpdateEntityTypesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.create":

type ProjectsAgentEntityTypesCreateCall struct {
	s          *Service
	parent     string
	entitytype *EntityType
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates an entity type in the specified agent.
func (r *ProjectsAgentEntityTypesService) Create(parent string, entitytype *EntityType) *ProjectsAgentEntityTypesCreateCall {
	c := &ProjectsAgentEntityTypesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.entitytype = entitytype
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of entity synonyms defined in `entity_type`. If not
// specified, the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesCreateCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesCreateCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesCreateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.entitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.create" call.
// Exactly one of *EntityType or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *EntityType.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesCreateCall) Do(opts ...googleapi.CallOption) (*EntityType, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &EntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an entity type in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language of entity synonyms defined in `entity_type`. If not\nspecified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to create a entity type for.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "request": {
	//     "$ref": "EntityType"
	//   },
	//   "response": {
	//     "$ref": "EntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.delete":

type ProjectsAgentEntityTypesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified entity type.
func (r *ProjectsAgentEntityTypesService) Delete(name string) *ProjectsAgentEntityTypesDeleteCall {
	c := &ProjectsAgentEntityTypesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesDeleteCall) Context(ctx context.Context) *ProjectsAgentEntityTypesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentEntityTypesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.entityTypes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the entity type to delete.\nFormat: `projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntityType ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.get":

type ProjectsAgentEntityTypesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified entity type.
func (r *ProjectsAgentEntityTypesService) Get(name string) *ProjectsAgentEntityTypesGetCall {
	c := &ProjectsAgentEntityTypesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to retrieve entity synonyms for. If not specified,
// the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesGetCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesGetCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesGetCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEntityTypesGetCall) IfNoneMatch(entityTag string) *ProjectsAgentEntityTypesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesGetCall) Context(ctx context.Context) *ProjectsAgentEntityTypesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.get" call.
// Exactly one of *EntityType or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *EntityType.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesGetCall) Do(opts ...googleapi.CallOption) (*EntityType, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &EntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.entityTypes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language to retrieve entity synonyms for. If not specified,\nthe agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required. The name of the entity type.\nFormat: `projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntityType ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "EntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.list":

type ProjectsAgentEntityTypesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all entity types in the specified agent.
func (r *ProjectsAgentEntityTypesService) List(parent string) *ProjectsAgentEntityTypesListCall {
	c := &ProjectsAgentEntityTypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to list entity synonyms for. If not specified,
// the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesListCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentEntityTypesListCall) PageSize(pageSize int64) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentEntityTypesListCall) PageToken(pageToken string) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesListCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentEntityTypesListCall) IfNoneMatch(entityTag string) *ProjectsAgentEntityTypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesListCall) Context(ctx context.Context) *ProjectsAgentEntityTypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.list" call.
// Exactly one of *ListEntityTypesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListEntityTypesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesListCall) Do(opts ...googleapi.CallOption) (*ListEntityTypesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListEntityTypesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all entity types in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.entityTypes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language to list entity synonyms for. If not specified,\nthe agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to list all entity types from.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "response": {
	//     "$ref": "ListEntityTypesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentEntityTypesListCall) Pages(ctx context.Context, f func(*ListEntityTypesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dialogflow.projects.agent.entityTypes.patch":

type ProjectsAgentEntityTypesPatchCall struct {
	s          *Service
	nameid     string
	entitytype *EntityType
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the specified entity type.
func (r *ProjectsAgentEntityTypesService) Patch(nameid string, entitytype *EntityType) *ProjectsAgentEntityTypesPatchCall {
	c := &ProjectsAgentEntityTypesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.entitytype = entitytype
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of entity synonyms defined in `entity_type`. If not
// specified, the agent's default language is used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentEntityTypesPatchCall) LanguageCode(languageCode string) *ProjectsAgentEntityTypesPatchCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentEntityTypesPatchCall) UpdateMask(updateMask string) *ProjectsAgentEntityTypesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesPatchCall) Context(ctx context.Context) *ProjectsAgentEntityTypesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.entitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.patch" call.
// Exactly one of *EntityType or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *EntityType.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesPatchCall) Do(opts ...googleapi.CallOption) (*EntityType, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &EntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.entityTypes.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "languageCode": {
	//       "description": "Optional. The language of entity synonyms defined in `entity_type`. If not\nspecified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required for all methods except `create` (`create` populates the name\nautomatically.\nThe unique identifier of the entity type. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "EntityType"
	//   },
	//   "response": {
	//     "$ref": "EntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.entities.batchCreate":

type ProjectsAgentEntityTypesEntitiesBatchCreateCall struct {
	s                          *Service
	parent                     string
	batchcreateentitiesrequest *BatchCreateEntitiesRequest
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
	header_                    http.Header
}

// BatchCreate: Creates multiple new entities in the specified entity
// type (extends the
// existing collection of entries).
//
// Operation <response: google.protobuf.Empty>
func (r *ProjectsAgentEntityTypesEntitiesService) BatchCreate(parent string, batchcreateentitiesrequest *BatchCreateEntitiesRequest) *ProjectsAgentEntityTypesEntitiesBatchCreateCall {
	c := &ProjectsAgentEntityTypesEntitiesBatchCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchcreateentitiesrequest = batchcreateentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesEntitiesBatchCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesEntitiesBatchCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchcreateentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entities:batchCreate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.entities.batchCreate" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesEntitiesBatchCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates multiple new entities in the specified entity type (extends the\nexisting collection of entries).\n\nOperation \u003cresponse: google.protobuf.Empty\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}/entities:batchCreate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.entities.batchCreate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the entity type to create entities in. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entities:batchCreate",
	//   "request": {
	//     "$ref": "BatchCreateEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.entities.batchDelete":

type ProjectsAgentEntityTypesEntitiesBatchDeleteCall struct {
	s                          *Service
	parent                     string
	batchdeleteentitiesrequest *BatchDeleteEntitiesRequest
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
	header_                    http.Header
}

// BatchDelete: Deletes entities in the specified entity
// type.
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesEntitiesService) BatchDelete(parent string, batchdeleteentitiesrequest *BatchDeleteEntitiesRequest) *ProjectsAgentEntityTypesEntitiesBatchDeleteCall {
	c := &ProjectsAgentEntityTypesEntitiesBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchdeleteentitiesrequest = batchdeleteentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesEntitiesBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Context(ctx context.Context) *ProjectsAgentEntityTypesEntitiesBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchdeleteentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entities:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.entities.batchDelete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesEntitiesBatchDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes entities in the specified entity type.\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}/entities:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.entities.batchDelete",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the entity type to delete entries for. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entities:batchDelete",
	//   "request": {
	//     "$ref": "BatchDeleteEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.entityTypes.entities.batchUpdate":

type ProjectsAgentEntityTypesEntitiesBatchUpdateCall struct {
	s                          *Service
	parent                     string
	batchupdateentitiesrequest *BatchUpdateEntitiesRequest
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
	header_                    http.Header
}

// BatchUpdate: Updates entities in the specified entity type (replaces
// the existing
// collection of entries).
//
// Operation <response: google.protobuf.Empty,
//            metadata: google.protobuf.Struct>
func (r *ProjectsAgentEntityTypesEntitiesService) BatchUpdate(parent string, batchupdateentitiesrequest *BatchUpdateEntitiesRequest) *ProjectsAgentEntityTypesEntitiesBatchUpdateCall {
	c := &ProjectsAgentEntityTypesEntitiesBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchupdateentitiesrequest = batchupdateentitiesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsAgentEntityTypesEntitiesBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Context(ctx context.Context) *ProjectsAgentEntityTypesEntitiesBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchupdateentitiesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entities:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.entityTypes.entities.batchUpdate" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentEntityTypesEntitiesBatchUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates entities in the specified entity type (replaces the existing\ncollection of entries).\n\nOperation \u003cresponse: google.protobuf.Empty,\n           metadata: google.protobuf.Struct\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/entityTypes/{entityTypesId}/entities:batchUpdate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.entityTypes.entities.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the entity type to update the entities in. Format:\n`projects/\u003cProject ID\u003e/agent/entityTypes/\u003cEntity Type ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entities:batchUpdate",
	//   "request": {
	//     "$ref": "BatchUpdateEntitiesRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.batchDelete":

type ProjectsAgentIntentsBatchDeleteCall struct {
	s                         *Service
	parent                    string
	batchdeleteintentsrequest *BatchDeleteIntentsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// BatchDelete: Deletes intents in the specified agent.
//
// Operation <response: google.protobuf.Empty>
func (r *ProjectsAgentIntentsService) BatchDelete(parent string, batchdeleteintentsrequest *BatchDeleteIntentsRequest) *ProjectsAgentIntentsBatchDeleteCall {
	c := &ProjectsAgentIntentsBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchdeleteintentsrequest = batchdeleteintentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsBatchDeleteCall) Context(ctx context.Context) *ProjectsAgentIntentsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsBatchDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchdeleteintentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.batchDelete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentIntentsBatchDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes intents in the specified agent.\n\nOperation \u003cresponse: google.protobuf.Empty\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.intents.batchDelete",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to delete all entities types for. Format:\n`projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents:batchDelete",
	//   "request": {
	//     "$ref": "BatchDeleteIntentsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.batchUpdate":

type ProjectsAgentIntentsBatchUpdateCall struct {
	s                         *Service
	parent                    string
	batchupdateintentsrequest *BatchUpdateIntentsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// BatchUpdate: Updates/Creates multiple intents in the specified
// agent.
//
// Operation <response: BatchUpdateIntentsResponse>
func (r *ProjectsAgentIntentsService) BatchUpdate(parent string, batchupdateintentsrequest *BatchUpdateIntentsRequest) *ProjectsAgentIntentsBatchUpdateCall {
	c := &ProjectsAgentIntentsBatchUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.batchupdateintentsrequest = batchupdateintentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsBatchUpdateCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsBatchUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsBatchUpdateCall) Context(ctx context.Context) *ProjectsAgentIntentsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsBatchUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchupdateintentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents:batchUpdate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.batchUpdate" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAgentIntentsBatchUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates/Creates multiple intents in the specified agent.\n\nOperation \u003cresponse: BatchUpdateIntentsResponse\u003e",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents:batchUpdate",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.intents.batchUpdate",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the agent to update or create intents in.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents:batchUpdate",
	//   "request": {
	//     "$ref": "BatchUpdateIntentsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.create":

type ProjectsAgentIntentsCreateCall struct {
	s          *Service
	parent     string
	intent     *Intent
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates an intent in the specified agent.
func (r *ProjectsAgentIntentsService) Create(parent string, intent *Intent) *ProjectsAgentIntentsCreateCall {
	c := &ProjectsAgentIntentsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.intent = intent
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsCreateCall) IntentView(intentView string) *ProjectsAgentIntentsCreateCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of training phrases, parameters and rich messages
// defined in `intent`. If not specified, the agent's default language
// is
// used. [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentIntentsCreateCall) LanguageCode(languageCode string) *ProjectsAgentIntentsCreateCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsCreateCall) Context(ctx context.Context) *ProjectsAgentIntentsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.intent)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.create" call.
// Exactly one of *Intent or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Intent.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentIntentsCreateCall) Do(opts ...googleapi.CallOption) (*Intent, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Intent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an intent in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.intents.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language of training phrases, parameters and rich messages\ndefined in `intent`. If not specified, the agent's default language is\nused. [More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to create a intent for.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents",
	//   "request": {
	//     "$ref": "Intent"
	//   },
	//   "response": {
	//     "$ref": "Intent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.delete":

type ProjectsAgentIntentsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified intent.
func (r *ProjectsAgentIntentsService) Delete(name string) *ProjectsAgentIntentsDeleteCall {
	c := &ProjectsAgentIntentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsDeleteCall) Context(ctx context.Context) *ProjectsAgentIntentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentIntentsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified intent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents/{intentsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.intents.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the intent to delete.\nFormat: `projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/intents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.get":

type ProjectsAgentIntentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified intent.
func (r *ProjectsAgentIntentsService) Get(name string) *ProjectsAgentIntentsGetCall {
	c := &ProjectsAgentIntentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsGetCall) IntentView(intentView string) *ProjectsAgentIntentsGetCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to retrieve training phrases, parameters and rich
// messages for. If not specified, the agent's default language is
// used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentIntentsGetCall) LanguageCode(languageCode string) *ProjectsAgentIntentsGetCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsGetCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentIntentsGetCall) IfNoneMatch(entityTag string) *ProjectsAgentIntentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsGetCall) Context(ctx context.Context) *ProjectsAgentIntentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.get" call.
// Exactly one of *Intent or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Intent.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentIntentsGetCall) Do(opts ...googleapi.CallOption) (*Intent, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Intent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified intent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents/{intentsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.intents.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language to retrieve training phrases, parameters and rich\nmessages for. If not specified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required. The name of the intent.\nFormat: `projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/intents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Intent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.intents.list":

type ProjectsAgentIntentsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all intents in the specified agent.
func (r *ProjectsAgentIntentsService) List(parent string) *ProjectsAgentIntentsListCall {
	c := &ProjectsAgentIntentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsListCall) IntentView(intentView string) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// to list training phrases, parameters and rich
// messages for. If not specified, the agent's default language is
// used.
// [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent before they can be used.
func (c *ProjectsAgentIntentsListCall) LanguageCode(languageCode string) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentIntentsListCall) PageSize(pageSize int64) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentIntentsListCall) PageToken(pageToken string) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsListCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentIntentsListCall) IfNoneMatch(entityTag string) *ProjectsAgentIntentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsListCall) Context(ctx context.Context) *ProjectsAgentIntentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/intents")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.list" call.
// Exactly one of *ListIntentsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListIntentsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentIntentsListCall) Do(opts ...googleapi.CallOption) (*ListIntentsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListIntentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all intents in the specified agent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.intents.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language to list training phrases, parameters and rich\nmessages for. If not specified, the agent's default language is used.\n[More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The agent to list all intents from.\nFormat: `projects/\u003cProject ID\u003e/agent`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/intents",
	//   "response": {
	//     "$ref": "ListIntentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentIntentsListCall) Pages(ctx context.Context, f func(*ListIntentsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dialogflow.projects.agent.intents.patch":

type ProjectsAgentIntentsPatchCall struct {
	s          *Service
	nameid     string
	intent     *Intent
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the specified intent.
func (r *ProjectsAgentIntentsService) Patch(nameid string, intent *Intent) *ProjectsAgentIntentsPatchCall {
	c := &ProjectsAgentIntentsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.intent = intent
	return c
}

// IntentView sets the optional parameter "intentView": The resource
// view to apply to the returned intent.
//
// Possible values:
//   "INTENT_VIEW_UNSPECIFIED"
//   "INTENT_VIEW_FULL"
func (c *ProjectsAgentIntentsPatchCall) IntentView(intentView string) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("intentView", intentView)
	return c
}

// LanguageCode sets the optional parameter "languageCode": The language
// of training phrases, parameters and rich messages
// defined in `intent`. If not specified, the agent's default language
// is
// used. [More than a
// dozen
// languages](https://dialogflow.com/docs/reference/language) are
// supported.
// Note: languages must be enabled in the agent, before they can be
// used.
func (c *ProjectsAgentIntentsPatchCall) LanguageCode(languageCode string) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("languageCode", languageCode)
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentIntentsPatchCall) UpdateMask(updateMask string) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentIntentsPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentIntentsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentIntentsPatchCall) Context(ctx context.Context) *ProjectsAgentIntentsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentIntentsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentIntentsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.intent)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.intents.patch" call.
// Exactly one of *Intent or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Intent.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentIntentsPatchCall) Do(opts ...googleapi.CallOption) (*Intent, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Intent{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified intent.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/intents/{intentsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.intents.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "intentView": {
	//       "description": "Optional. The resource view to apply to the returned intent.",
	//       "enum": [
	//         "INTENT_VIEW_UNSPECIFIED",
	//         "INTENT_VIEW_FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "languageCode": {
	//       "description": "Optional. The language of training phrases, parameters and rich messages\ndefined in `intent`. If not specified, the agent's default language is\nused. [More than a dozen\nlanguages](https://dialogflow.com/docs/reference/language) are supported.\nNote: languages must be enabled in the agent, before they can be used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Required for all methods except `create` (`create` populates the name\nautomatically.\nThe unique identifier of this intent.\nFormat: `projects/\u003cProject ID\u003e/agent/intents/\u003cIntent ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/intents/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "Intent"
	//   },
	//   "response": {
	//     "$ref": "Intent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.deleteContexts":

type ProjectsAgentSessionsDeleteContextsCall struct {
	s          *Service
	parent     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteContexts: Deletes all active contexts in the specified session.
func (r *ProjectsAgentSessionsService) DeleteContexts(parent string) *ProjectsAgentSessionsDeleteContextsCall {
	c := &ProjectsAgentSessionsDeleteContextsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsDeleteContextsCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsDeleteContextsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsDeleteContextsCall) Context(ctx context.Context) *ProjectsAgentSessionsDeleteContextsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsDeleteContextsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsDeleteContextsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.deleteContexts" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentSessionsDeleteContextsCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes all active contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.sessions.deleteContexts",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the session to delete all contexts from. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.detectIntent":

type ProjectsAgentSessionsDetectIntentCall struct {
	s                   *Service
	sessionid           string
	detectintentrequest *DetectIntentRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// DetectIntent: Processes a natural language query and returns
// structured, actionable data
// as a result. This method is not idempotent, because it may cause
// contexts
// and session entity types to be updated, which in turn might
// affect
// results of future queries.
func (r *ProjectsAgentSessionsService) DetectIntent(sessionid string, detectintentrequest *DetectIntentRequest) *ProjectsAgentSessionsDetectIntentCall {
	c := &ProjectsAgentSessionsDetectIntentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.sessionid = sessionid
	c.detectintentrequest = detectintentrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsDetectIntentCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsDetectIntentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsDetectIntentCall) Context(ctx context.Context) *ProjectsAgentSessionsDetectIntentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsDetectIntentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsDetectIntentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.detectintentrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+session}:detectIntent")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"session": c.sessionid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.detectIntent" call.
// Exactly one of *DetectIntentResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *DetectIntentResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsDetectIntentCall) Do(opts ...googleapi.CallOption) (*DetectIntentResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &DetectIntentResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Processes a natural language query and returns structured, actionable data\nas a result. This method is not idempotent, because it may cause contexts\nand session entity types to be updated, which in turn might affect\nresults of future queries.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}:detectIntent",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.sessions.detectIntent",
	//   "parameterOrder": [
	//     "session"
	//   ],
	//   "parameters": {
	//     "session": {
	//       "description": "Required. The name of the session this query is sent to. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`.\nIt's up to the API caller to choose an appropriate session ID. It can be\na random number or some type of user identifier (preferably hashed).\nThe length of the session ID must not exceed 36 bytes.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+session}:detectIntent",
	//   "request": {
	//     "$ref": "DetectIntentRequest"
	//   },
	//   "response": {
	//     "$ref": "DetectIntentResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.create":

type ProjectsAgentSessionsContextsCreateCall struct {
	s          *Service
	parent     string
	context    *Context
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a context.
func (r *ProjectsAgentSessionsContextsService) Create(parent string, context *Context) *ProjectsAgentSessionsContextsCreateCall {
	c := &ProjectsAgentSessionsContextsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.context = context
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsCreateCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.create" call.
// Exactly one of *Context or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Context.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentSessionsContextsCreateCall) Do(opts ...googleapi.CallOption) (*Context, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.sessions.contexts.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a context for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "request": {
	//     "$ref": "Context"
	//   },
	//   "response": {
	//     "$ref": "Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.delete":

type ProjectsAgentSessionsContextsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified context.
func (r *ProjectsAgentSessionsContextsService) Delete(name string) *ProjectsAgentSessionsContextsDeleteCall {
	c := &ProjectsAgentSessionsContextsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsDeleteCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentSessionsContextsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.sessions.contexts.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.get":

type ProjectsAgentSessionsContextsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified context.
func (r *ProjectsAgentSessionsContextsService) Get(name string) *ProjectsAgentSessionsContextsGetCall {
	c := &ProjectsAgentSessionsContextsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsGetCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsContextsGetCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsContextsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsGetCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.get" call.
// Exactly one of *Context or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Context.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentSessionsContextsGetCall) Do(opts ...googleapi.CallOption) (*Context, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.contexts.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.contexts.list":

type ProjectsAgentSessionsContextsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all contexts in the specified session.
func (r *ProjectsAgentSessionsContextsService) List(parent string) *ProjectsAgentSessionsContextsListCall {
	c := &ProjectsAgentSessionsContextsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentSessionsContextsListCall) PageSize(pageSize int64) *ProjectsAgentSessionsContextsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentSessionsContextsListCall) PageToken(pageToken string) *ProjectsAgentSessionsContextsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsListCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsContextsListCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsContextsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsListCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/contexts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.list" call.
// Exactly one of *ListContextsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListContextsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsContextsListCall) Do(opts ...googleapi.CallOption) (*ListContextsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListContextsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all contexts in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.contexts.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all contexts from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/contexts",
	//   "response": {
	//     "$ref": "ListContextsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentSessionsContextsListCall) Pages(ctx context.Context, f func(*ListContextsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dialogflow.projects.agent.sessions.contexts.patch":

type ProjectsAgentSessionsContextsPatchCall struct {
	s          *Service
	nameid     string
	context    *Context
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates the specified context.
func (r *ProjectsAgentSessionsContextsService) Patch(nameid string, context *Context) *ProjectsAgentSessionsContextsPatchCall {
	c := &ProjectsAgentSessionsContextsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.context = context
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentSessionsContextsPatchCall) UpdateMask(updateMask string) *ProjectsAgentSessionsContextsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsContextsPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsContextsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsContextsPatchCall) Context(ctx context.Context) *ProjectsAgentSessionsContextsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsContextsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsContextsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.context)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.contexts.patch" call.
// Exactly one of *Context or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Context.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentSessionsContextsPatchCall) Do(opts ...googleapi.CallOption) (*Context, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Context{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified context.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/contexts/{contextsId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.sessions.contexts.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of the context. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/contexts/\u003cContext ID\u003e`.\nNote: The Context ID is always converted to lowercase.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/contexts/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "Context"
	//   },
	//   "response": {
	//     "$ref": "Context"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.create":

type ProjectsAgentSessionsEntityTypesCreateCall struct {
	s                 *Service
	parent            string
	sessionentitytype *SessionEntityType
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Create: Creates a session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Create(parent string, sessionentitytype *SessionEntityType) *ProjectsAgentSessionsEntityTypesCreateCall {
	c := &ProjectsAgentSessionsEntityTypesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.sessionentitytype = sessionentitytype
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sessionentitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.create" call.
// Exactly one of *SessionEntityType or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SessionEntityType.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsEntityTypesCreateCall) Do(opts ...googleapi.CallOption) (*SessionEntityType, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes",
	//   "httpMethod": "POST",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The session to create a session entity type for.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "request": {
	//     "$ref": "SessionEntityType"
	//   },
	//   "response": {
	//     "$ref": "SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.delete":

type ProjectsAgentSessionsEntityTypesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes the specified session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Delete(name string) *ProjectsAgentSessionsEntityTypesDeleteCall {
	c := &ProjectsAgentSessionsEntityTypesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAgentSessionsEntityTypesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "DELETE",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the entity type to delete. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.get":

type ProjectsAgentSessionsEntityTypesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the specified session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Get(name string) *ProjectsAgentSessionsEntityTypesGetCall {
	c := &ProjectsAgentSessionsEntityTypesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsEntityTypesGetCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsEntityTypesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.get" call.
// Exactly one of *SessionEntityType or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SessionEntityType.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsEntityTypesGetCall) Do(opts ...googleapi.CallOption) (*SessionEntityType, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the session entity type. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.agent.sessions.entityTypes.list":

type ProjectsAgentSessionsEntityTypesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of all session entity types in the specified
// session.
func (r *ProjectsAgentSessionsEntityTypesService) List(parent string) *ProjectsAgentSessionsEntityTypesListCall {
	c := &ProjectsAgentSessionsEntityTypesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of items to return in a single page. By
// default 100 and at most 1000.
func (c *ProjectsAgentSessionsEntityTypesListCall) PageSize(pageSize int64) *ProjectsAgentSessionsEntityTypesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// next_page_token value returned from a previous list request.
func (c *ProjectsAgentSessionsEntityTypesListCall) PageToken(pageToken string) *ProjectsAgentSessionsEntityTypesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesListCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAgentSessionsEntityTypesListCall) IfNoneMatch(entityTag string) *ProjectsAgentSessionsEntityTypesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesListCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+parent}/entityTypes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.list" call.
// Exactly one of *ListSessionEntityTypesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListSessionEntityTypesResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsEntityTypesListCall) Do(opts ...googleapi.CallOption) (*ListSessionEntityTypesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListSessionEntityTypesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of all session entity types in the specified session.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of items to return in a single page. By\ndefault 100 and at most 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. The next_page_token value returned from a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The session to list all session entity types from.\nFormat: `projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+parent}/entityTypes",
	//   "response": {
	//     "$ref": "ListSessionEntityTypesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAgentSessionsEntityTypesListCall) Pages(ctx context.Context, f func(*ListSessionEntityTypesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "dialogflow.projects.agent.sessions.entityTypes.patch":

type ProjectsAgentSessionsEntityTypesPatchCall struct {
	s                 *Service
	nameid            string
	sessionentitytype *SessionEntityType
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Patch: Updates the specified session entity type.
func (r *ProjectsAgentSessionsEntityTypesService) Patch(nameid string, sessionentitytype *SessionEntityType) *ProjectsAgentSessionsEntityTypesPatchCall {
	c := &ProjectsAgentSessionsEntityTypesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.nameid = nameid
	c.sessionentitytype = sessionentitytype
	return c
}

// UpdateMask sets the optional parameter "updateMask": The mask to
// control which fields get updated.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) UpdateMask(updateMask string) *ProjectsAgentSessionsEntityTypesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Fields(s ...googleapi.Field) *ProjectsAgentSessionsEntityTypesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Context(ctx context.Context) *ProjectsAgentSessionsEntityTypesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAgentSessionsEntityTypesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sessionentitytype)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.nameid,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.agent.sessions.entityTypes.patch" call.
// Exactly one of *SessionEntityType or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SessionEntityType.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAgentSessionsEntityTypesPatchCall) Do(opts ...googleapi.CallOption) (*SessionEntityType, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SessionEntityType{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified session entity type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/agent/sessions/{sessionsId}/entityTypes/{entityTypesId}",
	//   "httpMethod": "PATCH",
	//   "id": "dialogflow.projects.agent.sessions.entityTypes.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The unique identifier of this session entity type. Format:\n`projects/\u003cProject ID\u003e/agent/sessions/\u003cSession ID\u003e/entityTypes/\u003cEntity Type\nDisplay Name\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/agent/sessions/[^/]+/entityTypes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. The mask to control which fields get updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "request": {
	//     "$ref": "SessionEntityType"
	//   },
	//   "response": {
	//     "$ref": "SessionEntityType"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dialogflow.projects.operations.get":

type ProjectsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *ProjectsOperationsService) Get(name string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsGetCall) Context(ctx context.Context) *ProjectsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "dialogflow.projects.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v2beta1/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "dialogflow.projects.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
