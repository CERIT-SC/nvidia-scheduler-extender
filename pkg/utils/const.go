package utils

const (
	ResourceName = "cerit.io/gpu-mem"
	CountName    = "cerit.io/gpu-count"

	EnvNVGPU              = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex      = "CERIT_IO_GPU_MEM_IDX"
	EnvResourceByPod      = "CERIT_IO_GPU_MEM_POD"
	EnvResourceByDev      = "CERIT_IO_GPU_MEM_DEV"
	EnvAssignedFlag       = "CERIT_IO_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime = "CERIT_IO_GPU_MEM_ASSUME_TIME"
)
