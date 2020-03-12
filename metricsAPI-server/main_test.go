package server_testing

import (
	"context"
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListingPods(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		t.Fatal("listing Pods should not raise an error")
	}

	for _, items := range pods.Items {
		fmt.Println(items.GetName())
	}
	expected := []string{"", "practice-deployment-b69b89546-nl8gv", "practice-deployment-b69b89546-xpjfm"}
	for i, items := range pods.Items {
		fmt.Println(i)
		fmt.Println("Say bye")
		if items.GetName() != expected[i] {
			t.Error("Expected" + expected[i] + "received" + items.GetName())
		}
	}
}
