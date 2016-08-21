package sakuracloud

import (
	"fmt"

	"bytes"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yamamoto-febc/libsacloud/api"
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

func resourceSakuraCloudVPCRouterDHCPServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSakuraCloudVPCRouterDHCPServerCreate,
		Read:   resourceSakuraCloudVPCRouterDHCPServerRead,
		Delete: resourceSakuraCloudVPCRouterDHCPServerDelete,
		Schema: map[string]*schema.Schema{
			"vpc_router_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_router_interface_index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"range_start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"range_stop": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  "target SakuraCloud zone",
				ValidateFunc: validateStringInWord([]string{"is1a", "is1b", "tk1a", "tk1v"}),
			},
		},
	}
}

func resourceSakuraCloudVPCRouterDHCPServerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	zone, ok := d.GetOk("zone")
	if ok {
		client.Zone = zone.(string)
	}

	routerID := d.Get("vpc_router_id").(string)
	sakuraMutexKV.Lock(routerID)
	defer sakuraMutexKV.Unlock(routerID)

	vpcRouter, err := client.VPCRouter.Read(routerID)
	if err != nil {
		return fmt.Errorf("Couldn't find SakuraCloud VPCRouter resource: %s", err)
	}

	dhcpServer := expandVPCRouterDHCPServer(d)
	if vpcRouter.Settings == nil {
		vpcRouter.InitVPCRouterSetting()
	}

	vpcRouter.Settings.Router.AddDHCPServer(d.Get("vpc_router_interface_index").(int), dhcpServer.RangeStart, dhcpServer.RangeStop)
	vpcRouter, err = client.VPCRouter.UpdateSetting(routerID, vpcRouter)
	if err != nil {
		return fmt.Errorf("Failed to enable SakuraCloud VPCRouterDHCPServer resource: %s", err)
	}
	_, err = client.VPCRouter.Config(routerID)
	if err != nil {
		return fmt.Errorf("Couldn'd apply SakuraCloud VPCRouter config: %s", err)
	}

	d.SetId(vpcRouterDHCPServerIDHash(routerID, dhcpServer))
	return resourceSakuraCloudVPCRouterDHCPServerRead(d, meta)
}

func resourceSakuraCloudVPCRouterDHCPServerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)
	zone, ok := d.GetOk("zone")
	if ok {
		client.Zone = zone.(string)
	}

	routerID := d.Get("vpc_router_id").(string)
	vpcRouter, err := client.VPCRouter.Read(routerID)
	if err != nil {
		return fmt.Errorf("Couldn't find SakuraCloud VPCRouter resource: %s", err)
	}

	dhcpServer := expandVPCRouterDHCPServer(d)
	if vpcRouter.Settings != nil && vpcRouter.Settings.Router != nil && vpcRouter.Settings.Router.DHCPServer != nil &&
		vpcRouter.Settings.Router.FindDHCPServer(d.Get("vpc_router_interface_index").(int), dhcpServer.RangeStart, dhcpServer.RangeStop) != nil {
		d.Set("range_start", dhcpServer.RangeStart)
		d.Set("range_stop", dhcpServer.RangeStop)
	}

	d.Set("zone", client.Zone)

	return nil
}

func resourceSakuraCloudVPCRouterDHCPServerDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*api.Client)
	zone, ok := d.GetOk("zone")
	if ok {
		client.Zone = zone.(string)
	}

	routerID := d.Get("vpc_router_id").(string)
	sakuraMutexKV.Lock(routerID)
	defer sakuraMutexKV.Unlock(routerID)

	vpcRouter, err := client.VPCRouter.Read(routerID)
	if err != nil {
		return fmt.Errorf("Couldn't find SakuraCloud VPCRouter resource: %s", err)
	}

	if vpcRouter.Settings.Router.DHCPServer != nil {

		dhcpServer := expandVPCRouterDHCPServer(d)
		vpcRouter.Settings.Router.RemoveDHCPServer(d.Get("vpc_router_interface_index").(int), dhcpServer.RangeStart, dhcpServer.RangeStop)

		vpcRouter, err = client.VPCRouter.UpdateSetting(routerID, vpcRouter)
		if err != nil {
			return fmt.Errorf("Failed to delete SakuraCloud VPCRouterDHCPServer resource: %s", err)
		}

		_, err = client.VPCRouter.Config(routerID)
		if err != nil {
			return fmt.Errorf("Couldn'd apply SakuraCloud VPCRouter config: %s", err)
		}
	}

	d.SetId("")
	return nil
}

func vpcRouterDHCPServerIDHash(routerID string, s *sacloud.VPCRouterDHCPServerConfig) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s-", routerID))
	buf.WriteString(fmt.Sprintf("%s-", s.Interface))
	buf.WriteString(fmt.Sprintf("%s-", s.RangeStart))
	buf.WriteString(fmt.Sprintf("%s", s.RangeStop))

	return fmt.Sprintf("%d", hashcode.String(buf.String()))
}

func expandVPCRouterDHCPServer(d *schema.ResourceData) *sacloud.VPCRouterDHCPServerConfig {

	var dhcpServer = &sacloud.VPCRouterDHCPServerConfig{
		Interface:  fmt.Sprintf("eth%d", d.Get("vpc_router_interface_index").(int)),
		RangeStart: d.Get("range_start").(string),
		RangeStop:  d.Get("range_stop").(string),
	}

	return dhcpServer
}