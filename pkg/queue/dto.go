package queue

type ContributionCreatedPayload struct {
	ContributionId int64  `json:"contributionId"`
	UserId         int64  `json:"userId"`
	UserName       string `json:"userName"`
	FacultyId      int64  `json:"facultyId"`
}

type ArticleUploadedPayload struct {
	ArticleId int64  `json:"articleId"`
	Link      string `json:"link"`
}
