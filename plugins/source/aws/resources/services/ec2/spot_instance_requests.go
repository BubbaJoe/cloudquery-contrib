package ec2

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v3/types"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func SpotInstanceRequests() *schema.Table {
	tableName := "aws_ec2_spot_instance_requests"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotInstanceRequest.html`,
		Resolver:    fetchEC2SpotInstanceRequests,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.SpotInstanceRequest{}, transformers.WithPrimaryKeys("SpotInstanceRequestId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEC2SpotInstanceRequests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2
	pag := ec2.NewDescribeSpotInstanceRequestsPaginator(svc, &ec2.DescribeSpotInstanceRequestsInput{})
	for pag.HasMorePages() {
		resp, err := pag.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp.SpotInstanceRequests
	}
	return nil
}
