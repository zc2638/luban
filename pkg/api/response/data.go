/**
 * Created by zc on 2020/8/8.
**/
package response

type Timestamp struct {
	CreatedTS int64 `json:"created_ts"`
	UpdatedTS int64 `json:"updated_ts"`
}

type LoginResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Host     string `json:"host"`
}

type UserInfoResult struct {
	Username string `json:"username"`
	Host     string `json:"host"`
}

type SpaceResultItem struct {
	SpaceID string `json:"space_id"`
	Name    string `json:"name"`
	Timestamp
}

type ResourceResultItem struct {
	ResourceID string `json:"resource_id"`
	SpaceID    string `json:"space_id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Format     string `json:"format"`
	Content    string `json:"content"`
	Label      string `json:"label"`
	Timestamp
}

type VersionResultItem struct {
	VersionID  string `json:"version_id"`
	ResourceID string `json:"resource_id"`
	Name       string `json:"name"`
	Format     string `json:"format"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	Timestamp
}
