package testdata

// GetOrganizationResponse example
const GetOrganizationResponse = `{
    "resource": "organization",
    "id": "org_12345678",
    "name": "Mollie B.V.",
    "email": "info@mollie.com",
    "address": {
        "streetAndNumber": "Keizersgracht 313",
        "postalCode": "1016 EE",
        "city": "Amsterdam",
        "country": "NL"
    },
    "registrationNumber": "30204462",
    "vatNumber": "NL815839091B01",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/organizations/org_12345678",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/organizations-api/get-organization",
            "type": "text/html"
        }
    }
}`

// GetCurrentOrganizationResponse example
const GetCurrentOrganizationResponse = `{
     "resource": "organization",
     "id": "org_12345678",
     "name": "Mollie B.V.",
     "email": "info@mollie.com",
     "address": {
        "streetAndNumber" : "Keizersgracht 313",
        "postalCode": "1016 EE",
         "city": "Amsterdam",
         "country": "NL"
     },
     "registrationNumber": "30204462",
     "vatNumber": "NL815839091B01",
     "_links": {
         "self": {
             "href": "https://api.mollie.com/v2/organizations/me",
             "type": "application/hal+json"
         },
         "chargebacks": {
             "href": "https://api.mollie.com/v2/chargebacks",
             "type": "application/hal+json"
         },
         "customers": {
             "href": "https://api.mollie.com/v2/customers",
             "type": "application/hal+json"
         },
         "invoices": {
             "href": "https://api.mollie.com/v2/invoices",
             "type": "application/hal+json"
         },
         "payments": {
             "href": "https://api.mollie.com/v2/payments",
             "type": "application/hal+json"
         },
         "profiles": {
             "href": "https://api.mollie.com/v2/profiles",
             "type": "application/hal+json"
         },
         "refunds": {
             "href": "https://api.mollie.com/v2/refunds",
             "type": "application/hal+json"
         },
         "settlements": {
             "href": "https://api.mollie.com/v2/settlements",
             "type": "application/hal+json"
         },
         "documentation": {
             "href": "https://docs.mollie.com/reference/v2/organizations-api/current-organization",
             "type": "text/html"
         }
     }
 }`
