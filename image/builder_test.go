// Copyright (c) 2019 Seams-CMS and contributors. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package image

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	var url string

	url = Build("foobar", "path.jpg").SkipCDN().Blur().Boxed(10, 20).Url()
	assert.Equal(t, "https://assets-nocdn.seams-cms.com/p/blur/boxed(10,20)/foobar/path.jpg", url)

	url = Build("foobar", "path.jpg").Url()
	assert.Equal(t, "https://assets.seams-cms.com/foobar/path.jpg", url)

	url = Build("foobar", "path.jpg").UseCDN().Url()
	assert.Equal(t, "https://assets.seams-cms.com/foobar/path.jpg", url)

	url = Build("foobar", "path.jpg").UseCDN().Colorize(10, 20, 30, 50).Blur().Blur().Crop(CROP_BOTTOM, 10, 20).Url()
	assert.Equal(t, "https://assets.seams-cms.com/p/blur/blur/colorize(10,20,30,50)/crop(bottom,10,20)/foobar/path.jpg", url)

	url = Build("foobar", "path.jpg").UseCDN().Colorize(10, 20, 30, 50).CropSides().Blur().Crop(CROP_BOTTOM, 10, 20).Blur().Url()
	assert.Equal(t, "https://assets.seams-cms.com/p/blur/blur/colorize(10,20,30,50)/crop(bottom,10,20)/cropsides()/foobar/path.jpg", url)

	url = Build("foobar", "path.jpg").Flip(FLIP_HORIZONTAL).Gray().Height(100).Rotate(104).Negate().Width(10).Url()
	assert.Equal(t, "https://assets.seams-cms.com/p/flip(horizontal)/gray()/height(100)/negate()/rotate(104)/width(10)/foobar/path.jpg", url)

}
