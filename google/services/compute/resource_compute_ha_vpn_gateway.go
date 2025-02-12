// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceComputeHaVpnGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeHaVpnGatewayCreate,
		Read:   resourceComputeHaVpnGatewayRead,
		Update: resourceComputeHaVpnGatewayUpdate,
		Delete: resourceComputeHaVpnGatewayDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeHaVpnGatewayImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		SchemaVersion: 1,

		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceComputeHaVpnGatewayResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: ResourceComputeHaVpnGatewayUpgradeV0,
				Version: 0,
			},
		},
		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The network this VPN gateway is accepting traffic for.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"gateway_ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"IPV4", "IPV6", ""}),
				Description:  `The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value: "IPV4" Possible values: ["IPV4", "IPV6"]`,
				Default:      "IPV4",
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels for this resource. These can only be added or modified by the setLabels method.
Each label key/value pair must comply with RFC1035. Label values may be empty.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The region this gateway should sit in.`,
			},
			"stack_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY", ""}),
				Description: `The stack type for this VPN gateway to identify the IP protocols that are enabled.
If not specified, IPV4_ONLY will be used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]`,
				Default: "IPV4_ONLY",
			},
			"vpn_interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `A list of interfaces on this VPN gateway.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: `The numeric ID of this VPN gateway interface.`,
						},
						"interconnect_attachment": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description: `URL of the interconnect attachment resource. When the value
of this field is present, the VPN Gateway will be used for
IPsec-encrypted Cloud Interconnect; all Egress or Ingress
traffic for this VPN Gateway interface will go through the
specified interconnect attachment resource.

Not currently available publicly.`,
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The external IP address for this VPN gateway interface.`,
						},
					},
				},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `A fingerprint for the labels being applied to this VpnGateway, which is essentially a hash
