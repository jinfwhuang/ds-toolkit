package ds

import (
	protods "github.com/jinfwhuang/ds-toolkit/proto/ds"
)

// Do I have permission to the DataBlob
func (u *User) checkPerm(data *protods.DataBlob) bool {
	panic("not implemented")
}

// Get the decrypted data in EncryptedData
func (u *User) extractData(data *protods.DataBlob) []byte {
	panic("not implemented")
}

// User add key to an existing Blob and creat a new DataBlob
func (u *User) addKey(blob *protods.DataBlob, newDataOwner *User) *protods.DataBlob {
	panic("not implemented")
}

// 1. Generate an AES-key
// 2. Add an entry to "Secrets"
// 3. Encrypt data with AES key
func (u *User) createDataBlob(data []byte) *protods.DataBlob {
	panic("not implemented")
}
