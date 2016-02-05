// Copyright 2016 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

// Plan represents a service plan
type Plan struct {
	Name        string
	Description string
}

func GetPlansByServiceName(serviceName string) ([]Plan, error) {
	s := Service{Name: serviceName}
	err := s.Get()
	if err != nil {
		return nil, err
	}
	endpoint, err := s.getClient("production")
	if err != nil {
		return []Plan{}, nil
	}
	plans, err := endpoint.Plans()
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func GetPlanByServiceNameAndPlanName(serviceName, planName string) (Plan, error) {
	plans, err := GetPlansByServiceName(serviceName)
	if err != nil {
		return Plan{}, err
	}
	for _, plan := range plans {
		if plan.Name == planName {
			return plan, nil
		}
	}
	return Plan{}, nil
}
