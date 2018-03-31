package restxml_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/private/protocol/xml/xmlutil"
	"github.com/aws/aws-sdk-go-v2/private/util"
)

var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = ioutil.Discard
var _ = util.Trim("")
var _ = url.Values{}
var _ = io.EOF
var _ = aws.String
var _ = fmt.Println
var _ = reflect.Value{}

func init() {
	protocol.RandReader = &awstesting.ZeroReader{}
}

// OutputService1ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService1ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService1ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService1ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService1ProtocolTest client from just a config.
//     svc := outputservice1protocoltest.New(myConfig)
func NewOutputService1ProtocolTest(config aws.Config) *OutputService1ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService1ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice1protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService1ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService1ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService1TestCaseOperation1 = "OperationName"

// OutputService1TestCaseOperation1Request is a API request type for the OutputService1TestCaseOperation1 API operation.
type OutputService1TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService1TestShapeOutputService1TestCaseOperation1Input
	Copy  func(*OutputService1TestShapeOutputService1TestCaseOperation1Input) OutputService1TestCaseOperation1Request
}

// Send marshals and sends the OutputService1TestCaseOperation1 API request.
func (r OutputService1TestCaseOperation1Request) Send() (*OutputService1TestShapeOutputService1TestCaseOperation2Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService1TestShapeOutputService1TestCaseOperation2Output), nil
}

// OutputService1TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService1TestCaseOperation1Request method.
//    req := client.OutputService1TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) OutputService1TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService1TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService1TestShapeOutputService1TestCaseOperation1Input{}
	}

	output := &OutputService1TestShapeOutputService1TestCaseOperation2Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService1TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService1TestCaseOperation1Request}
}

const opOutputService1TestCaseOperation2 = "OperationName"

// OutputService1TestCaseOperation2Request is a API request type for the OutputService1TestCaseOperation2 API operation.
type OutputService1TestCaseOperation2Request struct {
	*aws.Request
	Input *OutputService1TestShapeOutputService1TestCaseOperation2Input
	Copy  func(*OutputService1TestShapeOutputService1TestCaseOperation2Input) OutputService1TestCaseOperation2Request
}

// Send marshals and sends the OutputService1TestCaseOperation2 API request.
func (r OutputService1TestCaseOperation2Request) Send() (*OutputService1TestShapeOutputService1TestCaseOperation2Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService1TestShapeOutputService1TestCaseOperation2Output), nil
}

// OutputService1TestCaseOperation2Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService1TestCaseOperation2Request method.
//    req := client.OutputService1TestCaseOperation2Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation2Request(input *OutputService1TestShapeOutputService1TestCaseOperation2Input) OutputService1TestCaseOperation2Request {
	op := &aws.Operation{
		Name: opOutputService1TestCaseOperation2,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService1TestShapeOutputService1TestCaseOperation2Input{}
	}

	output := &OutputService1TestShapeOutputService1TestCaseOperation2Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService1TestCaseOperation2Request{Request: req, Input: input, Copy: c.OutputService1TestCaseOperation2Request}
}

type OutputService1TestShapeOutputService1TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService1TestShapeOutputService1TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService1TestShapeOutputService1TestCaseOperation2Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService1TestShapeOutputService1TestCaseOperation2Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService1TestShapeOutputService1TestCaseOperation2Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	Char *string `type:"character"`

	Double *float64 `type:"double"`

	FalseBool *bool `type:"boolean"`

	Float *float64 `type:"float"`

	ImaHeader *string `location:"header" type:"string"`

	ImaHeaderLocation *string `location:"header" locationName:"X-Foo" type:"string"`

	Long *int64 `type:"long"`

	Num *int64 `locationName:"FooNum" type:"integer"`

	Str *string `type:"string"`

	Timestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	TrueBool *bool `type:"boolean"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService1TestShapeOutputService1TestCaseOperation2Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService1TestShapeOutputService1TestCaseOperation2Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.Char != nil {
		v := *s.Char

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Char", protocol.StringValue(v), metadata)
	}
	if s.Double != nil {
		v := *s.Double

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Double", protocol.Float64Value(v), metadata)
	}
	if s.FalseBool != nil {
		v := *s.FalseBool

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FalseBool", protocol.BoolValue(v), metadata)
	}
	if s.Float != nil {
		v := *s.Float

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Float", protocol.Float64Value(v), metadata)
	}
	if s.Long != nil {
		v := *s.Long

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Long", protocol.Int64Value(v), metadata)
	}
	if s.Num != nil {
		v := *s.Num

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FooNum", protocol.Int64Value(v), metadata)
	}
	if s.Str != nil {
		v := *s.Str

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Str", protocol.StringValue(v), metadata)
	}
	if s.Timestamp != nil {
		v := *s.Timestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Timestamp", protocol.TimeValue{V: v, Format: protocol.ISO8601TimeFormat}, metadata)
	}
	if s.TrueBool != nil {
		v := *s.TrueBool

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TrueBool", protocol.BoolValue(v), metadata)
	}
	if s.ImaHeader != nil {
		v := *s.ImaHeader

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ImaHeader", protocol.StringValue(v), metadata)
	}
	if s.ImaHeaderLocation != nil {
		v := *s.ImaHeaderLocation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Foo", protocol.StringValue(v), metadata)
	}
	return nil
}

// OutputService2ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService2ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService2ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService2ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService2ProtocolTest client from just a config.
//     svc := outputservice2protocoltest.New(myConfig)
func NewOutputService2ProtocolTest(config aws.Config) *OutputService2ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService2ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice2protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService2ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService2ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService2TestCaseOperation1 = "OperationName"

// OutputService2TestCaseOperation1Request is a API request type for the OutputService2TestCaseOperation1 API operation.
type OutputService2TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService2TestShapeOutputService2TestCaseOperation1Input
	Copy  func(*OutputService2TestShapeOutputService2TestCaseOperation1Input) OutputService2TestCaseOperation1Request
}

// Send marshals and sends the OutputService2TestCaseOperation1 API request.
func (r OutputService2TestCaseOperation1Request) Send() (*OutputService2TestShapeOutputService2TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService2TestShapeOutputService2TestCaseOperation1Output), nil
}

// OutputService2TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService2TestCaseOperation1Request method.
//    req := client.OutputService2TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1Request(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) OutputService2TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService2TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService2TestShapeOutputService2TestCaseOperation1Input{}
	}

	output := &OutputService2TestShapeOutputService2TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService2TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService2TestCaseOperation1Request}
}

type OutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService2TestShapeOutputService2TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService2TestShapeOutputService2TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Blob is automatically base64 encoded/decoded by the SDK.
	Blob []byte `type:"blob"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService2TestShapeOutputService2TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService2TestShapeOutputService2TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.Blob != nil {
		v := s.Blob

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Blob", protocol.BytesValue(v), metadata)
	}
	return nil
}

