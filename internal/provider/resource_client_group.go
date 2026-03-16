package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ubiquiti-community/go-unifi/unifi"
)

func resourceClientGroup() *schema.Resource {
	return &schema.Resource{
		Description: "`unifi_user_group` manages a user group (called \"client group\" in the UI), which can be used " +
			"to limit bandwidth for groups of users.",

		CreateContext: resourceClientGroupCreate,
		ReadContext:   resourceClientGroupRead,
		UpdateContext: resourceClientGroupUpdate,
		DeleteContext: resourceClientGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: importSiteAndID,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the user group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"site": {
				Description: "The name of the site to associate the user group with.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "The name of the user group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"qos_rate_max_down": {
				Description: "The QOS maximum download rate.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     -1,
				// TODO: validate does not equal 0,1
			},
			"qos_rate_max_up": {
				Description: "The QOS maximum upload rate.",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     -1,
				// TODO: validate does not equal 0,1
			},
		},
	}
}

func resourceClientGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	req, err := resourceClientGroupGetResourceData(d)
	if err != nil {
		return diag.FromErr(err)
	}

	site := d.Get("site").(string)
	if site == "" {
		site = c.site
	}

	resp, err := c.c.CreateClientGroup(context.TODO(), site, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.ID)

	return resourceClientGroupSetResourceData(resp, d)
}

func resourceClientGroupGetResourceData(d *schema.ResourceData) (*unifi.ClientGroup, error) {
	qos_down := int64(d.Get("qos_rate_max_down").(int))
	qos_up := int64(d.Get("qos_rate_max_up").(int))
	return &unifi.ClientGroup{
		Name:           d.Get("name").(string),
		QOSRateMaxDown: &qos_down,
		QOSRateMaxUp:   &qos_up,
	}, nil
}

func resourceClientGroupSetResourceData(resp *unifi.ClientGroup, d *schema.ResourceData) diag.Diagnostics {
	d.Set("name", resp.Name)
	d.Set("qos_rate_max_down", resp.QOSRateMaxDown)
	d.Set("qos_rate_max_up", resp.QOSRateMaxUp)

	return nil
}

func resourceClientGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	id := d.Id()

	site := d.Get("site").(string)
	if site == "" {
		site = c.site
	}

	resp, err := c.c.GetClientGroup(context.TODO(), site, id)
	if _, ok := err.(*unifi.NotFoundError); ok {
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceClientGroupSetResourceData(resp, d)
}

func resourceClientGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	req, err := resourceClientGroupGetResourceData(d)
	if err != nil {
		return diag.FromErr(err)
	}

	req.ID = d.Id()

	site := d.Get("site").(string)
	if site == "" {
		site = c.site
	}
	req.SiteID = site

	resp, err := c.c.UpdateClientGroup(context.TODO(), site, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceClientGroupSetResourceData(resp, d)
}

func resourceClientGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client)

	id := d.Id()

	site := d.Get("site").(string)
	if site == "" {
		site = c.site
	}
	err := c.c.DeleteClientGroup(context.TODO(), site, id)
	if _, ok := err.(*unifi.NotFoundError); ok {
		return nil
	}
	return diag.FromErr(err)
}
