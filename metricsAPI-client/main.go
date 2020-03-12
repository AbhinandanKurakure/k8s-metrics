package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	metricspb "k8s-api/protofiles"
)

func main() {
	//variable declaration
	var nameSpace, searchPodName string
	var choice int
	//Dial creates a client connection to the given target.
	cc, err := grpc.Dial(":9090", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error While connecting to Server:", err)
	}

	defer cc.Close()
	//client api
	metricsClient := metricspb.NewClientAPIServiceClient(cc)

	//Take user input and run the selected API
	fmt.Println("To print all the pods in a particular NameSpace Press 1")
	fmt.Println("To print all the NameSpace in the Present Cluster Press 2")
	fmt.Println("To Search for a Particular Pod in a given NameSpace Press 3")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		//Take the input from the user
		fmt.Print("Enter the Namespace:")
		fmt.Scanln(&nameSpace)
		if nameSpace == "" {
			log.Fatal("cannot have empty NameSpace")
		}

		reqPods := &metricspb.InputValues{
			NameSpace: nameSpace,
		}

		//call is made to ListPods() function which returns all the pods present in a given cluster
		res, err := metricsClient.ListPods(context.Background(), reqPods)
		if err != nil {
			log.Fatal("Error from Client Side:", err)
		}

		fmt.Println("The Pods present in the given Namespace are")
		for _, pods := range res.Pods {
			fmt.Println(pods)
		}

	case 2:
		reqNameSpace := &metricspb.ListNameSpaceInput{}
		fmt.Println("The Total NameSpaces available in the Cluster are")
		//call is made to ListNameSpace() function which returns all the Namespaces
		res, err := metricsClient.ListNameSpace(context.Background(), reqNameSpace)
		if err != nil {
			log.Fatal(err.Error())
		}
		//loop to print all the namespaces
		for _, ns := range res.NameSpaceList {
			fmt.Println(ns)
		}
	case 3:
		//Taking the Inputs from the user for nameSpace and the pod named to be searched
		fmt.Print("Enter the Namespace:")
		fmt.Scanln(&nameSpace)

		if nameSpace == "" {
			log.Fatal("cannot have empty NameSpace")
		}

		fmt.Print("Enter the PodName which needs to be searched:")
		fmt.Scanln(&searchPodName)

		if nameSpace == "" {
			log.Fatal("cannot have empty NameSpace")
		}
		reqPodName := &metricspb.CheckPodInput{
			NameSpace: nameSpace,
			PodName:   searchPodName,
		}

		//call is made to CheckPod() which returns a string containing the message or error
		res, err := metricsClient.CheckPod(context.Background(), reqPodName)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(res.CheckPodOutput)
	default:
		//triggered in case user pressed any other key except (1,2,3)
		log.Fatal("invalid choice.")

	}

}
