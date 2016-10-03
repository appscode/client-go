package client

import (
	credential "github.com/appscode/api/credential/v1beta1"
	db "github.com/appscode/api/db/v0.1"
	"github.com/appscode/api/health"
	k8s "github.com/appscode/api/kubernetes/v1beta1"
	loadbalancer "github.com/appscode/api/loadbalancer/v1beta1"
	mailinglist "github.com/appscode/api/mailinglist/v1beta1"
	namespace "github.com/appscode/api/namespace/v1beta1"
	operation "github.com/appscode/api/operation/v1beta1"
	ssh "github.com/appscode/api/ssh/v1beta1"
	"google.golang.org/grpc"
)

// single client service in api. returned directly the api client.
type loneClientInterface interface {
	CloudCredential() credential.CloudCredentialsClient
	DB() db.DatabasesClient
	Event() k8s.EventsClient
	Health() health.HealthClient
	LoadBalancer() loadbalancer.LoadBalancersClient
	MailingList() mailinglist.MailingListClient
	Team() namespace.TeamsClient
	SSH() ssh.SSHClient
	Operation() operation.OperationsClient
}

type loneClientServices struct {
	cloudClient        credential.CloudCredentialsClient
	dbClient           db.DatabasesClient
	eventClient        k8s.EventsClient
	healthClient       health.HealthClient
	teamClient         namespace.TeamsClient
	loadBalancerClient loadbalancer.LoadBalancersClient
	sshClient          ssh.SSHClient
	mailingListClient  mailinglist.MailingListClient
	operationClient    operation.OperationsClient
}

func newLoneClientService(conn *grpc.ClientConn) loneClientInterface {
	return &loneClientServices{
		cloudClient:        credential.NewCloudCredentialsClient(conn),
		dbClient:           db.NewDatabasesClient(conn),
		eventClient:        k8s.NewEventsClient(conn),
		healthClient:       health.NewHealthClient(conn),
		loadBalancerClient: loadbalancer.NewLoadBalancersClient(conn),
		sshClient:          ssh.NewSSHClient(conn),
		mailingListClient:  mailinglist.NewMailingListClient(conn),
		operationClient:    operation.NewOperationsClient(conn),
	}
}

func (s *loneClientServices) CloudCredential() credential.CloudCredentialsClient {
	return s.cloudClient
}

func (s *loneClientServices) DB() db.DatabasesClient {
	return s.dbClient
}

func (s *loneClientServices) Event() k8s.EventsClient {
	return s.eventClient
}

func (s *loneClientServices) Health() health.HealthClient {
	return s.healthClient
}

func (s *loneClientServices) Team() namespace.TeamsClient {
	return s.teamClient
}

func (s *loneClientServices) LoadBalancer() loadbalancer.LoadBalancersClient {
	return s.loadBalancerClient
}

func (s *loneClientServices) SSH() ssh.SSHClient {
	return s.sshClient
}

func (s *loneClientServices) MailingList() mailinglist.MailingListClient {
	return s.mailingListClient
}

func (s *loneClientServices) Operation() operation.OperationsClient {
	return s.operationClient
}
