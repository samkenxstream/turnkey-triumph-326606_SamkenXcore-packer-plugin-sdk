// This is an example plugin. In packer 0.9.0 and up, core plugins are compiled
// into the main binary so these files are no longer necessary for the packer
// project.
//
// However, it is still possible to create a third-party plugin for packer that
// is distributed independently from the packer distribution. These continue to
// work in the same way. They will be loaded from the same directory as packer
// by looking for packer-[builder|provisioner|post-processor]-plugin-name. For
// example:
//
//    packer-builder-docker
//
// Look at command/plugin.go to see how the core plugins are loaded now, but the
// format below was used for packer <= 0.8.6 and is forward-compatible.
package main

import (
	"github.com/hashicorp/packer-plugin-amazon/builder/chroot"
	dockerpush "github.com/hashicorp/packer-plugin-docker/post-processor/docker-push"
	pluginsdk "github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer/provisioner/powershell"
)

func main() {
	server, err := pluginsdk.Server()
	if err != nil {
		panic(err)
	}
	// Choose the appropriate type of plugin. You should only use one of these
	// at a time, which means you will have a separate plugin for each builder,
	// provisioner, or post-processor.
	server.RegisterBuilder(new(chroot.Builder))
	server.RegisterPostProcessor(new(dockerpush.PostProcessor))
	server.RegisterProvisioner(new(powershell.Provisioner))
	server.Serve()
}