// OutputService3ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService3ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService3ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService3ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService3ProtocolTest client from just a config.
//     svc := outputservice3protocoltest.New(myConfig)
func NewOutputService3ProtocolTest(config aws.Config) *OutputService3ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService3ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice3protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService3ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService3ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService3TestCaseOperation1 = "OperationName"

// OutputService3TestCaseOperation1Request is a API request type for the OutputService3TestCaseOperation1 API operation.
type OutputService3TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService3TestShapeOutputService3TestCaseOperation1Input
	Copy  func(*OutputService3TestShapeOutputService3TestCaseOperation1Input) OutputService3TestCaseOperation1Request
}

// Send marshals and sends the OutputService3TestCaseOperation1 API request.
func (r OutputService3TestCaseOperation1Request) Send() (*OutputService3TestShapeOutputService3TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService3TestShapeOutputService3TestCaseOperation1Output), nil
}

// OutputService3TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService3TestCaseOperation1Request method.
//    req := client.OutputService3TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1Request(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) OutputService3TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService3TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService3TestShapeOutputService3TestCaseOperation1Input{}
	}

	output := &OutputService3TestShapeOutputService3TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService3TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService3TestCaseOperation1Request}
}

type OutputService3TestShapeOutputService3TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService3TestShapeOutputService3TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService3TestShapeOutputService3TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	ListMember []string `type:"list"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService3TestShapeOutputService3TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService3TestShapeOutputService3TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ListMember) > 0 {
		v := s.ListMember

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ListMember", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	return nil
}

// OutputService4ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService4ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService4ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService4ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService4ProtocolTest client from just a config.
//     svc := outputservice4protocoltest.New(myConfig)
func NewOutputService4ProtocolTest(config aws.Config) *OutputService4ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService4ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice4protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService4ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService4ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService4TestCaseOperation1 = "OperationName"

// OutputService4TestCaseOperation1Request is a API request type for the OutputService4TestCaseOperation1 API operation.
type OutputService4TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService4TestShapeOutputService4TestCaseOperation1Input
	Copy  func(*OutputService4TestShapeOutputService4TestCaseOperation1Input) OutputService4TestCaseOperation1Request
}

// Send marshals and sends the OutputService4TestCaseOperation1 API request.
func (r OutputService4TestCaseOperation1Request) Send() (*OutputService4TestShapeOutputService4TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService4TestShapeOutputService4TestCaseOperation1Output), nil
}

// OutputService4TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService4TestCaseOperation1Request method.
//    req := client.OutputService4TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1Request(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) OutputService4TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService4TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService4TestShapeOutputService4TestCaseOperation1Input{}
	}

	output := &OutputService4TestShapeOutputService4TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService4TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService4TestCaseOperation1Request}
}

type OutputService4TestShapeOutputService4TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService4TestShapeOutputService4TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService4TestShapeOutputService4TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	ListMember []string `locationNameList:"item" type:"list"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService4TestShapeOutputService4TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService4TestShapeOutputService4TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ListMember) > 0 {
		v := s.ListMember

		metadata := protocol.Metadata{ListLocationName: "item"}
		ls0 := e.List(protocol.BodyTarget, "ListMember", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	return nil
}

// OutputService5ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService5ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService5ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService5ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService5ProtocolTest client from just a config.
//     svc := outputservice5protocoltest.New(myConfig)
func NewOutputService5ProtocolTest(config aws.Config) *OutputService5ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService5ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice5protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService5ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService5ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService5TestCaseOperation1 = "OperationName"

// OutputService5TestCaseOperation1Request is a API request type for the OutputService5TestCaseOperation1 API operation.
type OutputService5TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService5TestShapeOutputService5TestCaseOperation1Input
	Copy  func(*OutputService5TestShapeOutputService5TestCaseOperation1Input) OutputService5TestCaseOperation1Request
}

