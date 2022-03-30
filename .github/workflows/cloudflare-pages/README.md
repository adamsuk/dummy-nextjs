## CloudFlare Pages

This directory is for automating the CloudFlare Pages projects utilising [cloudflare-go](https://github.com/cloudflare/cloudflare-go).

### Require parameters

It is intended this functionality will be incorporated into CI at some stage so most (not all just yet, some are still hardcoded) variables are pulled from environment variables:

- CF_API_KEY
- CF_API_EMAIL
- CF_ACCOUNT_ID

### TODO

- Update Page projects functionality - not sure if it's me or `cloudflare-go`
- Github Actions wrapper
- Testing!!
- General tidy up
