import OpenLogin from "@toruslabs/openlogin";

// clientId can be any string for localhost
const openlogin = new OpenLogin({ clientId: "YOUR_PROJECT_ID", network: "testnet" });

await openlogin.init();

// if openlogin instance has private key then user is already logged in
if (openlogin.privKey) {
   console.log("User is already logged in. Private key: " + openlogin.privKey);
} else {
    await openlogin.login({
        loginProvider: "google",
        redirectUrl: "https://example.com/home",
    });
}


console.log("ffff")