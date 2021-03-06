// +build tools

package tools

import (
	_ "knative.dev/hack"
	_ "knative.dev/pkg/hack"

	_ "k8s.io/code-generator/cmd/deepcopy-gen"

	// For chaos testing the leaderelection stuff.
	_ "knative.dev/pkg/leaderelection/chaosduck"

	// To get the ingress CRD
	_ "knative.dev/networking/config"

	_ "knative.dev/networking/test/conformance/ingress"
	_ "knative.dev/networking/test/test_images/grpc-ping"
	_ "knative.dev/networking/test/test_images/httpproxy"
	_ "knative.dev/networking/test/test_images/retry"
	_ "knative.dev/networking/test/test_images/runtime"
	_ "knative.dev/networking/test/test_images/timeout"
	_ "knative.dev/networking/test/test_images/wsserver"
)
