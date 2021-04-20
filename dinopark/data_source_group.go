package dinopark

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	person_api "go.mozilla.org/person-api"
)

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Returns information about specific group.",
		ReadContext: dataSourceDinoParkGroups,
		Schema:      map[string]*schema.Schema{},
	}
}

func dataSourceDinoParkGroups(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	personClient := m.(*person_api.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	personClient.GetAllUsers()

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
