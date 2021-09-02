package model

type EngagementRate struct {
	UserName   string  `json:"userName"`
	Percentage float64 `json:"percentage"`
}

type Instagram struct {
	EntryData struct {
		ProfilePage []struct {
			Graphql struct {
				User struct {
					EdgeFollowedBy struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					EdgeOwnerToTimelineMedia struct {
						Count int `json:"count"`
						Edges []struct {
							Node struct {
								DisplayURL  string `json:"display_url"`
								EdgeLikedBy struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								EdgeMediaToCaption struct {
									Edges []interface{} `json:"edges"`
								} `json:"edge_media_to_caption"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								EdgeMediaToTaggedUser struct {
									Edges []interface{} `json:"edges"`
								} `json:"edge_media_to_tagged_user"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}
