package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	openshiftClusterAPINamespaceFileName = "05-openshift-cluster-api-namespace.yaml"
)

var _ asset.WritableAsset = (*OpenshiftClusterAPINamespace)(nil)

// OpenshiftClusterAPINamespace is the constant to represent contents of Openshift_ClusterApiNamespace.yaml file
type OpenshiftClusterAPINamespace struct {
	fileName string
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *OpenshiftClusterAPINamespace) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *OpenshiftClusterAPINamespace) Name() string {
	return "OpenshiftClusterAPINamespace"
}

// Generate generates the actual files by this asset
func (t *OpenshiftClusterAPINamespace) Generate(parents asset.Parents) error {
	t.fileName = openshiftClusterAPINamespaceFileName
	data, err := content.GetBootkubeTemplate(t.fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, t.fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *OpenshiftClusterAPINamespace) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *OpenshiftClusterAPINamespace) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, openshiftClusterAPINamespaceFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}