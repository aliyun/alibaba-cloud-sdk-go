/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package endpoints

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var readXmlOnce sync.Once
var v = Endpoints{}

type LocalXmlResolver struct {
}

func (resolver *LocalXmlResolver) TryResolve(param *ResolveParam) (endpoint string, support bool, err error) {
	readXmlOnce.Do(func() {
		dir, _ := os.Getwd()

		file, err := os.Open(dir + "/endpoints/endpoints.xml")
		if err != nil {
			fmt.Printf("error: %v", err)
			support = false
			return
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("error: %v", err)
			support = false
			return
		}

		err = xml.Unmarshal(data, &v)
		if err != nil {
			fmt.Printf("error: %v", err)
			support = false
			return
		}
	})

	for _, xmlEndpoint := range v.EndpointList {
		for _, xmlRegionId := range xmlEndpoint.RegionIds.Id {
			if xmlRegionId == param.RegionId {
				for _, xmlProduct := range xmlEndpoint.Products.ProductList {
					if xmlProduct.ProductName == param.Product {
						endpoint = xmlProduct.DomainName
						support = true
						return
					}
				}
			}
		}
	}

	support = false
	return
}

type Endpoints struct {
	XMLName      xml.Name   `xml:"Endpoints"`
	EndpointList []Endpoint `xml:"Endpoint"`
}

type Endpoint struct {
	XMLName   xml.Name  `xml:"Endpoint"`
	Name      string    `xml:"name,attr"`
	RegionIds RegionIds `xml:"RegionIds"`
	Products  Products  `xml:"Products"`
}

type RegionIds struct {
	Id []string `xml:"RegionId"`
}

type Products struct {
	ProductList []Product `xml:"Product"`
}

type RegionId struct {
	XMLName string `xml:"RegionId"`
}

type Product struct {
	ProductName string `xml:"ProductName"`
	DomainName  string `xml:"DomainName"`
}
