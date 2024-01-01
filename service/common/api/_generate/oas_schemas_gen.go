// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

type ApiKeyAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *ApiKeyAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *ApiKeyAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// Ref: #/components/schemas/ChannelResponse
type ChannelResponse struct {
	ID         OptString                    `json:"id"`
	Snippet    OptChannelSnippetResponse    `json:"snippet"`
	Statistics OptChannelStatisticsResponse `json:"statistics"`
}

// GetID returns the value of ID.
func (s *ChannelResponse) GetID() OptString {
	return s.ID
}

// GetSnippet returns the value of Snippet.
func (s *ChannelResponse) GetSnippet() OptChannelSnippetResponse {
	return s.Snippet
}

// GetStatistics returns the value of Statistics.
func (s *ChannelResponse) GetStatistics() OptChannelStatisticsResponse {
	return s.Statistics
}

// SetID sets the value of ID.
func (s *ChannelResponse) SetID(val OptString) {
	s.ID = val
}

// SetSnippet sets the value of Snippet.
func (s *ChannelResponse) SetSnippet(val OptChannelSnippetResponse) {
	s.Snippet = val
}

// SetStatistics sets the value of Statistics.
func (s *ChannelResponse) SetStatistics(val OptChannelStatisticsResponse) {
	s.Statistics = val
}

// Ref: #/components/schemas/ChannelSnippetResponse
type ChannelSnippetResponse struct {
	CustomUrl   OptString             `json:"customUrl"`
	Description OptString             `json:"description"`
	PublishedAt OptString             `json:"publishedAt"`
	Thumbnails  OptThumbnailsResponse `json:"thumbnails"`
	Title       OptString             `json:"title"`
}

// GetCustomUrl returns the value of CustomUrl.
func (s *ChannelSnippetResponse) GetCustomUrl() OptString {
	return s.CustomUrl
}

// GetDescription returns the value of Description.
func (s *ChannelSnippetResponse) GetDescription() OptString {
	return s.Description
}

// GetPublishedAt returns the value of PublishedAt.
func (s *ChannelSnippetResponse) GetPublishedAt() OptString {
	return s.PublishedAt
}

// GetThumbnails returns the value of Thumbnails.
func (s *ChannelSnippetResponse) GetThumbnails() OptThumbnailsResponse {
	return s.Thumbnails
}

// GetTitle returns the value of Title.
func (s *ChannelSnippetResponse) GetTitle() OptString {
	return s.Title
}

// SetCustomUrl sets the value of CustomUrl.
func (s *ChannelSnippetResponse) SetCustomUrl(val OptString) {
	s.CustomUrl = val
}

// SetDescription sets the value of Description.
func (s *ChannelSnippetResponse) SetDescription(val OptString) {
	s.Description = val
}

// SetPublishedAt sets the value of PublishedAt.
func (s *ChannelSnippetResponse) SetPublishedAt(val OptString) {
	s.PublishedAt = val
}

// SetThumbnails sets the value of Thumbnails.
func (s *ChannelSnippetResponse) SetThumbnails(val OptThumbnailsResponse) {
	s.Thumbnails = val
}

// SetTitle sets the value of Title.
func (s *ChannelSnippetResponse) SetTitle(val OptString) {
	s.Title = val
}

// Ref: #/components/schemas/ChannelStatisticsResponse
type ChannelStatisticsResponse struct {
	HiddenSubscriberCount OptBool   `json:"hiddenSubscriberCount"`
	SubscriberCount       OptString `json:"subscriberCount"`
	VideoCount            OptString `json:"videoCount"`
	ViewCount             OptString `json:"viewCount"`
}

// GetHiddenSubscriberCount returns the value of HiddenSubscriberCount.
func (s *ChannelStatisticsResponse) GetHiddenSubscriberCount() OptBool {
	return s.HiddenSubscriberCount
}

// GetSubscriberCount returns the value of SubscriberCount.
func (s *ChannelStatisticsResponse) GetSubscriberCount() OptString {
	return s.SubscriberCount
}

// GetVideoCount returns the value of VideoCount.
func (s *ChannelStatisticsResponse) GetVideoCount() OptString {
	return s.VideoCount
}

