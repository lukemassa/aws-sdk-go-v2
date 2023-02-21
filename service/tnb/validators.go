// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelSolNetworkOperation struct {
}

func (*validateOpCancelSolNetworkOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelSolNetworkOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelSolNetworkOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelSolNetworkOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSolNetworkInstance struct {
}

func (*validateOpCreateSolNetworkInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSolNetworkInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSolNetworkInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSolNetworkInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSolFunctionPackage struct {
}

func (*validateOpDeleteSolFunctionPackage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSolFunctionPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSolFunctionPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSolFunctionPackageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSolNetworkInstance struct {
}

func (*validateOpDeleteSolNetworkInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSolNetworkInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSolNetworkInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSolNetworkInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSolNetworkPackage struct {
}

func (*validateOpDeleteSolNetworkPackage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSolNetworkPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSolNetworkPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSolNetworkPackageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolFunctionInstance struct {
}

func (*validateOpGetSolFunctionInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolFunctionInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolFunctionInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolFunctionInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolFunctionPackageContent struct {
}

func (*validateOpGetSolFunctionPackageContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolFunctionPackageContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolFunctionPackageContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolFunctionPackageContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolFunctionPackageDescriptor struct {
}

func (*validateOpGetSolFunctionPackageDescriptor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolFunctionPackageDescriptor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolFunctionPackageDescriptorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolFunctionPackageDescriptorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolFunctionPackage struct {
}

func (*validateOpGetSolFunctionPackage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolFunctionPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolFunctionPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolFunctionPackageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolNetworkInstance struct {
}

func (*validateOpGetSolNetworkInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolNetworkInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolNetworkInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolNetworkInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolNetworkOperation struct {
}

func (*validateOpGetSolNetworkOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolNetworkOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolNetworkOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolNetworkOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolNetworkPackageContent struct {
}

func (*validateOpGetSolNetworkPackageContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolNetworkPackageContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolNetworkPackageContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolNetworkPackageContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolNetworkPackageDescriptor struct {
}

func (*validateOpGetSolNetworkPackageDescriptor) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolNetworkPackageDescriptor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolNetworkPackageDescriptorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolNetworkPackageDescriptorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSolNetworkPackage struct {
}

func (*validateOpGetSolNetworkPackage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSolNetworkPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSolNetworkPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSolNetworkPackageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInstantiateSolNetworkInstance struct {
}

func (*validateOpInstantiateSolNetworkInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInstantiateSolNetworkInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InstantiateSolNetworkInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInstantiateSolNetworkInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutSolFunctionPackageContent struct {
}

func (*validateOpPutSolFunctionPackageContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutSolFunctionPackageContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutSolFunctionPackageContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutSolFunctionPackageContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutSolNetworkPackageContent struct {
}

func (*validateOpPutSolNetworkPackageContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutSolNetworkPackageContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutSolNetworkPackageContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutSolNetworkPackageContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTerminateSolNetworkInstance struct {
}

func (*validateOpTerminateSolNetworkInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTerminateSolNetworkInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TerminateSolNetworkInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTerminateSolNetworkInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSolFunctionPackage struct {
}

func (*validateOpUpdateSolFunctionPackage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSolFunctionPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSolFunctionPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSolFunctionPackageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSolNetworkInstance struct {
}

func (*validateOpUpdateSolNetworkInstance) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSolNetworkInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSolNetworkInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSolNetworkInstanceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSolNetworkPackage struct {
}

func (*validateOpUpdateSolNetworkPackage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSolNetworkPackage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSolNetworkPackageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSolNetworkPackageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpValidateSolFunctionPackageContent struct {
}

func (*validateOpValidateSolFunctionPackageContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpValidateSolFunctionPackageContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ValidateSolFunctionPackageContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpValidateSolFunctionPackageContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpValidateSolNetworkPackageContent struct {
}

func (*validateOpValidateSolNetworkPackageContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpValidateSolNetworkPackageContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ValidateSolNetworkPackageContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpValidateSolNetworkPackageContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelSolNetworkOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelSolNetworkOperation{}, middleware.After)
}

func addOpCreateSolNetworkInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSolNetworkInstance{}, middleware.After)
}

func addOpDeleteSolFunctionPackageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSolFunctionPackage{}, middleware.After)
}

func addOpDeleteSolNetworkInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSolNetworkInstance{}, middleware.After)
}

func addOpDeleteSolNetworkPackageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSolNetworkPackage{}, middleware.After)
}

func addOpGetSolFunctionInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolFunctionInstance{}, middleware.After)
}

func addOpGetSolFunctionPackageContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolFunctionPackageContent{}, middleware.After)
}

func addOpGetSolFunctionPackageDescriptorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolFunctionPackageDescriptor{}, middleware.After)
}

func addOpGetSolFunctionPackageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolFunctionPackage{}, middleware.After)
}

func addOpGetSolNetworkInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolNetworkInstance{}, middleware.After)
}

func addOpGetSolNetworkOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolNetworkOperation{}, middleware.After)
}

func addOpGetSolNetworkPackageContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolNetworkPackageContent{}, middleware.After)
}

func addOpGetSolNetworkPackageDescriptorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolNetworkPackageDescriptor{}, middleware.After)
}

func addOpGetSolNetworkPackageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSolNetworkPackage{}, middleware.After)
}

func addOpInstantiateSolNetworkInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInstantiateSolNetworkInstance{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutSolFunctionPackageContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutSolFunctionPackageContent{}, middleware.After)
}

func addOpPutSolNetworkPackageContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutSolNetworkPackageContent{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpTerminateSolNetworkInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTerminateSolNetworkInstance{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateSolFunctionPackageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSolFunctionPackage{}, middleware.After)
}

func addOpUpdateSolNetworkInstanceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSolNetworkInstance{}, middleware.After)
}

func addOpUpdateSolNetworkPackageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSolNetworkPackage{}, middleware.After)
}

func addOpValidateSolFunctionPackageContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpValidateSolFunctionPackageContent{}, middleware.After)
}

func addOpValidateSolNetworkPackageContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpValidateSolNetworkPackageContent{}, middleware.After)
}

func validateUpdateSolNetworkModify(v *types.UpdateSolNetworkModify) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSolNetworkModify"}
	if v.VnfInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfInstanceId"))
	}
	if v.VnfConfigurableProperties == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfConfigurableProperties"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelSolNetworkOperationInput(v *CancelSolNetworkOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelSolNetworkOperationInput"}
	if v.NsLcmOpOccId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsLcmOpOccId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSolNetworkInstanceInput(v *CreateSolNetworkInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSolNetworkInstanceInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if v.NsName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSolFunctionPackageInput(v *DeleteSolFunctionPackageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSolFunctionPackageInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSolNetworkInstanceInput(v *DeleteSolNetworkInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSolNetworkInstanceInput"}
	if v.NsInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSolNetworkPackageInput(v *DeleteSolNetworkPackageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSolNetworkPackageInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolFunctionInstanceInput(v *GetSolFunctionInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolFunctionInstanceInput"}
	if v.VnfInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolFunctionPackageContentInput(v *GetSolFunctionPackageContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolFunctionPackageContentInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if len(v.Accept) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Accept"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolFunctionPackageDescriptorInput(v *GetSolFunctionPackageDescriptorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolFunctionPackageDescriptorInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if len(v.Accept) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Accept"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolFunctionPackageInput(v *GetSolFunctionPackageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolFunctionPackageInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolNetworkInstanceInput(v *GetSolNetworkInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolNetworkInstanceInput"}
	if v.NsInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolNetworkOperationInput(v *GetSolNetworkOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolNetworkOperationInput"}
	if v.NsLcmOpOccId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsLcmOpOccId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolNetworkPackageContentInput(v *GetSolNetworkPackageContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolNetworkPackageContentInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if len(v.Accept) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Accept"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolNetworkPackageDescriptorInput(v *GetSolNetworkPackageDescriptorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolNetworkPackageDescriptorInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSolNetworkPackageInput(v *GetSolNetworkPackageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSolNetworkPackageInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpInstantiateSolNetworkInstanceInput(v *InstantiateSolNetworkInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InstantiateSolNetworkInstanceInput"}
	if v.NsInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutSolFunctionPackageContentInput(v *PutSolFunctionPackageContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutSolFunctionPackageContentInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if v.File == nil {
		invalidParams.Add(smithy.NewErrParamRequired("File"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutSolNetworkPackageContentInput(v *PutSolNetworkPackageContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutSolNetworkPackageContentInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if v.File == nil {
		invalidParams.Add(smithy.NewErrParamRequired("File"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTerminateSolNetworkInstanceInput(v *TerminateSolNetworkInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TerminateSolNetworkInstanceInput"}
	if v.NsInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsInstanceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSolFunctionPackageInput(v *UpdateSolFunctionPackageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSolFunctionPackageInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if len(v.OperationalState) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OperationalState"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSolNetworkInstanceInput(v *UpdateSolNetworkInstanceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSolNetworkInstanceInput"}
	if v.NsInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsInstanceId"))
	}
	if len(v.UpdateType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("UpdateType"))
	}
	if v.ModifyVnfInfoData != nil {
		if err := validateUpdateSolNetworkModify(v.ModifyVnfInfoData); err != nil {
			invalidParams.AddNested("ModifyVnfInfoData", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSolNetworkPackageInput(v *UpdateSolNetworkPackageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSolNetworkPackageInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if len(v.NsdOperationalState) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("NsdOperationalState"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpValidateSolFunctionPackageContentInput(v *ValidateSolFunctionPackageContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ValidateSolFunctionPackageContentInput"}
	if v.VnfPkgId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VnfPkgId"))
	}
	if v.File == nil {
		invalidParams.Add(smithy.NewErrParamRequired("File"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpValidateSolNetworkPackageContentInput(v *ValidateSolNetworkPackageContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ValidateSolNetworkPackageContentInput"}
	if v.NsdInfoId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NsdInfoId"))
	}
	if v.File == nil {
		invalidParams.Add(smithy.NewErrParamRequired("File"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
