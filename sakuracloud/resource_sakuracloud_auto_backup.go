// Copyright 2016-2020 terraform-provider-sakuracloud authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sakuracloud

import (
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func resourceSakuraCloudAutoBackup() *schema.Resource {
	resourceName := "AutoBackup"
	return &schema.Resource{
		Create: resourceSakuraCloudAutoBackupCreate,
		Read:   resourceSakuraCloudAutoBackupRead,
		Update: resourceSakuraCloudAutoBackupUpdate,
		Delete: resourceSakuraCloudAutoBackupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": schemaResourceName(resourceName),
			"disk_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSakuracloudIDType,
				Description:  "The disk id to backed up",
			},
			"weekdays": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Description: descf(
					"A list of weekdays to backed up. The values in the list must be in [%s]",
					types.BackupWeekdayStrings,
				),
			},
			"max_backup_num": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 10),
				Description:  descf("The number backup files to keep. %s", descRange(1, 10)),
			},
			"icon_id":     schemaResourceIconID(resourceName),
			"description": schemaResourceDescription(resourceName),
			"tags":        schemaResourceTags(resourceName),
			"zone":        schemaResourceZone(resourceName),
		},
	}
}

func resourceSakuraCloudAutoBackupCreate(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutCreate)
	defer cancel()

	autoBackupOp := sacloud.NewAutoBackupOp(client)

	if err := validateBackupWeekdays(d, "weekdays"); err != nil {
		return err
	}

	autoBackup, err := autoBackupOp.Create(ctx, zone, expandAutoBackupCreateRequest(d))
	if err != nil {
		return fmt.Errorf("creating SakuraCloud AutoBackup is failed: %s", err)
	}

	d.SetId(autoBackup.ID.String())
	return resourceSakuraCloudAutoBackupRead(d, meta)
}

func resourceSakuraCloudAutoBackupRead(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutRead)
	defer cancel()

	autoBackupOp := sacloud.NewAutoBackupOp(client)

	autoBackup, err := autoBackupOp.Read(ctx, zone, sakuraCloudID(d.Id()))
	if err != nil {
		if sacloud.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("could not find SakuraCloud AutoBackup[%s]: %s", d.Id(), err)
	}
	return setAutoBackupResourceData(d, client, autoBackup)
}

func resourceSakuraCloudAutoBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutUpdate)
	defer cancel()

	autoBackupOp := sacloud.NewAutoBackupOp(client)

	autoBackup, err := autoBackupOp.Read(ctx, zone, sakuraCloudID(d.Id()))
	if err != nil {
		return fmt.Errorf("could not read SakuraCloud AutoBackup[%s]: %s", d.Id(), err)
	}

	if err := validateBackupWeekdays(d, "weekdays"); err != nil {
		return err
	}

	if _, err = autoBackupOp.Update(ctx, zone, autoBackup.ID, expandAutoBackupUpdateRequest(d, autoBackup)); err != nil {
		return fmt.Errorf("updating SakuraCloud AutoBackup[%s] is failed: %s", d.Id(), err)
	}

	return resourceSakuraCloudAutoBackupRead(d, meta)
}

func resourceSakuraCloudAutoBackupDelete(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutDelete)
	defer cancel()

	autoBackupOp := sacloud.NewAutoBackupOp(client)

	autoBackup, err := autoBackupOp.Read(ctx, zone, sakuraCloudID(d.Id()))
	if err != nil {
		if sacloud.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("could not read SakuraCloud AutoBackup[%s]: %s", d.Id(), err)
	}

	if err := autoBackupOp.Delete(ctx, zone, autoBackup.ID); err != nil {
		return fmt.Errorf("deleting SakuraCloud AutoBackup[%s] is failed: %s", d.Id(), err)
	}

	d.SetId("")
	return nil
}

func setAutoBackupResourceData(d *schema.ResourceData, client *APIClient, data *sacloud.AutoBackup) error {
	d.Set("name", data.Name)                              // nolint
	d.Set("disk_id", data.DiskID.String())                // nolint
	d.Set("max_backup_num", data.MaximumNumberOfArchives) // nolint
	d.Set("icon_id", data.IconID.String())                // nolint
	d.Set("description", data.Description)                // nolint
	d.Set("zone", getZone(d, client))                     // nolint
	if err := d.Set("weekdays", flattenBackupWeekdays(data.BackupSpanWeekdays)); err != nil {
		return err
	}
	return d.Set("tags", flattenTags(data.Tags))
}
