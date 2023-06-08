# MeldAPI

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/PropertyMeld-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"MeldAPI"
	"MeldAPI/pkg/models/operations"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.Building.BuildingList(ctx, operations.BuildingListRequest{
        XPmOrg: 548814,
        Limit: meldapi.Int64(592845),
        Offset: meldapi.Int64(715190),
        Ordering: meldapi.String("quibusdam"),
    }, operations.BuildingListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedBuildingSerializerListList != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Building](docs/building/README.md)

* [BuildingList](docs/building/README.md#buildinglist)
* [BuildingRetrieve](docs/building/README.md#buildingretrieve)

### [CheckToken](docs/checktoken/README.md)

* [CheckTokenRetrieve](docs/checktoken/README.md#checktokenretrieve) - Used to check the validity of an oauth2 token

### [Estimates](docs/estimates/README.md)

* [EstimatesList](docs/estimates/README.md#estimateslist)
* [EstimatesRetrieve](docs/estimates/README.md#estimatesretrieve)

### [Expenditure](docs/expenditure/README.md)

* [ExpenditureList](docs/expenditure/README.md#expenditurelist)
* [ExpenditureRetrieve](docs/expenditure/README.md#expenditureretrieve)

### [Floor](docs/floor/README.md)

* [FloorList](docs/floor/README.md#floorlist)
* [FloorRetrieve](docs/floor/README.md#floorretrieve)

### [Invoice](docs/invoice/README.md)

* [InvoiceAttachmentRetrieve](docs/invoice/README.md#invoiceattachmentretrieve)
* [InvoiceList](docs/invoice/README.md#invoicelist)
* [InvoiceRetrieve](docs/invoice/README.md#invoiceretrieve)

### [ManagementAgent](docs/managementagent/README.md)

* [ManagementAgentList](docs/managementagent/README.md#managementagentlist)
* [ManagementAgentRetrieve](docs/managementagent/README.md#managementagentretrieve)

### [ManagerFile](docs/managerfile/README.md)

* [ManagerFileList](docs/managerfile/README.md#managerfilelist)

### [Meld](docs/meld/README.md)

* [MeldList](docs/meld/README.md#meldlist)
* [MeldManagerFilesList](docs/meld/README.md#meldmanagerfileslist)
* [MeldRetrieve](docs/meld/README.md#meldretrieve)
* [MeldTenantFilesList](docs/meld/README.md#meldtenantfileslist)
* [MeldVendorFilesList](docs/meld/README.md#meldvendorfileslist)

### [Owner](docs/owner/README.md)

* [OwnerCreate](docs/owner/README.md#ownercreate)
* [OwnerDestroy](docs/owner/README.md#ownerdestroy)
* [OwnerList](docs/owner/README.md#ownerlist)
* [OwnerPartialUpdate](docs/owner/README.md#ownerpartialupdate)
* [OwnerRetrieve](docs/owner/README.md#ownerretrieve)
* [OwnerUpdate](docs/owner/README.md#ownerupdate)

### [Ping](docs/ping/README.md)

* [PingRetrieve](docs/ping/README.md#pingretrieve) - Used to double check that the api is up an running.

### [Project](docs/project/README.md)

* [ProjectList](docs/project/README.md#projectlist)
* [ProjectRetrieve](docs/project/README.md#projectretrieve)

### [Property](docs/property/README.md)

* [PropertyCreate](docs/property/README.md#propertycreate)
* [PropertyDestroy](docs/property/README.md#propertydestroy)
* [PropertyList](docs/property/README.md#propertylist)
* [PropertyPartialUpdate](docs/property/README.md#propertypartialupdate)
* [PropertyRetrieve](docs/property/README.md#propertyretrieve)
* [PropertyUpdate](docs/property/README.md#propertyupdate)

### [PropertyGroup](docs/propertygroup/README.md)

* [PropertyGroupList](docs/propertygroup/README.md#propertygrouplist)
* [PropertyGroupRetrieve](docs/propertygroup/README.md#propertygroupretrieve)

### [Resident](docs/resident/README.md)

* [ResidentCreateForm](docs/resident/README.md#residentcreateform)
* [ResidentCreateJSON](docs/resident/README.md#residentcreatejson)
* [ResidentCreateMultipart](docs/resident/README.md#residentcreatemultipart)
* [ResidentDestroy](docs/resident/README.md#residentdestroy)
* [ResidentList](docs/resident/README.md#residentlist)
* [ResidentPartialUpdateForm](docs/resident/README.md#residentpartialupdateform)
* [ResidentPartialUpdateJSON](docs/resident/README.md#residentpartialupdatejson)
* [ResidentPartialUpdateMultipart](docs/resident/README.md#residentpartialupdatemultipart)
* [ResidentRetrieve](docs/resident/README.md#residentretrieve)
* [ResidentUpdateForm](docs/resident/README.md#residentupdateform)
* [ResidentUpdateJSON](docs/resident/README.md#residentupdatejson)
* [ResidentUpdateMultipart](docs/resident/README.md#residentupdatemultipart)

### [ResidentFile](docs/residentfile/README.md)

* [ResidentFileList](docs/residentfile/README.md#residentfilelist)

### [Schema](docs/schema/README.md)

* [SchemaRetrieve](docs/schema/README.md#schemaretrieve) - OpenApi3 schema for this API. Format can be selected via content negotiation.

- YAML: application/vnd.oai.openapi
- JSON: application/vnd.oai.openapi+json

### [Unit](docs/unit/README.md)

* [UnitCreate](docs/unit/README.md#unitcreate)
* [UnitDestroy](docs/unit/README.md#unitdestroy)
* [UnitList](docs/unit/README.md#unitlist)
* [UnitPartialUpdate](docs/unit/README.md#unitpartialupdate)
* [UnitRetrieve](docs/unit/README.md#unitretrieve)
* [UnitUpdate](docs/unit/README.md#unitupdate)

### [Vendor](docs/vendor/README.md)

* [VendorDestroy](docs/vendor/README.md#vendordestroy)
* [VendorList](docs/vendor/README.md#vendorlist)
* [VendorRetrieve](docs/vendor/README.md#vendorretrieve)

### [VendorFile](docs/vendorfile/README.md)

* [VendorFileList](docs/vendorfile/README.md#vendorfilelist)

### [VendorInvite](docs/vendorinvite/README.md)

* [VendorInviteCreate](docs/vendorinvite/README.md#vendorinvitecreate)

### [WorkLog](docs/worklog/README.md)

* [WorkLogList](docs/worklog/README.md#workloglist)
* [WorkLogRetrieve](docs/worklog/README.md#worklogretrieve)
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
