package utility

import (
	"context"
	"io"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const (
	containerName = "media"
)

type storage struct {
	container azblob.ContainerClient
}

var onceStorage sync.Once
var instanceStorage *storage

func GetStorage() *storage {
	onceStorage.Do(func() {
		connStr := GetConfig().GetStorage()
		service, _ := azblob.NewServiceClientFromConnectionString(connStr, nil)
		container := service.NewContainerClient(containerName)

		instanceStorage = &storage{container}
	})

	return instanceStorage
}

func (obj storage) GetURL() string {
	return obj.container.URL() + "/"
}

func (obj storage) Upload(name string, data io.ReadSeekCloser) {

	ctx := context.Background()
	blockBlob := obj.container.NewBlockBlobClient(name)
	blockBlob.Upload(ctx, data, nil)

}

func (obj storage) Delete(name string) {

	ctx := context.Background()
	blockBlob := obj.container.NewBlockBlobClient(name)
	blockBlob.Delete(ctx, nil)

}
