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
func (h *UploadImage) Handler(ctx context.Context, imageFile, imageName, dir string) (string, error) {
	result, err := h.cloudinary.Upload.Upload(ctx, imageFile, uploader.UploadParams{PublicID: imageName, Folder: "foodiebuddie/" + dir})
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println("upload success")
	log.Println(result.SecureURL)
	return result.SecureURL, nil
}
