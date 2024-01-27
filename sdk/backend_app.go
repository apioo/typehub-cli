// backend_app automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import "time"

type BackendApp struct {
	Id         int               `json:"id"`
	UserId     int               `json:"userId"`
	Status     int               `json:"status"`
	Name       string            `json:"name"`
	Url        string            `json:"url"`
	Parameters string            `json:"parameters"`
	AppKey     string            `json:"appKey"`
	AppSecret  string            `json:"appSecret"`
	Metadata   CommonMetadata    `json:"metadata"`
	Date       time.Time         `json:"date"`
	Scopes     []string          `json:"scopes"`
	Tokens     []BackendAppToken `json:"tokens"`
}
