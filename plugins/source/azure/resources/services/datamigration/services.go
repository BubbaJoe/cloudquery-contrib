package datamigration

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:                 "azure_datamigration_services",
		Resolver:             fetchServices,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/datamigration/services/list?tabs=HTTP#datamigrationservice",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_datamigration_services", client.Namespacemicrosoft_datamigration),
		Transform:            transformers.TransformWithStruct(&armdatamigration.Service{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdatamigration.NewServicesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
