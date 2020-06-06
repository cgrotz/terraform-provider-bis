package bis

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// AuthSuccess is ...
type AuthSuccess struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// Authenticate requests a token from the IoT Suite by using the OAuth Client Credentials grant
func Authenticate(clientID string, clientSecret string, thingsSolutionID string) (AuthSuccess, error) {
	client := resty.New()
	requestBody := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=service:iot-things-eu-1:%s/full-access", clientID, clientSecret, thingsSolutionID)

	log.Printf("[INFO] Trying to authenticate using request %s", requestBody)
	var result AuthSuccess
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetBody(requestBody).
		SetResult(&result).
		Post("https://access.bosch-iot-suite.com/token")
	if err != nil {
		return result, err
	}
	if resp.StatusCode() != 200 {
		return result, fmt.Errorf("Unable to authenticate expected 200 got %d status=%s body=%s",
			resp.StatusCode(), resp.Status(), resp.Body())

	}
	return result, nil
}

// NamespaceEntry Struct used for the API calls to IoT Things
type NamespaceEntry struct {
	State   string `json:"state"`
	Default bool   `json:"default"`
}

// NamespaceUpdate Struct used for the API calls to IoT Things
type NamespaceUpdate struct {
	Default bool `json:"default"`
}

// RetrieveNamespace retrieves information about an existing namespace
func RetrieveNamespace(config Config, apiToken string, thingsSolutionID string, namespace string) (NamespaceEntry, error) {
	var namespaceEntry NamespaceEntry
	auth, err := Authenticate(config.ClientID, config.ClientSecret, thingsSolutionID)
	if err != nil {
		log.Printf("[ERROR] failed to authenticate %s", err)
		return namespaceEntry, err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken)).
		SetHeader("x-cr-api-token", apiToken).
		SetPathParams(map[string]string{
			"solutionId":  thingsSolutionID,
			"namespaceId": namespace,
		}).
		SetResult(&namespaceEntry).
		Get("https://things.eu-1.bosch-iot-suite.com/api/2/solutions/{solutionId}/namespaces/{namespaceId}")

	if err != nil {
		log.Printf("[ERROR] failed reading namespace from IoT Things API %s", err)
		return namespaceEntry, err
	}
	if resp.StatusCode() != 200 {
		return namespaceEntry, fmt.Errorf("Unable to make call to IoT Things API expected 200 got %d", resp.StatusCode())
	}
	return namespaceEntry, nil
}

// CreateNamespace tries to create a Namespace in IoT Things
func CreateNamespace(config Config, apiToken string, thingsSolutionID string, namespace string, defaultNamespace bool) error {
	var namespaceEntry NamespaceUpdate
	namespaceEntry.Default = defaultNamespace
	auth, err := Authenticate(config.ClientID, config.ClientSecret, thingsSolutionID)
	if err != nil {
		log.Printf("[ERROR] failed to authenticate %s", err)
		return err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", auth.AccessToken)).
		SetHeader("x-cr-api-token", apiToken).
		SetPathParams(map[string]string{
			"solutionId":  thingsSolutionID,
			"namespaceId": namespace,
		}).
		SetBody(namespaceEntry).
		Put("https://things.eu-1.bosch-iot-suite.com/api/2/solutions/{solutionId}/namespaces/{namespaceId}")

	if err != nil {
		log.Printf("[ERROR] failed creating namespace from IoT Things API %s", err)
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("Unable to make call to IoT Things API expected 200 got %d %s", resp.StatusCode(), resp.Body())
	}
	return nil
}
