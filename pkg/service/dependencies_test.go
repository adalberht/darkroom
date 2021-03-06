package service

import (
	"testing"

	"github.com/gojek/darkroom/pkg/config"
	"github.com/gojek/darkroom/pkg/storage/aws/cloudfront"
	"github.com/gojek/darkroom/pkg/storage/aws/s3"
	"github.com/gojek/darkroom/pkg/storage/webfolder"
	"github.com/stretchr/testify/assert"
)

func TestNewDependencies(t *testing.T) {
	deps := NewDependencies()
	assert.NotNil(t, deps)
	assert.Nil(t, deps.Storage)
}

func TestNewDependenciesWithWebFolderStorage(t *testing.T) {
	v := config.Viper()
	v.Set("source.kind", "WebFolder")
	v.Set("source.baseURL", "https://example.com/path/to/folder")
	config.Update()

	deps := NewDependencies()
	assert.NotNil(t, deps)
	assert.IsType(t, &webfolder.Storage{}, deps.Storage)
}

func TestNewDependenciesWithS3Storage(t *testing.T) {
	v := config.Viper()
	v.Set("source.kind", "S3")
	config.Update()

	deps := NewDependencies()
	assert.NotNil(t, deps)
	assert.IsType(t, &s3.Storage{}, deps.Storage)
}

func TestNewDependenciesWithCloudfrontStorage(t *testing.T) {
	v := config.Viper()
	v.Set("source.kind", "Cloudfront")
	v.Set("source.secureProtocol", "true")
	config.Update()

	deps := NewDependencies()
	assert.NotNil(t, deps)
	assert.IsType(t, &cloudfront.Storage{}, deps.Storage)
}
