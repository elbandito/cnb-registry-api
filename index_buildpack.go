package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

const MetadataLabel = "io.buildpacks.buildpackage.metadata"

//go:generate mockery --all --output=internal/mocks --case=underscore

type ImageFunction func(name.Reference, ...remote.Option) (v1.Image, error)

type Entry struct {
	Namespace string `json:"ns"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Address   string `json:"addr"`
}

type Metadata struct {
	ID       string
	Version  string
	Homepage string
	Stacks   []stack
}

type stack struct {
	ID string
}

func main() {
	fmt.Println("at=main msg='Building cacher cache'")

	if len(os.Args) != 2 {
		panic("invalid inputs: expected entry json")
	}
	rawEntry := os.Args[1]

	var e Entry
	if err := json.Unmarshal([]byte(rawEntry), &e); err != nil {
		panic("invalid inputs: unable to parse entry json")
	}

	m, err := FetchBuildpackConfig(e, remote.Image)
	if err != nil {
		panic(err)
	}

	err = UpdateOrInsertConfig(m)
	if err != nil {
		panic(err)
	}
}

func FetchBuildpackConfig(e Entry, imageFn ImageFunction) (Metadata, error) {
	ref, err := name.ParseReference(e.Address)
	if err != nil {
		return Metadata{}, err
	}

	if _, ok := ref.(name.Digest); !ok {

		return Metadata{}, errors.New(fmt.Sprintf("address is not a digest: %s", e.Address))
	}

	image, err := imageFn(ref)
	if err != nil {
		return Metadata{}, err
	}

	configFile, err := image.ConfigFile()
	if err != nil {
		return Metadata{}, err
	}

	raw, ok := configFile.Config.Labels[MetadataLabel]
	if !ok {
		return Metadata{}, errors.New(fmt.Sprintf("could not find metadata label for %s", e.Address))
	}

	var m Metadata
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		return Metadata{}, err
	}

	if fmt.Sprintf("%s/%s", e.Namespace, e.Name) != m.ID {
		return Metadata{}, errors.New(fmt.Sprintf("invalid ID for %s", e.Address))
	}

	if e.Version != m.Version {
		return Metadata{}, errors.New(fmt.Sprintf("invalid version for %s", e.Address))

	}

	var stacks []string
	for _, s := range m.Stacks {
		stacks = append(stacks, s.ID)
	}

	return m, nil
}

func UpdateOrInsertConfig(m Metadata) error {

	// TODO

	return nil
}
