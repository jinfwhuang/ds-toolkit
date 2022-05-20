package user

/**



 */
type DsData struct {
	// Header
	EncryptionKey []byte  // AES128, the key itself is encrypted
	userKeys []UserKey // Each entry could be used to recover AES_key

	// Data
	EncryptedData []byte  // Encrypted by AES_key
}

type UserKey struct {
	counterPartyPubkey []byte  // This should be ephemeral; H_b, p_b
	pubkey []byte  // H_a, private key is p_a

	// Encrypted with shared secret: s = p_a * p_b * G = p_a * H_b
	// Anyone with knowledge of p_a is able to calculate s
	// encryptedKey = AES_encrypt(ase_key, a)
	encryptedAesKey []byte // Be able to recover aes_key

}
func (u *DsData) unmarshall() []byte {
	panic("not implemented")
}

func (u *DsData) marshall([]byte) {
	panic("not implemented")
}


// Do I have permission to the DsMata
func (u *User) checkPerm(data *DsData) bool {
	panic("not implemented")
}

func (u *User) extractRawData(data *DsData) []byte {
	panic("not implemented")
}

func (u *User) addKey(data *DsData, owner, user User) {
	panic("not implemented")
}

// 1. Create an AES-key
// 2. Add itself to the Userkeys
// 3. Encrypt data with AES key
func (u *User) createDsData(data []byte) *DsData {
	panic("not implemented")
}