// GetViewCount returns the value of ViewCount.
func (s *ChannelStatisticsResponse) GetViewCount() OptString {
	return s.ViewCount
}

// SetHiddenSubscriberCount sets the value of HiddenSubscriberCount.
func (s *ChannelStatisticsResponse) SetHiddenSubscriberCount(val OptBool) {
	s.HiddenSubscriberCount = val
}

// SetSubscriberCount sets the value of SubscriberCount.
func (s *ChannelStatisticsResponse) SetSubscriberCount(val OptString) {
	s.SubscriberCount = val
}

// SetVideoCount sets the value of VideoCount.
func (s *ChannelStatisticsResponse) SetVideoCount(val OptString) {
	s.VideoCount = val
}

// SetViewCount sets the value of ViewCount.
func (s *ChannelStatisticsResponse) SetViewCount(val OptString) {
	s.ViewCount = val
}

// ChannelsGetBadRequest is response for ChannelsGet operation.
type ChannelsGetBadRequest struct{}

func (*ChannelsGetBadRequest) channelsGetRes() {}

// ChannelsGetForbidden is response for ChannelsGet operation.
type ChannelsGetForbidden struct{}

func (*ChannelsGetForbidden) channelsGetRes() {}

// ChannelsGetInternalServerError is response for ChannelsGet operation.
type ChannelsGetInternalServerError struct{}

func (*ChannelsGetInternalServerError) channelsGetRes() {}

// ChannelsGetNotFound is response for ChannelsGet operation.
type ChannelsGetNotFound struct{}

func (*ChannelsGetNotFound) channelsGetRes() {}

// ChannelsGetUnauthorized is response for ChannelsGet operation.
type ChannelsGetUnauthorized struct{}

func (*ChannelsGetUnauthorized) channelsGetRes() {}

// ChannelsPostBadRequest is response for ChannelsPost operation.
type ChannelsPostBadRequest struct{}

func (*ChannelsPostBadRequest) channelsPostRes() {}

// ChannelsPostForbidden is response for ChannelsPost operation.
type ChannelsPostForbidden struct{}

func (*ChannelsPostForbidden) channelsPostRes() {}

// ChannelsPostInternalServerError is response for ChannelsPost operation.
type ChannelsPostInternalServerError struct{}

func (*ChannelsPostInternalServerError) channelsPostRes() {}

// ChannelsPostNotFound is response for ChannelsPost operation.
type ChannelsPostNotFound struct{}

func (*ChannelsPostNotFound) channelsPostRes() {}

type ChannelsPostOKApplicationJSON string

func (*ChannelsPostOKApplicationJSON) channelsPostRes() {}

type ChannelsPostReq struct {
	// Array of YouTube channel IDs.
	Ids []string `json:"ids"`
}

// GetIds returns the value of Ids.
func (s *ChannelsPostReq) GetIds() []string {
	return s.Ids
}

// SetIds sets the value of Ids.
func (s *ChannelsPostReq) SetIds(val []string) {
	s.Ids = val
}

// ChannelsPostUnauthorized is response for ChannelsPost operation.
type ChannelsPostUnauthorized struct{}

func (*ChannelsPostUnauthorized) channelsPostRes() {}

// ChannelsPutBadRequest is response for ChannelsPut operation.
type ChannelsPutBadRequest struct{}

func (*ChannelsPutBadRequest) channelsPutRes() {}

// ChannelsPutForbidden is response for ChannelsPut operation.
type ChannelsPutForbidden struct{}

func (*ChannelsPutForbidden) channelsPutRes() {}

// ChannelsPutInternalServerError is response for ChannelsPut operation.
type ChannelsPutInternalServerError struct{}

func (*ChannelsPutInternalServerError) channelsPutRes() {}

// ChannelsPutNotFound is response for ChannelsPut operation.
type ChannelsPutNotFound struct{}

func (*ChannelsPutNotFound) channelsPutRes() {}

type ChannelsPutOKApplicationJSON string

func (*ChannelsPutOKApplicationJSON) channelsPutRes() {}

type ChannelsPutReq struct {
	// Array of YouTube channel IDs.
	Ids []string `json:"ids"`
}

