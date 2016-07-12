// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package querylog

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var _ = json.RawMessage{} // suppress unused package warning

func easyjson_320612c2_decode_github_com_abh_geodns_querylog_FileLogger(in *jlexer.Lexer, out *FileLogger) {
	if in.IsNull() {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_320612c2_encode_github_com_abh_geodns_querylog_FileLogger(out *jwriter.Writer, in FileLogger) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
func (v FileLogger) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_320612c2_encode_github_com_abh_geodns_querylog_FileLogger(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v FileLogger) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_320612c2_encode_github_com_abh_geodns_querylog_FileLogger(w, v)
}
func (v *FileLogger) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_320612c2_decode_github_com_abh_geodns_querylog_FileLogger(&r, v)
	return r.Error()
}
func (v *FileLogger) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_320612c2_decode_github_com_abh_geodns_querylog_FileLogger(l, v)
}
func easyjson_320612c2_decode_github_com_abh_geodns_querylog_Entry(in *jlexer.Lexer, out *Entry) {
	if in.IsNull() {
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
		case "Time":
			out.Time = int64(in.Int64())
		case "Origin":
			out.Origin = string(in.String())
		case "Name":
			out.Name = string(in.String())
		case "Qtype":
			out.Qtype = uint16(in.Uint16())
		case "Rcode":
			out.Rcode = int(in.Int())
		case "Answers":
			out.Answers = int(in.Int())
		case "Targets":
			in.Delim('[')
			if !in.IsDelim(']') {
				out.Targets = make([]string, 0, 4)
			} else {
				out.Targets = nil
			}
			for !in.IsDelim(']') {
				var v1 string
				v1 = string(in.String())
				out.Targets = append(out.Targets, v1)
				in.WantComma()
			}
			in.Delim(']')
		case "LabelName":
			out.LabelName = string(in.String())
		case "RemoteAddr":
			out.RemoteAddr = string(in.String())
		case "ClientAddr":
			out.ClientAddr = string(in.String())
		case "HasECS":
			out.HasECS = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
}
func easyjson_320612c2_encode_github_com_abh_geodns_querylog_Entry(out *jwriter.Writer, in Entry) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Time\":")
	out.Int64(int64(in.Time))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Origin\":")
	out.String(string(in.Origin))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Qtype\":")
	out.Uint16(uint16(in.Qtype))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Rcode\":")
	out.Int(int(in.Rcode))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Answers\":")
	out.Int(int(in.Answers))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Targets\":")
	out.RawByte('[')
	for v2, v3 := range in.Targets {
		if v2 > 0 {
			out.RawByte(',')
		}
		out.String(string(v3))
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"LabelName\":")
	out.String(string(in.LabelName))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"RemoteAddr\":")
	out.String(string(in.RemoteAddr))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ClientAddr\":")
	out.String(string(in.ClientAddr))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"HasECS\":")
	out.Bool(bool(in.HasECS))
	out.RawByte('}')
}
func (v Entry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson_320612c2_encode_github_com_abh_geodns_querylog_Entry(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}
func (v Entry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson_320612c2_encode_github_com_abh_geodns_querylog_Entry(w, v)
}
func (v *Entry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson_320612c2_decode_github_com_abh_geodns_querylog_Entry(&r, v)
	return r.Error()
}
func (v *Entry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson_320612c2_decode_github_com_abh_geodns_querylog_Entry(l, v)
}