package client

import (
	alert "github.com/appscode/api/alert/v0.1"
	artifactory "github.com/appscode/api/artifactory/v0.1"
	auth "github.com/appscode/api/auth/v0.1"
	backup "github.com/appscode/api/backup/v0.1"
	billing "github.com/appscode/api/billing/v0.1"
	ca "github.com/appscode/api/certificate/v0.1"
	ci "github.com/appscode/api/ci/v0.1"
	credential "github.com/appscode/api/credential/v0.1"
	glusterfs "github.com/appscode/api/glusterfs/v0.1"
	kubernetes "github.com/appscode/api/kubernetes/v0.1"
	monitoring "github.com/appscode/api/monitoring/v0.1"
	pv "github.com/appscode/api/pv/v0.1"
	"google.golang.org/grpc"
)

// multi client services are grouped by there main client. the api service
// clients are wrapped around with sub-service.
type multiClientInterface interface {
	Alert() *alertService
	Artifactory() *artifactoryService
	Authentication() *authenticationService
	Backup() *backupService
	Billing() *billingService
	CA() *caService
	CI() *ciService
	Credential() *credentialService
	GlusterFS() *glusterFSService
	Kubernetes() *kubernetesService
	Monitoring() *monitoringService
	PV() *pvService
}

type multiClientServices struct {
	alertClient          *alertService
	artifactoryClient    *artifactoryService
	authenticationClient *authenticationService
	backupClient         *backupService
	billingClient        *billingService
	caClient             *caService
	ciClient             *ciService
	credentialClient     *credentialService
	glusterfsClient      *glusterFSService
	kubernetesClient     *kubernetesService
	monitoringClient     *monitoringService
	pvClient             *pvService
}

func newMultiClientService(conn *grpc.ClientConn) multiClientInterface {
	return &multiClientServices{
		alertClient: &alertService{
			alertClient:        alert.NewAlertsClient(conn),
			notificationClient: alert.NewNotificationsClient(conn),
		},
		artifactoryClient: &artifactoryService{
			artifactClient: artifactory.NewArtifactsClient(conn),
			versionClient:  artifactory.NewVersionsClient(conn),
		},
		authenticationClient: &authenticationService{
			authenticationClient: auth.NewAuthenticationClient(conn),
			conduitClient:        auth.NewConduitClient(conn),
		},
		backupClient: &backupService{
			backupServerClient: backup.NewServerClient(conn),
			backupClientClient: backup.NewClientClient(conn),
		},
		billingClient: &billingService{
			paymentMethodClient: billing.NewPaymentMethodClient(conn),
			subscriptionClient:  billing.NewSubscriptionClient(conn),
			purchaseClient:      billing.NewPurchaseClient(conn),
			chargeClient:        billing.NewChargeClient(conn),
			invoiceClient:       billing.NewInvoiceClient(conn),
			quotaClient:         billing.NewQuotaClient(conn),
		},
		caClient: &caService{
			caClient:          ca.NewCAsClient(conn),
			acmeUsersClient:   ca.NewAcmeUsersClient(conn),
			certificateClient: ca.NewCertificatesClient(conn),
		},
		ciClient: &ciService{
			buildClient:  ci.NewBuildsClient(conn),
			jobClient:    ci.NewJobsClient(conn),
			masterClient: ci.NewMasterClient(conn),
			agentClient:  ci.NewAgentsClient(conn),
		},
		credentialClient: &credentialService{
			cloudClient: credential.NewCloudCredentialClient(conn),
		},
		glusterfsClient: &glusterFSService{
			clusterClient: glusterfs.NewClustersClient(conn),
			volumeClient:  glusterfs.NewVolumesClient(conn),
		},
		kubernetesClient: &kubernetesService{
			kubernetesClient: kubernetes.NewClientsClient(conn),
			clusterClient:    kubernetes.NewClustersClient(conn),
		},
		monitoringClient: &monitoringService{
			dashboardClient: monitoring.NewDashboardClient(conn),
		},
		pvClient: &pvService{
			disk: pv.NewDisksClient(conn),
			pv:   pv.NewPersistentVolumesClient(conn),
			pvc:  pv.NewPersistentVolumeClaimsClient(conn),
		},
	}
}

func (s *multiClientServices) Alert() *alertService {
	return s.alertClient
}

func (s *multiClientServices) Artifactory() *artifactoryService {
	return s.artifactoryClient
}

func (s *multiClientServices) Authentication() *authenticationService {
	return s.authenticationClient
}

func (s *multiClientServices) Backup() *backupService {
	return s.backupClient
}

func (s *multiClientServices) Billing() *billingService {
	return s.billingClient
}

func (s *multiClientServices) CA() *caService {
	return s.caClient
}

func (s *multiClientServices) CI() *ciService {
	return s.ciClient
}

