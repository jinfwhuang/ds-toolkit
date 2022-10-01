# Ceramic

## Summary

Ceramic had relatively poor documentation. Javascript SDK, HTTP client and CLI are provided for interacting with the protocol.
Ceramic is based on Streams, where you can update the data in each stream and a log of the changes is kept. Each change of data in the stream requires the commit to be signed. However, by using the HTTP client there is no guidance how this should happen. There is an attempt for this to be done in `ceramic.go`.

The experience with the CLI was really easy though, there were good examples provided for it and everything worked out of the box.

The Javascript SDK was not tested.

Useful links:

- [Docs](https://developers.ceramic.network/learn/welcome/)
- [Getting started with CLI](https://developers.ceramic.network/build/cli/quick-start/)

## How to run - CLI

You can follow the steps in [Getting started with CLI](https://developers.ceramic.network/build/cli/quick-start/). The guide is well made and there were no caveats when following it.

1. `npm install --global @ceramicnetwork/cli @glazed/cli`
2. `ceramic daemon`, this should start the ceramic daemon used to process the payloads
3. `glaze did:create`, this should output a key - `did:key:abc...123` and seed - `ab...zx`
4. `glaze tile:create --key YOUR_KEY --content 'CONTENT'`, where `YOUR_KEY` should be the did key created from the previous step (without the leading `did:key:`), the output should be the newly created stream's id
5. `glaze tile:show STREAM_ID`, where `STREAM_ID` is the stream id from previous command
