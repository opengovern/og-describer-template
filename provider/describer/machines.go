package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-fly/pkg/sdk/models"
	"github.com/opengovern/og-describer-fly/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListMachines(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	flyChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	apps, err := ListApps(ctx, handler, stream)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(flyChan)
		defer close(errorChan)
		for _, app := range apps {
			if err := processMachines(ctx, handler, app.Name, flyChan, &wg); err != nil {
				errorChan <- err // Send error to the error channel
			}
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-flyChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func processMachines(ctx context.Context, handler *resilientbridge.ResilientBridge, appName string, flyChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var machines []model.MachineJSON
	baseURL := "/v1/apps/"

	finalURL := fmt.Sprintf("%s%s/machines", baseURL, appName)

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("fly", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &machines); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, machine := range machines {
		wg.Add(1)
		go func(machine model.MachineJSON) {
			defer wg.Done()
			configChecks := model.ConfigCheck{
				GracePeriod:   machine.Config.Checks.GracePeriod,
				Headers:       machine.Config.Checks.Headers,
				Interval:      machine.Config.Checks.Interval,
				Kind:          machine.Config.Checks.Kind,
				Method:        machine.Config.Checks.Method,
				Path:          machine.Config.Checks.Path,
				Port:          machine.Config.Checks.Port,
				Protocol:      machine.Config.Checks.Protocol,
				Type:          machine.Config.Checks.Type,
				Timeout:       machine.Config.Checks.Timeout,
				TLSServerName: machine.Config.Checks.TLSServerName,
				TLSSkipVerify: machine.Config.Checks.TLSSkipVerify,
			}
			var configContainers []model.Container
			for _, container := range machine.Config.Containers {
				configContainers = append(configContainers, model.Container{
					Cmd:        container.Cmd,
					DependsOn:  container.DependsOn,
					Env:        container.Env,
					Entrypoint: container.Entrypoint,
					Exec:       container.Exec,
					EnvFrom:    container.EnvFrom,
					Files:      container.Files,
					Image:      container.Image,
					Name:       container.Name,
					Restart:    container.Restart,
					Secrets:    container.Secrets,
					Stop:       container.Stop,
					User:       container.User,
				})
			}
			var dnsForwardRules []model.DNSForwardRule
			for _, dnsForwardRule := range machine.Config.DNS.DNSForwardRules {
				dnsForwardRules = append(dnsForwardRules, model.DNSForwardRule{
					Addr:     dnsForwardRule.Addr,
					Basename: dnsForwardRule.Basename,
				})
			}
			var dnsOptions []model.DNSOption
			for _, dnsOption := range machine.Config.DNS.Options {
				dnsOptions = append(dnsOptions, model.DNSOption{
					Name:  dnsOption.Name,
					Value: dnsOption.Value,
				})
			}
			dns := model.DNS{
				DNSForwardRules:  dnsForwardRules,
				Hostname:         machine.Config.DNS.Hostname,
				HostnameFQDN:     machine.Config.DNS.HostnameFQDN,
				Nameservers:      machine.Config.DNS.Nameservers,
				Options:          dnsOptions,
				Searches:         machine.Config.DNS.Searches,
				SkipRegistration: machine.Config.DNS.SkipRegistration,
			}
			var files []model.File
			for _, file := range machine.Config.Files {
				files = append(files, model.File{
					GuestPath:  file.GuestPath,
					Mode:       file.Mode,
					RawValue:   file.RawValue,
					SecretName: file.SecretName,
				})
			}
			guest := model.Guest{
				CPUKind:          machine.Config.Guest.CPUKind,
				CPUs:             machine.Config.Guest.CPUs,
				GPUKind:          machine.Config.Guest.GPUKind,
				GPUs:             machine.Config.Guest.GPUs,
				HostDedicationID: machine.Config.Guest.HostDedicationID,
				KernelArgs:       machine.Config.Guest.KernelArgs,
				MemoryMB:         machine.Config.Guest.MemoryMB,
			}
			init := model.Init{
				Cmd:        machine.Config.Init.Cmd,
				Exec:       machine.Config.Init.Exec,
				Entrypoint: machine.Config.Init.Entrypoint,
				KernelArgs: machine.Config.Init.KernelArgs,
				SwapSizeMB: machine.Config.Init.SwapSizeMB,
				TTY:        machine.Config.Init.TTY,
			}
			metric := model.Metric{
				HTTPS: machine.Config.Metrics.HTTPS,
				Port:  machine.Config.Metrics.Port,
				Path:  machine.Config.Metrics.Path,
			}
			var mounts []model.Mount
			for _, mount := range machine.Config.Mounts {
				mounts = append(mounts, model.Mount{
					AddSizeGB:              mount.AddSizeGB,
					Encrypted:              mount.Encrypted,
					ExtendThresholdPercent: mount.ExtendThresholdPercent,
					Name:                   mount.Name,
					Path:                   mount.Path,
					SizeGB:                 mount.SizeGB,
					SizeGBLimit:            mount.SizeGBLimit,
					Volume:                 mount.Volume,
				})
			}
			var processes []model.Process
			for _, process := range machine.Config.Processes {
				processes = append(processes, model.Process{
					Cmd:              process.Cmd,
					Exec:             process.Exec,
					Env:              process.Env,
					Entrypoint:       process.Entrypoint,
					EnvFrom:          process.EnvFrom,
					IgnoreAppSecrets: process.IgnoreAppSecrets,
					Secrets:          process.Secrets,
					User:             process.User,
				})
			}
			restart := model.Restart{
				GPUBidPrice: machine.Config.Restart.GPUBidPrice,
				MaxRetries:  machine.Config.Restart.MaxRetries,
				Policy:      machine.Config.Restart.Policy,
			}
			var services []model.Service
			for _, service := range machine.Config.Services {
				concurrency := model.ServiceConcurrency{
					HardLimit: service.Concurrency.HardLimit,
					SoftLimit: service.Concurrency.SoftLimit,
					Type:      service.Concurrency.Type,
				}
				services = append(services, model.Service{
					Autostart:                service.Autostart,
					Autostop:                 service.Autostop,
					Checks:                   service.Checks,
					Concurrency:              concurrency,
					ForceInstanceDescription: service.ForceInstanceDescription,
					ForceInstanceKey:         service.ForceInstanceKey,
					InternalPort:             service.InternalPort,
					MinMachinesRunning:       service.MinMachinesRunning,
					Ports:                    service.Ports,
					Protocol:                 service.Protocol,
				})
			}
			var statics []model.Static
			for _, static := range machine.Config.Statics {
				statics = append(statics, model.Static{
					GuestPath:     static.GuestPath,
					IndexDocument: static.IndexDocument,
					TigrisBucket:  static.TigrisBucket,
					URLPrefix:     static.URLPrefix,
				})
			}
			stopConfig := model.StopConfig{
				Signal: machine.Config.StopConfig.Signal,
				Timeout: struct{ TimeDuration int }{
					TimeDuration: machine.Config.StopConfig.Timeout.TimeDuration,
				},
			}
			config := model.Config{
				AutoDestroy:             machine.Config.AutoDestroy,
				Checks:                  configChecks,
				Containers:              configContainers,
				DNS:                     dns,
				DisableMachineAutostart: machine.Config.DisableMachineAutostart,
				Env:                     machine.Config.Env,
				Files:                   files,
				Guest:                   guest,
				Image:                   machine.Config.Image,
				Init:                    init,
				Metadata:                machine.Config.Metadata,
				Metrics:                 metric,
				Mounts:                  mounts,
				Processes:               processes,
				Restart:                 restart,
				Size:                    machine.Config.Size,
				Schedule:                machine.Config.Schedule,
				Services:                services,
				Standbys:                machine.Config.Standbys,
				Statics:                 statics,
				StopConfig:              stopConfig,
			}
			var checks []model.Check
			for _, check := range machine.Checks {
				checks = append(checks, model.Check{
					Name:      check.Name,
					Output:    check.Output,
					Status:    check.Status,
					UpdatedAt: check.UpdatedAt,
				})
			}
			var events []model.Event
			for _, event := range machine.Events {
				events = append(events, model.Event{
					ID:        event.ID,
					Request:   event.Request,
					Status:    event.Status,
					Source:    event.Source,
					Type:      event.Type,
					Timestamp: event.Timestamp,
				})
			}
			imageRef := model.ImageRef{
				Digest:     machine.ImageRef.Digest,
				Labels:     machine.ImageRef.Labels,
				Registry:   machine.ImageRef.Registry,
				Repository: machine.ImageRef.Repository,
				Tag:        machine.ImageRef.Tag,
			}
			incompleteConfigChecks := model.ConfigCheck{
				GracePeriod:   machine.IncompleteConfig.Checks.GracePeriod,
				Headers:       machine.IncompleteConfig.Checks.Headers,
				Interval:      machine.IncompleteConfig.Checks.Interval,
				Kind:          machine.IncompleteConfig.Checks.Kind,
				Method:        machine.IncompleteConfig.Checks.Method,
				Path:          machine.IncompleteConfig.Checks.Path,
				Port:          machine.IncompleteConfig.Checks.Port,
				Protocol:      machine.IncompleteConfig.Checks.Protocol,
				Type:          machine.IncompleteConfig.Checks.Type,
				Timeout:       machine.IncompleteConfig.Checks.Timeout,
				TLSServerName: machine.IncompleteConfig.Checks.TLSServerName,
				TLSSkipVerify: machine.IncompleteConfig.Checks.TLSSkipVerify,
			}
			var incompleteConfigContainers []model.Container
			for _, container := range machine.IncompleteConfig.Containers {
				incompleteConfigContainers = append(incompleteConfigContainers, model.Container{
					Cmd:        container.Cmd,
					DependsOn:  container.DependsOn,
					Env:        container.Env,
					Entrypoint: container.Entrypoint,
					Exec:       container.Exec,
					EnvFrom:    container.EnvFrom,
					Files:      container.Files,
					Image:      container.Image,
					Name:       container.Name,
					Restart:    container.Restart,
					Secrets:    container.Secrets,
					Stop:       container.Stop,
					User:       container.User,
				})
			}
			var incompleteDNSForwardRules []model.DNSForwardRule
			for _, dnsForwardRule := range machine.IncompleteConfig.DNS.DNSForwardRules {
				incompleteDNSForwardRules = append(incompleteDNSForwardRules, model.DNSForwardRule{
					Addr:     dnsForwardRule.Addr,
					Basename: dnsForwardRule.Basename,
				})
			}
			var incompleteDNSOptions []model.DNSOption
			for _, dnsOption := range machine.IncompleteConfig.DNS.Options {
				incompleteDNSOptions = append(incompleteDNSOptions, model.DNSOption{
					Name:  dnsOption.Name,
					Value: dnsOption.Value,
				})
			}
			incompleteDNS := model.DNS{
				DNSForwardRules:  incompleteDNSForwardRules,
				Hostname:         machine.Config.DNS.Hostname,
				HostnameFQDN:     machine.Config.DNS.HostnameFQDN,
				Nameservers:      machine.Config.DNS.Nameservers,
				Options:          incompleteDNSOptions,
				Searches:         machine.Config.DNS.Searches,
				SkipRegistration: machine.Config.DNS.SkipRegistration,
			}
			var incompleteFiles []model.File
			for _, file := range machine.IncompleteConfig.Files {
				incompleteFiles = append(incompleteFiles, model.File{
					GuestPath:  file.GuestPath,
					Mode:       file.Mode,
					RawValue:   file.RawValue,
					SecretName: file.SecretName,
				})
			}
			incompleteGuest := model.Guest{
				CPUKind:          machine.IncompleteConfig.Guest.CPUKind,
				CPUs:             machine.IncompleteConfig.Guest.CPUs,
				GPUKind:          machine.IncompleteConfig.Guest.GPUKind,
				GPUs:             machine.IncompleteConfig.Guest.GPUs,
				HostDedicationID: machine.IncompleteConfig.Guest.HostDedicationID,
				KernelArgs:       machine.IncompleteConfig.Guest.KernelArgs,
				MemoryMB:         machine.IncompleteConfig.Guest.MemoryMB,
			}
			incompleteInit := model.Init{
				Cmd:        machine.IncompleteConfig.Init.Cmd,
				Exec:       machine.IncompleteConfig.Init.Exec,
				Entrypoint: machine.IncompleteConfig.Init.Entrypoint,
				KernelArgs: machine.IncompleteConfig.Init.KernelArgs,
				SwapSizeMB: machine.IncompleteConfig.Init.SwapSizeMB,
				TTY:        machine.IncompleteConfig.Init.TTY,
			}
			incompleteMetric := model.Metric{
				HTTPS: machine.IncompleteConfig.Metrics.HTTPS,
				Port:  machine.IncompleteConfig.Metrics.Port,
				Path:  machine.IncompleteConfig.Metrics.Path,
			}
			var incompleteMounts []model.Mount
			for _, mount := range machine.IncompleteConfig.Mounts {
				incompleteMounts = append(incompleteMounts, model.Mount{
					AddSizeGB:              mount.AddSizeGB,
					Encrypted:              mount.Encrypted,
					ExtendThresholdPercent: mount.ExtendThresholdPercent,
					Name:                   mount.Name,
					Path:                   mount.Path,
					SizeGB:                 mount.SizeGB,
					SizeGBLimit:            mount.SizeGBLimit,
					Volume:                 mount.Volume,
				})
			}
			var incompleteProcesses []model.Process
			for _, process := range machine.IncompleteConfig.Processes {
				incompleteProcesses = append(incompleteProcesses, model.Process{
					Cmd:              process.Cmd,
					Exec:             process.Exec,
					Env:              process.Env,
					Entrypoint:       process.Entrypoint,
					EnvFrom:          process.EnvFrom,
					IgnoreAppSecrets: process.IgnoreAppSecrets,
					Secrets:          process.Secrets,
					User:             process.User,
				})
			}
			incompleteRestart := model.Restart{
				GPUBidPrice: machine.IncompleteConfig.Restart.GPUBidPrice,
				MaxRetries:  machine.IncompleteConfig.Restart.MaxRetries,
				Policy:      machine.IncompleteConfig.Restart.Policy,
			}
			var incompleteServices []model.Service
			for _, service := range machine.IncompleteConfig.Services {
				concurrency := model.ServiceConcurrency{
					HardLimit: service.Concurrency.HardLimit,
					SoftLimit: service.Concurrency.SoftLimit,
					Type:      service.Concurrency.Type,
				}
				incompleteServices = append(incompleteServices, model.Service{
					Autostart:                service.Autostart,
					Autostop:                 service.Autostop,
					Checks:                   service.Checks,
					Concurrency:              concurrency,
					ForceInstanceDescription: service.ForceInstanceDescription,
					ForceInstanceKey:         service.ForceInstanceKey,
					InternalPort:             service.InternalPort,
					MinMachinesRunning:       service.MinMachinesRunning,
					Ports:                    service.Ports,
					Protocol:                 service.Protocol,
				})
			}
			var incompleteStatics []model.Static
			for _, static := range machine.IncompleteConfig.Statics {
				incompleteStatics = append(incompleteStatics, model.Static{
					GuestPath:     static.GuestPath,
					IndexDocument: static.IndexDocument,
					TigrisBucket:  static.TigrisBucket,
					URLPrefix:     static.URLPrefix,
				})
			}
			incompleteStopConfig := model.StopConfig{
				Signal: machine.IncompleteConfig.StopConfig.Signal,
				Timeout: struct{ TimeDuration int }{
					TimeDuration: machine.IncompleteConfig.StopConfig.Timeout.TimeDuration,
				},
			}
			incompleteConfig := model.Config{
				AutoDestroy:             machine.Config.AutoDestroy,
				Checks:                  incompleteConfigChecks,
				Containers:              incompleteConfigContainers,
				DNS:                     incompleteDNS,
				DisableMachineAutostart: machine.Config.DisableMachineAutostart,
				Env:                     machine.Config.Env,
				Files:                   incompleteFiles,
				Guest:                   incompleteGuest,
				Image:                   machine.Config.Image,
				Init:                    incompleteInit,
				Metadata:                machine.Config.Metadata,
				Metrics:                 incompleteMetric,
				Mounts:                  incompleteMounts,
				Processes:               incompleteProcesses,
				Restart:                 incompleteRestart,
				Size:                    machine.Config.Size,
				Schedule:                machine.Config.Schedule,
				Services:                incompleteServices,
				Standbys:                machine.Config.Standbys,
				Statics:                 incompleteStatics,
				StopConfig:              incompleteStopConfig,
			}
			value := models.Resource{
				ID:   machine.ID,
				Name: machine.Name,
				Description: model.MachineDescription{
					Config:           config,
					CreatedAt:        machine.CreatedAt,
					Checks:           checks,
					Events:           events,
					HostStatus:       machine.HostStatus,
					ID:               machine.ID,
					ImageRef:         imageRef,
					IncompleteConfig: incompleteConfig,
					InstanceID:       machine.InstanceID,
					Name:             machine.Name,
					Nonce:            machine.Nonce,
					PrivateIP:        machine.PrivateIP,
					Region:           machine.Region,
					State:            machine.State,
					UpdatedAt:        machine.UpdatedAt,
				},
			}
			flyChan <- value
		}(machine)
	}
	return nil
}
