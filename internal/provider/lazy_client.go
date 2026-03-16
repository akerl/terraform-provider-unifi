package provider

import (
	"context"
	"fmt"
	"sync"

	"github.com/ubiquiti-community/go-unifi/unifi"
)

type lazyClient struct {
	baseURL   string
	user      string
	pass      string
	insecure  bool
	subsystem string

	once  sync.Once
	inner *unifi.ApiClient
}

var initErr error

func (c *lazyClient) init(ctx context.Context) error {
	c.once.Do(func() {
		cfg := unifi.Config{
			BaseURL:       c.baseURL,
			Username:      c.user,
			Password:      c.pass,
			AllowInsecure: c.insecure,
		}
		c.inner, initErr = unifi.New(ctx, &cfg)
	})
	return initErr
}

func (c *lazyClient) Version() string {
	if err := c.init(context.Background()); err != nil {
		panic(fmt.Sprintf("client not initialized: %s", err))
	}
	return c.inner.Version()
}
func (c *lazyClient) ListClientGroup(ctx context.Context, site string) ([]unifi.ClientGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListClientGroup(ctx, site)
}
func (c *lazyClient) ListWLANGroup(ctx context.Context, site string) ([]unifi.WLANGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListWLANGroup(ctx, site)
}
func (c *lazyClient) ListAPGroup(ctx context.Context, site string) ([]unifi.APGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListAPGroup(ctx, site)
}
func (c *lazyClient) DeleteNetwork(ctx context.Context, site, id, name string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteNetwork(ctx, site, id, name)
}
func (c *lazyClient) CreateNetwork(ctx context.Context, site string, d *unifi.Network) (*unifi.Network, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateNetwork(ctx, site, d)
}
func (c *lazyClient) GetNetwork(ctx context.Context, site, id string) (*unifi.Network, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetNetwork(ctx, site, id)
}
func (c *lazyClient) ListNetwork(ctx context.Context, site string) ([]unifi.Network, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListNetwork(ctx, site)
}
func (c *lazyClient) UpdateNetwork(ctx context.Context, site string, d *unifi.Network) (*unifi.Network, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateNetwork(ctx, site, d)
}
func (c *lazyClient) DeleteWLAN(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteWLAN(ctx, site, id)
}
func (c *lazyClient) CreateWLAN(ctx context.Context, site string, d *unifi.WLAN) (*unifi.WLAN, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateWLAN(ctx, site, d)
}
func (c *lazyClient) GetWLAN(ctx context.Context, site, id string) (*unifi.WLAN, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetWLAN(ctx, site, id)
}
func (c *lazyClient) UpdateWLAN(ctx context.Context, site string, d *unifi.WLAN) (*unifi.WLAN, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateWLAN(ctx, site, d)
}
func (c *lazyClient) DeleteClientGroup(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteClientGroup(ctx, site, id)
}
func (c *lazyClient) CreateClientGroup(ctx context.Context, site string, d *unifi.ClientGroup) (*unifi.ClientGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateClientGroup(ctx, site, d)
}
func (c *lazyClient) GetClientGroup(ctx context.Context, site, id string) (*unifi.ClientGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetClientGroup(ctx, site, id)
}
func (c *lazyClient) UpdateClientGroup(ctx context.Context, site string, d *unifi.ClientGroup) (*unifi.ClientGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateClientGroup(ctx, site, d)
}
func (c *lazyClient) GetDevice(ctx context.Context, site, id string) (*unifi.Device, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetDevice(ctx, site, id)
}
func (c *lazyClient) GetDeviceByMAC(ctx context.Context, site, mac string) (*unifi.Device, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetDeviceByMAC(ctx, site, mac)
}
func (c *lazyClient) CreateDevice(ctx context.Context, site string, d *unifi.Device) (*unifi.Device, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateDevice(ctx, site, d)
}
func (c *lazyClient) UpdateDevice(ctx context.Context, site string, d *unifi.Device) (*unifi.Device, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateDevice(ctx, site, d)
}
func (c *lazyClient) DeleteDevice(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteDevice(ctx, site, id)
}
func (c *lazyClient) ListDevice(ctx context.Context, site string) ([]unifi.Device, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListDevice(ctx, site)
}
func (c *lazyClient) AdoptDevice(ctx context.Context, site, mac string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.AdoptDevice(ctx, site, mac)
}
func (c *lazyClient) ForgetDevice(ctx context.Context, site, mac string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.ForgetDevice(ctx, site, mac)
}
func (c *lazyClient) GetClient(ctx context.Context, site, id string) (*unifi.Client, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetClient(ctx, site, id)
}
func (c *lazyClient) GetClientByMAC(ctx context.Context, site, mac string) (*unifi.Client, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetClientByMAC(ctx, site, mac)
}
func (c *lazyClient) CreateClient(ctx context.Context, site string, d *unifi.Client) (*unifi.Client, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateClient(ctx, site, d)
}
func (c *lazyClient) UpdateClient(ctx context.Context, site string, d *unifi.Client) (*unifi.Client, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateClient(ctx, site, d)
}
func (c *lazyClient) DeleteClientByMAC(ctx context.Context, site, mac string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteClientByMAC(ctx, site, mac)
}
func (c *lazyClient) BlockClientByMAC(ctx context.Context, site, mac string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.BlockClientByMAC(ctx, site, mac)
}
func (c *lazyClient) UnblockClientByMAC(ctx context.Context, site, mac string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.UnblockClientByMAC(ctx, site, mac)
}
func (c *lazyClient) ListFirewallGroup(ctx context.Context, site string) ([]unifi.FirewallGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListFirewallGroup(ctx, site)
}
func (c *lazyClient) DeleteFirewallGroup(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteFirewallGroup(ctx, site, id)
}
func (c *lazyClient) CreateFirewallGroup(ctx context.Context, site string, d *unifi.FirewallGroup) (*unifi.FirewallGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateFirewallGroup(ctx, site, d)
}
func (c *lazyClient) GetFirewallGroup(ctx context.Context, site, id string) (*unifi.FirewallGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetFirewallGroup(ctx, site, id)
}
func (c *lazyClient) UpdateFirewallGroup(ctx context.Context, site string, d *unifi.FirewallGroup) (*unifi.FirewallGroup, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateFirewallGroup(ctx, site, d)
}
func (c *lazyClient) ListFirewallRule(ctx context.Context, site string) ([]unifi.FirewallRule, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListFirewallRule(ctx, site)
}
func (c *lazyClient) DeleteFirewallRule(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteFirewallRule(ctx, site, id)
}
func (c *lazyClient) CreateFirewallRule(ctx context.Context, site string, d *unifi.FirewallRule) (*unifi.FirewallRule, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateFirewallRule(ctx, site, d)
}
func (c *lazyClient) GetFirewallRule(ctx context.Context, site, id string) (*unifi.FirewallRule, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetFirewallRule(ctx, site, id)
}
func (c *lazyClient) UpdateFirewallRule(ctx context.Context, site string, d *unifi.FirewallRule) (*unifi.FirewallRule, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateFirewallRule(ctx, site, d)
}
func (c *lazyClient) GetPortForward(ctx context.Context, site, id string) (*unifi.PortForward, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetPortForward(ctx, site, id)
}
func (c *lazyClient) DeletePortForward(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeletePortForward(ctx, site, id)
}
func (c *lazyClient) CreatePortForward(ctx context.Context, site string, d *unifi.PortForward) (*unifi.PortForward, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreatePortForward(ctx, site, d)
}
func (c *lazyClient) UpdatePortForward(ctx context.Context, site string, d *unifi.PortForward) (*unifi.PortForward, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdatePortForward(ctx, site, d)
}
func (c *lazyClient) ListRADIUSProfile(ctx context.Context, site string) ([]unifi.RADIUSProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListRADIUSProfile(ctx, site)
}
func (c *lazyClient) GetRADIUSProfile(ctx context.Context, site, id string) (*unifi.RADIUSProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetRADIUSProfile(ctx, site, id)
}
func (c *lazyClient) DeleteRADIUSProfile(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteRADIUSProfile(ctx, site, id)
}
func (c *lazyClient) CreateRADIUSProfile(ctx context.Context, site string, d *unifi.RADIUSProfile) (*unifi.RADIUSProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateRADIUSProfile(ctx, site, d)
}
func (c *lazyClient) UpdateRADIUSProfile(ctx context.Context, site string, d *unifi.RADIUSProfile) (*unifi.RADIUSProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateRADIUSProfile(ctx, site, d)
}
func (c *lazyClient) ListAccounts(ctx context.Context, site string) ([]unifi.Account, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListAccount(ctx, site)
}
func (c *lazyClient) GetAccount(ctx context.Context, site, id string) (*unifi.Account, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetAccount(ctx, site, id)
}
func (c *lazyClient) DeleteAccount(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteAccount(ctx, site, id)
}
func (c *lazyClient) CreateAccount(ctx context.Context, site string, d *unifi.Account) (*unifi.Account, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateAccount(ctx, site, d)
}
func (c *lazyClient) UpdateAccount(ctx context.Context, site string, d *unifi.Account) (*unifi.Account, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateAccount(ctx, site, d)
}
func (c *lazyClient) GetSite(ctx context.Context, id string) (*unifi.Site, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetSite(ctx, id)
}
func (c *lazyClient) ListSites(ctx context.Context) ([]unifi.Site, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListSites(ctx)
}
func (c *lazyClient) CreateSite(ctx context.Context, description string) ([]unifi.Site, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateSite(ctx, description)
}
func (c *lazyClient) DeleteSite(ctx context.Context, id string) ([]unifi.Site, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.DeleteSite(ctx, id)
}
func (c *lazyClient) UpdateSite(ctx context.Context, name, description string) ([]unifi.Site, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateSite(ctx, name, description)
}

