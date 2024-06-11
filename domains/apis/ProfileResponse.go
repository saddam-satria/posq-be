package apis

type ProfileResponse struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Username       string  `json:"username"`
	PhoneNumber    string  `json:"phone_number"`
	ProfilePicture *string `json:"profile_picture"`
	Role           string  `json:"role"`
	RoleName       string  `json:"role_name"`
	BusinessId     string  `json:"business_id"`
}