of the labels set used for optimistic locking. The fingerprint is initially generated by
Compute Engine and changes after every request to modify or update labels.
You must always provide an up-to-date fingerprint hash in order to update or change labels,
otherwise the request will fail with error 412 conditionNotMet.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeHaVpnGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeHaVpnGatewayDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeHaVpnGatewayName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeHaVpnGatewayNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	stackTypeProp, err := expandComputeHaVpnGatewayStackType(d.Get("stack_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("stack_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(stackTypeProp)) && (ok || !reflect.DeepEqual(v, stackTypeProp)) {
		obj["stackType"] = stackTypeProp
	}
	gatewayIpVersionProp, err := expandComputeHaVpnGatewayGatewayIpVersion(d.Get("gateway_ip_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gateway_ip_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(gatewayIpVersionProp)) && (ok || !reflect.DeepEqual(v, gatewayIpVersionProp)) {
		obj["gatewayIpVersion"] = gatewayIpVersionProp
	}
	vpnInterfacesProp, err := expandComputeHaVpnGatewayVpnInterfaces(d.Get("vpn_interfaces"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vpn_interfaces"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpnInterfacesProp)) && (ok || !reflect.DeepEqual(v, vpnInterfacesProp)) {
		obj["vpnInterfaces"] = vpnInterfacesProp
	}
	labelFingerprintProp, err := expandComputeHaVpnGatewayLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	labelsProp, err := expandComputeHaVpnGatewayEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	regionProp, err := expandComputeHaVpnGatewayRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnGateways")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new HaVpnGateway: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HaVpnGateway: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating HaVpnGateway: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating HaVpnGateway", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create HaVpnGateway: %s", err)
	}

	log.Printf("[DEBUG] Finished creating HaVpnGateway %q: %#v", d.Id(), res)

	return resourceComputeHaVpnGatewayRead(d, meta)
}

func resourceComputeHaVpnGatewayRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HaVpnGateway: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputeHaVpnGateway %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}

	if err := d.Set("description", flattenComputeHaVpnGatewayDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("name", flattenComputeHaVpnGatewayName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("network", flattenComputeHaVpnGatewayNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("stack_type", flattenComputeHaVpnGatewayStackType(res["stackType"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("gateway_ip_version", flattenComputeHaVpnGatewayGatewayIpVersion(res["gatewayIpVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("vpn_interfaces", flattenComputeHaVpnGatewayVpnInterfaces(res["vpnInterfaces"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("labels", flattenComputeHaVpnGatewayLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeHaVpnGatewayLabelFingerprint(res["labelFingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("terraform_labels", flattenComputeHaVpnGatewayTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("effective_labels", flattenComputeHaVpnGatewayEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("region", flattenComputeHaVpnGatewayRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading HaVpnGateway: %s", err)
	}

	return nil
}

func resourceComputeHaVpnGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HaVpnGateway: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("label_fingerprint") || d.HasChange("effective_labels") {
		obj := make(map[string]interface{})

		labelFingerprintProp, err := expandComputeHaVpnGatewayLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}
		labelsProp, err := expandComputeHaVpnGatewayEffectiveLabels(d.Get("effective_labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}/setLabels")
		if err != nil {
			return err
		}

		headers := make(http.Header)

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "POST",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})
		if err != nil {
			return fmt.Errorf("Error updating HaVpnGateway %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating HaVpnGateway %q: %#v", d.Id(), res)
		}

		err = ComputeOperationWaitTime(
			config, res, project, "Updating HaVpnGateway", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeHaVpnGatewayRead(d, meta)
}

func resourceComputeHaVpnGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for HaVpnGateway: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting HaVpnGateway %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "HaVpnGateway")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting HaVpnGateway", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting HaVpnGateway %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeHaVpnGatewayImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/vpnGateways/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/vpnGateways/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeHaVpnGatewayDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHaVpnGatewayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHaVpnGatewayNetwork(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeHaVpnGatewayStackType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return "IPV4_ONLY"
	}

	return v
}

func flattenComputeHaVpnGatewayGatewayIpVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil || tpgresource.IsEmptyValue(reflect.ValueOf(v)) {
		return "IPV4"
	}

	return v
}

func flattenComputeHaVpnGatewayVpnInterfaces(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"id":                      flattenComputeHaVpnGatewayVpnInterfacesId(original["id"], d, config),
			"ip_address":              flattenComputeHaVpnGatewayVpnInterfacesIpAddress(original["ipAddress"], d, config),
			"interconnect_attachment": flattenComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(original["interconnectAttachment"], d, config),
		})
	}
	return transformed
}
func flattenComputeHaVpnGatewayVpnInterfacesId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeHaVpnGatewayVpnInterfacesIpAddress(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertSelfLinkToV1(v.(string))
}

func flattenComputeHaVpnGatewayLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeHaVpnGatewayLabelFingerprint(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHaVpnGatewayTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenComputeHaVpnGatewayEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeHaVpnGatewayRegion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func expandComputeHaVpnGatewayDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeHaVpnGatewayStackType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayGatewayIpVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfaces(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedId, err := expandComputeHaVpnGatewayVpnInterfacesId(original["id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["id"] = transformedId
		}

		transformedIpAddress, err := expandComputeHaVpnGatewayVpnInterfacesIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		transformedInterconnectAttachment, err := expandComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(original["interconnect_attachment"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInterconnectAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["interconnectAttachment"] = transformedInterconnectAttachment
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeHaVpnGatewayVpnInterfacesId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfacesIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayVpnInterfacesInterconnectAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("interconnectAttachments", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for interconnect_attachment: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeHaVpnGatewayLabelFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeHaVpnGatewayEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeHaVpnGatewayRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}

func resourceComputeHaVpnGatewayResourceV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateGCEName,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035.  Specifically, the name must be 1-63 characters long and
match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
the first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The network this VPN gateway is accepting traffic for.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"gateway_ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"IPV4", "IPV6", ""}),
				Description:  `The IP family of the gateway IPs for the HA-VPN gateway interfaces. If not specified, IPV4 will be used. Default value: "IPV4" Possible values: ["IPV4", "IPV6"]`,
				Default:      "IPV4",
			},
			"region": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The region this gateway should sit in.`,
			},
			"stack_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY", ""}),
				Description: `The stack type for this VPN gateway to identify the IP protocols that are enabled.
If not specified, IPV4_ONLY will be used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]`,
				Default: "IPV4_ONLY",
			},
			"vpn_interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `A list of interfaces on this VPN gateway.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: `The numeric ID of this VPN gateway interface.`,
						},
						"interconnect_attachment": {
							Type:             schema.TypeString,
							Optional:         true,
							ForceNew:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description: `URL of the interconnect attachment resource. When the value
of this field is present, the VPN Gateway will be used for
IPsec-encrypted Cloud Interconnect; all Egress or Ingress
traffic for this VPN Gateway interface will go through the
specified interconnect attachment resource.

Not currently available publicly.`,
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The external IP address for this VPN gateway interface.`,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func ResourceComputeHaVpnGatewayUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("[DEBUG] Attributes before migration: %#v", rawState)

	// Check if "gateway_ip_version" already exists
	if _, ok := rawState["gateway_ip_version"]; !ok {
		// Add the missing attribute with the default value
		rawState["gateway_ip_version"] = "IPV4"
	} else {
		log.Printf("[DEBUG] 'gateway_ip_version' already exists: %#v", rawState["gateway_ip_version"])
	}

	log.Printf("[DEBUG] Attributes after migration: %#v", rawState)
	return rawState, nil
}
