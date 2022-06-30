package dto

type UmsAdminParam struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Email    string `json:"email,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Note     string `json:"note,omitempty"`
}
