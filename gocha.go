package gocha

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

var baseURL string = "https://quickchart.io/chart?%s"

type chartFormat string

const (
	PNG    chartFormat = "png"
	WEBP   chartFormat = "webp"
	JPG    chartFormat = "jpg"
	SVG    chartFormat = "svg"
	PDF    chartFormat = "pdf"
	BASE64 chartFormat = "base64"
)

type QuickChart struct {
	config           string
	width            string
	height           string
	devicePixelRatio string
	backgroundColor  string
	version          string
	format           chartFormat
}

func (qc *QuickChart) SetConfig(config string) {
	qc.config = config
}

func (qc *QuickChart) SetWidth(width int) {
	qc.width = strconv.Itoa(width)
}

func (qc *QuickChart) SetHeight(height int) {
	qc.height = strconv.Itoa(height)
}

func (qc *QuickChart) SetBackgroundColor(color string) {
	qc.backgroundColor = color
}

func (qc *QuickChart) SetDevicePixelRatio(ratio string) {
	qc.devicePixelRatio = ratio
}

func (qc *QuickChart) SetVersion(version string) {
	qc.version = version
}

func (qc *QuickChart) SetFormat(format chartFormat) {
	qc.format = format
}

func (qc *QuickChart) GetURL() (string, error) {
	if qc.config == "" {
		return "", errors.New("Config must be set")
	}

	params := url.Values{}
	params.Add("c", qc.config)
	params.Add("w", qc.width)
	params.Add("h", qc.height)
	params.Add("bkg", qc.backgroundColor)
	params.Add("devicePixelRatio", qc.devicePixelRatio)
	params.Add("f", string(qc.format))
	params.Add("v", qc.version)

	return fmt.Sprintf(baseURL, params.Encode()), nil
}

func NewQuickChart() QuickChart {
	return QuickChart{
		width:            "500",
		height:           "500",
		version:          "3",
		format:           PNG,
		devicePixelRatio: "1.0",
		backgroundColor:  "white",
	}

}
