package cryptox

import "github.com/Rockup-Consulting/go_std/core/secrets"

func NewDevRotationService() RotationService {
	secretMap, err := secrets.MapFromBase64String("ewogICAgIjEiOiAidGhpc2hhc3RvYmUzMmJ5dGVzZm9yaXR0b3dvcmshOikiLAogICAgIjIiOiAidGhpc2hhc3RvYmUzMmFzd2VsbGZvcml0dG93b3JrOikiCn0=")
	if err != nil {
		panic(err)
	}
	authService, err := NewRotationService(secretMap)
	if err != nil {
		panic(err)
	}

	return authService
}
