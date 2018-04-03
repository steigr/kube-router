// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opStartOutboundVoiceContact = "StartOutboundVoiceContact"

// StartOutboundVoiceContactRequest generates a "aws/request.Request" representing the
// client's request for the StartOutboundVoiceContact operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See StartOutboundVoiceContact for more information on using the StartOutboundVoiceContact
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the StartOutboundVoiceContactRequest method.
//    req, resp := client.StartOutboundVoiceContactRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/StartOutboundVoiceContact
func (c *Connect) StartOutboundVoiceContactRequest(input *StartOutboundVoiceContactInput) (req *request.Request, output *StartOutboundVoiceContactOutput) {
	op := &request.Operation{
		Name:       opStartOutboundVoiceContact,
		HTTPMethod: "PUT",
		HTTPPath:   "/contact/outbound-voice",
	}

	if input == nil {
		input = &StartOutboundVoiceContactInput{}
	}

	output = &StartOutboundVoiceContactOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StartOutboundVoiceContact API operation for Amazon Connect Service.
//
// The StartOutboundVoiceContact operation initiates a contact flow to place
// an outbound call to a customer.
//
// There is a throttling limit placed on usage of the API that includes a RateLimit
// of 2 per second, and a BurstLimit of 5 per second.
//
// If you are using an IAM account, it must have permissions to the connect:StartOutboundVoiceContact
// action.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Connect Service's
// API operation StartOutboundVoiceContact for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   One or more of the parameters provided to the operation are not valid.
//
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The specified resource was not found.
//
//   * ErrCodeInternalServiceException "InternalServiceException"
//   Request processing failed due to an error or failure with the service.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//   The limit exceeded the maximum allowed active calls in a queue.
//
//   * ErrCodeDestinationNotAllowedException "DestinationNotAllowedException"
//   Outbound calls to the destination number are not allowed for your instance.
//   You can request that the country be included in the allowed countries for
//   your instance by submitting a Service Limit Increase (https://console.aws.amazon.com/support/v1#/case/create?issueType=service-limit-increase).
//
//   * ErrCodeOutboundContactNotPermittedException "OutboundContactNotPermittedException"
//   The contact is not permitted because outbound calling is not enabled for
//   the instance.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/StartOutboundVoiceContact
func (c *Connect) StartOutboundVoiceContact(input *StartOutboundVoiceContactInput) (*StartOutboundVoiceContactOutput, error) {
	req, out := c.StartOutboundVoiceContactRequest(input)
	return out, req.Send()
}

// StartOutboundVoiceContactWithContext is the same as StartOutboundVoiceContact with the addition of
// the ability to pass a context and additional request options.
//
// See StartOutboundVoiceContact for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Connect) StartOutboundVoiceContactWithContext(ctx aws.Context, input *StartOutboundVoiceContactInput, opts ...request.Option) (*StartOutboundVoiceContactOutput, error) {
	req, out := c.StartOutboundVoiceContactRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopContact = "StopContact"

// StopContactRequest generates a "aws/request.Request" representing the
// client's request for the StopContact operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See StopContact for more information on using the StopContact
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the StopContactRequest method.
//    req, resp := client.StopContactRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/StopContact
func (c *Connect) StopContactRequest(input *StopContactInput) (req *request.Request, output *StopContactOutput) {
	op := &request.Operation{
		Name:       opStopContact,
		HTTPMethod: "POST",
		HTTPPath:   "/contact/stop",
	}

	if input == nil {
		input = &StopContactInput{}
	}

	output = &StopContactOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StopContact API operation for Amazon Connect Service.
//
// Ends the contact initiated by the StartOutboundVoiceContact operation.
//
// If you are using an IAM account, it must have permissions to the connect:StopContact
// operation.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Connect Service's
// API operation StopContact for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request is not valid.
//
//   * ErrCodeContactNotFoundException "ContactNotFoundException"
//   The contact with the specified ID is not active or does not exist.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   One or more of the parameters provided to the operation are not valid.
//
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The specified resource was not found.
//
//   * ErrCodeInternalServiceException "InternalServiceException"
//   Request processing failed due to an error or failure with the service.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/StopContact
func (c *Connect) StopContact(input *StopContactInput) (*StopContactOutput, error) {
	req, out := c.StopContactRequest(input)
	return out, req.Send()
}

// StopContactWithContext is the same as StopContact with the addition of
// the ability to pass a context and additional request options.
//
// See StopContact for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Connect) StopContactWithContext(ctx aws.Context, input *StopContactInput, opts ...request.Option) (*StopContactOutput, error) {
	req, out := c.StopContactRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartOutboundVoiceContactInput struct {
	_ struct{} `type:"structure"`

	// Specify a custom key-value pair using an attribute map. The attributes are
	// standard Amazon Connect attributes, and can be accessed in contact flows
	// just like any other contact attributes.
	//
	// There can be up to 32,768 UTF-8 bytes across all key-value pairs. Attribute
	// keys can include only alphanumeric, dash, and underscore characters.
	//
	// For example, to play a greeting when the customer answers the call, you can
	// pass the customer name in attributes similar to the following:
	Attributes map[string]*string `type:"map"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. The token is valid for 7 days after creation. If a contact
	// is already started, the contact ID is returned. If the contact is disconnected,
	// a new contact is started.
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The identifier for the contact flow to execute for the outbound call. This
	// is a GUID value only. Amazon Resource Name (ARN) values are not supported.
	//
	// To find the ContactFlowId, open the contact flow to use in the Amazon Connect
	// contact flow designer. The ID for the contact flow is displayed in the address
	// bar as part of the URL. For example, an address displayed when you open a
	// contact flow is similar to the following: https://myconnectinstance.awsapps.com/connect/contact-flows/edit?id=arn:aws:connect:us-east-1:361814831152:instance/2fb42df9-78a2-4b99-b484-f5cf80dc300c/contact-flow/b0b8f2dd-ed1b-4c44-af36-ce189a178181.
	// At the end of the URL, you see contact-flow/b0b8f2dd-ed1b-4c44-af36-ce189a178181.
	// The ContactFlowID for this contact flow is b0b8f2dd-ed1b-4c44-af36-ce189a178181.
	// Make sure to include only the GUID after the "contact-flow/" in your requests.
	//
	// ContactFlowId is a required field
	ContactFlowId *string `type:"string" required:"true"`

	// The phone number, in E.164 format, of the customer to call with the outbound
	// contact.
	//
	// DestinationPhoneNumber is a required field
	DestinationPhoneNumber *string `type:"string" required:"true"`

	// The identifier for your Amazon Connect instance. To find the InstanceId value
	// for your Amazon Connect instance, open the Amazon Connect console (https://console.aws.amazon.com/connect/).
	// Select the instance alias of the instance and view the instance ID in the
	// Overview section. For example, the instance ID is the set of characters at
	// the end of the instance ARN, after "instance/", such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The queue to which to add the call. If you specify a queue, the phone displayed
	// for caller ID is the phone number defined for the queue. If you do not specify
	// a queue, the queue used is the queue defined in the contact flow specified
	// by ContactFlowId.
	//
	// To find the QueueId, open the queue to use in the Amazon Connect queue editor.
	// The ID for the queue is displayed in the address bar as part of the URL.
	// For example, the QueueId value is the set of characters at the end of the
	// URL, after "queue/", such as aeg40574-2d01-51c3-73d6-bf8624d2168c.
	QueueId *string `type:"string"`

	// The phone number, in E.164 format, associated with your Amazon Connect instance
	// to use to place the outbound call.
	SourcePhoneNumber *string `type:"string"`
}

// String returns the string representation
func (s StartOutboundVoiceContactInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartOutboundVoiceContactInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartOutboundVoiceContactInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartOutboundVoiceContactInput"}
	if s.ContactFlowId == nil {
		invalidParams.Add(request.NewErrParamRequired("ContactFlowId"))
	}
	if s.DestinationPhoneNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationPhoneNumber"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *StartOutboundVoiceContactInput) SetAttributes(v map[string]*string) *StartOutboundVoiceContactInput {
	s.Attributes = v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *StartOutboundVoiceContactInput) SetClientToken(v string) *StartOutboundVoiceContactInput {
	s.ClientToken = &v
	return s
}

// SetContactFlowId sets the ContactFlowId field's value.
func (s *StartOutboundVoiceContactInput) SetContactFlowId(v string) *StartOutboundVoiceContactInput {
	s.ContactFlowId = &v
	return s
}

// SetDestinationPhoneNumber sets the DestinationPhoneNumber field's value.
func (s *StartOutboundVoiceContactInput) SetDestinationPhoneNumber(v string) *StartOutboundVoiceContactInput {
	s.DestinationPhoneNumber = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *StartOutboundVoiceContactInput) SetInstanceId(v string) *StartOutboundVoiceContactInput {
	s.InstanceId = &v
	return s
}

// SetQueueId sets the QueueId field's value.
func (s *StartOutboundVoiceContactInput) SetQueueId(v string) *StartOutboundVoiceContactInput {
	s.QueueId = &v
	return s
}

// SetSourcePhoneNumber sets the SourcePhoneNumber field's value.
func (s *StartOutboundVoiceContactInput) SetSourcePhoneNumber(v string) *StartOutboundVoiceContactInput {
	s.SourcePhoneNumber = &v
	return s
}

type StartOutboundVoiceContactOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of this contact within your Amazon Connect instance.
	ContactId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartOutboundVoiceContactOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartOutboundVoiceContactOutput) GoString() string {
	return s.String()
}

// SetContactId sets the ContactId field's value.
func (s *StartOutboundVoiceContactOutput) SetContactId(v string) *StartOutboundVoiceContactOutput {
	s.ContactId = &v
	return s
}

type StopContactInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the contact to end. This is the ContactId value
	// returned from the StartOutboundVoiceContact operation.
	//
	// ContactId is a required field
	ContactId *string `min:"1" type:"string" required:"true"`

	// The identifier of the Amazon Connect instance in which the contact is active.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StopContactInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopContactInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopContactInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StopContactInput"}
	if s.ContactId == nil {
		invalidParams.Add(request.NewErrParamRequired("ContactId"))
	}
	if s.ContactId != nil && len(*s.ContactId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ContactId", 1))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContactId sets the ContactId field's value.
func (s *StopContactInput) SetContactId(v string) *StopContactInput {
	s.ContactId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *StopContactInput) SetInstanceId(v string) *StopContactInput {
	s.InstanceId = &v
	return s
}

type StopContactOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopContactOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopContactOutput) GoString() string {
	return s.String()
}
