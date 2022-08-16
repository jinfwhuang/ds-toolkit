# Storj

## Summary

Storj worked quite well, with the provided Uplink Go SDK. The documentation was good and the provided examples as well. There is also CLI but it was not tested. Alongside with Arweave it was the most effortless implementation from the five candidates.

Useful links:

- [Uplink - Go SDK](https://github.com/storj/uplink)

## How to run - Go

1. Create account on [storj](https://eu1.storj.io/signup)
2. Create an [api key](https://eu1.storj.io/access-grants)
3. Change the `myAccessGrant` variable in the Go file
4. Run `storj.Write`, providing key and data.
5. Run `storj.Read` providing a key, to retrieve the data
