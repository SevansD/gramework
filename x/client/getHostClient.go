// Copyright 2017 Kirill Danshin and Gramework contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//

package client

import (
	"net/url"

	"github.com/valyala/fasthttp"
)

func (client *Instance) getHostClient(addr *url.URL) *fasthttp.HostClient {
	client.clientsMu.RLock()
	if hostClient, ok := client.clients[addr.Host]; ok {
		client.clientsMu.RUnlock()
		return hostClient
	}

	client.clientsMu.RUnlock()
	hostClient := &fasthttp.HostClient{
		Addr:  addr.Host,
		IsTLS: addr.Scheme == httpsScheme,
	}

	client.clientsMu.Lock()
	client.clients[addr.Host] = hostClient
	client.clientsMu.Unlock()
	return hostClient
}
