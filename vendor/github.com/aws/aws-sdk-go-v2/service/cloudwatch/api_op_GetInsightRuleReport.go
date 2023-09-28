// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation returns the time series data collected by a Contributor Insights
// rule. The data includes the identity and number of contributors to the log
// group. You can also optionally return one or more statistics about each data
// point in the time series. These statistics can include the following:
//   - UniqueContributors -- the number of unique contributors for each data point.
//   - MaxContributorValue -- the value of the top contributor for each data point.
//     The identity of the contributor might change for each data point in the graph.
//     If this rule aggregates by COUNT, the top contributor for each data point is the
//     contributor with the most occurrences in that period. If the rule aggregates by
//     SUM, the top contributor is the contributor with the highest sum in the log
//     field specified by the rule's Value , during that period.
//   - SampleCount -- the number of data points matched by the rule.
//   - Sum -- the sum of the values from all contributors during the time period
//     represented by that data point.
//   - Minimum -- the minimum value from a single observation during the time
//     period represented by that data point.
//   - Maximum -- the maximum value from a single observation during the time
//     period represented by that data point.
//   - Average -- the average value from all contributors during the time period
//     represented by that data point.
func (c *Client) GetInsightRuleReport(ctx context.Context, params *GetInsightRuleReportInput, optFns ...func(*Options)) (*GetInsightRuleReportOutput, error) {
	if params == nil {
		params = &GetInsightRuleReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInsightRuleReport", params, optFns, c.addOperationGetInsightRuleReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInsightRuleReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInsightRuleReportInput struct {

	// The end time of the data to use in the report. When used in a raw HTTP Query
	// API, it is formatted as yyyy-MM-dd'T'HH:mm:ss . For example, 2019-07-01T23:59:59
	// .
	//
	// This member is required.
	EndTime *time.Time

	// The period, in seconds, to use for the statistics in the
	// InsightRuleMetricDatapoint results.
	//
	// This member is required.
	Period *int32

	// The name of the rule that you want to see data from.
	//
	// This member is required.
	RuleName *string

	// The start time of the data to use in the report. When used in a raw HTTP Query
	// API, it is formatted as yyyy-MM-dd'T'HH:mm:ss . For example, 2019-07-01T23:59:59
	// .
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of contributors to include in the report. The range is 1 to
	// 100. If you omit this, the default of 10 is used.
	MaxContributorCount *int32

	// Specifies which metrics to use for aggregation of contributor values for the
	// report. You can specify one or more of the following metrics:
	//   - UniqueContributors -- the number of unique contributors for each data point.
	//   - MaxContributorValue -- the value of the top contributor for each data point.
	//   The identity of the contributor might change for each data point in the graph.
	//   If this rule aggregates by COUNT, the top contributor for each data point is the
	//   contributor with the most occurrences in that period. If the rule aggregates by
	//   SUM, the top contributor is the contributor with the highest sum in the log
	//   field specified by the rule's Value , during that period.
	//   - SampleCount -- the number of data points matched by the rule.
	//   - Sum -- the sum of the values from all contributors during the time period
	//   represented by that data point.
	//   - Minimum -- the minimum value from a single observation during the time
	//   period represented by that data point.
	//   - Maximum -- the maximum value from a single observation during the time
	//   period represented by that data point.
	//   - Average -- the average value from all contributors during the time period
	//   represented by that data point.
	Metrics []string

	// Determines what statistic to use to rank the contributors. Valid values are Sum
	// and Maximum .
	OrderBy *string

	noSmithyDocumentSerde
}

type GetInsightRuleReportOutput struct {

	// The sum of the values from all individual contributors that match the rule.
	AggregateValue *float64

	// Specifies whether this rule aggregates contributor data by COUNT or SUM.
	AggregationStatistic *string

	// An approximate count of the unique contributors found by this rule in this time
	// period.
	ApproximateUniqueCount *int64

	// An array of the unique contributors found by this rule in this time period. If
	// the rule contains multiple keys, each combination of values for the keys counts
	// as a unique contributor.
	Contributors []types.InsightRuleContributor

	// An array of the strings used as the keys for this rule. The keys are the
	// dimensions used to classify contributors. If the rule contains more than one
	// key, then each unique combination of values for the keys is counted as a unique
	// contributor.
	KeyLabels []string

	// A time series of metric data points that matches the time period in the rule
	// request.
	MetricDatapoints []types.InsightRuleMetricDatapoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInsightRuleReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetInsightRuleReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetInsightRuleReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addGetInsightRuleReportResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetInsightRuleReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInsightRuleReport(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetInsightRuleReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetInsightRuleReport",
	}
}

type opGetInsightRuleReportResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetInsightRuleReportResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetInsightRuleReportResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "monitoring"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "monitoring"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("monitoring")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addGetInsightRuleReportResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetInsightRuleReportResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}