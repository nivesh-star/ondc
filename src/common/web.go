package common

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nivesh-star/ondc/src/types"
)

func SetupHTTPServer(httpPort int, httpRouter *gin.Engine) (*http.Server, *gin.Engine) {
	if httpRouter == nil {
		httpRouter = gin.Default()
	}
	addr := fmt.Sprintf(":%d", httpPort)
	// address := ResolveAddress(addr)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           httpRouter,
		ReadHeaderTimeout: 30 * time.Second}
	return httpServer, httpRouter
}

func RunHTTPServer(serverName string, httpServer *http.Server) {
	slog.Info("%s server Listening and serving HTTP on %s\n", serverName, httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("%s httpServer listen error: %s\n", serverName, err)
	}
	slog.Info("%s server exiting", serverName)
}

func SendOndcGWRequest(action string, payload []byte) error {

	authHeader, err := GetAuthHeader(payload, os.Getenv("SIGN_PRIVATE_KEY"))
	if err != nil {
		return errors.New("GetAuthHeader failed: " + err.Error())
	}

	gwAddress := os.Getenv("ONDC_GW_ADDRESS")
	req, err := http.NewRequest(http.MethodPost, gwAddress+action, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)
	slog.Info(gwAddress)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer resp.Body.Close()

	// var responseDefault types.InlineResponseDefault
	// err = json.NewDecoder(resp.Body).Decode(&responseDefault)
	// if err != nil {
	// 	return errors.New("Fail to decode ONDC gateway response:" + err.Error())
	// }
	content, _ := io.ReadAll(resp.Body)
	slog.Info("Sandeep" + string(content))
	return nil
}

func SendOndcGWErrorResponse(c *gin.Context, err string) {
	response := types.InlineResponseDefault{
		Message: &types.InlineResponseDefaultMessage{
			Ack: &types.AllOfinlineResponseDefaultMessageAck{
				Status: "NACK",
			},
		},
		Error_: &types.ModelError{
			Message: err,
		},
	}
	c.JSON(http.StatusBadRequest, response)
}

func SendOndcGWErrorSuccess(c *gin.Context) {
	response := types.InlineResponseDefault{
		Message: &types.InlineResponseDefaultMessage{
			Ack: &types.AllOfinlineResponseDefaultMessageAck{
				Status: "NACK",
			},
		},
		Error_: &types.ModelError{
			Message: "Success",
		},
	}
	c.JSON(http.StatusBadRequest, response)
}
