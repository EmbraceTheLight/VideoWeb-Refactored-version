package domain

type FavoritesInfo struct {
	UserId        int64
	FavoritesId   int64
	FavoritesName *string
	IsPrivate     *int32
	Description   *string
}

// ToMap converts the FavoritesInfo struct to a map[string]any
func (f *FavoritesInfo) ToMap() map[string]any {
	ret := make(map[string]any)
	// id != zero value of int64
	if f.UserId != 0 {
		ret["user_id"] = f.UserId
	}
	if f.FavoritesId != 0 {
		ret["favorites_id"] = f.FavoritesId
	}

	if f.FavoritesName != nil {
		ret["favorites_name"] = f.FavoritesName
	}

	if f.IsPrivate != nil {
		ret["is_private"] = f.IsPrivate
	}

	if f.Description != nil {
		ret["description"] = f.Description
	}
	return ret
}
