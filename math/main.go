package main

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"encoding/json"
	"math/util/numberutil"
)
//Json string will be unmrashalled to Request struct
type Request struct {
	Method string
	Points struct {
		X float64
		Y float64
	}
}

//Response will be marshalled to Response struct
type Response struct {
	Result float64
	IsSuccess bool
	Err  string
}

func main(){
	response := Response{}
	response, err := run()
	if err != nil {
		response.IsSuccess = false
		response.Err = err.Error()
		response.Result = 0.0
	}
	data, errs := json.Marshal(response)
	if errs != nil{
		panic(errs)
	}
	fmt.Println(string(data))
}

func run() (Response,error) {
	response := Response{}

	if len(os.Args) != 2{
		return response, errors.New("Invalid json string as an input.")
	}

	//Get JSON string from commandline args
	args := os.Args[1]

	request := &Request{}
	//Unmarshal json string to Request struct
	err := json.Unmarshal([]byte(args), &request)
	if err != nil {
		return response, err
	} else{
		x := request.Points.X
		y := request.Points.Y
		switch method := strings.ToLower(request.Method); method {
		case "add":
			result, err := numberutil.Add(x, y)
			if err != nil{
				return response, err
			}

			response.Result = result//numberutil.Add(x, y)
		case "subtract":
			result, err := numberutil.Subtract(x, y)
			if err != nil{
				return response, err
			}
			response.Result = result//numberutil.Subtract(x, y)
			//TODO:More functionality can be add here	
		default:
			return response, errors.New("Invalid method.")
		}
	}
	response.IsSuccess = true
	response.Err = ""
	return response, nil
}
