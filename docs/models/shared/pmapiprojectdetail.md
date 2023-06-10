# PmAPIProjectDetail


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Coordinators`                                              | [][ManagementAgent](../../models/shared/managementagent.md) | :heavy_check_mark:                                          | N/A                                                         |
| `Created`                                                   | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |
| `Description`                                               | **string*                                                   | :heavy_minus_sign:                                          | N/A                                                         |
| `DueDate`                                                   | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |
| `ID`                                                        | *int64*                                                     | :heavy_check_mark:                                          | N/A                                                         |
| `Melds`                                                     | []*int64*                                                   | :heavy_check_mark:                                          | N/A                                                         |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | N/A                                                         |
| `ProjectType`                                               | [*ProjectTypeEnum](../../models/shared/projecttypeenum.md)  | :heavy_minus_sign:                                          | N/A                                                         |
| `StartDate`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |
| `Unit`                                                      | [*PmAPIUnit](../../models/shared/pmapiunit.md)              | :heavy_minus_sign:                                          | N/A                                                         |
| `Updated`                                                   | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | N/A                                                         |