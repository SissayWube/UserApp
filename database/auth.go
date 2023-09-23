package db

import "UserApp/model"

func RevokeRefreshToken(userID uint) error {

	// Update the refresh tokens to mark them as revoked for the specified user
	return db.Model(&model.RefreshToken{}).Where("user_id = ?", userID).Update("revoked", true).Error
}