// Send marshals and sends the OutputService5TestCaseOperation1 API request.
func (r OutputService5TestCaseOperation1Request) Send() (*OutputService5TestShapeOutputService5TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService5TestShapeOutputService5TestCaseOperation1Output), nil
}

// OutputService5TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService5TestCaseOperation1Request method.
//    req := client.OutputService5TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1Request(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) OutputService5TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService5TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService5TestShapeOutputService5TestCaseOperation1Input{}
	}

	output := &OutputService5TestShapeOutputService5TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService5TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService5TestCaseOperation1Request}
}

type OutputService5TestShapeOutputService5TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService5TestShapeOutputService5TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService5TestShapeOutputService5TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	ListMember []string `type:"list" flattened:"true"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService5TestShapeOutputService5TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService5TestShapeOutputService5TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ListMember) > 0 {
		v := s.ListMember

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "ListMember", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	return nil
}

// OutputService6ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService6ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService6ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService6ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService6ProtocolTest client from just a config.
//     svc := outputservice6protocoltest.New(myConfig)
func NewOutputService6ProtocolTest(config aws.Config) *OutputService6ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService6ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice6protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService6ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService6ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService6TestCaseOperation1 = "OperationName"

// OutputService6TestCaseOperation1Request is a API request type for the OutputService6TestCaseOperation1 API operation.
type OutputService6TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService6TestShapeOutputService6TestCaseOperation1Input
	Copy  func(*OutputService6TestShapeOutputService6TestCaseOperation1Input) OutputService6TestCaseOperation1Request
}

// Send marshals and sends the OutputService6TestCaseOperation1 API request.
func (r OutputService6TestCaseOperation1Request) Send() (*OutputService6TestShapeOutputService6TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService6TestShapeOutputService6TestCaseOperation1Output), nil
}

// OutputService6TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService6TestCaseOperation1Request method.
//    req := client.OutputService6TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1Request(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) OutputService6TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService6TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService6TestShapeOutputService6TestCaseOperation1Input{}
	}

	output := &OutputService6TestShapeOutputService6TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService6TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService6TestCaseOperation1Request}
}

type OutputService6TestShapeOutputService6TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService6TestShapeOutputService6TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService6TestShapeOutputService6TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	Map map[string]OutputService6TestShapeSingleStructure `type:"map"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService6TestShapeOutputService6TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService6TestShapeOutputService6TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Map) > 0 {
		v := s.Map

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "Map", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	return nil
}

type OutputService6TestShapeSingleStructure struct {
	_ struct{} `type:"structure"`

	Foo *string `locationName:"foo" type:"string"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService6TestShapeSingleStructure) MarshalFields(e protocol.FieldEncoder) error {
	if s.Foo != nil {
		v := *s.Foo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "foo", protocol.StringValue(v), metadata)
	}
	return nil
}

// OutputService7ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService7ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService7ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService7ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService7ProtocolTest client from just a config.
//     svc := outputservice7protocoltest.New(myConfig)
func NewOutputService7ProtocolTest(config aws.Config) *OutputService7ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService7ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice7protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService7ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService7ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService7TestCaseOperation1 = "OperationName"

// OutputService7TestCaseOperation1Request is a API request type for the OutputService7TestCaseOperation1 API operation.
type OutputService7TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService7TestShapeOutputService7TestCaseOperation1Input
	Copy  func(*OutputService7TestShapeOutputService7TestCaseOperation1Input) OutputService7TestCaseOperation1Request
}

// Send marshals and sends the OutputService7TestCaseOperation1 API request.
func (r OutputService7TestCaseOperation1Request) Send() (*OutputService7TestShapeOutputService7TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService7TestShapeOutputService7TestCaseOperation1Output), nil
}

// OutputService7TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService7TestCaseOperation1Request method.
//    req := client.OutputService7TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1Request(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) OutputService7TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService7TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService7TestShapeOutputService7TestCaseOperation1Input{}
	}

	output := &OutputService7TestShapeOutputService7TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService7TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService7TestCaseOperation1Request}
}

type OutputService7TestShapeOutputService7TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService7TestShapeOutputService7TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService7TestShapeOutputService7TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	Map map[string]string `type:"map" flattened:"true"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService7TestShapeOutputService7TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService7TestShapeOutputService7TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Map) > 0 {
		v := s.Map

		metadata := protocol.Metadata{Flatten: true}
		ms0 := e.Map(protocol.BodyTarget, "Map", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	return nil
}

// OutputService8ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService8ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService8ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService8ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService8ProtocolTest client from just a config.
//     svc := outputservice8protocoltest.New(myConfig)
func NewOutputService8ProtocolTest(config aws.Config) *OutputService8ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService8ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice8protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService8ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService8ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService8TestCaseOperation1 = "OperationName"

// OutputService8TestCaseOperation1Request is a API request type for the OutputService8TestCaseOperation1 API operation.
type OutputService8TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService8TestShapeOutputService8TestCaseOperation1Input
	Copy  func(*OutputService8TestShapeOutputService8TestCaseOperation1Input) OutputService8TestCaseOperation1Request
}

// Send marshals and sends the OutputService8TestCaseOperation1 API request.
func (r OutputService8TestCaseOperation1Request) Send() (*OutputService8TestShapeOutputService8TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService8TestShapeOutputService8TestCaseOperation1Output), nil
}

