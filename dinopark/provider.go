package dinopark

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	person_api "go.mozilla.org/person-api"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DINOPARK_CLIENT_ID", nil),
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DINOPARK_CLIENT_SECRET", nil),
			},
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DINOPARK_BASE_URL", nil),
			},
			"auth0_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DINOPARK_AUTH0_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"dinopark_group": dataSourceGroups(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	baseURL := d.Get("base_url").(string)
	auth0URL := d.Get("auth0_url").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	client, err := person_api.NewClient(clientId, clientSecret, baseURL, auth0URL)

	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}
