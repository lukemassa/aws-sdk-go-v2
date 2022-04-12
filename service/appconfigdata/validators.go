// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfigdata

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpGetLatestConfiguration struct {
}

func (*validateOpGetLatestConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLatestConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLatestConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLatestConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartConfigurationSession struct {
}

func (*validateOpStartConfigurationSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartConfigurationSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartConfigurationSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartConfigurationSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetLatestConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLatestConfiguration{}, middleware.After)
}

func addOpStartConfigurationSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartConfigurationSession{}, middleware.After)
}

func validateOpGetLatestConfigurationInput(v *GetLatestConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLatestConfigurationInput"}
	if v.ConfigurationToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartConfigurationSessionInput(v *StartConfigurationSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartConfigurationSessionInput"}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ConfigurationProfileIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ConfigurationProfileIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}