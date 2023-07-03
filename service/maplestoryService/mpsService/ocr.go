package mpsService

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	ocrErrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"qyyh-go/model"
)

func OCR(parm model.OcrParm, _ *gin.Context) (data any, err error) {
	img := parm.Img

	credential := common.NewCredential("AKIDLmJsNiUp4KPYyEQVcgnohdO4cxSKv23w", "hcaSBmIsnMsUu9E7rTAoIbmrnogKsFRx")
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-hongkong", cpf)

	request := ocr.NewRecognizeTableAccurateOCRRequest()

	request.ImageBase64 = common.StringPtr(img)

	response, err := client.RecognizeTableAccurateOCR(request)
	if _, ok := err.(*ocrErrors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if response.Response.TableDetections == nil || len(response.Response.TableDetections) == 0 || response.Response.TableDetections[0].Cells == nil {
		return nil, errors.New("text is null")
	}
	return response.Response.TableDetections[0].Cells, err
}
