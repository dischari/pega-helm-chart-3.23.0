package test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	k8score "k8s.io/api/core/v1"
)

var volumeDefaultMode int32 = 420
var volumeDefaultModePtr = &volumeDefaultMode

// VerifyCredentialsSecret - Verifies the credential secret deployed with the values as provided in default values.yaml
func VerifyCredentialsSecret(t *testing.T, helmChartPath string, options *helm.Options) {

	secretOutput := helm.RenderTemplate(t, options, helmChartPath, []string{"templates/pega-credentials-secret.yaml"})
	var secretobj k8score.Secret
	helm.UnmarshalK8SYaml(t, secretOutput, &secretobj)
	secretData := secretobj.Data
	require.Equal(t, string(secretData["DB_USERNAME"]), "YOUR_JDBC_USERNAME")
	require.Equal(t, string(secretData["DB_PASSWORD"]), "YOUR_JDBC_PASSWORD")
}

// VerfiyRegistrySecret - Verifies the registry secret deployed with the values as provided in default values.yaml
func VerfiyRegistrySecret(t *testing.T, helmChartPath string, options *helm.Options) {

	registrySecret := helm.RenderTemplate(t, options, helmChartPath, []string{"templates/pega-registry-secret.yaml"})
	var registrySecretObj k8score.Secret
	helm.UnmarshalK8SYaml(t, registrySecret, &registrySecretObj)
	reqgistrySecretData := registrySecretObj.Data
	require.Contains(t, string(reqgistrySecretData[".dockerconfigjson"]), "YOUR_DOCKER_REGISTRY")
	require.Contains(t, string(reqgistrySecretData[".dockerconfigjson"]), "WU9VUl9ET0NLRVJfUkVHSVNUUllfVVNFUk5BTUU6WU9VUl9ET0NLRVJfUkVHSVNUUllfUEFTU1dPUkQ=")
}

// compareConfigMapData - Compares the config map deployed for each kind of tier with the excepted xml's
func compareConfigMapData(t *testing.T, actualFileData string, expectedFileName string) {
	expectedFile, err := ioutil.ReadFile(expectedFileName)
	require.Empty(t, err)
	expectedFileData := string(expectedFile)
	expectedFileData = strings.Replace(expectedFileData, "\r", "", -1)

	equal := false
	if expectedFileData == actualFileData {
		equal = true
	}
	require.Equal(t, true, equal)
}
