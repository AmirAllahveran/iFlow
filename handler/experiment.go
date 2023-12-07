package handler

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) RunExperiment(c echo.Context) error {
	fmt.Println("inside function")
	//userID := extractID(c)
	//var experiment model.Experiment
	//var module model.Module
	//var experimentPackage model.Package

	// Define a variable to hold the request data
	var req ExperimentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	fmt.Println(req.Experiment)
	fmt.Println(req.Packages)
	networkConfig := types.NetworkCreate{
		CheckDuplicate: true,
		Driver:         "bridge",
		Options: map[string]string{
			"com.experiment.name": req.Experiment.Title,
		},
	}

	net, err := h.dck.CreateNetwork(networkConfig, req.Experiment.Title)
	if err != nil {
		fmt.Println("net error")
		fmt.Println(err)
		return err
	}

	networkID := net.ID
	fmt.Println("netID: " + networkID)

	for i, module := range req.Packages.Modules {
		var envVars []string
		for key, env := range req.Experiment.ModuleConfigs[i].Config {
			envVars[i] = key + "=" + env
		}
		config := container.Config{
			Image: module.Image,
			Env:   envVars,
			//Cmd:   strings.Split(req.Experiment.ModuleConfigs[i].Cmd, " "),
		}
		for counter := 0; counter < req.Experiment.ModuleConfigs[i].Count; counter++ {
			err := h.dck.ImagePull(module.Image)
			if err != nil {
				fmt.Println(err)
				return err
			}
			createResponse, err := h.dck.ContainerCreate(config, networkID, module.Name+"-"+strconv.Itoa(i))
			fmt.Println(createResponse.ID)
			if err != nil {
				fmt.Println("create error")
				fmt.Println(err)
				return err
			}
			err = h.dck.ContainerStart(createResponse.ID)
			if err != nil {
				fmt.Println()
				fmt.Println("error start")
				fmt.Println(err)
				return err
			}
		}
	}

	return c.JSON(http.StatusCreated, NewResponseData("experiment created"))
}
