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


### [Building](docs/sdks/building/README.md)

* [BuildingList](docs/sdks/building/README.md#buildinglist)
* [BuildingRetrieve](docs/sdks/building/README.md#buildingretrieve)

### [CheckToken](docs/sdks/checktoken/README.md)

* [CheckTokenRetrieve](docs/sdks/checktoken/README.md#checktokenretrieve) - Used to check the validity of an oauth2 token

### [Estimates](docs/sdks/estimates/README.md)

* [EstimatesList](docs/sdks/estimates/README.md#estimateslist)
* [EstimatesRetrieve](docs/sdks/estimates/README.md#estimatesretrieve)

### [Expenditure](docs/sdks/expenditure/README.md)

* [ExpenditureList](docs/sdks/expenditure/README.md#expenditurelist)
* [ExpenditureRetrieve](docs/sdks/expenditure/README.md#expenditureretrieve)

### [Floor](docs/sdks/floor/README.md)

* [FloorList](docs/sdks/floor/README.md#floorlist)
* [FloorRetrieve](docs/sdks/floor/README.md#floorretrieve)

### [Invoice](docs/sdks/invoice/README.md)

* [InvoiceAttachmentRetrieve](docs/sdks/invoice/README.md#invoiceattachmentretrieve)
* [InvoiceList](docs/sdks/invoice/README.md#invoicelist)
* [InvoiceRetrieve](docs/sdks/invoice/README.md#invoiceretrieve)

### [ManagementAgent](docs/sdks/managementagent/README.md)

* [ManagementAgentList](docs/sdks/managementagent/README.md#managementagentlist)
* [ManagementAgentRetrieve](docs/sdks/managementagent/README.md#managementagentretrieve)

### [ManagerFile](docs/sdks/managerfile/README.md)

* [ManagerFileList](docs/sdks/managerfile/README.md#managerfilelist)

### [Meld](docs/sdks/meld/README.md)

* [MeldList](docs/sdks/meld/README.md#meldlist)
* [MeldManagerFilesList](docs/sdks/meld/README.md#meldmanagerfileslist)
* [MeldRetrieve](docs/sdks/meld/README.md#meldretrieve)
* [MeldTenantFilesList](docs/sdks/meld/README.md#meldtenantfileslist)
* [MeldVendorFilesList](docs/sdks/meld/README.md#meldvendorfileslist)

### [Owner](docs/sdks/owner/README.md)

* [OwnerCreate](docs/sdks/owner/README.md#ownercreate)
* [OwnerDestroy](docs/sdks/owner/README.md#ownerdestroy)
* [OwnerList](docs/sdks/owner/README.md#ownerlist)
* [OwnerPartialUpdate](docs/sdks/owner/README.md#ownerpartialupdate)
* [OwnerRetrieve](docs/sdks/owner/README.md#ownerretrieve)
* [OwnerUpdate](docs/sdks/owner/README.md#ownerupdate)

### [Ping](docs/sdks/ping/README.md)

* [PingRetrieve](docs/sdks/ping/README.md#pingretrieve) - Used to double check that the api is up an running.

### [Project](docs/sdks/project/README.md)

* [ProjectList](docs/sdks/project/README.md#projectlist)
* [ProjectRetrieve](docs/sdks/project/README.md#projectretrieve)

### [Property](docs/sdks/property/README.md)

* [PropertyCreate](docs/sdks/property/README.md#propertycreate)
* [PropertyDestroy](docs/sdks/property/README.md#propertydestroy)
* [PropertyList](docs/sdks/property/README.md#propertylist)
* [PropertyPartialUpdate](docs/sdks/property/README.md#propertypartialupdate)
* [PropertyRetrieve](docs/sdks/property/README.md#propertyretrieve)
* [PropertyUpdate](docs/sdks/property/README.md#propertyupdate)

### [PropertyGroup](docs/sdks/propertygroup/README.md)

* [PropertyGroupList](docs/sdks/propertygroup/README.md#propertygrouplist)
* [PropertyGroupRetrieve](docs/sdks/propertygroup/README.md#propertygroupretrieve)

### [Resident](docs/sdks/resident/README.md)

* [ResidentCreateForm](docs/sdks/resident/README.md#residentcreateform)
* [ResidentCreateJSON](docs/sdks/resident/README.md#residentcreatejson)
* [ResidentCreateMultipart](docs/sdks/resident/README.md#residentcreatemultipart)
* [ResidentDestroy](docs/sdks/resident/README.md#residentdestroy)
* [ResidentList](docs/sdks/resident/README.md#residentlist)
* [ResidentPartialUpdateForm](docs/sdks/resident/README.md#residentpartialupdateform)
* [ResidentPartialUpdateJSON](docs/sdks/resident/README.md#residentpartialupdatejson)
* [ResidentPartialUpdateMultipart](docs/sdks/resident/README.md#residentpartialupdatemultipart)
* [ResidentRetrieve](docs/sdks/resident/README.md#residentretrieve)
* [ResidentUpdateForm](docs/sdks/resident/README.md#residentupdateform)
* [ResidentUpdateJSON](docs/sdks/resident/README.md#residentupdatejson)
* [ResidentUpdateMultipart](docs/sdks/resident/README.md#residentupdatemultipart)

### [ResidentFile](docs/sdks/residentfile/README.md)

* [ResidentFileList](docs/sdks/residentfile/README.md#residentfilelist)

### [Schema](docs/sdks/schema/README.md)

* [SchemaRetrieve](docs/sdks/schema/README.md#schemaretrieve) - OpenApi3 schema for this API. Format can be selected via content negotiation.

- YAML: application/vnd.oai.openapi
- JSON: application/vnd.oai.openapi+json

### [Unit](docs/sdks/unit/README.md)

* [UnitCreate](docs/sdks/unit/README.md#unitcreate)
* [UnitDestroy](docs/sdks/unit/README.md#unitdestroy)
* [UnitList](docs/sdks/unit/README.md#unitlist)
* [UnitPartialUpdate](docs/sdks/unit/README.md#unitpartialupdate)
* [UnitRetrieve](docs/sdks/unit/README.md#unitretrieve)
* [UnitUpdate](docs/sdks/unit/README.md#unitupdate)

### [Vendor](docs/sdks/vendor/README.md)

* [VendorDestroy](docs/sdks/vendor/README.md#vendordestroy)
* [VendorList](docs/sdks/vendor/README.md#vendorlist)
* [VendorRetrieve](docs/sdks/vendor/README.md#vendorretrieve)

### [VendorFile](docs/sdks/vendorfile/README.md)

* [VendorFileList](docs/sdks/vendorfile/README.md#vendorfilelist)

### [VendorInvite](docs/sdks/vendorinvite/README.md)

* [VendorInviteCreate](docs/sdks/vendorinvite/README.md#vendorinvitecreate)

### [WorkLog](docs/sdks/worklog/README.md)

* [WorkLogList](docs/sdks/worklog/README.md#workloglist)
* [WorkLogRetrieve](docs/sdks/worklog/README.md#worklogretrieve)
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
