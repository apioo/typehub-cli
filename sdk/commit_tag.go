// CommitTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type CommitTag struct {
	internal *sdkgen.TagAbstract
}

// Get Returns a commit
func (client *CommitTag) Get(user string, document string, id string) (Commit, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/commit/:id", pathParams))
	if err != nil {
		return Commit{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Commit{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Commit{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Commit{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Commit
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Commit{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Commit{}, err
		}

		return Commit{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Commit{}, err
		}

		return Commit{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Commit{}, err
		}

		return Commit{}, &MessageException{
			Payload: response,
		}
	default:
		return Commit{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll Returns all commits for a document
func (client *CommitTag) GetAll(user string, document string, startIndex int, count int, search string) (CommitCollection, error) {
	pathParams := make(map[string]interface{})
	pathParams["user"] = user
	pathParams["document"] = document

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/document/:user/:document/commit", pathParams))
	if err != nil {
		return CommitCollection{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return CommitCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommitCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommitCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommitCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommitCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommitCollection{}, err
		}

		return CommitCollection{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommitCollection{}, err
		}

		return CommitCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommitCollection{}, err
		}

		return CommitCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return CommitCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewCommitTag(httpClient *http.Client, parser *sdkgen.Parser) *CommitTag {
	return &CommitTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
