package testdata

// GetOnboardingStatusResponse example.
const GetOnboardingStatusResponse = `{
    "resource": "onboarding",
    "name": "Mollie B.V.",
    "signedUpAt": "2018-12-20T10:49:08+00:00",
    "status": "completed",
    "canReceivePayments": true,
    "canReceiveSettlements": true,
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/onboarding/me",
            "type": "application/hal+json"
        },
        "dashboard": {
            "href": "https://www.mollie.com/dashboard/onboarding",
            "type": "text/html"
        },
        "organization": {
            "href": "https://api.mollie.com/v2/organization/org_12345",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/onboarding-api/get-onboarding-status",
            "type": "text/html"
        }
    }
}
`
