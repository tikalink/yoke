package content

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tikafog/of/buffers/content"
	"github.com/tikafog/of/version"
)

const CurrentDataVersion = 2

const (
	// ErrUnsupportedExtType ...
	//ErrUnsupportedExtType = errors.New("unsupported ext type")
	// ErrMustBePointer ...
	//ErrMustBePointer = errors.New("the interface parameter must be a pointer type")
	WrongVersionType = "wrong version type"
)

var (
	ErrWrongVersionType = Error(WrongVersionType)
)

// ExtType ...
type ExtType = content.ExtType

//Ext ...
//@Description:
type Ext struct {
	ExtType content.ExtType `json:"ext_type,omitempty"`
	Length  int             `json:"length,omitempty"`
	Data    []byte          `json:"data,omitempty"`
}

// Type ...
type Type = content.Type

// Content Content
type Content struct {
	Version string       `json:"version,omitempty"`
	From    string       `json:"from,omitempty"`
	Message *Message     `json:"message,omitempty"`
	Exts    []Ext        `json:"ext,omitempty"`
	Type    content.Type `json:"type,omitempty"`
}

func (c *Content) SetExts(exts ...Ext) *Content {
	c.Exts = exts
	return c
}

func (c *Content) AddExts(exts ...Ext) *Content {
	c.Exts = append(c.Exts, exts...)
	return c
}

func (c *Content) SetMessage(m *Message) *Content {
	c.Message = m
	return c
}

func (c *Content) NewMessage(data []byte) *Content {
	c.Message = NewContentMessage(data)
	return c
}

func (c *Content) NewMessageDetail(data []byte, index, last int64) *Content {
	c.Message = NewContentMessageWithDetail(data, index, last)
	return c
}

func (c *Content) NewMessageLast(last int64) *Content {
	c.Message = NewContentMessageLast(last)
	return c
}

func (c *Content) NewMessageAndLast(data []byte, last int64) *Content {
	c.Message = NewContentMessageAndLast(data, last)
	return c
}

func (c *Content) SetType(p content.Type) *Content {
	c.Type = p
	return c
}

func (c *Content) SetFrom(s string) *Content {
	c.From = s
	return c
}

// ParseJSONContent ...
// @Description: parse version 1 json data
// @param []byte
// @return *Content
// @return error
func ParseJSONContent(bytes []byte) (*Content, error) {
	var c Content
	err := json.Unmarshal(bytes, &c)
	if err != nil {
		return nil, err
	}
	if string(c.Version) != version.VersionOne {
		return nil, ErrWrongVersionType
	}
	return &c, err
}

func ParseJSONContentFromReader(reader io.Reader) (*Content, error) {
	var c Content
	err := json.NewDecoder(reader).Decode(&c)
	if err != nil {
		return nil, err
	}
	if string(c.Version) != version.VersionOne {
		return nil, ErrWrongVersionType
	}
	return &c, err
}

//ParseContent ...
//@Description: parse version 2 flatbuffer data
//@param []byte
//@return retC
//@return err
func ParseContent(bytes []byte) (retC *Content, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			if rerr == ErrWrongVersionType {
				err = ErrWrongVersionType
				return
			}
			err = fmt.Errorf("parse content error: %v", rerr)
		}
	}()
	c := content.GetRootAsContent(bytes, 0)
	if string(c.Version()) != version.VersionTwo {
		return nil, ErrWrongVersionType
	}
	return parseContent(c), err
}

func parseContent(cc *content.Content) *Content {
	//var node oc.Node
	message := cc.Message(nil)
	ext := make([]content.Ext, cc.ExtLength())
	var exts []Ext
	for i := 0; i < cc.ExtLength(); i++ {
		if cc.Ext(&ext[i], i) {
			exts = append(exts, Ext{
				ExtType: ext[i].Type(),
				Length:  len(ext[i].Data()),
				Data:    ext[i].Data(),
			})
		}
	}

	oc := NewContent(cc.Type())
	if message != nil {
		oc.Message = &Message{
			Last:    message.Last(),
			Index:   message.Index(),
			Version: int(message.Version()),
			Length:  len(message.Data()),
			Data:    message.Data(),
		}
	}
	oc.Exts = exts
	return oc
}

