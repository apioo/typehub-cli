// commit_previous automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import "time"

type CommitPrevious struct {
	Id         int       `json:"id"`
	Message    string    `json:"message"`
	CommitHash string    `json:"commitHash"`
	Spec       any       `json:"spec"`
	InsertDate time.Time `json:"insertDate"`
}
