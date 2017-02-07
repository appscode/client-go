package client

import (
	"fmt"
	"testing"

	auth "github.com/appscode/api/auth/v1beta1"
	kv1beta1 "github.com/appscode/api/kubernetes/v1beta1"
	kv1beta2 "github.com/appscode/api/kubernetes/v1beta2"
	"github.com/appscode/client/credential"
)

func TestKubernetesClient(t *testing.T) {
	c := getClient()
	//Clients client Test
	reqV1beta2 := &kv1beta2.ListResourceRequest{
		Cluster: "qa",
		Type:    "pods",
	}
	respV1beta2Client, err := c.Kubernetes().V1beta2().Client().List(c.Context(), reqV1beta2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Kubernetes V1beta2 Client: ", respV1beta2Client.Status)
	reqV1beta1Client := &kv1beta1.ClientRequest{}
	respV1beta1, err := c.Kubernetes().V1beta1().Client().Pods(c.Context(), reqV1beta1Client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Kubernetes V1beta1 Client:", respV1beta1.Status)
	//LoadBalancer client test
	req := &kv1beta1.ListRequest{
		Cluster: "qa",
	}
	resp, err := c.Kubernetes().V1beta1().LoadBalancer().List(c.Context(), req)
	fmt.Println("Status Loadbalancer v1beta1: ", resp.Status)
	//Disk client Test
	reqV1beta2Disk := &kv1beta2.DiskListRequest{
		Cluster: "qa",
	}
	respV1beta2Disk, err := c.Kubernetes().V1beta2().Disk().List(c.Context(), reqV1beta2Disk)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Kubernetes V1beta2 Disk:", respV1beta2Disk.Status)

}

func TestAuthClient(t *testing.T) {
	c := getClient()
	req := &auth.ProjectListRequest{
		true,
		[]string{""},
	}
	resp, err := c.Authentication().Project().List(c.Context(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Status Auth project:", resp.Status)
}

func getClient() *Client {
	c, err := New(NewOptionWithAuth("localhost:50077", credential.NewBearerAuth("appscode", "api-vtodsh6y7klyrfhenaemcnd55qng")))
	if err != nil {
		fmt.Println(err)
	}
	return c
}
