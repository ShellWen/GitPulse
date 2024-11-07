// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Developer struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	Bio       string `json:"bio"`
	Blog      string `json:"blog"`
	Email     string `json:"email"`
	Followers int64  `json:"followers"`
	Following int64  `json:"following"`
	Stars     int64  `json:"stars"`
	Repos     int64  `json:"repos"`
	Gists     int64  `json:"gists"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DeveloperWithPulsePoint struct {
	Developer  Developer  `json:"developer"`
	PulsePoint PulsePoint `json:"pulse_point"`
}

type GetLanguageUsageReq struct {
	Login string `path:"login"`
}

type GetLanguageUsageResp struct {
	LanguageUsage LanguageUsage `json:"languages"`
}

type GetLanguages struct {
}

type GetPulsePointRankReq struct {
	Language string `form:"language,optional"`
	Region   string `form:"region,optional"`
	Limit    int64  `form:"limit,optional,default=100,range=[1:100]"`
}

type GetPulsePointReq struct {
	Login string `path:"login"`
}

type GetPulsePointResp struct {
	PulsePoint PulsePoint `json:"pulse_point"`
}

type GetRegionReq struct {
	Login string `path:"login"`
}

type GetRegionResp struct {
	Region Region `json:"region"`
}

type Language struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type LanguageUsage struct {
	Id        int64                    `json:"id"`
	Languages []LanguageWithPercentage `json:"languages"`
	UpdatedAt string                   `json:"updated_at"`
}

type LanguageWithPercentage struct {
	Language   Language `json:"language"`
	Percentage float64  `json:"percentage"`
}

type PulsePoint struct {
	Id         int64   `json:"id"`
	PulsePoint float64 `json:"pulse_point"`
	UpdatedAt  string  `json:"updated_at"`
}

type Region struct {
	Id         int64   `json:"id"`
	Region     string  `json:"region"`
	Confidence float64 `json:"confidence"`
}
