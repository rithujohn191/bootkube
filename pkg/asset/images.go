package asset

// DefaultImages are the defualt images bootkube components use.
var DefaultImages = ImageVersions{
	Etcd:            "quay.io/coreos/etcd:v3.1.8",
	EtcdOperator:    "quay.io/coreos/etcd-operator:v0.3.2",
	Flannel:         "quay.io/coreos/flannel:v0.7.1-amd64",
	Calico:          "quay.io/calico/node:v1.2.1",
	CalicoCNI:       "quay.io/calico/cni:v1.8.3",
	Hyperkube:       "quay.io/coreos/hyperkube:v1.6.4_coreos.0",
	Kenc:            "quay.io/coreos/kenc:8f6e2e885f790030fbbb0496ea2a2d8830e58b8f",
	KubeDNS:         "gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.1",
	KubeDNSMasq:     "gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.1",
	KubeDNSSidecar:  "gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.1",
	PodCheckpointer: "quay.io/coreos/pod-checkpointer:4e7a7dab10bc4d895b66c21656291c6e0b017248",
}

// DefaultCNIRelease is the default version of cni release
var DefaultCNIRelease = "0799f5732f2a11b329d9e3d51b9c8f2e3759f2ff"
