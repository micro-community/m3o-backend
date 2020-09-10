package main

import (
	"fmt"
	"strings"

	"github.com/micro/micro/v3/service/logger"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	k8s "github.com/scaleway/scaleway-sdk-go/api/k8s/v1beta4"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/slack-go/slack"
)

func checkInfraUsage() {
	// issues is a string slice containing any wastage. If there are any elements in this slice
	// by the end of this function, it will be reported via Slack to the team to investigate
	var issues []string

	logger.Infof("Starting Infra Usage Check")

	// load the clusters
	clRsp, err := k8sAPI.ListClusters(&k8s.ListClustersRequest{Region: scalewayRegion})
	if err != nil {
		logger.Errorf("Error listing clusters: %v", err)
		return
	}

	logger.Infof("We're running %v clusters", clRsp.TotalCount)
	clusterIDs := make(map[string]bool, clRsp.TotalCount)
	for _, c := range clRsp.Clusters {
		clusterIDs[c.ID] = true
	}

	// load the load balancers
	lbRsp, err := lbAPI.ListLbs(&lb.ListLbsRequest{Region: scalewayRegion})
	if err != nil {
		logger.Errorf("Error listing load balancers: %v", err)
		return
	}

	// check the load balancers for wastage by ensuring the cluster they belong to still exists
	logger.Infof("We're running %v load balancers", lbRsp.TotalCount)
lbLoop:
	for _, l := range lbRsp.Lbs {
		logger.Infof("Inspecting load balancer %v", l.ID)

		for _, t := range l.Tags {
			// tag does not contain cluster id
			if !strings.HasPrefix(t, "cluster=") {
				continue
			}

			// cluster id does not exist (the cluster was probably removed, but the option to delete the
			// associated load balancers was left unchecked)
			cID := strings.TrimPrefix(t, "cluster=")
			if _, ok := clusterIDs[cID]; !ok {
				issues = append(issues, fmt.Sprintf("Load Balancer #%v belongs to Cluster %v which doesn't exist", l.ID, cID))
				continue
			}

			continue lbLoop
		}

		// we're not manually creating load balancers in scaleway, so they shouldn't exist without an
		// associated cluster
		issues = append(issues, fmt.Sprintf("Load Balancer #%v does not belong to any cluster", l.ID))
	}

	// load the servers
	svrRsp, err := inAPI.ListServers(&instance.ListServersRequest{Zone: scalewayZone})
	if err != nil {
		logger.Errorf("Error listing servers: %v", err)
		return
	}

	// check the servers for wastage by ensuring the cluster they belong to still exists
	logger.Infof("We're running %v servers", svrRsp.TotalCount)
	serverIDs := make(map[string]bool, svrRsp.TotalCount)
	for _, c := range svrRsp.Servers {
		serverIDs[c.ID] = true
	}

svrLoop:
	for _, s := range svrRsp.Servers {
		logger.Infof("Inspecting server %v", s.ID)

		for _, t := range s.Tags {
			// tag does not contain cluster id
			if !strings.HasPrefix(t, "kapsule=") {
				continue
			}

			// cluster id does not exist (the cluster was probably removed, but the option to delete the
			// associated load balancers was left unchecked)
			cID := strings.TrimPrefix(t, "kapsule=")
			if _, ok := clusterIDs[cID]; !ok {
				issues = append(issues, fmt.Sprintf("Load Balancer #%v belongs to Cluster %v which doesn't exist", s.ID, cID))
			}
			continue svrLoop
		}

		// we're not manually creating server in scaleway, so they shouldn't exist without an
		// associated cluster
		issues = append(issues, fmt.Sprintf("Server #%v does not belong to any cluster", s.ID))
	}

	// load the volumes
	volRsp, err := inAPI.ListVolumes(&instance.ListVolumesRequest{Zone: scalewayZone})
	if err != nil {
		logger.Errorf("Error listing volumes: %v", err)
		return
	}

	// check the volumes for wastage by ensuring the cluster they belong to still exists
	logger.Infof("We're running %v volumes", volRsp.TotalCount)
	for _, v := range volRsp.Volumes {
		logger.Infof("Inspecting volume %v", v.ID)

		if v.Server == nil {
			issues = append(issues, fmt.Sprintf("Volume #%v has no associated server", v.ID))
			continue
		}

		if _, ok := serverIDs[v.Server.ID]; !ok {
			issues = append(issues, fmt.Sprintf("Volume #%v's belongs to Server %v which doesn't exist", v.ID, v.Server.ID))
		}
	}

	logger.Infof("Infra Usage Check Complete. %v issues have been found.", len(issues))
	for _, i := range issues {
		fmt.Printf("\t - %v\n", i)
	}

	if len(issues) == 0 {
		return
	}

	msg := fmt.Sprintf("*Possible infrastructure wastage detected*:")
	for _, i := range issues {
		msg += fmt.Sprintf("\n- %v", i)
	}

	slackbot.SendMessage("team-important",
		slack.MsgOptionUsername("Infrastructure Service"),
		slack.MsgOptionText(msg, false),
	)
}
