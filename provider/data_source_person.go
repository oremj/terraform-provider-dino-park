package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	person_api "go.mozilla.org/person-api"
)

func dataSourcePerson() *schema.Resource {
	return &schema.Resource{
		Description: "Returns information about specific person.",
		ReadContext: dataSourceDinoParkPerson,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDinoParkPerson(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	personClient := m.(*person_api.Client)
	email := d.Get("email").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	person, err := personClient.GetPersonByEmail(email)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching person: %s", err))
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	var groups []string
	for group := range person.AccessInformation.LDAP.Values {
		groups = append(groups, group)
	}

	if err := d.Set("groups", groups); err != nil {
		return diag.FromErr(fmt.Errorf("error setting groups: %s", err))
	}

	return diags
}
