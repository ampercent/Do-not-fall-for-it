package plugins

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/sudarshantk/honeypot/parser"
	"github.com/sudarshantk/honeypot/tracer"
)

func TestBuildSendEventFailValidation(t *testing.T) {
	beelzebubCloud := InitBeelzebubCloud("", "")

	_, err := beelzebubCloud.SendEvent(tracer.Event{})

	assert.Equal(t, "authToken is empty", err.Error())
}

func TestBuildSendEventWithResults(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	uri := "localhost:8081"

	// Given
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/events", uri),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, &tracer.Event{})
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	beelzebubCloud := InitBeelzebubCloud(uri, "fake_test_key_for_testing_only")
	beelzebubCloud.client = client

	//When
	result, err := beelzebubCloud.SendEvent(tracer.Event{})

	//Then
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func TestBuildSendEventErro(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	uri := "localhost:8081/events"

	// Given
	httpmock.RegisterResponder("POST", uri,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(500, ""), nil
		},
	)

	beelzebubCloud := InitBeelzebubCloud(uri, "fake_test_key_for_testing_only")
	beelzebubCloud.client = client

	//When
	result, _ := beelzebubCloud.SendEvent(tracer.Event{})

	//Then
	assert.Equal(t, false, result)
}

func TestGetHoneypotsConfigurationsWithResults(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	uri := "localhost:8081"

	// Given
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/honeypots", uri),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, &[]HoneypotConfigResponseDTO{
				{
					ID:      "fake_test_key_for_testing_only56",
					Config:  "apiVersion: \"v1\"\nprotocol: \"ssh\"\naddress: \":2222\"\ndescription: \"SSH interactive ChatGPT\"\ncommands:\n  - regex: \"^(.+)$\"\n    plugin: \"LLMHoneypot\"\nserverVersion: \"OpenSSH\"\nserverName: \"ubuntu\"\npasswordRegex: \"^(root|qwerty|Smoker666|fake_test_key_for_testing_only56|jenkins|minecraft|sinus|alex|postgres|Lyfake_test_key_for_testing_only56)$\"\ndeadlineTimeoutSeconds: 60\nplugin:\n  llmModel: \"gpt-4o\"\n  openAISecretKey: \"fake_test_key_for_testing_only\"\n",
					TokenID: "fake_test_key_for_testing_only567",
				},
			})
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	beelzebubCloud := InitBeelzebubCloud(uri, "fake_test_key_for_testing_only")
	beelzebubCloud.client = client

	//When
	result, err := beelzebubCloud.GetHoneypotsConfigurations()

	//Then
	assert.Equal(t, &[]parser.BeelzebubServiceConfiguration{
		{
			ApiVersion:  "v1",
			Protocol:    "ssh",
			Address:     ":2222",
			Description: "SSH interactive ChatGPT",
			Commands: []parser.Command{
				{
					RegexStr: "^(.+)$",
					Regex:    regexp.MustCompile("^(.+)$"),
					Plugin:   "LLMHoneypot",
				},
			},
			ServerVersion:          "OpenSSH",
			ServerName:             "ubuntu",
			PasswordRegex:          "^(root|qwerty|Smoker666|fake_test_key_for_testing_only56|jenkins|minecraft|sinus|alex|postgres|Lyfake_test_key_for_testing_only56)$",
			DeadlineTimeoutSeconds: 60,
			Plugin: parser.Plugin{
				LLMModel:        "gpt-4o",
				OpenAISecretKey: "fake_test_key_for_testing_only",
			},
		},
	}, &result)
	assert.Nil(t, err)
}

func TestGetHoneypotsConfigurationsWithErrorValidation(t *testing.T) {
	//Given
	beelzebubCloud := InitBeelzebubCloud("", "")

	//When
	result, err := beelzebubCloud.GetHoneypotsConfigurations()

	//Then
	assert.Nil(t, result)
	assert.Equal(t, "authToken is empty", err.Error())
}

func TestGetHoneypotsConfigurationsWithErrorAPI(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	uri := "localhost:8081"

	// Given
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/honeypots", uri),
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(500, ""), nil
		},
	)

	beelzebubCloud := InitBeelzebubCloud(uri, "fake_test_key_for_testing_only")
	beelzebubCloud.client = client

	//When
	result, err := beelzebubCloud.GetHoneypotsConfigurations()

	//Then
	assert.Nil(t, result)
	assert.Equal(t, "Response code: 500, error: ", err.Error())
}

func TestGetHoneypotsConfigurationsWithErrorUnmarshal(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	uri := "localhost:8081"

	// Given
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/honeypots", uri),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, "error")
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	beelzebubCloud := InitBeelzebubCloud(uri, "fake_test_key_for_testing_only")
	beelzebubCloud.client = client

	//When
	result, err := beelzebubCloud.GetHoneypotsConfigurations()

	//Then
	assert.Nil(t, result)
	assert.Equal(t, "json: cannot unmarshal string into Go value of type []plugins.HoneypotConfigResponseDTO", err.Error())
}

func TestGetHoneypotsConfigurationsWithErrorDeserializeYaml(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	uri := "localhost:8081"

	// Given
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/honeypots", uri),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, &[]HoneypotConfigResponseDTO{
				{
					ID:      "fake_test_key_for_testing_only56",
					Config:  "error",
					TokenID: "fake_test_key_for_testing_only567",
				},
			})
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	beelzebubCloud := InitBeelzebubCloud(uri, "fake_test_key_for_testing_only")
	beelzebubCloud.client = client

	//When
	result, err := beelzebubCloud.GetHoneypotsConfigurations()

	//Then
	assert.Nil(t, result)
	assert.Equal(t, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `error` into parser.BeelzebubServiceConfiguration", err.Error())
}
