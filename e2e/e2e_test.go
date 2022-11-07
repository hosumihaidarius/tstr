package e2e

import (
	"encoding/json"
	"fmt"
	commonv1 "github.com/nanzhong/tstr/api/common/v1"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

const (
	Default                = "default"
	TestFile               = "../testdata/requests/register_test.json"
	NamespaceFlag          = "namespace"
	NameFlag               = "name"
	IdFlag                 = "id"
	ScopesFlag             = "scopes"
	NameSpaceSelectorsFlag = "namespace-selectors"
	FileFlag               = "file"
)

var accessToken = "dev"
var grpcAddr = "localhost:9000"

type AccessTokens struct {
	AccessTokens []commonv1.AccessToken `json:"accessTokens"`
	AccessToken  commonv1.AccessToken   `json:"accessToken"`
}

type Tests struct {
	Test  commonv1.Test   `json:"test"`
	Tests []commonv1.Test `json:"tests"`
}

var accessTokens AccessTokens
var tests Tests
var testId string

func TestCtlAccessTokens(t *testing.T) {
	var accessTokenId string
	var accessTokenName = "test"
	var scopesFlagAdmin = "ADMIN"
	t.Run("TestIssueAccessToken", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl access-token issue %s %s %s",
			createFlag(NameFlag, accessTokenName),
			createFlag(ScopesFlag, scopesFlagAdmin),
			createFlag(NameSpaceSelectorsFlag, Default)), accessToken, grpcAddr))
		_ = json.Unmarshal(output, &accessTokens)
		if !isValidJson(output) {
			output = removeRunesUntilStartOfJson(output)
			_ = json.Unmarshal(output, &accessTokens)
		}
		assert.True(t, isValidJson(output), "Test Failed. The output is not a valid json format")
		assert.NotEqual(t, accessTokens.AccessToken.Id, "", "The access token was not issued")
		assert.Equal(t, accessTokens.AccessToken.Name, accessTokenName,
			"The access token received does not match with the one provided in the name flag")
		accessTokenId = accessTokens.AccessToken.Id
	})
	t.Run("TestGetAccessToken", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl access-token get %s", accessTokenId), accessToken, grpcAddr))
		_ = json.Unmarshal(output, &accessTokens)
		assert.True(t, isValidJson(output), "Test Failed. The output is not a valid json format")
		assert.Equal(t, accessTokens.AccessToken.Id, accessTokenId,
			"The access token id does not match with the issued one")
		assert.Equal(t, accessTokens.AccessToken.Name, accessTokenName,
			"The access token received does not match with the one provided in the name flag")
	})
	t.Run("TestListAccessTokens", func(t *testing.T) {
		output := runCmd(createTstrCommand("ctl access-token list", accessToken, grpcAddr))
		_ = json.Unmarshal(output, &accessTokens)
		assert.True(t, isValidJson(output), "Test Failed. The output is not a valid json format")
		assert.Greater(t, len(accessTokens.AccessTokens), 0, "The access token list is empty")
	})
}

func TestCtlTest(t *testing.T) {
	t.Run("TestRegisterTest", func(t *testing.T) {
		command := createTstrCommand(fmt.Sprintf("ctl test register %s %s ",
			createFlag(FileFlag, TestFile),
			createFlag(NamespaceFlag, Default)), accessToken, grpcAddr)
		output := runCmd(command)
		err := json.Unmarshal(output, &tests)
		if !isValidJson(output) {
			output = removeRunesUntilStartOfJson(output)
			err = json.Unmarshal(output, &tests)
		}
		assert.Nil(t, err, "Test Failed, unable to unmarshal the output")
		assert.NotEmpty(t, tests.Test.Id, "The test id is empty")
		assert.Equal(t, tests.Test.Name, parseJsonFile(TestFile)["name"].(string), fmt.Sprintf("The test name from ctl command output does not match the one from %s", TestFile))
		testId = tests.Test.Id
	})
	t.Run("TestGetRegisteredTest", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test get %s %s", testId, createFlag(NamespaceFlag, Default)), accessToken, grpcAddr))
		err := json.Unmarshal(output, &tests)
		if !isValidJson(output) {
			output = removeRunesUntilStartOfJson(output)
			err = json.Unmarshal(output, &tests)
		}
		assert.Nil(t, err, "Test Failed, unable to unmarshal the output")
		assert.NotEmpty(t, tests.Test.Id, "The test id is empty")
		assert.Equal(t, tests.Test.Name, parseJsonFile(TestFile)["name"].(string), fmt.Sprintf("The test name from ctl command output does not match the one from %s", TestFile))
	})
	t.Run("TestUpdateRegisteredTest", func(t *testing.T) {
		updatedTestName := "updatedTest"
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test register %s %s %s",
			createFlag(IdFlag, testId),
			createFlag(NameFlag, updatedTestName),
			createFlag(NamespaceFlag, Default)), accessToken, grpcAddr))
		assert.Contains(t, string(output), fmt.Sprintf("Updating existing test: %s", updatedTestName),
			"The test was not successfully updated")
	})
	t.Run("TestGetTestList", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test list %s", createFlag(NamespaceFlag, Default)), accessToken, grpcAddr))
		err := json.Unmarshal(output, &tests)
		if !isValidJson(output) {
			output = removeRunesUntilStartOfJson(output)
			err = json.Unmarshal(output, &tests)
		}
		assert.Greater(t, len(tests.Tests), 0, "The test list is empty")
		assert.Nil(t, err, "Test Failed, unable to unmarshal the output")
	})
	t.Run("TestDeleteRegisteredTest", func(t *testing.T) {
		output := runCmd(createTstrCommand(fmt.Sprintf("ctl test delete %s %s", testId,
			createFlag(NamespaceFlag, Default)), accessToken, grpcAddr))
		assert.Contains(t, string(output), fmt.Sprintf("Deleting registered test %s", testId), "The test was not deleted")
	})
}

func removeRunesUntilStartOfJson(output []byte) []byte {
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

func parseJsonFile(filePath string) map[string]interface{} {
	var jsonObject map[string]interface{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &jsonObject)
	if err != nil {
		panic(err)
	}
	return jsonObject
}
