package aaptparse

import (
	"fmt"
	"testing"
)

func TestParseApk(t *testing.T) {
	apk := Parse("eventbrite.apk")

	if apk == nil {
		t.Fail()
	}
	fmt.Println("Apk = ", "\n",
		"Label:", apk.AppLabel, "\n",
		"FeaturesNotRequired:", apk.FeaturesNotRequired, "\n",
		"Gl Use: ", apk.GlUse, "\n",
		"LibsNotRequired", apk.LibsNotRequired, "\n",
		"PackageName: ", apk.PackageName, "\n",
		"VersionCode: ", apk.VersionCode, "\n",
		"VersionName:", apk.VersionName, "\n",
		"SdkVersion:", apk.SdkVersion, "\n",
		"TargetSdkVersion:", apk.TargetSdkVersion, "\n",
		"UsesPermissions:", apk.Permissions, "\n",
		"NativeCode:", apk.NativeCode, "\n",
		"FeaturesRequired:", apk.FeaturesRequired)
}
