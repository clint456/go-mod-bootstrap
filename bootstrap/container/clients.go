//
// Copyright (c) 2022 Intel Corporation
// Copyright (C) 2024-2025 IOTech Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"

	"github.com/edgexfoundry/go-mod-bootstrap/v4/di"
)

// EventClientName contains the name of the EventClient's implementation in the DIC.
var EventClientName = di.TypeInstanceToName((*interfaces.EventClient)(nil))

// EventClientFrom helper function queries the DIC and returns the EventClient's implementation.
func EventClientFrom(get di.Get) interfaces.EventClient {
	if get(EventClientName) == nil {
		return nil
	}

	return get(EventClientName).(interfaces.EventClient)
}

// ReadingClientName contains the name of the ReadingClient instance in the DIC.
var ReadingClientName = di.TypeInstanceToName((*interfaces.ReadingClient)(nil))

// ReadingClientFrom helper function queries the DIC and returns the ReadingClient instance.
func ReadingClientFrom(get di.Get) interfaces.ReadingClient {
	client, ok := get(ReadingClientName).(interfaces.ReadingClient)
	if !ok {
		return nil
	}

	return client
}

// CommandClientName contains the name of the CommandClient's implementation in the DIC.
var CommandClientName = di.TypeInstanceToName((*interfaces.CommandClient)(nil))

// CommandClientFrom helper function queries the DIC and returns the CommandClient's implementation.
func CommandClientFrom(get di.Get) interfaces.CommandClient {
	if get(CommandClientName) == nil {
		return nil
	}

	return get(CommandClientName).(interfaces.CommandClient)
}

// NotificationClientName contains the name of the NotificationClient's implementation in the DIC.
var NotificationClientName = di.TypeInstanceToName((*interfaces.NotificationClient)(nil))

// NotificationClientFrom helper function queries the DIC and returns the NotificationClient's implementation.
func NotificationClientFrom(get di.Get) interfaces.NotificationClient {
	if get(NotificationClientName) == nil {
		return nil
	}

	return get(NotificationClientName).(interfaces.NotificationClient)
}

// SubscriptionClientName contains the name of the SubscriptionClient's implementation in the DIC.
var SubscriptionClientName = di.TypeInstanceToName((*interfaces.SubscriptionClient)(nil))

// SubscriptionClientFrom helper function queries the DIC and returns the SubscriptionClient's implementation.
func SubscriptionClientFrom(get di.Get) interfaces.SubscriptionClient {
	if get(SubscriptionClientName) == nil {
		return nil
	}

	return get(SubscriptionClientName).(interfaces.SubscriptionClient)
}

// NodeServiceClientName contains the name of the NodeServiceClient's implementation in the DIC.
var NodeServiceClientName = di.TypeInstanceToName((*interfaces.NodeServiceClient)(nil))

// NodeServiceClientFrom helper function queries the DIC and returns the NodeServiceClient's implementation.
func NodeServiceClientFrom(get di.Get) interfaces.NodeServiceClient {
	if get(NodeServiceClientName) == nil {
		return nil
	}

	return get(NodeServiceClientName).(interfaces.NodeServiceClient)
}

// ProfileClientName contains the name of the ProfileClient's implementation in the DIC.
var ProfileClientName = di.TypeInstanceToName((*interfaces.ProfileClient)(nil))

// ProfileClientFrom helper function queries the DIC and returns the ProfileClient's implementation.
func ProfileClientFrom(get di.Get) interfaces.ProfileClient {
	if get(ProfileClientName) == nil {
		return nil
	}

	return get(ProfileClientName).(interfaces.ProfileClient)
}

// NodeClientName contains the name of the NodeClient's implementation in the DIC.
var NodeClientName = di.TypeInstanceToName((*interfaces.NodeClient)(nil))

// NodeClientFrom helper function queries the DIC and returns the NodeClient's implementation.
func NodeClientFrom(get di.Get) interfaces.NodeClient {
	if get(NodeClientName) == nil {
		return nil
	}

	return get(NodeClientName).(interfaces.NodeClient)
}

// ProvisionWatcherClientName contains the name of the ProvisionWatcherClient's implementation in the DIC.
var ProvisionWatcherClientName = di.TypeInstanceToName((*interfaces.ProvisionWatcherClient)(nil))

// ProvisionWatcherClientFrom helper function queries the DIC and returns the ProvisionWatcherClient's implementation.
func ProvisionWatcherClientFrom(get di.Get) interfaces.ProvisionWatcherClient {
	if get(ProvisionWatcherClientName) == nil {
		return nil
	}

	return get(ProvisionWatcherClientName).(interfaces.ProvisionWatcherClient)
}

// NodeServiceCommandClientName contains the name of the NodeServiceCommandClient instance in the DIC.
var NodeServiceCommandClientName = di.TypeInstanceToName((*interfaces.NodeServiceCommandClient)(nil))

// NodeServiceCommandClientFrom helper function queries the DIC and returns the NodeServiceCommandClient instance.
func NodeServiceCommandClientFrom(get di.Get) interfaces.NodeServiceCommandClient {
	client, ok := get(NodeServiceCommandClientName).(interfaces.NodeServiceCommandClient)
	if !ok {
		return nil
	}

	return client
}

// ScheduleJobClientName contains the name of the ScheduleJobClient's implementation in the DIC.
var ScheduleJobClientName = di.TypeInstanceToName((*interfaces.ScheduleJobClient)(nil))

// ScheduleJobClientFrom helper function queries the DIC and returns the ScheduleJobClient's implementation.
func ScheduleJobClientFrom(get di.Get) interfaces.ScheduleJobClient {
	if get(ScheduleJobClientName) == nil {
		return nil
	}

	return get(ScheduleJobClientName).(interfaces.ScheduleJobClient)
}

// ScheduleActionRecordClientName contains the name of the ScheduleActionRecordClient's implementation in the DIC.
var ScheduleActionRecordClientName = di.TypeInstanceToName((*interfaces.ScheduleActionRecordClient)(nil))

// ScheduleActionRecordClientFrom helper function queries the DIC and returns the ScheduleActionRecordClient's implementation.
func ScheduleActionRecordClientFrom(get di.Get) interfaces.ScheduleActionRecordClient {
	if get(ScheduleActionRecordClientName) == nil {
		return nil
	}

	return get(ScheduleActionRecordClientName).(interfaces.ScheduleActionRecordClient)
}

// SecurityProxyAuthClientName contains the name of the AuthClient's implementation in the DIC.
var SecurityProxyAuthClientName = di.TypeInstanceToName((*interfaces.AuthClient)(nil))

// SecurityProxyAuthClientFrom helper function queries the DIC and returns the AuthClient's implementation.
func SecurityProxyAuthClientFrom(get di.Get) interfaces.AuthClient {
	if get(SecurityProxyAuthClientName) == nil {
		return nil
	}

	return get(SecurityProxyAuthClientName).(interfaces.AuthClient)
}