// OutputService8TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService8TestCaseOperation1Request method.
//    req := client.OutputService8TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1Request(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) OutputService8TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService8TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService8TestShapeOutputService8TestCaseOperation1Input{}
	}

	output := &OutputService8TestShapeOutputService8TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService8TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService8TestCaseOperation1Request}
}

type OutputService8TestShapeOutputService8TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService8TestShapeOutputService8TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService8TestShapeOutputService8TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	Map map[string]string `locationNameKey:"foo" locationNameValue:"bar" type:"map"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService8TestShapeOutputService8TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService8TestShapeOutputService8TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Map) > 0 {
		v := s.Map

		metadata := protocol.Metadata{MapLocationNameKey: "foo", MapLocationNameValue: "bar"}
		ms0 := e.Map(protocol.BodyTarget, "Map", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.StringValue(v1))
		}
		ms0.End()

	}
	return nil
}

// OutputService9ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService9ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService9ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService9ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService9ProtocolTest client from just a config.
//     svc := outputservice9protocoltest.New(myConfig)
func NewOutputService9ProtocolTest(config aws.Config) *OutputService9ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService9ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice9protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService9ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService9ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService9TestCaseOperation1 = "OperationName"

// OutputService9TestCaseOperation1Request is a API request type for the OutputService9TestCaseOperation1 API operation.
type OutputService9TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService9TestShapeOutputService9TestCaseOperation1Input
	Copy  func(*OutputService9TestShapeOutputService9TestCaseOperation1Input) OutputService9TestCaseOperation1Request
}

// Send marshals and sends the OutputService9TestCaseOperation1 API request.
func (r OutputService9TestCaseOperation1Request) Send() (*OutputService9TestShapeOutputService9TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService9TestShapeOutputService9TestCaseOperation1Output), nil
}

// OutputService9TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService9TestCaseOperation1Request method.
//    req := client.OutputService9TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1Request(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) OutputService9TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService9TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService9TestShapeOutputService9TestCaseOperation1Input{}
	}

	output := &OutputService9TestShapeOutputService9TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService9TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService9TestCaseOperation1Request}
}

type OutputService9TestShapeOutputService9TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService9TestShapeOutputService9TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService9TestShapeOutputService9TestCaseOperation1Output struct {
	_ struct{} `type:"structure" payload:"Data"`

	responseMetadata aws.Response

	Data *OutputService9TestShapeSingleStructure `type:"structure"`

	Header *string `location:"header" locationName:"X-Foo" type:"string"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService9TestShapeOutputService9TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService9TestShapeOutputService9TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.Header != nil {
		v := *s.Header

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Foo", protocol.StringValue(v), metadata)
	}
	if s.Data != nil {
		v := s.Data

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Data", v, metadata)
	}
	return nil
}

type OutputService9TestShapeSingleStructure struct {
	_ struct{} `type:"structure"`

	Foo *string `type:"string"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService9TestShapeSingleStructure) MarshalFields(e protocol.FieldEncoder) error {
	if s.Foo != nil {
		v := *s.Foo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Foo", protocol.StringValue(v), metadata)
	}
	return nil
}

// OutputService10ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService10ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService10ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService10ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService10ProtocolTest client from just a config.
//     svc := outputservice10protocoltest.New(myConfig)
func NewOutputService10ProtocolTest(config aws.Config) *OutputService10ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService10ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice10protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService10ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService10ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService10TestCaseOperation1 = "OperationName"

// OutputService10TestCaseOperation1Request is a API request type for the OutputService10TestCaseOperation1 API operation.
type OutputService10TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService10TestShapeOutputService10TestCaseOperation1Input
	Copy  func(*OutputService10TestShapeOutputService10TestCaseOperation1Input) OutputService10TestCaseOperation1Request
}

// Send marshals and sends the OutputService10TestCaseOperation1 API request.
func (r OutputService10TestCaseOperation1Request) Send() (*OutputService10TestShapeOutputService10TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService10TestShapeOutputService10TestCaseOperation1Output), nil
}

// OutputService10TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService10TestCaseOperation1Request method.
//    req := client.OutputService10TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1Request(input *OutputService10TestShapeOutputService10TestCaseOperation1Input) OutputService10TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService10TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService10TestShapeOutputService10TestCaseOperation1Input{}
	}

	output := &OutputService10TestShapeOutputService10TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService10TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService10TestCaseOperation1Request}
}

type OutputService10TestShapeOutputService10TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService10TestShapeOutputService10TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService10TestShapeOutputService10TestCaseOperation1Output struct {
	_ struct{} `type:"structure" payload:"Stream"`

	responseMetadata aws.Response

	Stream []byte `type:"blob"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService10TestShapeOutputService10TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService10TestShapeOutputService10TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.Stream != nil {
		v := s.Stream

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Stream", protocol.BytesStream(v), metadata)
	}
	return nil
}

// OutputService11ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService11ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService11ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService11ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService11ProtocolTest client from just a config.
//     svc := outputservice11protocoltest.New(myConfig)
func NewOutputService11ProtocolTest(config aws.Config) *OutputService11ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService11ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice11protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService11ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService11ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService11TestCaseOperation1 = "OperationName"

