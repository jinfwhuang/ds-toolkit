# High Level

- Service provider means the "backend service" that Digital Green operates internally
- Each user is assumed to control private key on client side
- An user use private key to sign in though message signing
- Service provider keeps a user DB, but the user DB use "elliptical curve" signing as the authentication mechanism


## private key management
- User could choose to manage its own private key.
- User could choose to use social login + Torus to manage a private key. The service provider still does not have access to the private key. This is recommended. See demo.
- User could choose to allow the service provider to manage the private key on behalf of the user.


## Torus Public Key Infrastructure
- Explain how does this work ...
- Web3Auth already provides a sufficient library for both web and mobile app applications


## Login system
- The service provider requests the user to sign a message
  - Signed message
  ```js
  addr: "0x4e6b0228a5bc0ca7f2a8bfac93b13aa9cc506f12"
  message: "Sign this message to prove you have access to this wallet and we will sign you in. This won't cost you any Ether. Timestamp: 1648430364297 "
  signature: "0x3b820b3098aab43a5d840cf836b873d4541b0188b529866ef102cea9667016066221261e53f10fecf75c21a49a6bccc350629b0c5677d6fd2e008c0ce132ca3b1b"
  ```
- The user is signed in as "0x4e6b0228a5bc0ca7f2a8bfac93b13aa9cc506f12"
  - The service provider could augment this login further with private data

- Examples:
  - https://thegraph.com/studio/

## Further improvements
- Accounts linkage
  - Different social logins leading to the same private key
    - See https://docs.tor.us/customauth/linking-accounts
  - Different private key leading to the same account
    - This should be handled in the service provider's backend
- 



# Demo
- Frontend app that login via metamask
- Frontend app that gets the private key from torus
- Backend service that accepts accepts a login

















