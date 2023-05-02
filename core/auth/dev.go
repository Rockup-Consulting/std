package auth

// NewDevService is a utility that creates an auth.Service useful for testing or devving.
func NewDevService() *Service {
	secretMap, err := SecretMapFromBase64String("ewogICAgIjEiOiAidGhpc2hhc3RvYmUzMmJ5dGVzZm9yaXR0b3dvcmshOikiLAogICAgIjIiOiAidGhpc2hhc3RvYmUzMmFzd2VsbGZvcml0dG93b3JrOikiCn0=")
	if err != nil {
		panic(err)
	}
	authService := NewService(secretMap)

	return authService
}
