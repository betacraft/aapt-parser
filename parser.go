package aaptparse

import (
	"log"
	"os/exec"
	"strings"
)

var (
	permissions         = "uses-permission"
	featuresNotRequired = "uses-feature-not-required"
	packageKey          = "package"
	usesGl              = "uses-gl-es"
	appLabel            = "application-label"
	libsNotRequired     = "uses-library-not-required"
	targetSdkVersion    = "targetSdkVersion"
	sdkVersion          = "sdkVersion"
)

type Apk struct {
	Permissions         []string
	FeaturesNotRequired []string
	FeaturesRequired    []string // TODO
	LibsNotRequired     []string
	LibsRequired        []string // TODO
	AppLabel            string
	PackageName         string
	VersionCode         int
	VersionName         string
	TargetSdkVersion    string
	SdkVersion          string
	GlUse               string
}

func Parse(apkPath string) *Apk {
	apk := new(Apk)
	op, err := exec.Command("aapt", "dump", "badging", apkPath).Output()
	if err == exec.ErrNotFound {
		log.Println("Install aapt first")
		return nil
	}

	if err != nil {
		log.Println("Check if path is correct, use absolute path")
		return nil
	}

	data := strings.TrimSpace(string(op))

	lines := strings.Split(data, "\n") // get all lines
	for _, line := range lines {
		if strings.Contains(line, permissions) {
			getPermissionInfo(line, apk)
			continue
		}
		if strings.Contains(line, featuresNotRequired) {
			getFeatureNotReqInfo(line, apk)
			continue
		}
		if strings.Contains(line, packageKey) {
			getPackageInfo(line, apk)
			continue
		}
		if strings.Contains(line, usesGl) {
			getGlInfo(line, apk)
			continue
		}
		if strings.Contains(line, appLabel) {
			getAppLabel(line, apk)
			continue
		}
		if strings.Contains(line, libsNotRequired) {
			getLibsNotReqInfo(line, apk)
			continue
		}
		if strings.Contains(line, targetSdkVersion) {
			getTargetSdkInfo(line, apk)
			continue
		}
		if strings.Contains(line, sdkVersion) {
			getSdkInfo(line, apk)
			continue
		}
	}

	return apk
}
