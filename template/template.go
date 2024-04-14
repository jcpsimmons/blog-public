package template

import (
	"embed"
	"strings"
)

//go:embed *.html
var templateFiles embed.FS

func SiteFrameWrap(title, content, pageName, pageID string) ([]byte, error) {

	siteFrame, err := templateFiles.ReadFile("site-frame.html")
	if err != nil {
		return nil, err
	}

	siteFrameStr := string(siteFrame)
	siteFrameStr = strings.ReplaceAll(siteFrameStr, "TEMPLATE_PAGE_TITLE", title)
	siteFrameStr = strings.ReplaceAll(siteFrameStr, "TEMPLATE_PAGE_NAME", pageName)
	siteFrameStr = strings.ReplaceAll(siteFrameStr, "TEMPLATE_PAGE_CONTENT", content)
	siteFrameStr = strings.ReplaceAll(siteFrameStr, "TEMPLATE_BODY_ID", pageID)
	siteFrame = []byte(siteFrameStr)
	return siteFrame, nil
}
