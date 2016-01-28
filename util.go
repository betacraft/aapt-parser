package aaptparse

import (
	"strconv"
	"strings"
)

var singleQuoteReplace = strings.NewReplacer("'", "")

func getSplitDataAndRemoveSignleQuotes(line, sep string) string {
	tokens := strings.Split(line, sep)
	data := singleQuoteReplace.Replace(strings.TrimSpace(tokens[1]))
	return data
}

func getPermissionInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.Permissions = append(apk.Permissions, data)
}

func getFeatureNotReqInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.FeaturesNotRequired = append(apk.FeaturesNotRequired, data)
}

func getGlInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.GlUse = data
}

func getAppLabel(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.AppLabel = data
}

func getLibsNotReqInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.LibsNotRequired = append(apk.LibsNotRequired, data)
}

func getTargetSdkInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.TargetSdkVersion = data
}

func getSdkInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	apk.SdkVersion = data
}

func getPackageInfo(line string, apk *Apk) {
	data := getSplitDataAndRemoveSignleQuotes(line, ":")
	lines := strings.Split(data, " ")
	for _, v := range lines {
		if strings.HasPrefix(v, "name") {
			apk.PackageName = getSplitDataAndRemoveSignleQuotes(v, "=")
			continue
		}
		if strings.HasPrefix(v, "versionCode") {
			vCode := getSplitDataAndRemoveSignleQuotes(v, "=")
			apk.VersionCode, _ = strconv.Atoi(vCode)
			continue
		}
		if strings.HasPrefix(v, "versionName") {
			apk.VersionName = getSplitDataAndRemoveSignleQuotes(v, "=")
			continue
		}
	}
}
