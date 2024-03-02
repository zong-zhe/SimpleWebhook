kubebuilder-init:
# 我们将使用 tutorial.kubebuilder.io 域，
# 所以所有的 API 组将是<group>.tutorial.kubebuilder.io.
	kubebuilder init --domain tutorial.kubebuilder.io

kubebuilder-create-api:
	kubebuilder create api --group batchdemo --version v1alpha --kind DemoCronJob

kubebuilder-create-webhook:
	kubebuilder create webhook --group batchdemo --version v1alpha --kind DemoCronJob --defaulting --programmatic-validation
