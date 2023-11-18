package imageuploader

import (
	"context"
	"log"

	cld "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/cloudinary"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type UploadImage struct {
	cloudinary *cloudinary.Cloudinary
}

func NewUploadImage() *UploadImage {
	return &UploadImage{
		cloudinary: cld.Cld,
	}
}
func (h *UploadImage) Handler(ctx context.Context, imageName, dir string, imageFile any) (string, error) {
	log.Println("File is", imageFile)
	result, err := h.cloudinary.Upload.Upload(
		ctx,
		imageFile,
		uploader.UploadParams{
			PublicID: imageName,
			Folder:   "foodiebuddie/" + dir,
		})
	log.Println("Result is", result)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println("upload success")
	log.Println(result.SecureURL)
	return result.SecureURL, nil
}
