package cld

import (
	"log"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/config"
	"github.com/cloudinary/cloudinary-go"
)

var Cld *cloudinary.Cloudinary

func ConnectCloudinary(c *config.ImgUploaderCfg) error {
	cld, err := cloudinary.NewFromURL(c.CloudUrl)
	if err != nil {
		log.Println("Failed to intialize Cloudinary", err)
		return err
	}
	log.Println("cloudinary Success connected!!")
	Cld = cld
	return nil
}
