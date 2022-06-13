package handlers

import (
	"log"

	"github.com/abhishekshree/cdn/utils"
	"github.com/gin-gonic/gin"
)

const MAX_SIZE = 500000

const upload_folder = "cdn"

const allowed_filetypes = "application/pdf"

func UploadFileHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if file.Size > MAX_SIZE {
		log.Printf("%v size is too large \n", file.Filename)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "File size is too large"})
		return
	}

	if file.Header.Get("Content-Type") != allowed_filetypes {
		log.Printf("%v is not allowed \n", file.Header.Get("Content-Type"))
		ctx.AbortWithStatusJSON(400, gin.H{"error": "File type is not allowed"})
		return
	}

	uuid, err := utils.GenerateUUID()
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.SaveUploadedFile(file, upload_folder+"/"+uuid+"_"+file.Filename)
	log.Println(file.Filename + " to " + uuid + "__" + file.Filename)
	ctx.JSON(200, gin.H{
		"message":  "uploaded",
		"filename": uuid + "_" + file.Filename,
	})
}

func ViewAllHandler(ctx *gin.Context) {
	files, err := utils.ListFiles(upload_folder)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"files": files,
	})
}

func ViewFileHandler(ctx *gin.Context) {
	filename := ctx.Param("filename")
	// send file
	ctx.File(upload_folder + "/" + filename)
}
