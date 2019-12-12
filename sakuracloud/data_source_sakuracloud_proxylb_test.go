// Copyright 2016-2019 terraform-provider-sakuracloud authors
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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSakuraCloudDataSourceProxyLB_Basic(t *testing.T) {
	randString1 := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	randString2 := acctest.RandStringFromCharSet(20, acctest.CharSetAlpha)
	name := fmt.Sprintf("%s_%s", randString1, randString2)

	if ip, ok := os.LookupEnv(envProxyLBRealServerIP0); ok {
		proxyLBRealServerIP0 = ip
	} else {
		t.Skipf("ENV %q is requilred. skip", envProxyLBRealServerIP0)
		return
	}
	if ip, ok := os.LookupEnv(envProxyLBRealServerIP1); ok {
		proxyLBRealServerIP1 = ip
	} else {
		t.Skipf("ENV %q is requilred. skip", envProxyLBRealServerIP1)
		return
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		PreventPostDestroyRefresh: true,
		CheckDestroy:              testAccCheckSakuraCloudProxyLBDestroy,

		Steps: []resource.TestStep{
			{
				Config: testAccCheckSakuraCloudDataSourceProxyLBBase(name, proxyLBRealServerIP0, proxyLBRealServerIP1),
				Check:  testAccCheckSakuraCloudDataSourceExists("sakuracloud_proxylb.foobar"),
			},
			{
				Config: testAccCheckSakuraCloudDataSourceProxyLBConfig(name, proxyLBRealServerIP0, proxyLBRealServerIP1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraCloudDataSourceExists("data.sakuracloud_proxylb.foobar"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "name", name),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "plan", "100"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "description", "description_test"),

					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "health_check.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "health_check.0.delay_loop", "20"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "bind_ports.0.proxy_mode", "http"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "bind_ports.0.port", "80"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "servers.0.ipaddress", proxyLBRealServerIP0),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "servers.0.port", "80"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "servers.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "servers.1.ipaddress", proxyLBRealServerIP1),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "servers.1.port", "80"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "servers.1.enabled", "true"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "tags.#", "3"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "tags.0", "tag1"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "tags.1", "tag2"),
					resource.TestCheckResourceAttr("data.sakuracloud_proxylb.foobar", "tags.2", "tag3"),
				),
			},
			{
				Config: testAccCheckSakuraCloudDataSourceProxyLBConfig_With_Tag(name, proxyLBRealServerIP0, proxyLBRealServerIP1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraCloudDataSourceExists("data.sakuracloud_proxylb.foobar"),
				),
			},
			{
				Config: testAccCheckSakuraCloudDataSourceProxyLBConfig_NotExists(name, proxyLBRealServerIP0, proxyLBRealServerIP1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraCloudDataSourceNotExists("data.sakuracloud_proxylb.foobar"),
				),
				Destroy: true,
			},
			{
				Config: testAccCheckSakuraCloudDataSourceProxyLBConfig_With_NotExists_Tag(name, proxyLBRealServerIP0, proxyLBRealServerIP1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSakuraCloudDataSourceNotExists("data.sakuracloud_proxylb.foobar"),
				),
				Destroy: true,
			},
		},
	})
}

func testAccCheckSakuraCloudDataSourceProxyLBBase(name, ip1, ip2 string) string {
	return fmt.Sprintf(`
resource "sakuracloud_proxylb" "foobar" {
  name = "%s"
  health_check {
    protocol   = "tcp"
    delay_loop = 20
  }
  bind_ports {
    proxy_mode = "http"
    port       = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  description = "description_test"
  tags        = ["tag1", "tag2", "tag3"]
}`, name, ip1, ip2)
}

func testAccCheckSakuraCloudDataSourceProxyLBConfig(name, ip1, ip2 string) string {
	return fmt.Sprintf(`
resource "sakuracloud_proxylb" "foobar" {
  name = "%s"
  health_check {
    protocol   = "tcp"
    delay_loop = 20
  }
  bind_ports {
    proxy_mode = "http"
    port       = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  description = "description_test"
  tags        = ["tag1", "tag2", "tag3"]
}

data "sakuracloud_proxylb" "foobar" {
  filters {
	names = ["%s"]
  }
}`, name, ip1, ip2, name)
}

func testAccCheckSakuraCloudDataSourceProxyLBConfig_With_Tag(name, ip1, ip2 string) string {
	return fmt.Sprintf(`
resource "sakuracloud_proxylb" "foobar" {
  name = "%s"
  health_check {
    protocol   = "tcp"
    delay_loop = 20
  }
  bind_ports {
    proxy_mode = "http"
    port       = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  description = "description_test"
  tags        = ["tag1", "tag2", "tag3"]
}

data "sakuracloud_proxylb" "foobar" {
  filters {
	tags = ["tag1","tag3"]
  }
}`, name, ip1, ip2)
}

func testAccCheckSakuraCloudDataSourceProxyLBConfig_With_NotExists_Tag(name, ip1, ip2 string) string {
	return fmt.Sprintf(`
resource "sakuracloud_proxylb" "foobar" {
  name = "%s"
  health_check {
    protocol   = "tcp"
    delay_loop = 20
  }
  bind_ports {
    proxy_mode = "http"
    port       = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  description = "description_test"
  tags        = ["tag1", "tag2", "tag3"]
}

data "sakuracloud_proxylb" "foobar" {
  filters {
	tags = ["tag1-xxxxxxx","tag3-xxxxxxxx"]
  }
}`, name, ip1, ip2)
}

func testAccCheckSakuraCloudDataSourceProxyLBConfig_NotExists(name, ip1, ip2 string) string {
	return fmt.Sprintf(`
resource "sakuracloud_proxylb" "foobar" {
  name = "%s"
  health_check {
    protocol   = "tcp"
    delay_loop = 20
  }
  bind_ports {
    proxy_mode = "http"
    port       = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  servers {
    ipaddress = "%s"
    port      = 80
  }
  description = "description_test"
  tags        = ["tag1", "tag2", "tag3"]
}
data "sakuracloud_proxylb" "foobar" {
  filters {
	names = ["xxxxxxxxxxxxxxxxxx"]
  }
}`, name, ip1, ip2)
}
