// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package bfapi

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonCaf77204DecodeGithubComTarbBfapi(in *jlexer.Lexer, out *VendorAccessTokenInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "access_token":
			out.AccessToken = string(in.String())
		case "expires_in":
			out.ExpiresIn = int64(in.Int64())
		case "refresh_token":
			out.RefreshToken = string(in.String())
		case "token_type":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.TokenType).UnmarshalJSON(data))
			}
		case "application_subscription":
			(out.ApplicationSubscription).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi(out *jwriter.Writer, in VendorAccessTokenInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"access_token\":"
		out.RawString(prefix[1:])
		out.String(string(in.AccessToken))
	}
	{
		const prefix string = ",\"expires_in\":"
		out.RawString(prefix)
		out.Int64(int64(in.ExpiresIn))
	}
	{
		const prefix string = ",\"refresh_token\":"
		out.RawString(prefix)
		out.String(string(in.RefreshToken))
	}
	{
		const prefix string = ",\"token_type\":"
		out.RawString(prefix)
		out.Raw((in.TokenType).MarshalJSON())
	}
	{
		const prefix string = ",\"application_subscription\":"
		out.RawString(prefix)
		(in.ApplicationSubscription).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v VendorAccessTokenInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v VendorAccessTokenInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *VendorAccessTokenInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *VendorAccessTokenInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi1(in *jlexer.Lexer, out *TokenArg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "client_id":
			out.ClientID = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "client_secret":
			out.ClientSecret = string(in.String())
		case "refresh_token":
			out.RefreshToken = string(in.String())
		case "grant_type":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.GrantType).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi1(out *jwriter.Writer, in TokenArg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"client_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ClientID))
	}
	if in.Code != "" {
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"client_secret\":"
		out.RawString(prefix)
		out.String(string(in.ClientSecret))
	}
	if in.RefreshToken != "" {
		const prefix string = ",\"refresh_token\":"
		out.RawString(prefix)
		out.String(string(in.RefreshToken))
	}
	{
		const prefix string = ",\"grant_type\":"
		out.RawString(prefix)
		out.Raw((in.GrantType).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TokenArg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TokenArg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TokenArg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TokenArg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi1(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi2(in *jlexer.Lexer, out *SubscriptionHistoryItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "subscriptionToken":
			out.SubscriptionToken = string(in.String())
		case "expiryDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiryDateTime).UnmarshalJSON(data))
			}
		case "expiredDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiredDateTime).UnmarshalJSON(data))
			}
		case "createdDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedDateTime).UnmarshalJSON(data))
			}
		case "activationDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ActivationDateTime).UnmarshalJSON(data))
			}
		case "cancellationDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CancellationDateTime).UnmarshalJSON(data))
			}
		case "subscriptionStatus":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.SubscriptionStatus).UnmarshalJSON(data))
			}
		case "clientReference":
			out.ClientReference = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi2(out *jwriter.Writer, in SubscriptionHistoryItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.SubscriptionToken != "" {
		const prefix string = ",\"subscriptionToken\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.SubscriptionToken))
	}
	if in.ExpiryDateTime != 0 {
		const prefix string = ",\"expiryDateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ExpiryDateTime))
	}
	if in.ExpiredDateTime != 0 {
		const prefix string = ",\"expiredDateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ExpiredDateTime))
	}
	if in.CreatedDateTime != 0 {
		const prefix string = ",\"createdDateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.CreatedDateTime))
	}
	if in.ActivationDateTime != 0 {
		const prefix string = ",\"activationDateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.ActivationDateTime))
	}
	if in.CancellationDateTime != 0 {
		const prefix string = ",\"cancellationDateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.CancellationDateTime))
	}
	if in.SubscriptionStatus != 0 {
		const prefix string = ",\"subscriptionStatus\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint8(uint8(in.SubscriptionStatus))
	}
	if in.ClientReference != "" {
		const prefix string = ",\"clientReference\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ClientReference))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubscriptionHistoryItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubscriptionHistoryItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubscriptionHistoryItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubscriptionHistoryItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi2(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi3(in *jlexer.Lexer, out *DeveloperAppVersion) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "owner":
			out.Owner = string(in.String())
		case "versionId":
			out.VersionID = int64(in.Int64())
		case "version":
			out.Version = string(in.String())
		case "applicationKey":
			out.ApplicationKey = string(in.String())
		case "delayData":
			out.DelayData = bool(in.Bool())
		case "subscriptionRequired":
			out.SubscriptionRequired = bool(in.Bool())
		case "ownerManaged":
			out.OwnerManaged = bool(in.Bool())
		case "active":
			out.Active = bool(in.Bool())
		case "vendorId":
			out.VendorID = string(in.String())
		case "vendorSecret":
			out.VendorSecret = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi3(out *jwriter.Writer, in DeveloperAppVersion) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"owner\":"
		out.RawString(prefix[1:])
		out.String(string(in.Owner))
	}
	{
		const prefix string = ",\"versionId\":"
		out.RawString(prefix)
		out.Int64(int64(in.VersionID))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"applicationKey\":"
		out.RawString(prefix)
		out.String(string(in.ApplicationKey))
	}
	{
		const prefix string = ",\"delayData\":"
		out.RawString(prefix)
		out.Bool(bool(in.DelayData))
	}
	{
		const prefix string = ",\"subscriptionRequired\":"
		out.RawString(prefix)
		out.Bool(bool(in.SubscriptionRequired))
	}
	{
		const prefix string = ",\"ownerManaged\":"
		out.RawString(prefix)
		out.Bool(bool(in.OwnerManaged))
	}
	{
		const prefix string = ",\"active\":"
		out.RawString(prefix)
		out.Bool(bool(in.Active))
	}
	{
		const prefix string = ",\"vendorId\":"
		out.RawString(prefix)
		out.String(string(in.VendorID))
	}
	{
		const prefix string = ",\"vendorSecret\":"
		out.RawString(prefix)
		out.String(string(in.VendorSecret))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeveloperAppVersion) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeveloperAppVersion) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeveloperAppVersion) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeveloperAppVersion) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi3(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi4(in *jlexer.Lexer, out *DeveloperApp) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "appName":
			out.AppName = string(in.String())
		case "appID":
			out.AppID = int64(in.Int64())
		case "appVersions":
			if in.IsNull() {
				in.Skip()
				out.AppVersions = nil
			} else {
				in.Delim('[')
				if out.AppVersions == nil {
					if !in.IsDelim(']') {
						out.AppVersions = make([]DeveloperAppVersion, 0, 1)
					} else {
						out.AppVersions = []DeveloperAppVersion{}
					}
				} else {
					out.AppVersions = (out.AppVersions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 DeveloperAppVersion
					(v1).UnmarshalEasyJSON(in)
					out.AppVersions = append(out.AppVersions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi4(out *jwriter.Writer, in DeveloperApp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"appName\":"
		out.RawString(prefix[1:])
		out.String(string(in.AppName))
	}
	{
		const prefix string = ",\"appID\":"
		out.RawString(prefix)
		out.Int64(int64(in.AppID))
	}
	{
		const prefix string = ",\"appVersions\":"
		out.RawString(prefix)
		if in.AppVersions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.AppVersions {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeveloperApp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeveloperApp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeveloperApp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeveloperApp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi4(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi5(in *jlexer.Lexer, out *CurrencyRate) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "currencyCode":
			out.CurrencyCode = string(in.String())
		case "rate":
			out.Rate = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi5(out *jwriter.Writer, in CurrencyRate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"currencyCode\":"
		out.RawString(prefix[1:])
		out.String(string(in.CurrencyCode))
	}
	{
		const prefix string = ",\"rate\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CurrencyRate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CurrencyRate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CurrencyRate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CurrencyRate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi5(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi6(in *jlexer.Lexer, out *ApplicationSubscription) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "subscriptionToken":
			out.SubscriptionToken = string(in.String())
		case "expiryDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiryDateTime).UnmarshalJSON(data))
			}
		case "expiredDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ExpiredDateTime).UnmarshalJSON(data))
			}
		case "createdDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedDateTime).UnmarshalJSON(data))
			}
		case "activationDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ActivationDateTime).UnmarshalJSON(data))
			}
		case "cancellationDateTime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CancellationDateTime).UnmarshalJSON(data))
			}
		case "subscriptionStatus":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.SubscriptionStatus).UnmarshalJSON(data))
			}
		case "clientReference":
			out.ClientReference = string(in.String())
		case "vendorClientId":
			out.VendorClientID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi6(out *jwriter.Writer, in ApplicationSubscription) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"subscriptionToken\":"
		out.RawString(prefix[1:])
		out.String(string(in.SubscriptionToken))
	}
	{
		const prefix string = ",\"expiryDateTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ExpiryDateTime))
	}
	{
		const prefix string = ",\"expiredDateTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ExpiredDateTime))
	}
	{
		const prefix string = ",\"createdDateTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CreatedDateTime))
	}
	{
		const prefix string = ",\"activationDateTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.ActivationDateTime))
	}
	{
		const prefix string = ",\"cancellationDateTime\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.CancellationDateTime))
	}
	{
		const prefix string = ",\"subscriptionStatus\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.SubscriptionStatus))
	}
	{
		const prefix string = ",\"clientReference\":"
		out.RawString(prefix)
		out.String(string(in.ClientReference))
	}
	{
		const prefix string = ",\"vendorClientId\":"
		out.RawString(prefix)
		out.String(string(in.VendorClientID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ApplicationSubscription) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ApplicationSubscription) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ApplicationSubscription) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ApplicationSubscription) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi6(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi7(in *jlexer.Lexer, out *AccountFundsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "availableToBetBalance":
			out.AvailableBalance = float64(in.Float64())
		case "exposure":
			out.Exposure = float64(in.Float64())
		case "retainedCommission":
			out.RetainedCommission = float64(in.Float64())
		case "exposureLimit":
			out.ExposureLimit = float64(in.Float64())
		case "discountRate":
			out.DiscountRate = float64(in.Float64())
		case "pointsBalance":
			out.PointsBalance = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi7(out *jwriter.Writer, in AccountFundsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"availableToBetBalance\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.AvailableBalance))
	}
	{
		const prefix string = ",\"exposure\":"
		out.RawString(prefix)
		out.Float64(float64(in.Exposure))
	}
	{
		const prefix string = ",\"retainedCommission\":"
		out.RawString(prefix)
		out.Float64(float64(in.RetainedCommission))
	}
	{
		const prefix string = ",\"exposureLimit\":"
		out.RawString(prefix)
		out.Float64(float64(in.ExposureLimit))
	}
	{
		const prefix string = ",\"discountRate\":"
		out.RawString(prefix)
		out.Float64(float64(in.DiscountRate))
	}
	{
		const prefix string = ",\"pointsBalance\":"
		out.RawString(prefix)
		out.Int(int(in.PointsBalance))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AccountFundsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccountFundsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AccountFundsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccountFundsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi7(l, v)
}
func easyjsonCaf77204DecodeGithubComTarbBfapi8(in *jlexer.Lexer, out *AccountDetailsResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "currencyCode":
			out.CurrencyCode = string(in.String())
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
		case "localeCode":
			out.LocaleCode = string(in.String())
		case "region":
			out.Region = string(in.String())
		case "timezone":
			out.Timezone = string(in.String())
		case "discountRate":
			out.DiscountRate = float64(in.Float64())
		case "pointsBalance":
			out.PointsBalance = int(in.Int())
		case "countryCode":
			out.CountryCode = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonCaf77204EncodeGithubComTarbBfapi8(out *jwriter.Writer, in AccountDetailsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"currencyCode\":"
		out.RawString(prefix[1:])
		out.String(string(in.CurrencyCode))
	}
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"localeCode\":"
		out.RawString(prefix)
		out.String(string(in.LocaleCode))
	}
	{
		const prefix string = ",\"region\":"
		out.RawString(prefix)
		out.String(string(in.Region))
	}
	{
		const prefix string = ",\"timezone\":"
		out.RawString(prefix)
		out.String(string(in.Timezone))
	}
	{
		const prefix string = ",\"discountRate\":"
		out.RawString(prefix)
		out.Float64(float64(in.DiscountRate))
	}
	{
		const prefix string = ",\"pointsBalance\":"
		out.RawString(prefix)
		out.Int(int(in.PointsBalance))
	}
	{
		const prefix string = ",\"countryCode\":"
		out.RawString(prefix)
		out.String(string(in.CountryCode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AccountDetailsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCaf77204EncodeGithubComTarbBfapi8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccountDetailsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCaf77204EncodeGithubComTarbBfapi8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AccountDetailsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCaf77204DecodeGithubComTarbBfapi8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccountDetailsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCaf77204DecodeGithubComTarbBfapi8(l, v)
}