// GetIds returns the value of Ids.
func (s *ChannelsPutReq) GetIds() []string {
	return s.Ids
}

// SetIds sets the value of Ids.
func (s *ChannelsPutReq) SetIds(val []string) {
	s.Ids = val
}

// ChannelsPutUnauthorized is response for ChannelsPut operation.
type ChannelsPutUnauthorized struct{}

func (*ChannelsPutUnauthorized) channelsPutRes() {}

// Ref: #/components/schemas/ChannelsResponse
type ChannelsResponse struct {
	Channels []ChannelResponse `json:"channels"`
}

// GetChannels returns the value of Channels.
func (s *ChannelsResponse) GetChannels() []ChannelResponse {
	return s.Channels
}

// SetChannels sets the value of Channels.
func (s *ChannelsResponse) SetChannels(val []ChannelResponse) {
	s.Channels = val
}

func (*ChannelsResponse) channelsGetRes() {}

// ClipsGetBadRequest is response for ClipsGet operation.
type ClipsGetBadRequest struct{}

func (*ClipsGetBadRequest) clipsGetRes() {}

// ClipsGetForbidden is response for ClipsGet operation.
type ClipsGetForbidden struct{}

func (*ClipsGetForbidden) clipsGetRes() {}

// ClipsGetInternalServerError is response for ClipsGet operation.
type ClipsGetInternalServerError struct{}

func (*ClipsGetInternalServerError) clipsGetRes() {}

// ClipsGetNotFound is response for ClipsGet operation.
type ClipsGetNotFound struct{}

func (*ClipsGetNotFound) clipsGetRes() {}

// ClipsGetUnauthorized is response for ClipsGet operation.
type ClipsGetUnauthorized struct{}

func (*ClipsGetUnauthorized) clipsGetRes() {}

// ClipsPutBadRequest is response for ClipsPut operation.
type ClipsPutBadRequest struct{}

func (*ClipsPutBadRequest) clipsPutRes() {}

// ClipsPutForbidden is response for ClipsPut operation.
type ClipsPutForbidden struct{}

func (*ClipsPutForbidden) clipsPutRes() {}

// ClipsPutInternalServerError is response for ClipsPut operation.
type ClipsPutInternalServerError struct{}

func (*ClipsPutInternalServerError) clipsPutRes() {}

// ClipsPutNotFound is response for ClipsPut operation.
type ClipsPutNotFound struct{}

func (*ClipsPutNotFound) clipsPutRes() {}

type ClipsPutOKApplicationJSON string

func (*ClipsPutOKApplicationJSON) clipsPutRes() {}

type ClipsPutReq struct {
	CronType OptClipsPutReqCronType `json:"cronType"`
}

// GetCronType returns the value of CronType.
func (s *ClipsPutReq) GetCronType() OptClipsPutReqCronType {
	return s.CronType
}

// SetCronType sets the value of CronType.
func (s *ClipsPutReq) SetCronType(val OptClipsPutReqCronType) {
	s.CronType = val
}

type ClipsPutReqCronType string

const (
	ClipsPutReqCronTypeDaily   ClipsPutReqCronType = "daily"
	ClipsPutReqCronTypeWeekly  ClipsPutReqCronType = "weekly"
	ClipsPutReqCronTypeMonthly ClipsPutReqCronType = "monthly"
)

