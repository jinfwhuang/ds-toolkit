

Goal: Build an example of how to load the data in user-registry as the backend for a user database

- https://django-web3-auth.readthedocs.io/en/latest/readme.html

## Web3 login client
- The client use "pubkey" login flow to acquire jwt token
- Subsequent requests to restricted routes uses the jwt token 

## Web3 login server
- The server accepts login from users that are in the UserRegistry. All `admin` public keys are accepted for each user.
- The server uses a "pubkey" login flow to issue jwt token
- JWT examples
  - https://pkg.go.dev/github.com/golang-jwt/jwt#example-New-Hmac
  - https://github.com/golang-jwt/jwt/blob/main/http_example_test.go


## Implementation
- Implement the go server
- Implement a simple UI app in react
