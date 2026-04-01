package util

import (
	"github.com/robinlg/iam-sdk-go/iam/service/iam"
	restclient "github.com/robinlg/iam-sdk-go/rest"
	"github.com/robinlg/iam-sdk-go/tools/clientcmd"
	"github.com/robinlg/iam/pkg/cli/genericclioptions"
)

type factoryImpl struct {
	clientGetter genericclioptions.RESTClientGetter
}

func NewFactory(clientGetter genericclioptions.RESTClientGetter) Factory {
	if clientGetter == nil {
		panic("attempt to instantiate client_access_factory with nil clientGetter")
	}

	f := &factoryImpl{
		clientGetter: clientGetter,
	}

	return f
}

func (f *factoryImpl) ToRESTConfig() (*restclient.Config, error) {
	return f.clientGetter.ToRESTConfig()
}

func (f *factoryImpl) ToRawIAMConfigLoader() clientcmd.ClientConfig {
	return f.clientGetter.ToRawIAMConfigLoader()
}

func (f *factoryImpl) IAMClient() (*iam.IamClient, error) {
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	return iam.NewForConfig(clientConfig)
}

func (f *factoryImpl) RESTClient() (*restclient.RESTClient, error) {
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	setIAMDefaults(clientConfig)
	return restclient.RESTClientFor(clientConfig)
}
