package imaging

import (
	"github.com/ErfannNikooee/onvif/xsd"
	"github.com/ErfannNikooee/onvif/xsd/onvif"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"timg:GetServiceCapabilities"`
}

type GetImagingSettings struct {
	XMLName          string               `xml:"timg:GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetImagingSettingsResponse struct {
	ImagingSettings  onvif.ImagingSettings20 `xml:"ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"ForcePersistence"`
}

type SetImagingSettings struct {
	XMLName          string                  `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.ImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"timg:ForcePersistence"`
}

type SetImagingSettingsResponse struct{}

type GetOptions struct {
	XMLName          string               `xml:"timg:GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Move struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMove      `xml:"timg:Focus"`
}

type AbsoluteMove struct {
	XMLName          string                  `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	Focus            onvif.AbsoluteFocusMove `xml:"timg:Focus"`
}

type RelativeMove struct {
	XMLName          string                  `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	Focus            onvif.RelativeFocusMove `xml:"timg:Focus"`
}

type ContinuousMove struct {
	XMLName          string                    `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken      `xml:"timg:VideoSourceToken"`
	Focus            onvif.ContinuousFocusMove `xml:"timg:Focus"`
}

type MoveResponse struct { // is empty
}

type GetMoveOptions struct {
	XMLName          string               `xml:"timg:GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetMoveOptionsResponse struct {
	MoveOptions onvif.MoveOptions20 `xml:"MoveOptions"`
}

type Stop struct {
	XMLName          string               `xml:"timg:Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type StopResponse struct{}

type GetStatus struct {
	XMLName          string               `xml:"timg:GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetStatusResponse struct {
	Status onvif.ImagingStatus20
}

type GetPresets struct {
	XMLName          string               `xml:"timg:GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"timg:GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"timg:SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"timg:PresetToken"`
}

type SetCurrentPresetResponse struct {
}
