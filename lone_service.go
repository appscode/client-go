package client

import (
	db "github.com/appscode/api/db/v0.1"
	"github.com/appscode/api/health"
	k8s "github.com/appscode/api/kubernetes/v0.1"
	loadbalancer "github.com/appscode/api/loadbalancer/v0.1"
	mailinglist "github.com/appscode/api/mailinglist/v0.1"
	namespace "github.com/appscode/api/namespace/v0.1"
	ssh "github.com/appscode/api/ssh/v0.1"
	"google.golang.org/grpc"
)

// single client service in api. returned directly the api client.
type loneClientInterface interface {
	DB() db.DatabasesClient
	Event() k8s.EventsClient
	Health() health.HealthClient
	LoadBalancer() loadbalancer.LoadBalancersClient
	MailingList() mailinglist.MailingListClient
	Namespace() namespace.NamespaceClient
	SSH() ssh.SSHClient
}

type loneClientServices struct {
	dbClient           db.DatabasesClient
	eventClient        k8s.EventsClient
	healthClient       health.HealthClient
	namespaceClient    namespace.NamespaceClient
	loadBalancerClient loadbalancer.LoadBalancersClient
	sshClient          ssh.SSHClient
	mailingListClient  mailinglist.MailingListClient
}

func newLoneClientService(conn *grpc.ClientConn) loneClientInterface {
	return &loneClientServices{
		dbClient:           db.NewDatabasesClient(conn),
		eventClient:        k8s.NewEventsClient(conn),
		healthClient:       health.NewHealthClient(conn),
		loadBalancerClient: loadbalancer.NewLoadBalancersClient(conn),
		sshClient:          ssh.NewSSHClient(conn),
		mailingListClient:  mailinglist.NewMailingListClient(conn),
	}
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

func (s *loneClientServices) Namespace() namespace.NamespaceClient {
	return s.namespaceClient
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
