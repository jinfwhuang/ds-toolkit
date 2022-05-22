# Encryption, Permission, and Sharing



## Distributed Key Management

- Each document gets a unique encryption key
- Each document has a "grant-policy". The state machine understands which identity could run programs or direct access to the documents.
- The unique encryption key is stored as many shares. See Shamir’s Secret Sharing. No one single key distribution node could have direct access to the document. Let say 3 shares are required to reconstruct the key.
- For a compute-node to run the computation on the document, the compute-node get "auth" from 3 different key nodes. The auth mechanism is proving that user is part of the "grant-policy"
- The auth mechanism could just be signing a message.

#### Allow short term encryption key
- This has to involve rekeying the encryption key
- One way to do this is
  - allow the "storage nodes" to ask for "primary key"
  - storage node create ephemeral encryption key and encrypt a copy of the data for the compute node
  - the ephemeral key could be requested and controlled by the key committee
  - data is stored encrypted as primary key


## 