// AllValues returns all ClipsPutReqCronType values.
func (ClipsPutReqCronType) AllValues() []ClipsPutReqCronType {
	return []ClipsPutReqCronType{
		ClipsPutReqCronTypeDaily,
		ClipsPutReqCronTypeWeekly,
		ClipsPutReqCronTypeMonthly,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ClipsPutReqCronType) MarshalText() ([]byte, error) {
	switch s {
	case ClipsPutReqCronTypeDaily:
		return []byte(s), nil
	case ClipsPutReqCronTypeWeekly:
		return []byte(s), nil
	case ClipsPutReqCronTypeMonthly:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ClipsPutReqCronType) UnmarshalText(data []byte) error {
	switch ClipsPutReqCronType(data) {
	case ClipsPutReqCronTypeDaily:
		*s = ClipsPutReqCronTypeDaily
		return nil
	case ClipsPutReqCronTypeWeekly:
		*s = ClipsPutReqCronTypeWeekly
		return nil
	case ClipsPutReqCronTypeMonthly:
		*s = ClipsPutReqCronTypeMonthly
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// ClipsPutUnauthorized is response for ClipsPut operation.
type ClipsPutUnauthorized struct{}

func (*ClipsPutUnauthorized) clipsPutRes() {}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptChannelSnippetResponse returns new OptChannelSnippetResponse with value set to v.
func NewOptChannelSnippetResponse(v ChannelSnippetResponse) OptChannelSnippetResponse {
	return OptChannelSnippetResponse{
		Value: v,
		Set:   true,
	}
}

// OptChannelSnippetResponse is optional ChannelSnippetResponse.
type OptChannelSnippetResponse struct {
	Value ChannelSnippetResponse
	Set   bool
}

// IsSet returns true if OptChannelSnippetResponse was set.
func (o OptChannelSnippetResponse) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChannelSnippetResponse) Reset() {
	var v ChannelSnippetResponse
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChannelSnippetResponse) SetTo(v ChannelSnippetResponse) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChannelSnippetResponse) Get() (v ChannelSnippetResponse, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChannelSnippetResponse) Or(d ChannelSnippetResponse) ChannelSnippetResponse {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptChannelStatisticsResponse returns new OptChannelStatisticsResponse with value set to v.
func NewOptChannelStatisticsResponse(v ChannelStatisticsResponse) OptChannelStatisticsResponse {
	return OptChannelStatisticsResponse{
		Value: v,
		Set:   true,
	}
}

// OptChannelStatisticsResponse is optional ChannelStatisticsResponse.
type OptChannelStatisticsResponse struct {
	Value ChannelStatisticsResponse
	Set   bool
}

// IsSet returns true if OptChannelStatisticsResponse was set.
func (o OptChannelStatisticsResponse) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChannelStatisticsResponse) Reset() {
	var v ChannelStatisticsResponse
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChannelStatisticsResponse) SetTo(v ChannelStatisticsResponse) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChannelStatisticsResponse) Get() (v ChannelStatisticsResponse, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChannelStatisticsResponse) Or(d ChannelStatisticsResponse) ChannelStatisticsResponse {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptClipsPutReqCronType returns new OptClipsPutReqCronType with value set to v.
func NewOptClipsPutReqCronType(v ClipsPutReqCronType) OptClipsPutReqCronType {
	return OptClipsPutReqCronType{
		Value: v,
		Set:   true,
	}
}

// OptClipsPutReqCronType is optional ClipsPutReqCronType.
type OptClipsPutReqCronType struct {
	Value ClipsPutReqCronType
	Set   bool
}

// IsSet returns true if OptClipsPutReqCronType was set.
func (o OptClipsPutReqCronType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptClipsPutReqCronType) Reset() {
	var v ClipsPutReqCronType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptClipsPutReqCronType) SetTo(v ClipsPutReqCronType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptClipsPutReqCronType) Get() (v ClipsPutReqCronType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptClipsPutReqCronType) Or(d ClipsPutReqCronType) ClipsPutReqCronType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptSongsPutReqCronType returns new OptSongsPutReqCronType with value set to v.
func NewOptSongsPutReqCronType(v SongsPutReqCronType) OptSongsPutReqCronType {
	return OptSongsPutReqCronType{
		Value: v,
		Set:   true,
	}
}

// OptSongsPutReqCronType is optional SongsPutReqCronType.
type OptSongsPutReqCronType struct {
	Value SongsPutReqCronType
	Set   bool
}

// IsSet returns true if OptSongsPutReqCronType was set.
func (o OptSongsPutReqCronType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSongsPutReqCronType) Reset() {
	var v SongsPutReqCronType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSongsPutReqCronType) SetTo(v SongsPutReqCronType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSongsPutReqCronType) Get() (v SongsPutReqCronType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSongsPutReqCronType) Or(d SongsPutReqCronType) SongsPutReqCronType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptThumbnailResponse returns new OptThumbnailResponse with value set to v.
func NewOptThumbnailResponse(v ThumbnailResponse) OptThumbnailResponse {
	return OptThumbnailResponse{
		Value: v,
		Set:   true,
	}
}

// OptThumbnailResponse is optional ThumbnailResponse.
type OptThumbnailResponse struct {
	Value ThumbnailResponse
	Set   bool
}

// IsSet returns true if OptThumbnailResponse was set.
func (o OptThumbnailResponse) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptThumbnailResponse) Reset() {
	var v ThumbnailResponse
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptThumbnailResponse) SetTo(v ThumbnailResponse) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptThumbnailResponse) Get() (v ThumbnailResponse, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptThumbnailResponse) Or(d ThumbnailResponse) ThumbnailResponse {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptThumbnailsResponse returns new OptThumbnailsResponse with value set to v.
func NewOptThumbnailsResponse(v ThumbnailsResponse) OptThumbnailsResponse {
	return OptThumbnailsResponse{
		Value: v,
		Set:   true,
	}
}

// OptThumbnailsResponse is optional ThumbnailsResponse.
type OptThumbnailsResponse struct {
	Value ThumbnailsResponse
	Set   bool
}

// IsSet returns true if OptThumbnailsResponse was set.
func (o OptThumbnailsResponse) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptThumbnailsResponse) Reset() {
	var v ThumbnailsResponse
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptThumbnailsResponse) SetTo(v ThumbnailsResponse) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptThumbnailsResponse) Get() (v ThumbnailsResponse, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptThumbnailsResponse) Or(d ThumbnailsResponse) ThumbnailsResponse {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// SongsGetBadRequest is response for SongsGet operation.
type SongsGetBadRequest struct{}

func (*SongsGetBadRequest) songsGetRes() {}

// SongsGetForbidden is response for SongsGet operation.
type SongsGetForbidden struct{}

func (*SongsGetForbidden) songsGetRes() {}

// SongsGetInternalServerError is response for SongsGet operation.
type SongsGetInternalServerError struct{}

func (*SongsGetInternalServerError) songsGetRes() {}

// SongsGetNotFound is response for SongsGet operation.
type SongsGetNotFound struct{}

func (*SongsGetNotFound) songsGetRes() {}

// SongsGetUnauthorized is response for SongsGet operation.
type SongsGetUnauthorized struct{}

func (*SongsGetUnauthorized) songsGetRes() {}

// SongsPostBadRequest is response for SongsPost operation.
type SongsPostBadRequest struct{}

func (*SongsPostBadRequest) songsPostRes() {}

// SongsPostForbidden is response for SongsPost operation.
type SongsPostForbidden struct{}

func (*SongsPostForbidden) songsPostRes() {}

// SongsPostInternalServerError is response for SongsPost operation.
type SongsPostInternalServerError struct{}

func (*SongsPostInternalServerError) songsPostRes() {}

// SongsPostNotFound is response for SongsPost operation.
type SongsPostNotFound struct{}

func (*SongsPostNotFound) songsPostRes() {}

type SongsPostOKApplicationJSON string

func (*SongsPostOKApplicationJSON) songsPostRes() {}

type SongsPostReq struct {
	// Array of Song IDs.
	Ids []string `json:"ids"`
}

// GetIds returns the value of Ids.
func (s *SongsPostReq) GetIds() []string {
	return s.Ids
}

// SetIds sets the value of Ids.
func (s *SongsPostReq) SetIds(val []string) {
	s.Ids = val
}

// SongsPostUnauthorized is response for SongsPost operation.
type SongsPostUnauthorized struct{}

func (*SongsPostUnauthorized) songsPostRes() {}

// SongsPutBadRequest is response for SongsPut operation.
type SongsPutBadRequest struct{}

func (*SongsPutBadRequest) songsPutRes() {}

// SongsPutForbidden is response for SongsPut operation.
type SongsPutForbidden struct{}

func (*SongsPutForbidden) songsPutRes() {}

// SongsPutInternalServerError is response for SongsPut operation.
type SongsPutInternalServerError struct{}

func (*SongsPutInternalServerError) songsPutRes() {}

// SongsPutNotFound is response for SongsPut operation.
type SongsPutNotFound struct{}

func (*SongsPutNotFound) songsPutRes() {}

type SongsPutOKApplicationJSON string

func (*SongsPutOKApplicationJSON) songsPutRes() {}

type SongsPutReq struct {
	CronType OptSongsPutReqCronType `json:"cronType"`
}

// GetCronType returns the value of CronType.
func (s *SongsPutReq) GetCronType() OptSongsPutReqCronType {
	return s.CronType
}

// SetCronType sets the value of CronType.
func (s *SongsPutReq) SetCronType(val OptSongsPutReqCronType) {
	s.CronType = val
}

type SongsPutReqCronType string

const (
	SongsPutReqCronTypeDaily   SongsPutReqCronType = "daily"
	SongsPutReqCronTypeWeekly  SongsPutReqCronType = "weekly"
	SongsPutReqCronTypeMonthly SongsPutReqCronType = "monthly"
)

// AllValues returns all SongsPutReqCronType values.
func (SongsPutReqCronType) AllValues() []SongsPutReqCronType {
	return []SongsPutReqCronType{
		SongsPutReqCronTypeDaily,
		SongsPutReqCronTypeWeekly,
		SongsPutReqCronTypeMonthly,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s SongsPutReqCronType) MarshalText() ([]byte, error) {
	switch s {
	case SongsPutReqCronTypeDaily:
		return []byte(s), nil
	case SongsPutReqCronTypeWeekly:
		return []byte(s), nil
	case SongsPutReqCronTypeMonthly:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *SongsPutReqCronType) UnmarshalText(data []byte) error {
	switch SongsPutReqCronType(data) {
	case SongsPutReqCronTypeDaily:
		*s = SongsPutReqCronTypeDaily
		return nil
	case SongsPutReqCronTypeWeekly:
		*s = SongsPutReqCronTypeWeekly
		return nil
	case SongsPutReqCronTypeMonthly:
		*s = SongsPutReqCronTypeMonthly
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// SongsPutUnauthorized is response for SongsPut operation.
type SongsPutUnauthorized struct{}

func (*SongsPutUnauthorized) songsPutRes() {}

// Ref: #/components/schemas/ThumbnailResponse
type ThumbnailResponse struct {
	Height OptInt    `json:"height"`
	URL    OptString `json:"url"`
	Width  OptInt    `json:"width"`
}

// GetHeight returns the value of Height.
func (s *ThumbnailResponse) GetHeight() OptInt {
	return s.Height
}

// GetURL returns the value of URL.
func (s *ThumbnailResponse) GetURL() OptString {
	return s.URL
}

// GetWidth returns the value of Width.
func (s *ThumbnailResponse) GetWidth() OptInt {
	return s.Width
}

// SetHeight sets the value of Height.
func (s *ThumbnailResponse) SetHeight(val OptInt) {
	s.Height = val
}

// SetURL sets the value of URL.
func (s *ThumbnailResponse) SetURL(val OptString) {
	s.URL = val
}

// SetWidth sets the value of Width.
func (s *ThumbnailResponse) SetWidth(val OptInt) {
	s.Width = val
}

// Ref: #/components/schemas/ThumbnailsResponse
type ThumbnailsResponse struct {
	Default  OptThumbnailResponse `json:"default"`
	High     OptThumbnailResponse `json:"high"`
	Maxres   OptThumbnailResponse `json:"maxres"`
	Medium   OptThumbnailResponse `json:"medium"`
	Standard OptThumbnailResponse `json:"standard"`
}

// GetDefault returns the value of Default.
func (s *ThumbnailsResponse) GetDefault() OptThumbnailResponse {
	return s.Default
}

// GetHigh returns the value of High.
func (s *ThumbnailsResponse) GetHigh() OptThumbnailResponse {
	return s.High
}

// GetMaxres returns the value of Maxres.
func (s *ThumbnailsResponse) GetMaxres() OptThumbnailResponse {
	return s.Maxres
}

// GetMedium returns the value of Medium.
func (s *ThumbnailsResponse) GetMedium() OptThumbnailResponse {
	return s.Medium
}

// GetStandard returns the value of Standard.
func (s *ThumbnailsResponse) GetStandard() OptThumbnailResponse {
	return s.Standard
}

// SetDefault sets the value of Default.
func (s *ThumbnailsResponse) SetDefault(val OptThumbnailResponse) {
	s.Default = val
}

// SetHigh sets the value of High.
func (s *ThumbnailsResponse) SetHigh(val OptThumbnailResponse) {
	s.High = val
}

// SetMaxres sets the value of Maxres.
func (s *ThumbnailsResponse) SetMaxres(val OptThumbnailResponse) {
	s.Maxres = val
}

// SetMedium sets the value of Medium.
func (s *ThumbnailsResponse) SetMedium(val OptThumbnailResponse) {
	s.Medium = val
}

// SetStandard sets the value of Standard.
func (s *ThumbnailsResponse) SetStandard(val OptThumbnailResponse) {
	s.Standard = val
}

// Ref: #/components/schemas/VideoResponse
type VideoResponse struct {
	ChannelId    OptString             `json:"channelId"`
	ChannelTitle OptString             `json:"channelTitle"`
	Description  OptString             `json:"description"`
	ID           OptString             `json:"id"`
	PublishedAt  OptString             `json:"publishedAt"`
	Tags         []string              `json:"tags"`
	Thumbnails   OptThumbnailsResponse `json:"thumbnails"`
	Title        OptString             `json:"title"`
}

// GetChannelId returns the value of ChannelId.
func (s *VideoResponse) GetChannelId() OptString {
	return s.ChannelId
}

// GetChannelTitle returns the value of ChannelTitle.
func (s *VideoResponse) GetChannelTitle() OptString {
	return s.ChannelTitle
}

// GetDescription returns the value of Description.
func (s *VideoResponse) GetDescription() OptString {
	return s.Description
}

// GetID returns the value of ID.
func (s *VideoResponse) GetID() OptString {
	return s.ID
}

// GetPublishedAt returns the value of PublishedAt.
func (s *VideoResponse) GetPublishedAt() OptString {
	return s.PublishedAt
}

// GetTags returns the value of Tags.
func (s *VideoResponse) GetTags() []string {
	return s.Tags
}

// GetThumbnails returns the value of Thumbnails.
func (s *VideoResponse) GetThumbnails() OptThumbnailsResponse {
	return s.Thumbnails
}

// GetTitle returns the value of Title.
func (s *VideoResponse) GetTitle() OptString {
	return s.Title
}

// SetChannelId sets the value of ChannelId.
func (s *VideoResponse) SetChannelId(val OptString) {
	s.ChannelId = val
}

// SetChannelTitle sets the value of ChannelTitle.
func (s *VideoResponse) SetChannelTitle(val OptString) {
	s.ChannelTitle = val
}

// SetDescription sets the value of Description.
func (s *VideoResponse) SetDescription(val OptString) {
	s.Description = val
}

// SetID sets the value of ID.
func (s *VideoResponse) SetID(val OptString) {
	s.ID = val
}

// SetPublishedAt sets the value of PublishedAt.
func (s *VideoResponse) SetPublishedAt(val OptString) {
	s.PublishedAt = val
}

// SetTags sets the value of Tags.
func (s *VideoResponse) SetTags(val []string) {
	s.Tags = val
}

// SetThumbnails sets the value of Thumbnails.
func (s *VideoResponse) SetThumbnails(val OptThumbnailsResponse) {
	s.Thumbnails = val
}

// SetTitle sets the value of Title.
func (s *VideoResponse) SetTitle(val OptString) {
	s.Title = val
}

// Ref: #/components/schemas/VideosResponse
type VideosResponse struct {
	Videos []VideoResponse `json:"videos"`
}

// GetVideos returns the value of Videos.
func (s *VideosResponse) GetVideos() []VideoResponse {
	return s.Videos
}

// SetVideos sets the value of Videos.
func (s *VideosResponse) SetVideos(val []VideoResponse) {
	s.Videos = val
}

func (*VideosResponse) clipsGetRes() {}
func (*VideosResponse) songsGetRes() {}
