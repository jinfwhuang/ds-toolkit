# Sia

## Summary

Sia had really thorough and well writen documentation. Unfortunately there weren't available SDKs, only GUI and HTTP client. The HTTP client worked quite well, GUI as well. The biggest drawback is that you have to spin up own node, as there are no nodes for public use provided. The current size of the chain is ~37 GB. After downloading the chain, an api password is saved on your machine, which is used for authentication.

Useful links:

- [Sia docs](https://api.sia.tech/#introduction)
- [Authentication](https://api.sia.tech/#authentication)
- [Developer resources](https://sia.tech/developers)

## How to run - Go/curl

1. Download either the Sia UI or the Sia Daemon (found in Developer resources)
2. Install the Sia UI or unzip the download for Sia Daemon
3. Run the Sia UI or run `siad` for Sia Daemon
4. Wait for the node to download
5. Copy your apiPassword (location of it is found in Authentication)
6. Change the `apiPassword` variable in the Go file
7. Run `sia.Write` providing desired file to upload's directory and also the destination directory in sia network (N.B.: there is a chance you have to adjust your settings, check the docs for more information)
8. Run `sia.Read` providing the destination directory to download your file
