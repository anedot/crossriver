# crossriver

> Go interface to the [Cross River Bank API](https://docs.crbcos.com/cos-api/docs/welcome-to-cos)

## Usage
```go
import (
    "context"
    "os"

    "github.com/anedot/crossriver"
)

func main() {
	client, _ := crossriver.NewClient(
		os.Getenv("CROSSRIVER_CLIENT_ID"),
		os.Getenv("CROSSRIVER_CLIENT_SECRET"),
		os.Getenv("CROSSRIVER_API_BASE"),
		os.Getenv("CROSSRIVER_PARTNER_ID"),
	)

	client.SetLog(os.Stdout)

	ctx := context.Background()
	client.GetToken(ctx)
}
```

## ACH
### Get ACH Payment
```go
client.GetAchPayment(ctx, paymentId)
```

### Create ACH Payment
```go
payment := crossriver.Payment{
    AccountNumber: "2714035231",
    Receiver: crossriver.BankAccount{
        RoutingNumber:  "021000021",
        AccountNumber:  "456789000",
        AccountType:    "Checking",
        Name:           "Bob Smith",
        Identification: "XYZ123",
    },
    SecCode:         "WEB",
    Description:     "Payment",
    TransactionType: "Push",
    Amount:          10000,
    ServiceType:     "SameDay",
}

client.CreateAchPayment(ctx, payment)
```

### Get ACH Batch
```go
client.GetAchBatch(ctx, batchId)
```

### Create ACH Batch
```go
payments := []Payment{
    {
        AccountNumber: "1234567890",
        Receiver: Receiver{
            RoutingNumber:  "021000021",
            AccountNumber:  "456789000",
            AccountType:    "Checking",
            Name:           "Bob Smith",
            Identification: "XYZ123",
        },
        SecCode:         "PPD",
        Description:     "Payment",
        TransactionType: "Push",
        Amount:          10000,
        ServiceType:     "Standard",
    },
    {
        AccountNumber: "1234567890",
        Receiver: Receiver{
            RoutingNumber:  "021000021",
            AccountNumber:  "123787777",
            AccountType:    "Checking",
            Name:           "Alice Smith",
            Identification: "ABC456",
        },
        SecCode:         "PPD",
        Description:     "Payment",
        TransactionType: "Push",
        Amount:          20000,
        ServiceType:     "Standard",
    },
}

client.CreateAchBatch(ctx, payments)
```

## Webhooks
### Create Webhook
```go
webhook := crossriver.Webhook{
    EventName:    "string",
    CallbackUrl:  "string",
    AuthUsername: "string",
    AuthPassword: "string",
    Type:         "Push",
    Format:       "Basic",
}

client.CreateWebhook(ctx, webhook)
```
