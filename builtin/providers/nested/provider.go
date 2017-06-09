package nested

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"nested": nestedSchema(),
		},
	}
}

func nestedSchema() *schema.Resource {
	return &schema.Resource{
		Create: nestedRead,
		Read:   nestedRead,
		Update: nestedRead,
		Delete: nestedRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"foo": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bar": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"baz": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func nestedRead(d *schema.ResourceData, meta interface{}) error {
	start := time.Now().UnixNano()
	_ = d.Get("foo")
	end := time.Now().UnixNano()

	delta := end - start

	return fmt.Errorf("Time to Get: %d\n\n", delta)
}