// OutputService11TestCaseOperation1Request is a API request type for the OutputService11TestCaseOperation1 API operation.
type OutputService11TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService11TestShapeOutputService11TestCaseOperation1Input
	Copy  func(*OutputService11TestShapeOutputService11TestCaseOperation1Input) OutputService11TestCaseOperation1Request
}

// Send marshals and sends the OutputService11TestCaseOperation1 API request.
func (r OutputService11TestCaseOperation1Request) Send() (*OutputService11TestShapeOutputService11TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService11TestShapeOutputService11TestCaseOperation1Output), nil
}

// OutputService11TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService11TestCaseOperation1Request method.
//    req := client.OutputService11TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService11ProtocolTest) OutputService11TestCaseOperation1Request(input *OutputService11TestShapeOutputService11TestCaseOperation1Input) OutputService11TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService11TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService11TestShapeOutputService11TestCaseOperation1Input{}
	}

	output := &OutputService11TestShapeOutputService11TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService11TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService11TestCaseOperation1Request}
}

type OutputService11TestShapeOutputService11TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService11TestShapeOutputService11TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService11TestShapeOutputService11TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	Char *string `location:"header" locationName:"x-char" type:"character"`

	Double *float64 `location:"header" locationName:"x-double" type:"double"`

	FalseBool *bool `location:"header" locationName:"x-false-bool" type:"boolean"`

	Float *float64 `location:"header" locationName:"x-float" type:"float"`

	Integer *int64 `location:"header" locationName:"x-int" type:"integer"`

	Long *int64 `location:"header" locationName:"x-long" type:"long"`

	Str *string `location:"header" locationName:"x-str" type:"string"`

	Timestamp *time.Time `location:"header" locationName:"x-timestamp" type:"timestamp" timestampFormat:"iso8601"`

	TrueBool *bool `location:"header" locationName:"x-true-bool" type:"boolean"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService11TestShapeOutputService11TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService11TestShapeOutputService11TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.Char != nil {
		v := *s.Char

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-char", protocol.StringValue(v), metadata)
	}
	if s.Double != nil {
		v := *s.Double

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-double", protocol.Float64Value(v), metadata)
	}
	if s.FalseBool != nil {
		v := *s.FalseBool

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-false-bool", protocol.BoolValue(v), metadata)
	}
	if s.Float != nil {
		v := *s.Float

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-float", protocol.Float64Value(v), metadata)
	}
	if s.Integer != nil {
		v := *s.Integer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-int", protocol.Int64Value(v), metadata)
	}
	if s.Long != nil {
		v := *s.Long

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-long", protocol.Int64Value(v), metadata)
	}
	if s.Str != nil {
		v := *s.Str

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-str", protocol.StringValue(v), metadata)
	}
	if s.Timestamp != nil {
		v := *s.Timestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-timestamp", protocol.TimeValue{V: v, Format: protocol.RFC822TimeFromat}, metadata)
	}
	if s.TrueBool != nil {
		v := *s.TrueBool

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-true-bool", protocol.BoolValue(v), metadata)
	}
	return nil
}

// OutputService12ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService12ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService12ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService12ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService12ProtocolTest client from just a config.
//     svc := outputservice12protocoltest.New(myConfig)
func NewOutputService12ProtocolTest(config aws.Config) *OutputService12ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService12ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice12protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService12ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService12ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService12TestCaseOperation1 = "OperationName"

// OutputService12TestCaseOperation1Request is a API request type for the OutputService12TestCaseOperation1 API operation.
type OutputService12TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService12TestShapeOutputService12TestCaseOperation1Input
	Copy  func(*OutputService12TestShapeOutputService12TestCaseOperation1Input) OutputService12TestCaseOperation1Request
}

// Send marshals and sends the OutputService12TestCaseOperation1 API request.
func (r OutputService12TestCaseOperation1Request) Send() (*OutputService12TestShapeOutputService12TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService12TestShapeOutputService12TestCaseOperation1Output), nil
}

// OutputService12TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService12TestCaseOperation1Request method.
//    req := client.OutputService12TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService12ProtocolTest) OutputService12TestCaseOperation1Request(input *OutputService12TestShapeOutputService12TestCaseOperation1Input) OutputService12TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService12TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService12TestShapeOutputService12TestCaseOperation1Input{}
	}

	output := &OutputService12TestShapeOutputService12TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService12TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService12TestCaseOperation1Request}
}

type OutputService12TestShapeOutputService12TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService12TestShapeOutputService12TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService12TestShapeOutputService12TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	Foo *string `type:"string"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService12TestShapeOutputService12TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService12TestShapeOutputService12TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if s.Foo != nil {
		v := *s.Foo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Foo", protocol.StringValue(v), metadata)
	}
	return nil
}

// OutputService13ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService13ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService13ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService13ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService13ProtocolTest client from just a config.
//     svc := outputservice13protocoltest.New(myConfig)
func NewOutputService13ProtocolTest(config aws.Config) *OutputService13ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService13ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice13protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService13ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService13ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService13TestCaseOperation1 = "OperationName"

// OutputService13TestCaseOperation1Request is a API request type for the OutputService13TestCaseOperation1 API operation.
type OutputService13TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService13TestShapeOutputService13TestCaseOperation1Input
	Copy  func(*OutputService13TestShapeOutputService13TestCaseOperation1Input) OutputService13TestCaseOperation1Request
}

