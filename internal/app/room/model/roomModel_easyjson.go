// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package roomModel

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

func easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel(in *jlexer.Lexer, out *RoomOrder) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "CostDesc":
			out.CostDesc = string(in.String())
		case "DataDesc":
			out.DataDesc = string(in.String())
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
func easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel(out *jwriter.Writer, in RoomOrder) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CostDesc\":"
		out.RawString(prefix[1:])
		out.String(string(in.CostDesc))
	}
	{
		const prefix string = ",\"DataDesc\":"
		out.RawString(prefix)
		out.String(string(in.DataDesc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RoomOrder) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoomOrder) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoomOrder) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoomOrder) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel(l, v)
}
func easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel1(in *jlexer.Lexer, out *RoomID) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "room_id":
			out.RoomID = int64(in.Int64())
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
func easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel1(out *jwriter.Writer, in RoomID) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"room_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.RoomID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RoomID) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoomID) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoomID) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoomID) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel1(l, v)
}
func easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel2(in *jlexer.Lexer, out *RoomAdd) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cost":
			out.Cost = int64(in.Int64())
		case "description":
			out.Description = string(in.String())
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
func easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel2(out *jwriter.Writer, in RoomAdd) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cost\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Cost))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RoomAdd) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RoomAdd) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RoomAdd) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RoomAdd) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel2(l, v)
}
func easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel3(in *jlexer.Lexer, out *Room) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "room_id":
			out.RoomID = int64(in.Int64())
		case "create":
			out.Created = string(in.String())
		case "cost":
			out.Cost = int64(in.Int64())
		case "description":
			out.Description = string(in.String())
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
func easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel3(out *jwriter.Writer, in Room) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"room_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.RoomID))
	}
	{
		const prefix string = ",\"create\":"
		out.RawString(prefix)
		out.String(string(in.Created))
	}
	{
		const prefix string = ",\"cost\":"
		out.RawString(prefix)
		out.Int64(int64(in.Cost))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Room) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Room) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9b1d4d46EncodeGithubComKostikansAvitoTestInternalAppRoomModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Room) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Room) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9b1d4d46DecodeGithubComKostikansAvitoTestInternalAppRoomModel3(l, v)
}