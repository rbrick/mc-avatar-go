package mojang

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	NameToUuidLookup    = "https://api.mojang.com/users/profiles/minecraft/%s"
	UuidToProfileLookup = "https://sessionserver.mojang.com/session/minecraft/profile/%s"
)

type PropertyHeader struct {
	ID                string `json:"profileId"`
	Name              string `json:"profileName"`
	Timestamp         uint64 `json:"timestamp"`
	SignatureRequired bool   `json:",omitempty"`
}

type TextureMetadata struct {
	Model string `json:"model,omitempty"`
}

type Texture struct {
	Url      string          `json:"url"`
	Metadata TextureMetadata `json:"metadata,omitempty"`
}

type Textures struct {
	PropertyHeader
	Textures map[string]Texture
}

type Property struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature,omitempty"`
}

func (p *Property) DecodeTo(i interface{}) error {
	b, err := base64.StdEncoding.DecodeString(p.Value)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, i)
}

type Profile struct {
	ID             Uuid       `json:"id"`
	Name           string     `json:"name"`
	Properties     []Property `json:"properties,omitempty"`
	ProfileActions []string   `json:"profileActions,omitempty"`
}

func LookupUuid(name string) (*Profile, error) {
	resp, err := http.Get(fmt.Sprintf(NameToUuidLookup, name))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var profile Profile

	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func LookupProfile(uuid string) (*Profile, error) {
	resp, err := http.Get(fmt.Sprintf(UuidToProfileLookup, uuid))

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var profile Profile

	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