func (c *lazyClient) ListPortProfile(ctx context.Context, site string) ([]unifi.PortProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListPortProfile(ctx, site)
}

func (c *lazyClient) GetPortProfile(ctx context.Context, site, id string) (*unifi.PortProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetPortProfile(ctx, site, id)
}

func (c *lazyClient) DeletePortProfile(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeletePortProfile(ctx, site, id)
}

func (c *lazyClient) CreatePortProfile(ctx context.Context, site string, d *unifi.PortProfile) (*unifi.PortProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreatePortProfile(ctx, site, d)
}

func (c *lazyClient) UpdatePortProfile(ctx context.Context, site string, d *unifi.PortProfile) (*unifi.PortProfile, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdatePortProfile(ctx, site, d)
}

func (c *lazyClient) ListRouting(ctx context.Context, site string) ([]unifi.Routing, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListRouting(ctx, site)
}

func (c *lazyClient) GetRouting(ctx context.Context, site, id string) (*unifi.Routing, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetRouting(ctx, site, id)
}

func (c *lazyClient) DeleteRouting(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteRouting(ctx, site, id)
}

func (c *lazyClient) CreateRouting(ctx context.Context, site string, d *unifi.Routing) (*unifi.Routing, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateRouting(ctx, site, d)
}

func (c *lazyClient) UpdateRouting(ctx context.Context, site string, d *unifi.Routing) (*unifi.Routing, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateRouting(ctx, site, d)
}

func (c *lazyClient) ListDynamicDNS(ctx context.Context, site string) ([]unifi.DynamicDNS, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.ListDynamicDNS(ctx, site)
}

func (c *lazyClient) GetDynamicDNS(ctx context.Context, site, id string) (*unifi.DynamicDNS, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.GetDynamicDNS(ctx, site, id)
}

func (c *lazyClient) DeleteDynamicDNS(ctx context.Context, site, id string) error {
	if err := c.init(ctx); err != nil {
		return err
	}
	return c.inner.DeleteDynamicDNS(ctx, site, id)
}

func (c *lazyClient) CreateDynamicDNS(ctx context.Context, site string, d *unifi.DynamicDNS) (*unifi.DynamicDNS, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.CreateDynamicDNS(ctx, site, d)
}

func (c *lazyClient) UpdateDynamicDNS(ctx context.Context, site string, d *unifi.DynamicDNS) (*unifi.DynamicDNS, error) {
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c.inner.UpdateDynamicDNS(ctx, site, d)
}
