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

func easyjsonAe34d1eDecodeGithubComTarbBfapi(in *jlexer.Lexer, out *LoginResult) {
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
		case "token":
			out.SessionToken = string(in.String())
		case "product":
			out.AppKey = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "error":
			out.Error = string(in.String())
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
func easyjsonAe34d1eEncodeGithubComTarbBfapi(out *jwriter.Writer, in LoginResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SessionToken))
	}
	{
		const prefix string = ",\"product\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AppKey))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoginResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe34d1eEncodeGithubComTarbBfapi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoginResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe34d1eEncodeGithubComTarbBfapi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoginResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe34d1eDecodeGithubComTarbBfapi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoginResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe34d1eDecodeGithubComTarbBfapi(l, v)
}
func easyjsonAe34d1eDecodeGithubComTarbBfapi1(in *jlexer.Lexer, out *CertLoginResult) {
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
		case "sessionToken":
			out.SessionToken = string(in.String())
		case "loginStatus":
			out.Status = string(in.String())
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
func easyjsonAe34d1eEncodeGithubComTarbBfapi1(out *jwriter.Writer, in CertLoginResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"sessionToken\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SessionToken))
	}
	{
		const prefix string = ",\"loginStatus\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CertLoginResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe34d1eEncodeGithubComTarbBfapi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CertLoginResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe34d1eEncodeGithubComTarbBfapi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CertLoginResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe34d1eDecodeGithubComTarbBfapi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CertLoginResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe34d1eDecodeGithubComTarbBfapi1(l, v)
}
