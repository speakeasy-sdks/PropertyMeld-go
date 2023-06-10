# BasePayment


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `Amount`                                        | *string*                                        | :heavy_check_mark:                              | N/A                                             |
| `AmountRefunded`                                | **string*                                       | :heavy_minus_sign:                              | N/A                                             |
| `CheckNo`                                       | **string*                                       | :heavy_minus_sign:                              | N/A                                             |
| `CreateBy`                                      | map[string]*interface{}*                        | :heavy_minus_sign:                              | N/A                                             |
| `Created`                                       | [time.Time](https://pkg.go.dev/time#Time)       | :heavy_check_mark:                              | N/A                                             |
| `ID`                                            | *int64*                                         | :heavy_check_mark:                              | N/A                                             |
| `Method`                                        | [MethodEnum](../../models/shared/methodenum.md) | :heavy_check_mark:                              | N/A                                             |
| `Notes`                                         | **string*                                       | :heavy_minus_sign:                              | N/A                                             |
| `StripeChargeID`                                | **string*                                       | :heavy_minus_sign:                              | N/A                                             |
| `UpdateBy`                                      | map[string]*interface{}*                        | :heavy_minus_sign:                              | N/A                                             |
| `Updated`                                       | [time.Time](https://pkg.go.dev/time#Time)       | :heavy_check_mark:                              | N/A                                             |