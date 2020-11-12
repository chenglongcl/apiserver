package userprofile

type UpsertRequest struct {
	RealName    string `json:"realName"`
	Sex         int64  `json:"sex"`
	DateOfBirth string `json:"dateOfBirth"`
}
