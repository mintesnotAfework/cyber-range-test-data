package models

func CustomErrorHandlerMiddleware(err error) error {
	return err
}
