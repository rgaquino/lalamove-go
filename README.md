# lalamove-go

Unofficial Go Client SDK for the [Lalamove APIs](https://developers.lalamove.com/#introduction)

## Installation

To install, please execute the following `go get` command.

```
go get github.com/rgaquino/lalamove-go
```

## Sample Usage

```go
package main

import (
	"context"
    "fmt"
	"log"

    "github.com/rgaquino/lalamove-go"
)

func main() {
    c, err := lalamove.NewClient(
        lalamove.WithBaseURL("https://sandbox-rest.lalamove.com"),
        lalamove.WithAPIKey("API_KEY"),
        lalamove.WithSecret("SECRET_KEY"),
    )
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }
    
    req := &lalamove.GetQuotationRequest{
        ServiceType: lalamove.ServiceTypeMotorcycle,
        RequesterContact: lalamove.Contact{
            Name:  "Peter Pan",
            Phone: "232",
        },
        Stops: []lalamove.Waypoint{
            {
                Location: lalamove.Location{
                    Lat: "-6.255431000000001",
                    Lng: "106.60114290000001",
                },
                Addresses: lalamove.AddressTranslations{
                    lalamove.LocaleIndonesiaEN: {
                        DisplayString: "Jl. Perum Dasana Indah No.SD 3/ 17-18, RT.3/RW.1, " +
                            "Bojong Nangka, Klp. Dua, Tangerang, Banten 15810, Indonesia",
                        Country:       lalamove.CityCodeIndonesiaJakarata.GetLLMCountry(),
                    },
                },
            },
            {
                Location: lalamove.Location{
                    Lat: "-6.404722800000001",
                    Lng: "106.81902130000003",
                },
                Addresses: lalamove.AddressTranslations{
                    lalamove.LocaleIndonesiaEN: {
                        DisplayString: "Jl. Kartini, Ruko No. 1E, Depok, Pancoran MAS, " +
                            "Kota Depok, Jawa Barat 16431, Indonesia",
                        Country:       lalamove.CityCodeIndonesiaJakarata.GetLLMCountry(),
                    },
                },
            },
        },
        Deliveries: []lalamove.DeliveryInfo{
            {
                ToStop: 1,
                Contact: lalamove.Contact{
                    Name:  "mm",
                    Phone: "9999999",
                },
            },
        },
    }

    resp, err := c.GetQuotation(context.Background(), lalamove.CityCodeSingaporeSingapore, req)
    if err != nil {
        log.Fatalf("fatal error: %s", err)
    }
    fmt.Println(resp)
}
```