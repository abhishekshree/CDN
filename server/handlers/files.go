package handlers

import (
	"log"

	"github.com/abhishekshree/cdn/utils"
	"github.com/gin-gonic/gin"
)

const MAX_SIZE = 5000000000

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

	ctx.SaveUploadedFile(file, upload_folder+"/"+uuid+"__"+file.Filename)
	log.Println(file.Filename + " to " + uuid + "__" + file.Filename)
	ctx.JSON(200, gin.H{
		"message":  "uploaded",
		"filename": uuid + "__" + file.Filename,
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

type DeleteRequest struct {
	Filename string `json:"filename"`
}

func DeleteFileHandler(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	//TOSO: can add some secret code before deletion later for security
	filename := req.Filename

	ok := utils.DeleteFile(upload_folder + "/" + filename)
	if !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Could not delete file"})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "deleted",
	})
}

type ZipRequest struct {
	Files   []string `json:"files"`
	OutFile string   `json:"outfile"`
}

func ZipFilesHandler(ctx *gin.Context) {
	var req ZipRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	filename, err := utils.ZipFiles(req.Files, req.OutFile)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "zipped",
		"filename": filename,
	})
}
