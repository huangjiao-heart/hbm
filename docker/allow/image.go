package allow

import (
	"fmt"
	"net/url"

	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/kassisol/hbm/docker/allow/types"
	policyobj "github.com/kassisol/hbm/object/policy"
	"github.com/kassisol/hbm/pkg/juliengk/go-docker/image"
	"github.com/kassisol/hbm/pkg/juliengk/go-log"
	"github.com/kassisol/hbm/pkg/juliengk/go-log/driver"
	"github.com/kassisol/hbm/pkg/juliengk/go-utils"
	"github.com/kassisol/hbm/version"
)

func ImageCreate(req authorization.Request, config *types.Config) *types.AllowResult {
	u, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		return &types.AllowResult{
			Allow: false,
			Msg: map[string]string{
				"text": fmt.Sprintf("Could not parse URL query"),
			},
		}
	}

	params := u.Query()

	if v, ok := params["fromImage"]; ok {
		if !AllowImage(v[0], config) {
			return &types.AllowResult{
				Allow: false,
				Msg: map[string]string{
					"text":           fmt.Sprintf("Image %s is not allowed to be pulled", v[0]),
					"resource_type":  "image",
					"resource_value": v[0],
				},
			}
		}
	}

	return &types.AllowResult{Allow: true}
}

func AllowImage(img string, config *types.Config) bool {
	defer utils.RecoverFunc()

	l, _ := log.NewDriver("standard", nil)

	p, err := policyobj.New("sqlite", config.AppPath)
	if err != nil {
		l.WithFields(driver.Fields{
			"storagedriver": "sqlite",
			"logdriver":     "standard",
			"version":       version.Version,
		}).Fatal(err)
	}
	defer p.End()

	i := image.NewImage(img)

	if i.Official {
		if p.Validate(config.Username, "config", "image_create_official", "") {
			return true
		}
	}

	if len(i.Registry) > 0 {
		if p.Validate(config.Username, "registry", i.Registry, "") {
			return true
		}
	}

	if p.Validate(config.Username, "image", i.String(), "") {
		return true
	}

	return false
}
