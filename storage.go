package aefire

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

func (a *AEFire) StorageBucket() *storage.BucketHandle {
	bucket, err := a.Storage.Bucket(GAEAppID() + ".appspot.com")

	PanicIfError(err)

	return bucket
}

func (a *AEFire) StorageObject(storageUrl string) *storage.ObjectHandle {
	if strings.HasPrefix(storageUrl, "gs://") {
		u, e := url.Parse(storageUrl)
		PanicIfError(e)

		b, e := a.Storage.Bucket(u.Hostname())
		PanicIfError(e)

		return b.Object(strings.TrimLeft(u.Path, "/"))
	} else {
		return a.StorageBucket().Object(strings.TrimLeft(storageUrl, "/"))
	}
}

func (a *AEFire) StoreReader(c context.Context, reader io.ReadCloser, filePath string) error {
	b, err := ioutil.ReadAll(reader)
	defer reader.Close()

	if err != nil {
		return err
	}

	return a.StoreBytes(c, b, filePath)
}

func (a *AEFire) StoreBytes(c context.Context, b []byte, filePath string) error {
	bucket, err := a.Storage.Bucket(GAEAppID() + ".appspot.com")
	if err != nil {
		return err
	}

	obj := bucket.Object(filePath)
	writer := obj.NewWriter(c)
	defer writer.Close()
	_, err = writer.Write(b)

	return err
}

func StorageObjectUrl(obj *storage.ObjectHandle) string {
	return fmt.Sprintf(
		"gs://%s/%s",
		obj.BucketName(),
		obj.ObjectName())
}
