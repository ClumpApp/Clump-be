package utility

import (
	"context"
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

func (obj storage) Upload(name string, data []byte) {

	ctx := context.Background()
	blockBlob := obj.container.NewBlockBlobClient(name)
	blockBlob.UploadBufferToBlockBlob(ctx, data, azblob.HighLevelUploadToBlockBlobOption{})

}

func (obj storage) Delete(name string) {

	ctx := context.Background()
	blockBlob := obj.container.NewBlockBlobClient(name)
	blockBlob.Delete(ctx, nil)

}