// JSON ...
// @Description:
// @receiver Content
// @return []byte
// @return error
func (c *Content) JSON() ([]byte, error) {
	c.Version = version.VersionOne
	return json.Marshal(c)
}

func (c *Content) Clear() {
	c.Message = EmptyMessage
	c.Exts = []Ext{}
}

// FinishBytes ...
// @Description:
// @receiver Content
// @return []byte
func (c *Content) FinishBytes() []byte {
	return contentToBytes(c, c.Message.IsEmpty())
}

func contentToBytes(c *Content, has bool) []byte {
	builder := flatbuffers.NewBuilder(0)

	var _message flatbuffers.UOffsetT
	if !has {
		_dataM := builder.CreateByteString(c.Message.Data)
		content.MessageStart(builder)
		content.MessageAddLast(builder, c.Message.Last)
		content.MessageAddIndex(builder, c.Message.Index)
		content.MessageAddVersion(builder, int32(c.Message.Version))
		content.MessageAddData(builder, _dataM)
		_message = content.MessageEnd(builder)
	}
	var _exts []flatbuffers.UOffsetT
	for i := range c.Exts {
		_dataE := builder.CreateByteString(c.Exts[i].Data)
		content.ExtStart(builder)
		content.ExtAddType(builder, c.Exts[i].ExtType)
		content.ExtAddData(builder, _dataE)
		_exts = append(_exts, content.ExtEnd(builder))
	}
	content.ContentStartExtVector(builder, len(_exts))
	for i := range _exts {
		builder.PrependUOffsetT(_exts[i])
	}
	_extsVec := builder.EndVector(len(_exts))

	_version := builder.CreateString(version.VersionTwo)
	_from := builder.CreateString(c.From)
	content.ContentStart(builder)
	content.ContentAddExt(builder, _extsVec)
	if !has {
		content.ContentAddMessage(builder, _message)
	}
	content.ContentAddFrom(builder, _from)
	content.ContentAddVersion(builder, _version)
	content.ContentAddType(builder, c.Type)
	builder.Finish(content.ContentEnd(builder))
	return builder.FinishedBytes()
}

// NewContentWithMessage ...
// @param content.Type
// @param *Message
// @return *Content
func NewContentWithMessage(p content.Type, message *Message) *Content {
	return &Content{
		Type:    p,
		Message: message,
	}
}

func NewContentWithExts(p content.Type, exts ...Ext) *Content {
	return &Content{
		Type: p,
		Exts: exts,
	}
}

// NewContentWithType ...
// @param content.Type
// @return *Content
func NewContentWithType(p content.Type) *Content {
	return &Content{
		Type: p,
	}
}

// NewContent ...
// @param content.Type
// @return *Content
func NewContent(p content.Type) *Content {
	return &Content{
		Type: p,
	}
}

// bytesToStringArray ...
// @Description: convert bytes to string array
// @param []byte
// @return []string
func bytesToStringArray(data []byte) []string {
	buf := bytes.NewBuffer(data)
	addr := struct {
		number int32
		lens   []int32
	}{}
	_ = binary.Read(buf, binary.BigEndian, &addr.number)
	addr.lens = make([]int32, addr.number)
	var ret []string
	for i := int32(0); i < addr.number; i++ {
		_ = binary.Read(buf, binary.BigEndian, &addr.lens[i])
		ret = append(ret, string(buf.Next(int(addr.lens[i]))))
	}
	return ret
}

// stringArrayToBytes ...
// @Description: convert string array to bytes
// @param ...string
// @return []byte
func stringArrayToBytes(strs ...string) []byte {
	buf := bytes.NewBuffer(nil)
	number := int32(len(strs))
	_ = binary.Write(buf, binary.BigEndian, number)
	var size int32
	for i := range strs {
		size = int32(len(strs[i]))
		_ = binary.Write(buf, binary.BigEndian, size)
		buf.WriteString(strs[i])
	}
	return buf.Bytes()
}
