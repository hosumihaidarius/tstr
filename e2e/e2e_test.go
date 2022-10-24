package e2e

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"
)

const (
	Default   = "default"
	Namespace = "namespace"
	TestFile  = "../test/register_test.json"
)

var accessToken = "dev"
var grpcAddr = "localhost:9000"

type AccessTokens struct {
	AccessTokens []AccessToken `json:"accessTokens"`
	AccessToken  AccessToken   `json:"accessToken"`
}

type AccessToken struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	NamespaceSelectors []string `json:"namespaceSelectors"`
	Scopes             []string `json:"scopes"`
	IssuedAt           string   `json:"issuedAt"`
	ExpiresAt          string   `json:"expiresAt"`
}

type Tests struct {
	Test Test `json:"test"`
}

type Test struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}

var accessTokens AccessTokens
var tests Tests
var testId string

func TestCtlAccessTokens(t *testing.T) {
	var accessTokenId string
	t.Run("TestListAccessTokens", func(t *testing.T) {
		output := runCmd(createTstrCommand("ctl access-token list", accessToken, grpcAddr))
		err := json.Unmarshal(output, &accessTokens)
		if err != nil {
			t.Error("Test Failed. The output is not a valid json format")
		}
		if len(accessTokens.AccessTokens) > 0 {
			t.Log("Test Passed")
		} else {
			t.Error("Test Failed, the access token is not created")
		}
		accessTokenId = accessTokens.AccessTokens[0].Id
	})
	t.Run("TestGetAccessToken", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl access-token get %s", accessTokenId), accessToken, grpcAddr))
		err := json.Unmarshal(output, &accessTokens)
		if err != nil {
			t.Error("Test Failed. The output is not a valid json format")
		}
		if accessTokens.AccessToken.Id == accessTokenId {
			t.Logf("Test Passed, the get request retrieved the token %s", accessTokenId)
		} else {
			t.Errorf("Test Failed, the access token was not created")
		}
	})
}

// The runCmdInBackground function is used to start the runner in the background
func TestRun(t *testing.T) {
	t.Run("TestRunner", func(t *testing.T) {
		command := createTstrCommand(fmt.Sprintf("run %s %s", createFlag("accept-label-selectors", "region=nyc"), createFlag("name", "C02FQ5M1MD6M")), accessToken, grpcAddr)
		// TODO : we need to check if docker is running
		runCmdInBackground(command)
	})
}

func TestCtlTest(t *testing.T) {

	t.Run("TestRegisterTest", func(t *testing.T) {
		command := createTstrCommand(fmt.Sprintf("ctl test register %s %s ",
			createFlag("file", TestFile),
			createFlag(Namespace, Default)), accessToken, grpcAddr)
		output := runCmd(command)
		err := json.Unmarshal(output, &tests)
		if !isValidJson(output) {
			output = removeCharsUntilValidJson(output)
			err = json.Unmarshal(output, &accessTokens)
		}
		if err != nil {
			t.Error("Test Failed. The output is not a valid json format")
		}
		if tests.Test.Id != "" {
			testId = tests.Test.Id
			t.Logf("Test Passed, the test was created %s", tests.Test.Id)
		} else {
			t.Errorf("Test Failed, the test was not registered")
		}
	})

	t.Run("TestGetRegisteredTest", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test get %s %s", testId, createFlag(Namespace, Default)), accessToken, grpcAddr))
		err := json.Unmarshal(output, &tests)
		if !isValidJson(output) {
			output = removeCharsUntilValidJson(output)
			err = json.Unmarshal(output, &accessTokens)
		}
		if err != nil {
			t.Error("Test Failed. The output is not a valid json format")
		}
		if tests.Test.Id != "" {
			t.Logf("Test Passed, the test was created %s", tests.Test.Id)
		} else {
			t.Errorf("Test Failed, the test was not registered")
		}
	})

	t.Run("TestGetTestList", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test list %s", createFlag(Namespace, Default)), accessToken, grpcAddr))
		err := json.Unmarshal(output, &accessTokens)
		if !isValidJson(output) {
			output = removeCharsUntilValidJson(output)
			err = json.Unmarshal(output, &accessTokens)
		}
		if err != nil {
			t.Error("Test Failed. The output is not a valid even after removing invalid characters until {")
		}
	})

	t.Run("TestDeleteRegisteredTest", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test delete %s", testId), accessToken, grpcAddr))
		err := json.Unmarshal(output, &tests)
		if !isValidJson(output) {
			output = removeCharsUntilValidJson(output)
			err = json.Unmarshal(output, &accessTokens)
		}
		if err != nil {
			t.Error("Test Failed. The output is not a valid json format")
		}
		if tests.Test.Id != "" {
			t.Logf("Test Passed, the test was created %s", tests.Test.Id)
		} else {
			t.Errorf("Test Failed, the test was not registered")
		}
	})

}

func removeCharsUntilValidJson(output []byte) []byte {
	for i := 0; i < len(output); i++ {
		if output[i] == '{' {
			return output[i:]
		}
	}
	return output
}

func isValidJson(output []byte) bool {
	var jsonObject map[string]interface{}
	return json.Unmarshal(output, &jsonObject) == nil
}

func createTstrCommand(cmd string, accessToken string, grpcAddr string) string {
	return fmt.Sprintf("../tstr %s --access-token %s --grpc-addr %s --insecure", cmd, accessToken, grpcAddr)
}

func createFlag(flag string, value string) string {
	return fmt.Sprintf("--%s %s", flag, value)
}

func runCmd(cmd string) []byte {
	command := exec.Command("bash", "-c", cmd)
	output, err := command.Output()
	if err != nil {
		panic(err)
	}
	return output
}

func runCmdInBackground(cmd string) {
	command := exec.Command("bash", "-c", cmd)
	err := command.Start()
	if err != nil {
		panic(err)
	}
}
