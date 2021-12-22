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

var once sync.Once
var instance *storage

func getInstance() *storage {
	once.Do(func() {

		service, _ := azblob.NewServiceClientFromConnectionString("DefaultEndpointsProtocol=https;AccountName=clumpstorage;AccountKey=Iy2lMypYjDPIi6Gg6h7jWRFHh6L9x/mZKfUJUua5aq5hfL8I78Kcwag64hO6HmMLiXJwT7TIolL5FtSAyDYn+g==;EndpointSuffix=core.windows.net", nil)
		container := service.NewContainerClient(containerName)

		instance = &storage{container}
	})

	return instance
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
