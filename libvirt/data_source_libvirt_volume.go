package libvirt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLibvirtVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLibvirtVolumeRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pool": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLibvirtVolumeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)

	poolName := d.Get("pool").(string)
	volName := d.Get("name").(string)

	pool, err := client.libvirt.StoragePoolLookupByName(poolName)
	if err != nil {
		return err
	}

	vol, err := client.libvirt.StorageVolLookupByName(pool, volName)
	if err != nil {
		return err
	}

	d.SetId(vol.Key)

	return nil
}
