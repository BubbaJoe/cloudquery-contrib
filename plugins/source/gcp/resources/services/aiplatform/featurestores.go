package aiplatform

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"google.golang.org/genproto/googleapis/cloud/location"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
)

func Featurestores() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_featurestores",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores#Featurestore`,
		Resolver:    fetchFeaturestores,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Featurestore{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

func fetchFeaturestores(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListFeaturestoresRequest{
		Parent: parent.Item.(*location.Location).Name,
	}
	if filterLocation(parent) {
		return nil
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewFeaturestoreClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListFeaturestores(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
