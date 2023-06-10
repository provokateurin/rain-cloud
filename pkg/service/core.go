package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strings"

	"github.com/provokateurin/rain-cloud/pkg/registration"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type coreService struct {
	kubernetesClientset *kubernetes.Clientset
}

//nolint:gochecknoglobals
var CoreService = &coreService{}

func (s *coreService) GenerateAvatar(name string, size int, dark bool) (avatar []byte, contentType string, err error) {
	var background, foreground color.Color
	if dark {
		background = color.Black
		foreground = color.White
	} else {
		background = color.White
		foreground = color.Black
	}

	img := image.NewRGBA(
		image.Rectangle{
			Min: image.Point{},
			Max: image.Point{
				X: size,
				Y: size,
			},
		},
	)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			img.Set(x, y, background)
		}
	}

	point := fixed.Point26_6{X: fixed.I(0), Y: fixed.I(size / 2)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(foreground),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(name)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return nil, "", fmt.Errorf("failed to encode png: %w", err)
	}

	return buf.Bytes(), "image/png", nil
}

func (s *coreService) getKubernetesClient() (*kubernetes.Clientset, error) {
	if s.kubernetesClientset == nil {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get in-cluster-config: %w", err)
		}

		s.kubernetesClientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create clientset: %w", err)
		}
	}

	return s.kubernetesClientset, nil
}

func (s *coreService) GetRegistrations(ctx context.Context) ([]*registration.AppRegistration, error) {
	kubernetesClient, err := s.getKubernetesClient()
	if err != nil {
		return nil, err
	}

	serviceList, err := kubernetesClient.CoreV1().Services("rain-cloud").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}

	var registrations []*registration.AppRegistration

	//nolint:gocritic
	for i := range serviceList.Items {
		service := &serviceList.Items[i]
		if strings.HasPrefix(service.Name, "app-") {
			address := fmt.Sprintf("http://%s:%d", service.Name, service.Spec.Ports[0].Port)
			client, err := registration.NewClientWithResponses(address)
			if err != nil {
				return nil, fmt.Errorf("failed to create client: %w", err)
			}

			r, err := client.GetRegistrationWithResponse(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get registration: %w", err)
			}
			if r.JSON200 == nil {
				//nolint:goerr113
				return nil, fmt.Errorf("failed to get registration: %d", r.StatusCode())
			}

			registrations = append(registrations, r.JSON200)
		}
	}

	return registrations, nil
}

func (s *coreService) GetCapabilities(ctx context.Context) (map[string]map[string]interface{}, error) {
	registrations, err := s.GetRegistrations(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get registrations: %w", err)
	}

	capabilities := map[string]map[string]interface{}{}
	for _, r := range registrations {
		if r.Capabilities != nil {
			for key1, value1 := range *r.Capabilities {
				if _, ok := capabilities[key1]; ok {
					for key2, value2 := range value1 {
						capabilities[key1][key2] = value2
					}
				} else {
					capabilities[key1] = value1
				}
			}
		}
	}

	return capabilities, nil
}

func (s *coreService) GetNavigations(ctx context.Context) ([]*registration.NavigationEntry, error) {
	registrations, err := s.GetRegistrations(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get registrations: %w", err)
	}

	var entries []*registration.NavigationEntry
	for _, r := range registrations {
		if r.Navigation != nil {
			entries = append(entries, r.Navigation)
		}
	}

	return entries, nil
}
