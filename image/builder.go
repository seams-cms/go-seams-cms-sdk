package image

import (
	"fmt"
	"sort"
	"strings"
)

const ASSET_BASE_URL_CDN = "https://assets.seams-cms.com"
const ASSET_BASE_URL_NOCDN = "https://assets-nocdn.seams-cms.com"

type Crop string

const (
	CROP_TOP_LEFT     = "topleft"
	CROP_TOP          = "top"
	CROP_TOP_RIGHT    = "topright"
	CROP_LEFT         = "left"
	CROP_CENTER       = "center"
	CROP_RIGHT        = "right"
	CROP_BOTTOM_LEFT  = "bottomleft"
	CROP_BOTTOM       = "bottom"
	CROP_BOTTOM_RIGHT = "bottomright"
)

type Flip string

const (
	FLIP_HORIZONTAL = "horizontal"
	FLIP_VERTICAL   = "vertical"
	FLIP_BOTH       = "both"
)

type Builder struct {
	Workspace string
	Path      string
	CDN       bool
	Filters   []string
}

// Build creates a new builder for the given asset
func Build(workspace string, path string) *Builder {
	b := Builder{
		workspace,
		path,
		true,
		[]string{},
	}

	return &b
}

// SkipCDN when you want to skip the CDN and use the asset directly (not recommended)
func (b *Builder) SkipCDN() *Builder {
	b.CDN = false
	return b
}

// UseCDN when you want to use the CDN (default)
func (b *Builder) UseCDN() *Builder {
	b.CDN = true
	return b
}

func (b *Builder) Blur() *Builder {
	b.Filters = append(b.Filters, "blur")
	return b
}

func (b *Builder) Boxed(height, width int) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("boxed(%d,%d)", height, width))
	return b
}

func (b *Builder) Colorize(red, green, blue, alpha int) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("colorize(%d,%d,%d,%d)", red, green, blue, alpha))
	return b
}

func (b *Builder) Crop(position Crop, width, height int) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("crop(%s,%d,%d)", position, width, height))
	return b
}

func (b *Builder) CropSides() *Builder {
	b.Filters = append(b.Filters, "cropsides()")
	return b
}

func (b *Builder) Flip(direction Flip) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("flip(%s)", direction))
	return b
}

func (b *Builder) Gray() *Builder {
	b.Filters = append(b.Filters, "gray()")
	return b
}

func (b *Builder) Height(height int) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("height(%d)", height))
	return b
}

func (b *Builder) Negate() *Builder {
	b.Filters = append(b.Filters, "negate()")
	return b
}

func (b *Builder) Rotate(angle int) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("rotate(%d)", angle))
	return b
}

func (b *Builder) Width(width int) *Builder {
	b.Filters = append(b.Filters, fmt.Sprintf("width(%d)", width))
	return b
}

// Builds and returns the actual URL for this image
func (b *Builder) Url() string {
	url := fmt.Sprintf("%s/%s", b.Workspace, b.Path)

	if len(b.Filters) > 0 {
		sort.Strings(b.Filters)
		url = fmt.Sprintf("p/%s/%s", strings.Join(b.Filters, "/"), url)
	}

	if b.CDN {
		return fmt.Sprintf("%s/%s", ASSET_BASE_URL_CDN, url)
	}

	return fmt.Sprintf("%s/%s", ASSET_BASE_URL_NOCDN, url)
}
