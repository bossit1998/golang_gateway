package v1

// // @Security ApiKeyAuth
// // @Router /v1/auth/platforms [GET]
// // @Summary Get All Platforms
// // @Description API for getting platforms
// // @Tags auth
// // @Accept  json
// // @Produce  json
// // @Param page query integer false "page"
// // @Param limit query integer false "limit"
// // @Success 200 {object} models.GetAllPlatformsModel
// // @Failure 400 {object} models.ResponseError
// // @Failure 404 {object} models.ResponseError
// // @Failure 500 {object} models.ResponseError
// func (h *handlerV1) GetAllPlatforms(c *gin.Context) {
// 	var (
// 		jspbMarshal jsonpb.Marshaler
// 	)

// 	jspbMarshal.OrigName = true
// 	jspbMarshal.EmitDefaults = true

// 	page, err := ParsePageQueryParam(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.ResponseError{
// 			Error: ErrorBadRequest,
// 		})
// 		return
// 	}

// 	limit, err := ParseLimitQueryParam(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.ResponseError{
// 			Error: ErrorBadRequest,
// 		})
// 		return
// 	}

// 	res, err := h.grpcClient.PlatformService().GetAll(
// 		context.Background(),
// 		&pba.GetAllRequest{
// 			Page:  page,
// 			Limit: limit,
// 		},
// 	)
// 	if handleGRPCErr(c, h.log, err) {
// 		return
// 	}
// 	js, err := jspbMarshal.MarshalToString(res)

// 	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
// 		return
// 	}

// 	c.Header("Content-Type", "application/json")
// 	c.String(http.StatusOK, js)
// }
