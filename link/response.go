package link

const RequestServerError = "oops! SERVER ERROR."
const RequestUrlError = "oops! NOT A CORRECT URL."
const RequestSuccess = "linked! HAHA."

const RequestSuccessCode = 1

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