// Send marshals and sends the OutputService13TestCaseOperation1 API request.
func (r OutputService13TestCaseOperation1Request) Send() (*OutputService13TestShapeOutputService13TestCaseOperation2Input, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService13TestShapeOutputService13TestCaseOperation2Input), nil
}

// OutputService13TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService13TestCaseOperation1Request method.
//    req := client.OutputService13TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService13ProtocolTest) OutputService13TestCaseOperation1Request(input *OutputService13TestShapeOutputService13TestCaseOperation1Input) OutputService13TestCaseOperation1Request {
	op := &aws.Operation{
		Name:       opOutputService13TestCaseOperation1,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &OutputService13TestShapeOutputService13TestCaseOperation1Input{}
	}

	output := &OutputService13TestShapeOutputService13TestCaseOperation2Input{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService13TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService13TestCaseOperation1Request}
}

const opOutputService13TestCaseOperation2 = "OperationName"

// OutputService13TestCaseOperation2Request is a API request type for the OutputService13TestCaseOperation2 API operation.
type OutputService13TestCaseOperation2Request struct {
	*aws.Request
	Input *OutputService13TestShapeOutputService13TestCaseOperation2Input
	Copy  func(*OutputService13TestShapeOutputService13TestCaseOperation2Input) OutputService13TestCaseOperation2Request
}

// Send marshals and sends the OutputService13TestCaseOperation2 API request.
func (r OutputService13TestCaseOperation2Request) Send() (*OutputService13TestShapeOutputService13TestCaseOperation2Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService13TestShapeOutputService13TestCaseOperation2Output), nil
}

// OutputService13TestCaseOperation2Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService13TestCaseOperation2Request method.
//    req := client.OutputService13TestCaseOperation2Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService13ProtocolTest) OutputService13TestCaseOperation2Request(input *OutputService13TestShapeOutputService13TestCaseOperation2Input) OutputService13TestCaseOperation2Request {
	op := &aws.Operation{
		Name:       opOutputService13TestCaseOperation2,
		HTTPMethod: "POST",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &OutputService13TestShapeOutputService13TestCaseOperation2Input{}
	}

	output := &OutputService13TestShapeOutputService13TestCaseOperation2Output{}
	req := c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService13TestCaseOperation2Request{Request: req, Input: input, Copy: c.OutputService13TestCaseOperation2Request}
}

type OutputService13TestShapeOutputService13TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService13TestShapeOutputService13TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService13TestShapeOutputService13TestCaseOperation2Input struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	FooEnum OutputService13TestShapeRESTJSONEnumType `type:"string" enum:"true"`

	HeaderEnum OutputService13TestShapeRESTJSONEnumType `location:"header" locationName:"x-amz-enum" type:"string" enum:"true"`

	ListEnums []OutputService13TestShapeRESTJSONEnumType `type:"list"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService13TestShapeOutputService13TestCaseOperation2Input) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService13TestShapeOutputService13TestCaseOperation2Input) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.FooEnum) > 0 {
		v := s.FooEnum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FooEnum", v, metadata)
	}
	if len(s.ListEnums) > 0 {
		v := s.ListEnums

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ListEnums", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.StringValue(v1))
		}
		ls0.End()

	}
	if len(s.HeaderEnum) > 0 {
		v := s.HeaderEnum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-enum", v, metadata)
	}
	return nil
}

type OutputService13TestShapeOutputService13TestCaseOperation2Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService13TestShapeOutputService13TestCaseOperation2Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService13TestShapeOutputService13TestCaseOperation2Output) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

type OutputService13TestShapeRESTJSONEnumType string

// Enum values for OutputService13TestShapeRESTJSONEnumType
const (
	RESTJSONEnumTypeFoo OutputService13TestShapeRESTJSONEnumType = "foo"
	RESTJSONEnumTypeBar OutputService13TestShapeRESTJSONEnumType = "bar"
	RESTJSONEnumType0   OutputService13TestShapeRESTJSONEnumType = "0"
	RESTJSONEnumType1   OutputService13TestShapeRESTJSONEnumType = "1"
)

func (enum OutputService13TestShapeRESTJSONEnumType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OutputService13TestShapeRESTJSONEnumType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// OutputService14ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService14ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService14ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService14ProtocolTest client with a config.
//
// Example:
//     // Create a OutputService14ProtocolTest client from just a config.
//     svc := outputservice14protocoltest.New(myConfig)
func NewOutputService14ProtocolTest(config aws.Config) *OutputService14ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService14ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice14protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService14ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService14ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService14TestCaseOperation1 = "OperationName"

// OutputService14TestCaseOperation1Request is a API request type for the OutputService14TestCaseOperation1 API operation.
type OutputService14TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService14TestShapeOutputService14TestCaseOperation1Input
	Copy  func(*OutputService14TestShapeOutputService14TestCaseOperation1Input) OutputService14TestCaseOperation1Request
}

// Send marshals and sends the OutputService14TestCaseOperation1 API request.
func (r OutputService14TestCaseOperation1Request) Send() (*OutputService14TestShapeOutputService14TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService14TestShapeOutputService14TestCaseOperation1Output), nil
}

// OutputService14TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService14TestCaseOperation1Request method.
//    req := client.OutputService14TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService14ProtocolTest) OutputService14TestCaseOperation1Request(input *OutputService14TestShapeOutputService14TestCaseOperation1Input) OutputService14TestCaseOperation1Request {
	op := &aws.Operation{
		Name:       opOutputService14TestCaseOperation1,
		HTTPMethod: "GET",
		HTTPPath:   "/path",
	}

	if input == nil {
		input = &OutputService14TestShapeOutputService14TestCaseOperation1Input{}
	}

	output := &OutputService14TestShapeOutputService14TestCaseOperation1Output{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return OutputService14TestCaseOperation1Request{Request: req, Input: input, Copy: c.OutputService14TestCaseOperation1Request}
}

type OutputService14TestShapeItemDetailShape struct {
	_ struct{} `type:"structure" xmlPrefix:"xsi" xmlURI:"http://www.w3.org/2001/XMLSchema-instance"`

	ID *string `type:"string"`

	// Type is a required field
	Type OutputService14TestShapeItemType `locationName:"xsi:type" type:"string" xmlAttribute:"true" required:"true" enum:"true"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService14TestShapeItemDetailShape) MarshalFields(e protocol.FieldEncoder) error {
	if s.ID != nil {
		v := *s.ID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ID", protocol.StringValue(v), metadata)
	}
	// Skipping Type XML Attribute.
	return nil
}