func (s *multiClientServices) Credential() *credentialService {
	return s.credentialClient
}

func (s *multiClientServices) GlusterFS() *glusterFSService {
	return s.glusterfsClient
}

func (s *multiClientServices) Kubernetes() *kubernetesService {
	return s.kubernetesClient
}

func (s *multiClientServices) Monitoring() *monitoringService {
	return s.monitoringClient
}

func (s *multiClientServices) PV() *pvService {
	return s.pvClient
}

// original service clients that needs to exposed under grouped wrapper
// services.
type alertService struct {
	alertClient        alert.AlertsClient
	notificationClient alert.NotificationsClient
}

func (a *alertService) Alert() alert.AlertsClient {
	return a.alertClient
}

func (a *alertService) Notification() alert.NotificationsClient {
	return a.notificationClient
}

type artifactoryService struct {
	artifactClient artifactory.ArtifactsClient
	versionClient  artifactory.VersionsClient
}

func (a *artifactoryService) Artifacts() artifactory.ArtifactsClient {
	return a.artifactClient
}

func (a *artifactoryService) Versions() artifactory.VersionsClient {
	return a.versionClient
}

type authenticationService struct {
	authenticationClient auth.AuthenticationClient
	conduitClient        auth.ConduitClient
}

func (a *authenticationService) Authentication() auth.AuthenticationClient {
	return a.authenticationClient
}

func (a *authenticationService) Conduit() auth.ConduitClient {
	return a.conduitClient
}

type backupService struct {
	backupServerClient backup.ServerClient
	backupClientClient backup.ClientClient
}

func (b *backupService) Server() backup.ServerClient {
	return b.backupServerClient
}

func (b *backupService) Client() backup.ClientClient {
	return b.backupClientClient
}

type billingService struct {
	paymentMethodClient billing.PaymentMethodClient
	subscriptionClient  billing.SubscriptionClient
	purchaseClient      billing.PurchaseClient
	chargeClient        billing.ChargeClient
	invoiceClient       billing.InvoiceClient
	quotaClient         billing.QuotaClient
}

func (b *billingService) Charge() billing.ChargeClient {
	return b.chargeClient
}

func (b *billingService) Subscription() billing.SubscriptionClient {
	return b.subscriptionClient
}

func (b *billingService) Quota() billing.QuotaClient {
	return b.quotaClient
}

type caService struct {
	caClient          ca.CAsClient
	acmeUsersClient   ca.AcmeUsersClient
	certificateClient ca.CertificatesClient
}

func (c *caService) CAsClient() ca.CAsClient {
	return c.caClient
}

func (c *caService) AcmeUsersClient() ca.AcmeUsersClient {
	return c.acmeUsersClient
}

func (c *caService) CertificatesClient() ca.CertificatesClient {
	return c.certificateClient
}

type ciService struct {
	buildClient  ci.BuildsClient
	jobClient    ci.JobsClient
	masterClient ci.MasterClient
	agentClient  ci.AgentsClient
}

func (c *ciService) Build() ci.BuildsClient {
	return c.buildClient
}

func (c *ciService) Job() ci.JobsClient {
	return c.jobClient
}

func (c *ciService) Master() ci.MasterClient {
	return c.masterClient
}

func (c *ciService) Slave() ci.AgentsClient {
	return c.agentClient
}

type credentialService struct {
	cloudClient credential.CloudCredentialClient
}

func (c *credentialService) Cloud() credential.CloudCredentialClient {
	return c.cloudClient
}

type glusterFSService struct {
	clusterClient glusterfs.ClustersClient
	volumeClient  glusterfs.VolumesClient
}

func (g *glusterFSService) Cluster() glusterfs.ClustersClient {
	return g.clusterClient
}

func (g *glusterFSService) Volume() glusterfs.VolumesClient {
	return g.volumeClient
}

type kubernetesService struct {
	kubernetesClient kubernetes.ClientsClient
	clusterClient    kubernetes.ClustersClient
}

func (k *kubernetesService) Client() kubernetes.ClientsClient {
	return k.kubernetesClient
}

func (k *kubernetesService) Cluster() kubernetes.ClustersClient {
	return k.clusterClient
}

type monitoringService struct {
	dashboardClient monitoring.DashboardClient
}

func (m *monitoringService) Dashboard() monitoring.DashboardClient {
	return m.dashboardClient
}

type pvService struct {
	disk pv.DisksClient
	pv   pv.PersistentVolumesClient
	pvc  pv.PersistentVolumeClaimsClient
}

func (p *pvService) Disk() pv.DisksClient {
	return p.disk
}

func (p *pvService) PersistentVolume() pv.PersistentVolumesClient {
	return p.pv
}

func (p *pvService) PersistentVolumeClaim() pv.PersistentVolumeClaimsClient {
	return p.pvc
}
