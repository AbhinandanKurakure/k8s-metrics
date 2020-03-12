package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	restService "k8s-api/metricsAPI-server/server-rest"
	metricspb "k8s-api/protofiles"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var kubeConfig *string

type server struct {
	GrpcServer *grpc.Server
	Listener   net.Listener
}

type PortConfig struct {
	PortType   string `json:"portType"`
	GRPCPortNo string `json:"GRPCPortNo"`
	HTTPPortNo string `json:"HTTPPortNo"`
}

func configuration() (string, string) {
	//Opening and reading from json file
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("Error While opening config.json file")
	}

	defer jsonFile.Close()

	byteValues, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	var config PortConfig

	err = json.Unmarshal(byteValues, &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	return config.GRPCPortNo, config.HTTPPortNo
}

func (s *server) CheckPod(ctx context.Context, req *metricspb.CheckPodInput) (*metricspb.CheckPodOutput, error) {

	//get the nameSpace and PodName
	nameSpace := req.GetNameSpace()
	podName := req.GetPodName()

	//check if the nameSpace given is Empty?
	if nameSpace == "" {
		log.Fatal("Namespace not defined(Empty string)")
	}

	//checkingPod is function which returns a string which contains whether the pod is present in the mentioned NameSpace or not.
	isPodPresent := checkingPod(podName, nameSpace)
	res := &metricspb.CheckPodOutput{
		CheckPodOutput: isPodPresent,
	}

	return res, nil

}

//ListPods function lists all the Pods present in the given NameSpace
func (s *server) ListPods(ctx context.Context, req *metricspb.InputValues) (*metricspb.OutputValues, error) {
	fmt.Println("Testing Function invoked")

	nameSpace := req.GetNameSpace()
	//check if given NameSpace is empty?
	if nameSpace == "" {
		log.Fatal("Namespace not defined(Empty string)")
	}

	//listingPods function returns the list of pods
	podLists, err := listingPods(nameSpace)
	if err != nil {
		res := &metricspb.OutputValues{
			Errors: err.Error(),
		}
		return res, err
	} else {
		res := &metricspb.OutputValues{
			Errors: "",
			Pods:   podLists,
		}
		return res, nil
	}

}

//ListNameSpace function lists all the NameSpace in the given cluster
func (c *server) ListNameSpace(ctx context.Context, in *metricspb.ListNameSpaceInput) (*metricspb.NameSpaceOutput, error) {
	nameSpaces, err := listingNameSpace()

	if err != nil {
		res := &metricspb.NameSpaceOutput{}
		return res, err
	} else {
		res := &metricspb.NameSpaceOutput{
			NameSpaceList: nameSpaces,
		}
		return res, err
	}

}

func main() {
	//homeDir returns minikube's default namespace config file
	if home := homeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(Optional) absolute path to kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	//get the port No from the configuration function
	grpcPortNo, httpPortNo := configuration()

	fmt.Println("Values of grpcPortNo is" + grpcPortNo + " http port no=" + httpPortNo)
	//startGRPC starts gRPC service at port 9090
	grpcServer, err := startGRPC(grpcPortNo)
	if err != nil {
		log.Fatal("failed to create gRPC server:", err)
		os.Exit(1)
	}

	go func() {
		//start HTTP service at port 5000
		restService.StartHTTP(grpcPortNo, httpPortNo)
	}()
	//grpcRun serves by listening at port 9090
	if err = grpcServer.grpcRun(); err != nil {
		log.Fatal("error While Running the Grpc server")
		os.Exit(1)
	}

}

func checkingPod(podName string, nameSpace string) string {
	//InClusterConfig returns a config object which uses the service account kubernetes gives to pods.
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Print("failed while obtaining Incluster configuration")
		//accessing default cluster configuration as InClusterConfig is not available
		log.Println("accessing default cluster configuration")
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {

			log.Fatal(err.Error(), " else there no running minikube or any cluster")

		}
	}

	//create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	pods, err := clientset.CoreV1().Pods(nameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var i int
	for _, ObtainedpodsList := range pods.Items {
		podsListName := ObtainedpodsList.GetName()

		for i = 0; i < len(podName); i++ {

			if podsListName[i] != podName[i] {
				break
			}

		}
		if i == len(podName) {
			return podName + " is present in the " + nameSpace + " NameSpace"
		}

	}
	return podName + " is not present in the " + nameSpace + "NameSpace"

}

func listingPods(nameSpace string) ([]string, error) {
	var podSlice []string

	//InClusterConfig returns a config object which uses the service account kubernetes gives to pods.
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Print("failed while obtaining Incluster configuration")
		//accessing default cluster configuration as InClusterConfig is not available
		log.Println("accessing default cluster configuration")
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {

			log.Fatal(err.Error(), " else there no running minikube or any cluster")
			return podSlice, err

		}
	}

	//create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return podSlice, err

	}

	//pods contains the list of all the pods available at the given namespace
	//clientset
	pods, err := clientset.CoreV1().Pods(nameSpace).List(context.Background(), metav1.ListOptions{})

	if err != nil {
		fmt.Println(err.Error())
		return podSlice, err

	}

	//The following code is used for displaying all the pod names
	for _, items := range pods.Items {
		fmt.Println(items.GetName())
		podSlice = append(podSlice, items.GetName())
	}

	return podSlice, nil
}

func listingNameSpace() ([]string, error) {
	var nameSpaceSlice []string
	//InClusterConfig returns a config object which uses the service account kubernetes gives to pods.
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Print("failed while obtaining Incluster configuration")

		//accessing default cluster configuration as InClusterConfig is not available
		log.Println("accessing default cluster configuration")
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {

			log.Fatal(err.Error(), " else there no running minikube or any cluster")
			return nameSpaceSlice, err

		}

	}

	//create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return nameSpaceSlice, err

	}

	nameSpace, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})

	for _, items := range nameSpace.Items {
		nameSpaceSlice = append(nameSpaceSlice, items.GetName())
	}

	return nameSpaceSlice, err

}

func startGRPC(gRPCPort string) (*server, error) {

	fmt.Println("gRPC Porttype=", "tcp", "PortNo=", gRPCPort)

	//The listen func creates the type of connection and port number
	lis, err := net.Listen("tcp", gRPCPort)

	if err != nil {
		log.Fatal("error in the Server.||Error is=", err)
	}

	fmt.Println("server has Started")
	grpcServer := grpc.NewServer()
	metricspb.RegisterClientAPIServiceServer(grpcServer, &server{})
	//metricspb.RegisterClientAPIServiceServer(grpcServer, &server{})

	//Serve() accepts incoming connections on the listener lis, creating a new ServerTransport and service goroutine for each.
	return &server{GrpcServer: grpcServer, Listener: lis}, nil

}

func (s *server) grpcRun() error {
	if err := s.GrpcServer.Serve(s.Listener); err != nil {
		fmt.Println("failed to start gRPC server at localhost:", err)
	}
	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return ""
}