type OutputService14TestShapeItemShape struct {
	_ struct{} `type:"structure"`

	ItemDetail *OutputService14TestShapeItemDetailShape `type:"structure" xmlPrefix:"xsi" xmlURI:"http://www.w3.org/2001/XMLSchema-instance"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService14TestShapeItemShape) MarshalFields(e protocol.FieldEncoder) error {
	if s.ItemDetail != nil {
		v := s.ItemDetail
		attrs := make([]protocol.Attribute, 0, 1)

		if len(s.ItemDetail.Type) > 0 {

			v := s.ItemDetail.Type
			attrs = append(attrs, protocol.Attribute{Name: "xsi:type", Value: v, Meta: protocol.Metadata{}})
		}
		metadata := protocol.Metadata{Attributes: attrs, XMLNamespacePrefix: "xsi", XMLNamespaceURI: "http://www.w3.org/2001/XMLSchema-instance"}
		e.SetFields(protocol.BodyTarget, "ItemDetail", v, metadata)
	}
	return nil
}

type OutputService14TestShapeOutputService14TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService14TestShapeOutputService14TestCaseOperation1Input) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type OutputService14TestShapeOutputService14TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	ListItems []OutputService14TestShapeItemShape `locationName:"ItemsList" locationNameList:"Item" type:"list"`
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s OutputService14TestShapeOutputService14TestCaseOperation1Output) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s OutputService14TestShapeOutputService14TestCaseOperation1Output) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ListItems) > 0 {
		v := s.ListItems

		metadata := protocol.Metadata{ListLocationName: "Item"}
		ls0 := e.List(protocol.BodyTarget, "ItemsList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type OutputService14TestShapeItemType string

// Enum values for OutputService14TestShapeItemType
const (
	ItemTypeType1 OutputService14TestShapeItemType = "Type1"
	ItemTypeType2 OutputService14TestShapeItemType = "Type2"
	ItemTypeType3 OutputService14TestShapeItemType = "Type3"
)

func (enum OutputService14TestShapeItemType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OutputService14TestShapeItemType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService1ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><Str>myname</Str><FooNum>123</FooNum><FalseBool>false</FalseBool><TrueBool>true</TrueBool><Float>1.2</Float><Double>1.3</Double><Long>200</Long><Char>a</Char><Timestamp>2015-01-25T08:00:00Z</Timestamp></OperationNameResponse>"))
	req := svc.OutputService1TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("ImaHeader", "test")
	req.HTTPResponse.Header.Set("X-Foo", "abc")

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService1TestShapeOutputService1TestCaseOperation2Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "a", *out.Char; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.3, *out.Double; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := false, *out.FalseBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.2, *out.Float; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("test"), *out.ImaHeader; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("abc"), *out.ImaHeaderLocation; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(200), *out.Long; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(123), *out.Num; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("myname"), *out.Str; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := time.Unix(1.4221728e+09, 0).UTC().String(), out.Timestamp.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, *out.TrueBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService1ProtocolTestScalarMembersCase2(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService1ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><Str></Str><FooNum>123</FooNum><FalseBool>false</FalseBool><TrueBool>true</TrueBool><Float>1.2</Float><Double>1.3</Double><Long>200</Long><Char>a</Char><Timestamp>2015-01-25T08:00:00Z</Timestamp></OperationNameResponse>"))
	req := svc.OutputService1TestCaseOperation2Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("ImaHeader", "test")
	req.HTTPResponse.Header.Set("X-Foo", "abc")

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService1TestShapeOutputService1TestCaseOperation2Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "a", *out.Char; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.3, *out.Double; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := false, *out.FalseBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.2, *out.Float; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("test"), *out.ImaHeader; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("abc"), *out.ImaHeaderLocation; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(200), *out.Long; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(123), *out.Num; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string(""), *out.Str; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := time.Unix(1.4221728e+09, 0).UTC().String(), out.Timestamp.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, *out.TrueBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService2ProtocolTestBlobCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService2ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><Blob>dmFsdWU=</Blob></OperationNameResult>"))
	req := svc.OutputService2TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService2TestShapeOutputService2TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "value", string(out.Blob); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService3ProtocolTestListsCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService3ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><ListMember><member>abc</member><member>123</member></ListMember></OperationNameResult>"))
	req := svc.OutputService3TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService3TestShapeOutputService3TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("123"), out.ListMember[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService4ProtocolTestListWithCustomMemberNameCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService4ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><ListMember><item>abc</item><item>123</item></ListMember></OperationNameResult>"))
	req := svc.OutputService4TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService4TestShapeOutputService4TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("123"), out.ListMember[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService5ProtocolTestFlattenedListCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService5ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><ListMember>abc</ListMember><ListMember>123</ListMember></OperationNameResult>"))
	req := svc.OutputService5TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService5TestShapeOutputService5TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("123"), out.ListMember[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService6ProtocolTestNormalMapCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService6ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><Map><entry><key>qux</key><value><foo>bar</foo></value></entry><entry><key>baz</key><value><foo>bam</foo></value></entry></Map></OperationNameResult>"))
	req := svc.OutputService6TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService6TestShapeOutputService6TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bam"), *out.Map["baz"].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("bar"), *out.Map["qux"].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService7ProtocolTestFlattenedMapCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService7ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><Map><key>qux</key><value>bar</value></Map><Map><key>baz</key><value>bam</value></Map></OperationNameResult>"))
	req := svc.OutputService7TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService7TestShapeOutputService7TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bam"), out.Map["baz"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("bar"), out.Map["qux"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService8ProtocolTestNamedMapCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService8ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResult><Map><entry><foo>qux</foo><bar>bar</bar></entry><entry><foo>baz</foo><bar>bam</bar></entry></Map></OperationNameResult>"))
	req := svc.OutputService8TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService8TestShapeOutputService8TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bam"), out.Map["baz"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("bar"), out.Map["qux"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService9ProtocolTestXMLPayloadCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService9ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><Foo>abc</Foo></OperationNameResponse>"))
	req := svc.OutputService9TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("X-Foo", "baz")

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService9TestShapeOutputService9TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), *out.Data.Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("baz"), *out.Header; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService10ProtocolTestStreamingPayloadCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService10ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("abc"))
	req := svc.OutputService10TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService10TestShapeOutputService10TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "abc", string(out.Stream); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService11ProtocolTestScalarMembersInHeadersCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService11ProtocolTest(cfg)

	buf := bytes.NewReader([]byte(""))
	req := svc.OutputService11TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("x-char", "a")
	req.HTTPResponse.Header.Set("x-double", "1.5")
	req.HTTPResponse.Header.Set("x-false-bool", "false")
	req.HTTPResponse.Header.Set("x-float", "1.5")
	req.HTTPResponse.Header.Set("x-int", "1")
	req.HTTPResponse.Header.Set("x-long", "100")
	req.HTTPResponse.Header.Set("x-str", "string")
	req.HTTPResponse.Header.Set("x-timestamp", "Sun, 25 Jan 2015 08:00:00 GMT")
	req.HTTPResponse.Header.Set("x-true-bool", "true")

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService11TestShapeOutputService11TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "a", *out.Char; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.5, *out.Double; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := false, *out.FalseBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.5, *out.Float; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(1), *out.Integer; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(100), *out.Long; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("string"), *out.Str; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := time.Unix(1.4221728e+09, 0).UTC().String(), out.Timestamp.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, *out.TrueBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService12ProtocolTestEmptyStringCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService12ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><Foo/><RequestId>requestid</RequestId></OperationNameResponse>"))
	req := svc.OutputService12TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService12TestShapeOutputService12TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string(""), *out.Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService13ProtocolTestEnumCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService13ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><FooEnum>foo</FooEnum><ListEnums><member>0</member><member>1</member></ListEnums></OperationNameResponse>"))
	req := svc.OutputService13TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	req.HTTPResponse.Header.Set("x-amz-enum", "baz")

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService13TestShapeOutputService13TestCaseOperation2Input)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := OutputService13TestShapeRESTJSONEnumType("foo"), out.FooEnum; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService13TestShapeRESTJSONEnumType("baz"), out.HeaderEnum; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService13TestShapeRESTJSONEnumType("0"), out.ListEnums[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService13TestShapeRESTJSONEnumType("1"), out.ListEnums[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService13ProtocolTestEnumCase2(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService13ProtocolTest(cfg)

	buf := bytes.NewReader([]byte(""))
	req := svc.OutputService13TestCaseOperation2Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService13TestShapeOutputService13TestCaseOperation2Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}

}

func TestOutputService14ProtocolTestXMLAttributesCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService14ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<SomeOutputDoc xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\"><ItemsList><Item><ItemDetail xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"Type1\"><ID>id1</ID></ItemDetail></Item><Item><ItemDetail xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"Type2\"><ID>id2</ID></ItemDetail></Item><Item><ItemDetail xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"Type3\"><ID>id3</ID></ItemDetail></Item></ItemsList></SomeOutputDoc>"))
	req := svc.OutputService14TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	restxml.UnmarshalMeta(req.Request)
	restxml.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService14TestShapeOutputService14TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("id1"), *out.ListItems[0].ItemDetail.ID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService14TestShapeItemType("Type1"), out.ListItems[0].ItemDetail.Type; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("id2"), *out.ListItems[1].ItemDetail.ID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService14TestShapeItemType("Type2"), out.ListItems[1].ItemDetail.Type; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("id3"), *out.ListItems[2].ItemDetail.ID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService14TestShapeItemType("Type3"), out.ListItems[2].ItemDetail.Type; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}