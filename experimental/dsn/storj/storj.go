package storj

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"storj.io/uplink"
)

const (
	myAccessGrant = "1cCYAHiNyogwPRkZbkpo8C6txEeJ4cktWZoHzwAKTwNzGeAKhXyJZojfb8jSGGFSAc5M3zGgswJQXVEkrJ8na6nj4aeRRdtwvjgGFC5GHNWSXaN5YSrXzsXkUwwdLmf7A3Mk1rBR2v2mk6KfQUjLY1sko72neXkC3kDpvKV91mtp3yntk33A2SMMrMcBdEcFeNEqgRkn8TPsWBrZqDUs6JAZQQitLAQpyd6LceZ8yebzZntDWxiSpHwRGSwKHgNXGjKTBqgCcnEm"
	myBucket      = "dgf-test"
)

func Write(key string, data []byte) error {
	// Parse the Access Grant.
	access, err := uplink.ParseAccess(myAccessGrant)
	if err != nil {
		return fmt.Errorf("could not parse accwess grant: %v", err)
	}

	// Open up the Project we will be working with.
	project, err := uplink.OpenProject(context.Background(), access)
	if err != nil {
		return fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(context.Background(), myBucket)
	if err != nil {
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Intitiate the upload of our Object to the specified bucket and key.
	upload, err := project.UploadObject(context.Background(), myBucket, key, nil)
	if err != nil {
		return fmt.Errorf("could not initiate upload: %v", err)
	}

	// Copy the data to the upload.
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		return fmt.Errorf("could not commit uploaded object: %v", err)
	}
	return nil
}

func Read(key string) ([]byte, error) {
	// Parse the Access Grant.
	access, err := uplink.ParseAccess(myAccessGrant)
	if err != nil {
		return nil, fmt.Errorf("could not parse accwess grant: %v", err)
	}

	// Open up the Project we will be working with.
	project, err := uplink.OpenProject(context.Background(), access)
	if err != nil {
		return nil, fmt.Errorf("could not open project: %v", err)
	}

	defer project.Close()
	// Initiate a download of the same object again
	download, err := project.DownloadObject(context.Background(), myBucket, key, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open object: %v", err)
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		return nil, fmt.Errorf("could not read data: %v", err)
	}

	return receivedContents, nil
}
