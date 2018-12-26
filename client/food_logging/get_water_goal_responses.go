// Code generated by go-swagger; DO NOT EDIT.

package food_logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetWaterGoalReader is a Reader for the GetWaterGoal structure.
type GetWaterGoalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWaterGoalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWaterGoalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWaterGoalOK creates a GetWaterGoalOK with default headers values
func NewGetWaterGoalOK() *GetWaterGoalOK {
	return &GetWaterGoalOK{}
}

/*GetWaterGoalOK handles this case with default header values.

A successful request.
*/
type GetWaterGoalOK struct {
}

func (o *GetWaterGoalOK) Error() string {
	return fmt.Sprintf("[GET /1/user/-/foods/log/water/goal.json][%d] getWaterGoalOK ", 200)
}

func (o *GetWaterGoalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
