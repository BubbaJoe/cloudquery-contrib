package treasury

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryReceivedCredits() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_received_credits",
		Description: `https://stripe.com/docs/api/treasury/received_credits`,
		Transform:   client.TransformWithStruct(&stripe.TreasuryReceivedCredit{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryReceivedCredits,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchTreasuryReceivedCredits(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryReceivedCreditListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryReceivedCredits.List(lp)
	for it.Next() {
		res <- it.TreasuryReceivedCredit()
	}

	return it.Err()
}
